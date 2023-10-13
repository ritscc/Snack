#![cfg_attr(
    all(not(debug_assertions), target_os = "windows"),
    windows_subsystem = "windows"
)]

use snack_desktop::services::{channel, message, user};
use snack_desktop::{plugins, system_tray};
use tauri::api::notification::Notification;

fn main() {
    let context = tauri::generate_context!();
    let identifier = context.config().tauri.bundle.identifier.clone();
    let app = tauri::Builder::default()
        .plugin(plugins::store::tauri_store_plugin())
        .plugin(plugins::log::tauri_log_plugin())
        .system_tray(system_tray::init_system_tray())
        .on_system_tray_event(system_tray::handle_system_tray_events())
        .invoke_handler(command_handler())
        .setup(move |app| {
            notification(identifier, "Snack".into(), "Spawned!".into()).unwrap();
            app_setup(app.handle());
            Ok(())
        });

    app.run(context)
        .expect("error while running the application");
}

fn command_handler() -> impl Fn(tauri::Invoke) {
    tauri::generate_handler![
        user::command::get_test_users_data,
        user::command::get_test_user_data,
        message::command::get_test_messages,
        channel::command::get_test_channels
    ]
}

async fn app_setup(app_handle: tauri::AppHandle) {
    let app_handle = app_handle.clone();
    let cache_dir = app_handle.path_resolver().app_cache_dir().unwrap();
}

fn notification(identifier: String, title: String, body: String) -> Result<(), tauri::api::Error> {
    Notification::new(identifier).title(title).body(body).show()
}
