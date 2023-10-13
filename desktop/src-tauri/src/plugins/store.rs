use tauri::{plugin::TauriPlugin, Wry};

// for all data: Users, Message, and so on.
pub fn tauri_store_plugin() -> TauriPlugin<Wry> {
    tauri_plugin_store::Builder::default().build()
}
