use snack_api_rs::api::user::v1::{user_service_client::UserServiceClient, CreateUserRequest};

use super::schema::{User, UserOtherServiceAccounts, UserStatus};

#[tauri::command]
pub fn get_test_users_data() -> Vec<User> {
    let users = vec![
        User {
            id: 1,
            username: "lufe".into(),
            nick: "lufe".into(),
            realname: "????".into(),
            avator: "null".into(),
            role: vec!["Admin".into(), "Cyse".into()],
            locale: "jp".into(),
            rits_id: "ed0000".into(),
            own_email: Some("contact@lufe.jp".into()),
            comment: "I'm professional trash code generator".into(),
            status: UserStatus::ONLINE,
            service: Some(UserOtherServiceAccounts {
                x: Some("lufe_t".into()),
                github: Some("1uf3".into()),
                discord: Some("1ufe".into()),
            }),
            is_deleted: false,
        },
        User {
            id: 2,
            username: "offline".into(),
            nick: "OFFLINE".into(),
            realname: "offlineくん".into(),
            avator: "null".into(),
            role: vec![],
            locale: "us".into(),
            rits_id: "eb0000".into(),
            own_email: None,
            comment: ":(".into(),
            status: UserStatus::OFFLINE,
            service: Some(UserOtherServiceAccounts {
                x: None,
                github: Some("offline".into()),
                discord: None,
            }),
            is_deleted: false,
        },
        User {
            id: 3,
            username: "none".into(),
            nick: "NONE".into(),
            realname: "noneくん".into(),
            avator: "null".into(),
            role: vec![],
            locale: "us".into(),
            rits_id: "eo0000".into(),
            own_email: None,
            comment: ":)".into(),
            status: UserStatus::OFFLINE,
            service: None,
            is_deleted: false,
        },
        User {
            id: 4,
            username: "deleted".into(),
            nick: "DELETED".into(),
            realname: "deletedくん".into(),
            avator: "null".into(),
            role: vec![],
            locale: "us".into(),
            rits_id: "es0000".into(),
            own_email: None,
            comment: ":)".into(),
            status: UserStatus::OFFLINE,
            service: Some(UserOtherServiceAccounts {
                x: None,
                github: Some("offline".into()),
                discord: None,
            }),
            is_deleted: true,
        },
    ];
    users
}

#[tauri::command]
pub fn get_test_user_data(_id: i64) -> User {
    User {
        id: 1,
        username: "lufe".into(),
        nick: "lufe".into(),
        realname: "????".into(),
        avator: "null".into(),
        role: vec!["Admin".into(), "Cyse".into()],
        locale: "jp".into(),
        rits_id: "ed0000".into(),
        own_email: Some("contact@lufe.jp".into()),
        comment: "I'm professional trash code generator".into(),
        status: UserStatus::ONLINE,
        service: Some(UserOtherServiceAccounts {
            x: Some("lufe_t".into()),
            github: Some("1uf3".into()),
            discord: Some("1ufe".into()),
        }),
        is_deleted: false,
    }
}

#[allow(dead_code)]
async fn test_grpc_connect() {
    let mut client = UserServiceClient::connect("http://[::1]:50051")
        .await
        .unwrap();
    let request = tonic::Request::new(CreateUserRequest {
        username: "1uf3".into(),
        email: "contact@lufe.jp".into(),
        password: "password".into(),
    });

    let response = client.create_user(request).await.unwrap();

    println!("RESPONSE={:?}", response);
}
