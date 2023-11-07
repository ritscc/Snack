import { invoke } from "@tauri-apps/api/tauri";
import { channel } from "diagnostics_channel";

export interface DefaultChannel {
  channel_id: number;
  position: number;
  name: string;
  private: boolean;
  nsfw: boolean;
  user_ids: number[]; // Rust の Vec<i64> は TypeScript で number[] に対応
}

export interface PrivateChannel {
  channel_id: number;
  user_id: number;
}

export interface DmChannel {
  channel_id: number;
  user_id: number;
  other_user_id: number;
}

// ChannelType ユニオン型を定義
export type ChannelType =
  | { DefaultChannel: DefaultChannel }
  | { PrivateChannel: PrivateChannel }
  | { DmChannel: DmChannel };

// Channel インターフェースを定義
export interface Channel {
  type: ChannelType;
}

// get_test_channels関数の戻り値の型を修正
async function get_test_channels(): Promise<Channel[] | null> {
  try {
    const response: Channel[] = await invoke("get_test_channels");
    return response;
  } catch (err) {
    console.error(err);
    return null;
  }
}

export default get_test_channels;
