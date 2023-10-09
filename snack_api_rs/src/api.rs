#[path = "api"]
pub mod google {
    #[path = "google.protobuf.rs"]
    pub mod protobuf;
}

#[path = "api"]
pub mod auth {
    #[path = "auth.v1.rs"]
    pub mod v1;
}

#[path = "api"]
pub mod event {
    #[path = "event.v1.rs"]
    pub mod v1;
}

#[path = "api"]
pub mod channel {
    #[path = "channel.v1.rs"]
    pub mod v1;
}

#[path = "api"]
pub mod message {
    #[path = "message.v1.rs"]
    pub mod v1;
}

#[path = "api"]
pub mod stamp {
    #[path = "stamp.v1.rs"]
    pub mod v1;
}

#[path = "api"]
pub mod user_group {
    #[path = "user_group.v1.rs"]
    pub mod v1;
}

#[path = "api"]
pub mod user {
    #[path = "user.v1.rs"]
    pub mod v1;
}
