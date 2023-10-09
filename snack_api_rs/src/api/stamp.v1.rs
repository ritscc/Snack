#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Stamp {
    #[prost(string, tag = "1")]
    pub id: ::prost::alloc::string::String,
    #[prost(string, tag = "2")]
    pub name: ::prost::alloc::string::String,
    #[prost(string, tag = "3")]
    pub creater_user_id: ::prost::alloc::string::String,
    #[prost(string, tag = "4")]
    pub file_id: ::prost::alloc::string::String,
    #[prost(bool, tag = "5")]
    pub is_unicode: bool,
    #[prost(message, optional, tag = "6")]
    pub created_at: ::core::option::Option<super::super::google::protobuf::Timestamp>,
    #[prost(message, optional, tag = "7")]
    pub updated_at: ::core::option::Option<super::super::google::protobuf::Timestamp>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct StampPalette {
    #[prost(string, tag = "1")]
    pub id: ::prost::alloc::string::String,
    #[prost(string, tag = "2")]
    pub name: ::prost::alloc::string::String,
    #[prost(string, tag = "3")]
    pub description: ::prost::alloc::string::String,
    #[prost(string, tag = "4")]
    pub creater_user_id: ::prost::alloc::string::String,
    #[prost(string, repeated, tag = "5")]
    pub stamp_id: ::prost::alloc::vec::Vec<::prost::alloc::string::String>,
    #[prost(message, optional, tag = "6")]
    pub created_at: ::core::option::Option<super::super::google::protobuf::Timestamp>,
    #[prost(message, optional, tag = "7")]
    pub updated_at: ::core::option::Option<super::super::google::protobuf::Timestamp>,
}
