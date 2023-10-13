use super::schema::{Channel, ChannelType, DefaultChannel, DmChannel, PrivateChannel};

#[tauri::command]
pub fn get_test_channels() -> Vec<Channel> {
    let channels = vec![
        Channel {
            r#type: ChannelType::DefaultChannel(DefaultChannel {
                channel_id: 1,
                position: 1,
                name: "default_channel".into(),
                private: false,
                nsfw: false,
                user_id: vec![1, 2, 3],
            }),
        },
        Channel {
            r#type: ChannelType::DefaultChannel(DefaultChannel {
                channel_id: 2,
                position: 2,
                name: "default_channel_2".into(),
                private: false,
                nsfw: false,
                user_id: vec![1, 2],
            }),
        },
        Channel {
            r#type: ChannelType::PrivateChannel(PrivateChannel {
                channel_id: 3,
                user_id: 1,
            }),
        },
        Channel {
            r#type: ChannelType::DmChannel(DmChannel {
                channel_id: 4,
                user_id: 1,
                other_user_id: 2,
            }),
        },
    ];
    channels
}
