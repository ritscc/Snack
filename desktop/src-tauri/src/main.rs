#![cfg_attr(
    all(not(debug_assertions), target_os = "windows"),
    windows_subsystem = "windows"
)]

use tauri::{
    api::notification::Notification, AppHandle, CustomMenuItem, Manager, SystemTray,
    SystemTrayEvent, SystemTrayMenu, SystemTrayMenuItem, Wry,
};

use snack_desktop::channel::query::{__cmd__get_test_channels, get_test_channels};
use snack_desktop::message::query::{__cmd__get_test_messages, get_test_messages};
use snack_desktop::user::query::{
    __cmd__get_test_user_data, __cmd__get_test_users_data, get_test_user_data, get_test_users_data,
};

fn main() {
    let context = tauri::generate_context!();
    let identifier = context.config().tauri.bundle.identifier.clone();
    tauri::Builder::default()
        .system_tray(create_systemtray())
        .on_system_tray_event(on_main_menu_event)
        .invoke_handler(tauri::generate_handler![
            get_test_users_data,
            get_test_user_data,
            get_test_messages,
            get_test_channels
        ])
        .setup(move |_app| {
            notification(identifier, "Snack".into(), "Spawned!".into()).unwrap();
            Ok(())
        })
        .run(context)
        .expect("error while running tauri application");
}

fn create_systemtray() -> SystemTray {
    SystemTray::new().with_menu(
        SystemTrayMenu::new()
            .add_item(CustomMenuItem::new("quit", "Quit"))
            .add_item(CustomMenuItem::new("show", "Show"))
            .add_native_item(SystemTrayMenuItem::Separator)
            .add_item(CustomMenuItem::new("hide", "Hide")),
    )
}

fn on_main_menu_event(app: &AppHandle<Wry>, event: SystemTrayEvent) {
    match event {
        SystemTrayEvent::MenuItemClick { ref id, .. } if id == "quit" => std::process::exit(0),
        SystemTrayEvent::MenuItemClick { ref id, .. } if id == "hide" => {
            let window = app.get_window("main").unwrap();
            window.hide().unwrap();
        }
        SystemTrayEvent::MenuItemClick { ref id, .. } if id == "show" => {
            let window = app.get_window("main").unwrap();
            window.show().unwrap();
        }
        _ => {}
    }
}

fn notification(identifier: String, title: String, body: String) -> Result<(), tauri::api::Error> {
    Notification::new(identifier).title(title).body(body).show()
}
