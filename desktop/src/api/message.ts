'use client'

import { invoke } from "@tauri-apps/api/tauri"

async function get_test_messages() {
    // Todo: Change Function: get_messages
    const users = await invoke('get_test_messages').catch(err => {
        console.error(err)
        return null
    })
}
