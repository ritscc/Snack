import { invoke } from "@tauri-apps/api/tauri";

export interface Mention {
  everyone: boolean;
  user_groups: number[]; // Option<Vec<i64>> は TypeScript で number[] または undefined
  member_roles: string[]; // Option<Vec<String>> は TypeScript で string[] または undefined
  channels: number[]; // Option<Vec<i64>> は TypeScript で number[] または undefined
}

export interface Message {
  message_id: number; // i64 は TypeScript で number
  channel_id: number;
  user_id: number;
  content: string;
  created_time: Date; // DateTime<Utc> は TypeScript で Date
  updated_time: Date; // DateTime<Utc> は TypeScript で Date
  mention?: Mention; // Option<Mention>
  pinned: boolean;
  stamp_id: number; // Vec<i64> は TypeScript で number[]
}

async function get_test_messages() {
  try {
    const messages = await invoke<Message[]>("get_test_messages");
    return messages;
  } catch (err) {
    console.error(err);
    return null;
  }
}

export default get_test_messages;
