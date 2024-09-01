import { defineStore, acceptHMRUpdate } from "pinia";
import { User } from "../types/index";

interface GetUser {
  getUser: User;
}

const EMPTY_USER: User = {
  username: "",
  vanity: "",
  avatar: null,
  bio: null,
  locale: null, // Only user locale on active user.
  followers: null,
  following: null,
  deleted: null,
  flags: null,
};

export const useUser = defineStore("user", {
  state: () => EMPTY_USER,

  actions: {
    /**
     * Fetch username, vanity and flags data from connected user.
     */
    async fetchMe(forceFetching: boolean = false): Promise<void> {
      // Get session
      const session: string = useCookie("session").value || "";
      if (!forceFetching && (session === "" || this.vanity !== "")) return;

      // Set header
      const headers = new Headers();
      headers.append("Authorization", session);

      // Prepare GraphQL query
      const query = gql`
        query GetUser($vanity: String!) {
          getUser(vanity: $vanity) {
            vanity
            username
            avatar
            flags
          }
        }
      `;

      // Make request
      const { data, error } = await useAsyncQuery<GetUser>(query, {
        vanity: "@me",
      });

      // Set user in data
      this.$patch(data.value?.getUser || EMPTY_USER);
    },

    /**
     * Logout user by removing cookie and set it to null.
     */
    logout(): void {
      useCookie("session").value = null;
      this.$patch(EMPTY_USER);
    },
  },
});

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useUser, import.meta.hot));
}
