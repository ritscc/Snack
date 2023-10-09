'use client'

import { invoke } from "@tauri-apps/api/tauri"

async function get_test_channels() {
    // Todo: Change Function: get_test_channels
    const users = await invoke('get_test_channels').catch(err => {
        console.error(err)
        return null
    })
}