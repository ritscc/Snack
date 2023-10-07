'use client'

import { invoke } from "@tauri-apps/api/tauri"

async function get_users() {
    // Todo: Change Function: get_users
    const users = await invoke('get_test_users').catch(err => {
        console.error(err)
        return null
    })
}

async function get_user(id: number) {
    // Todo: Change Function: get_user
    const users = await invoke('get_test_user', { id }).catch(err => {
        console.error(err)
        return null
    })
}