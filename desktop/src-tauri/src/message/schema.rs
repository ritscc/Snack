use chrono::{DateTime, Utc};
use serde::Serialize;

#[derive(Serialize)]
pub struct Message {
    pub message_id: i64,
    pub channel_id: i64,
    pub user_id: i64,
    pub content: String,
    pub created_time: DateTime<Utc>,
    pub updated_time: DateTime<Utc>,
    pub mention: Option<Mention>,
    pub pinned: bool,
    pub stamp_id: Vec<i64>,
}

#[derive(Serialize)]
pub struct Mention {
    pub everyone: bool,
    pub user_groups: Option<Vec<i64>>,
    pub member_roles: Option<Vec<String>>,
    pub channels: Option<Vec<i64>>,
}
