// channels.ts
"use client";
import { invoke } from "@tauri-apps/api/tauri";

export type Channel = DefaultChannel | PrivateChannel | DmChannel;

export interface DefaultChannel {
  type: "DefaultChannel"; // Discriminated Union を使用
  channel_id: number;
  position: number;
  name: string;
  private: boolean;
  nsfw: boolean;
  user_ids: number[]; // TypeScriptではVecを配列に対応させる
}

export interface PrivateChannel {
  type: "PrivateChannel"; // Discriminated Union を使用
  channel_id: number;
  user_id: number;
}

export interface DmChannel {
  type: "DmChannel"; // Discriminated Union を使用
  channel_id: number;
  user_id: number;
  other_user_id: number;
}

// get_test_channels関数の戻り値の型を修正
async function get_test_channels(): Promise<DefaultChannel[] | null> {
  try {
    const response = await invoke("get_test_channels");
    return response as DefaultChannel[];
  } catch (err) {
    console.error(err);
    return null;
  }
}

export default get_test_channels;
