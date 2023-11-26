// テスト
import { invoke } from "@tauri-apps/api/tauri";
import useSWR from "swr";

export type UserStatus = "ONLINE" | "OFF";

export interface UserOtherServiceAccounts {
  x: string;
  github: string;
  discord: string;
}

export interface User {
  id: number;
  username: string;
  nick: string;
  realname: string;
  avatar: string;
  role: string[];
  locale: string;
  ritsId: string;
  ownEmail: string | null;
  comment: string;
  status: UserStatus;
  service?: UserOtherServiceAccounts;
  isDeleted: boolean;
}

//ユーザーとユーザーズのfetcher定義
const fetchUsers = (): Promise<User[]> => invoke("get_test_users_data");
const fetchUser = (id: number): Promise<User> => invoke("get_test_user_data", { id });

// 全ユーザーを取得するカスタムフック
export function useUsers() {
  const { data, error } = useSWR<User[]>("get_test_users_data", fetchUsers);

  return {
    users: data,
    isLoading: !error && !data,
    isError: !!error,
  };
}

// 特定のユーザーを取得するカスタムフック
export function useUser(id: number) {
  const { data, error } = useSWR<User>(`get_test_user_data/${id}`, () => fetchUser(id));

  return {
    user: data,
    isLoading: !error && !data,
    isError: !!error,
  };
}

//useEffect用
// export async function get_test_users(): Promise<User[] | null> {
//   try {
//     const users: User[] = await invoke("get_test_users");
//     return users;
//   } catch (err) {
//     console.error(err);
//     return null;
//   }
// }

// export async function get_test_user_data(id: number): Promise<User | null> {
//   try {
//     const user: User = await invoke("get_test_user_data", { id });
//     return user;
//   } catch (err) {
//     console.error(err);
//     return null;
//   }
// }

// async function get_users() {
//   // Todo: Change Function: get_users
//   const users = await invoke("get_test_users").catch((err) => {
//     console.error(err);
//     return null;
//   });
// }

// async function get_user(id: number) {
//   // Todo: Change Function: get_user
//   const users = await invoke("get_test_user_data", { id }).catch((err) => {
//     console.error(err);
//     return null;
//   });
// }
