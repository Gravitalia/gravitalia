export interface User {
  username: string;
  vanity: string;
  avatar?: string | null;
  bio?: string | null;
  locale?: string | null;
  followers?: number | null;
  following?: string | null;
  deleted?: boolean | null;
  flags?: number | null;
}
