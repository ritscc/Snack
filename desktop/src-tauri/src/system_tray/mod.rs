use tauri::{
    AppHandle, CustomMenuItem, Manager, Runtime, SystemTray, SystemTrayEvent, SystemTrayMenu,
    SystemTrayMenuItem,
};

pub fn init_system_tray() -> SystemTray {
    SystemTray::new().with_menu(
        SystemTrayMenu::new()
            .add_item(CustomMenuItem::new("quit", "Quit"))
            .add_native_item(SystemTrayMenuItem::Separator)
            .add_item(CustomMenuItem::new("hide", "Hide")),
    )
}

pub fn handle_system_tray_events<R: Runtime>() -> impl Fn(&AppHandle<R>, SystemTrayEvent) {
    |app, event| {
        match event {
            SystemTrayEvent::MenuItemClick { ref id, .. } if id == "quit" => std::process::exit(0),
            SystemTrayEvent::MenuItemClick { ref id, .. } if id == "hide" => {
                let window = app.get_window("main").unwrap();
                window.hide().unwrap();
            }
            SystemTrayEvent::LeftClick { .. } => {
                let window = app.get_window("main").unwrap();
                window.show().unwrap();
                window.set_focus().unwrap();
            }
            _ => {}
        };
    }
}
