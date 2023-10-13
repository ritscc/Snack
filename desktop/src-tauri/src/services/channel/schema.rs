use serde::Serialize;

#[derive(Serialize)]
pub enum ChannelType {
    DefaultChannel(DefaultChannel),
    PrivateChannel(PrivateChannel),
    DmChannel(DmChannel),
}

#[derive(Serialize)]
pub struct Channel {
    pub r#type: ChannelType,
}

#[derive(Serialize)]
pub struct DefaultChannel {
    pub channel_id: i64,
    pub position: i64,
    pub name: String,
    pub private: bool,
    pub nsfw: bool,
    pub user_id: Vec<i64>,
}

#[derive(Serialize)]
pub struct PrivateChannel {
    pub channel_id: i64,
    pub user_id: i64,
}

#[derive(Serialize)]
pub struct DmChannel {
    pub channel_id: i64,
    pub user_id: i64,
    pub other_user_id: i64,
}
