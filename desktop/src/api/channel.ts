"use client";

import { invoke } from "@tauri-apps/api/tauri";
import { get } from "http";

async function get_test_channels() {
  try {
    const users = await invoke("get_test_channels");
    return users; // データを返すように修正
  } catch (err) {
    console.error(err);
    return null;
  }
}

export default get_test_channels;
