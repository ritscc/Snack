use chrono::Utc;

use crate::message::schema::Mention;

use super::schema::Message;

#[tauri::command]
pub fn get_test_messages() -> Vec<Message> {
    let time = Utc::now();
    let messages = vec![
        Message {
            message_id: 1,
            channel_id: 1,
            user_id: 1,
            content: "Hello Snack!".into(),
            created_time: time.clone(),
            updated_time: time.clone(),
            mention: Some(Mention {
                everyone: true,
                user_groups: None,
                member_roles: None,
                channels: None,
            }),
            pinned: true,
            stamp_id: vec![],
        },
        Message {
            message_id: 2,
            channel_id: 1,
            user_id: 1,
            content: "This is Second Chat!".into(),
            created_time: time.clone(),
            updated_time: time.clone(),
            mention: None,
            pinned: false,
            stamp_id: vec![],
        },
        Message {
            message_id: 3,
            channel_id: 2,
            user_id: 1,
            content: "This is Second Channel First Chat!".into(),
            created_time: time.clone(),
            updated_time: time.clone(),
            mention: None,
            pinned: false,
            stamp_id: vec![],
        },
    ];
    messages
}
