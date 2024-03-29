package database

import (
	"context"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Gravitalia/gravitalia/helpers"
	"github.com/Gravitalia/gravitalia/model"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var (
	ctx     = context.Background()
	Session neo4j.SessionWithContext
)

// Init create the main variable for neo4j connection
func Init() {
	driver, _ := neo4j.NewDriverWithContext(os.Getenv("GRAPH_URL"), neo4j.BasicAuth(os.Getenv("GRAPH_USERNAME"), os.Getenv("GRAPH_PASSWORD"), ""))
	Session = driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	Mem = memcache.New(os.Getenv("MEM_URL"))

	_, err := Session.Run(ctx, "CREATE CONSTRAINT ON (u:User) ASSERT u.name IS UNIQUE;", nil)
	if err != nil {
		log.Printf("Cannot create constraints on User: %v", err)
	}

	_, err = Session.Run(ctx, "CREATE CONSTRAINT ON (p:Post) ASSERT p.id IS UNIQUE;", nil)
	if err != nil {
		log.Printf("Cannot create constraints on Post: %v", err)
	}

	_, err = Session.Run(ctx, "CREATE CONSTRAINT ON (c:Comment) ASSERT c.id IS UNIQUE;", nil)
	if err != nil {
		log.Printf("Cannot create constraints on Comment: %v", err)
	}

	_, err = Session.Run(ctx, "CREATE INDEX ON :User(name);", nil)
	if err != nil {
		log.Printf("Cannot create index on User: %v", err)
	}

	_, err = Session.Run(ctx, "CREATE INDEX ON :Post(id);", nil)
	if err != nil {
		log.Printf("Cannot create index on Post: %v", err)
	}
}

// MakeRequest is a simple way to send a query
func MakeRequest(query string, params map[string]any) (any, error) {
	data, err := Session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		result, err := transaction.Run(ctx,
			query,
			params)
		if err != nil {
			return nil, err
		}

		if result.Next(ctx) {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return nil, err
	}

	return data, nil
}

// CreateUser allows to create a new user into the graph database
func CreateUser(id string) (bool, error) {
	_, err := MakeRequest("MERGE (:User {name: $id, public: true, suspended: false});",
		map[string]any{"id": id})
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetProfile returns followers, following and other account data of the desired user
func GetProfile(id string) (model.Profile, error) {
	var profile model.Profile

	_, err := Session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		result, err := transaction.Run(ctx,
			"MATCH (n:User {name: $id}) OPTIONAL MATCH (n)-[:SUBSCRIBER]->(d:User) OPTIONAL MATCH (n)<-[:SUBSCRIBER]-(u:User) OPTIONAL MATCH (n)-[:CREATE]->(p:Post) RETURN count(DISTINCT u) AS followers, count(DISTINCT d) AS following, n.public, n.suspended, count(DISTINCT p) as postNumber;",
			map[string]any{"id": id})
		if err != nil {
			return nil, err
		}

		for result.Next(ctx) {
			if result.Record().Values[2] == nil {
				return nil, errors.New("invalid user")
			}

			profile.Followers = uint32(result.Record().Values[0].(int64))
			profile.Following = uint32(result.Record().Values[1].(int64))
			profile.Public = result.Record().Values[2].(bool)
			profile.Suspended = result.Record().Values[3].(bool)
			profile.PostCount = uint16(result.Record().Values[4].(int64))
		}

		return profile, nil
	})
	if err != nil {
		return model.Profile{Followers: 0, Following: 0}, err
	}

	return profile, nil
}

// GetBasicProfile returns public and suspended
func GetBasicProfile(id string) (model.Profile, error) {
	var profile model.Profile

	_, err := Session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		result, err := transaction.Run(ctx,
			"MATCH (u:User {name: $id}) RETURN u.public, u.suspended;",
			map[string]any{"id": id})
		if err != nil {
			return nil, err
		}

		for result.Next(ctx) {
			profile.Public = result.Record().Values[0].(bool)
			profile.Suspended = result.Record().Values[1].(bool)
		}

		return profile, nil
	})
	if err != nil {
		return model.Profile{Followers: 0, Following: 0}, err
	}

	return profile, nil
}

// GetUserPost is a function for getting every posts of a user
// and see their likes
func GetUserPost(id string, skip uint8) ([]model.Post, error) {
	list := make([]model.Post, 0)

	_, err := Session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		result, err := transaction.Run(ctx,
			"MATCH (u:User {name: $id})-[:CREATE]->(p:Post)-[:CONTAINS]->(m:Media) OPTIONAL MATCH (p)<-[l:LIKE]-(liker:User) RETURN p.id as id, collect(m.hash), p.description, p.text, count(DISTINCT l) ORDER BY id DESC SKIP 0 LIMIT 12;",
			map[string]any{"id": id, "skip": skip * 12})
		if err != nil {
			return nil, err
		}

		pos := 0
		for result.Next(ctx) {
			if result.Record().Values[0] == nil {
				return list, nil
			}

			record := result.Record()
			list = append(list, model.Post{})

			list[pos].Id = record.Values[0].(string)
			list[pos].Hash = record.Values[1].([]any)
			list[pos].Description = record.Values[2].(string)
			list[pos].Text = record.Values[3].(string)
			list[pos].Like = record.Values[4].(int64)

			pos++
		}

		return list, nil
	})
	if err != nil {
		return list, err
	}

	return list, nil
}

// UserRelation create a new relation (edge) between two nodes
func UserRelation(id string, to string, relationType string) (bool, error) {
	var content string
	switch relationType {
	case "SUBSCRIBER", "BLOCK", "REQUEST":
		content = "User"
	case "LIKE", "VIEW":
		content = "Post"
	case "LOVE":
		content = "Comment"
	}

	var identifier string
	if content == "User" {
		identifier = "name"
	} else {
		identifier = "id"
	}

	res, err := MakeRequest("MATCH (a:User {name: $id}) MATCH (b:"+content+"{"+identifier+": $to) OPTIONAL MATCH (a)-[r:"+relationType+"]->(b) DELETE r WITH a, b, count(r) AS deleted_count WHERE deleted_count = 0 CREATE (a)-[:BLOCK]->(b)",
		map[string]any{"id": id, "to": to})
	if err != nil {
		return false, err
	} else if res == nil {
		return false, errors.New("invalid " + content)
	}

	return true, nil
}

// UserUnRelation delete a relation (edge) between two nodes
func UserUnRelation(id string, to string, relationType string) (bool, error) {
	var content string
	switch relationType {
	case "SUBSCRIBER", "BLOCK", "REQUEST":
		content = "User"
	case "LIKE", "VIEW":
		content = "Post"
	case "LOVE":
		content = "Comment"
	}

	var identifier string
	if content == "User" {
		identifier = "name"
	} else {
		identifier = "id"
	}

	_, err := MakeRequest("MATCH (a:User {name: $id})-[r:"+relationType+"]->(b:"+content+" {"+identifier+": $to}) DELETE r QUERY MEMORY LIMIT 1 KB;",
		map[string]any{"id": id, "to": to})
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetPost allows to get data of a post
func GetPost(id string, user string) (model.Post, error) {
	var post model.Post

	_, err := Session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		result, err := transaction.Run(ctx,
			"MATCH (author:User)-[:CREATE]->(p:Post {id: $id}) OPTIONAL MATCH (p)-[:CONTAINS]-(m:Media) OPTIONAL MATCH (p)<-[:LIKE]-(likeUser:User) OPTIONAL MATCH (p)<-[:COMMENT]-(c:Comment)<-[:WROTE]-(u:User) OPTIONAL MATCH (c)-[love:LOVE]-(lover:User {name: $user}) WITH author, p, lover, COLLECT(m.hash) AS hash, COUNT(DISTINCT likeUser) AS numLikes, c, u, COUNT(DISTINCT love) AS loveComment WITH author, p, hash, numLikes, COLLECT({id: c.id, text: c.text, timestamp: c.timestamp, user: u.name, love: loveComment, me_loved: lover.name IS NOT NULL})[..20] AS comments RETURN p.id, hash, p.description, p.text, numLikes, author.name, comments;",
			map[string]any{"id": id, "user": user})
		if err != nil {
			return nil, err
		}

		if result.Next(ctx) {
			if result.Record().Values[0] == nil {
				return nil, errors.New("invalid post")
			}
			record := result.Record()

			post.Id = record.Values[0].(string)
			post.Hash = record.Values[1].([]any)
			post.Description = record.Values[2].(string)
			post.Text = record.Values[3].(string)
			post.Like = record.Values[4].(int64)
			post.Author = record.Values[5].(string)
			post.Comments = record.Values[6].([]any)

			return post, nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return model.Post{}, err
	}

	return post, nil
}

// IsUserSubscrirerTo check if a user (id) is subscrired to another one (user)
// and respond with true if a relation (edge) exists
// or with false if no relation exists
func IsUserSubscrirerTo(id string, user string) (bool, error) {
	res, err := MakeRequest("MATCH (a:User {name: $id})-[:SUBSCRIBER]->(b:User {name: $to}) RETURN a;",
		map[string]any{"id": id, "to": user})
	if err != nil {
		return false, err
	}

	if res != nil {
		return true, nil
	} else {
		return false, nil
	}
}

// CommentPost allows to post a comment on a post
func CommentPost(id string, user string, content string) (string, error) {
	comment_id := helpers.Generate()

	_, err := MakeRequest("CREATE (c:Comment {id: $comment_id, text: $content, timestamp: "+strconv.FormatInt(time.Now().Unix(), 10)+"}) WITH c MATCH (p:Post {id: $to}) MATCH (u:User {name: $id}) CREATE (c)-[:COMMENT]->(p) CREATE (u)-[:WROTE]->(c);", map[string]any{"id": user, "to": id, "comment_id": comment_id, "content": content})
	if err != nil {
		return "", err
	}

	return comment_id, nil
}

// CommentReply allows to post a comment on another comment
func CommentReply(id string, user string, content string, original_comment string) (string, error) {
	comment_id := helpers.Generate()

	_, err := MakeRequest("CREATE (new_comment:Comment {id: $comment_id, text: $content, timestamp: "+strconv.FormatInt(time.Now().Unix(), 10)+"}) WITH new_comment MATCH (:Comment {id: $to})<-[:WROTE]-(u:User) SET new_comment.replied_to = u.name WITH new_comment MATCH (u:User {name: $id}) WITH new_comment, u MATCH (o_comment:Comment {id: $original_comment}) CREATE (new_comment)-[:REPLY]->(o_comment) CREATE (u)-[:WROTE]->(new_comment);", map[string]any{"id": user, "to": id, "comment_id": comment_id, "content": content, "original_comment": original_comment})
	if err != nil {
		return "", err
	}

	return comment_id, nil
}

// GetComments sends 20 comments of a post
func GetComments(id string, skip int, user string) ([]any, error) {
	res, err := MakeRequest("MATCH (:Post {id: $id})<-[:COMMENT]-(c:Comment)<-[:WROTE]-(u:User) OPTIONAL MATCH (c:Comment)-[love:LOVE]-(lover:User) WITH  u, lover, c, count(DISTINCT love) as loveComment WITH collect({id: c.id, text: c.text, timestamp: c.timestamp, user: u.name, love: loveComment, me_loved: lover.name = $user }) as comments SKIP $skip LIMIT 20 RETURN comments;",
		map[string]any{"id": id, "skip": skip, "user": user})
	if err != nil {
		return nil, err
	}

	if res != nil {
		return res.([]any), nil
	} else {
		return nil, nil
	}
}

// GetReply sends 20 replies of a comment
func GetReply(post_id string, id string, skip int, user string) ([]any, error) {
	res, err := MakeRequest("MATCH (:Post {id: $post_id})<-[:COMMENT]-(:Comment {id: $id})<-[:REPLY]-(c:Comment)<-[:WROTE]-(u:User) OPTIONAL MATCH (c:Comment)-[love:LOVE]-(lover:User) WITH  u, lover, c, count(DISTINCT love) as loveComment WITH collect({id: c.id, text: c.text, timestamp: c.timestamp, user: u.name, love: loveComment, me_loved: lover.name = $user }) as comments SKIP $skip LIMIT 20 RETURN comments;",
		map[string]any{"post_id": post_id, "id": id, "skip": skip, "user": user})
	if err != nil {
		return nil, err
	}

	if res != nil {
		return res.([]any), nil
	} else {
		return nil, nil
	}
}

// CreatePost allows to create a new post into database
func CreatePost(user string, tag string, legend string, hash []string) (string, error) {
	id := helpers.Generate()

	_, err := MakeRequest("CREATE (p:Post {id: $id, text: $text, description: ''}) FOREACH (	hash IN $hashArray | MERGE (m:Media {type: 'image', hash: hash}) CREATE (p)-[:CONTAINS]->(m)	) WITH p MERGE (t:Tag {name: $tag}) CREATE (p)-[r:SHOW]->(t) WITH p MATCH (u:User {name: $user}) CREATE (u)-[r:CREATE]->(p);",
		map[string]any{"id": id, "user": user, "tag": tag, "text": legend, "hashArray": hash})
	if err != nil {
		return "", err
	}

	return id, nil
}
