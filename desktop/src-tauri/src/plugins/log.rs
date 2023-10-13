use chrono::Utc;
use log::LevelFilter;
use tauri::{plugin::TauriPlugin, Wry};
use tauri_plugin_log::LogTarget;

pub fn tauri_log_plugin() -> TauriPlugin<Wry> {
    tauri_plugin_log::Builder::default()
        .format(move |out, message, record| {
            let format: &str = "[%Y-%m-%d][%H:%M:%S]";
            out.finish(format_args!(
                "{} [{}]  {}",
                Utc::now().format(format),
                record.level(),
                message
            ))
        })
        .level(LevelFilter::Info)
        .targets([LogTarget::LogDir, LogTarget::Stdout, LogTarget::Webview])
        .build()
}
