use serde::Serialize;

#[derive(Serialize)]
pub struct UserGroup {
    pub id: i64,
    pub name: String,
    pub description: String,
    pub r#type: String,
    pub icon: String,
    pub role: String,
    pub member: Vec<Member>,
}

#[derive(Serialize)]
pub struct Member {
    pub group_id: i64,
    pub user_id: i64,
    pub grant: String,
}
