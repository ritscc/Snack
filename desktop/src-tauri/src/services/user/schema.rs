use serde::Serialize;

#[derive(Serialize)]
pub enum UserStatus {
    ONLINE,
    OFFLINE,
}

#[derive(Serialize)]
pub struct UserOtherServiceAccounts {
    pub x: Option<String>,
    pub github: Option<String>,
    pub discord: Option<String>,
}

#[derive(Serialize)]
pub struct User {
    pub id: i64,
    pub username: String,
    pub nick: String,
    pub realname: String,
    pub avator: String,
    pub role: Vec<String>,
    pub locale: String,
    pub rits_id: String,
    pub own_email: Option<String>,
    pub comment: String,
    pub status: UserStatus,
    pub service: Option<UserOtherServiceAccounts>,
    pub is_deleted: bool,
}
