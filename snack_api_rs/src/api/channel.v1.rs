#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ChannelId {
    #[prost(int64, tag = "1")]
    pub channel_id: i64,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Channel {
    #[prost(oneof = "channel::ChannelGroup", tags = "1, 2, 3")]
    pub channel_group: ::core::option::Option<channel::ChannelGroup>,
}
/// Nested message and enum types in `Channel`.
pub mod channel {
    #[allow(clippy::derive_partial_eq_without_eq)]
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum ChannelGroup {
        #[prost(message, tag = "1")]
        DefaultChannel(super::DefaultChannel),
        #[prost(message, tag = "2")]
        PrivateChannel(super::PrivateChannel),
        #[prost(message, tag = "3")]
        DmChannel(super::DmChannel),
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct DefaultChannel {
    #[prost(int64, tag = "1")]
    pub channel_id: i64,
    #[prost(enumeration = "ChannelType", tag = "2")]
    pub channel_type: i32,
    #[prost(int64, tag = "3")]
    pub position: i64,
    #[prost(string, tag = "4")]
    pub name: ::prost::alloc::string::String,
    #[prost(bool, tag = "5")]
    pub private: bool,
    #[prost(bool, tag = "6")]
    pub nsfw: bool,
    #[prost(int64, repeated, tag = "7")]
    pub user_id: ::prost::alloc::vec::Vec<i64>,
}
/// 個人チャンネル
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PrivateChannel {
    #[prost(int64, tag = "2")]
    pub channel_id: i64,
    #[prost(int64, tag = "1")]
    pub user_id: i64,
}
/// 1v1のダイレクトメッセージチャンネル
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct DmChannel {
    #[prost(int64, tag = "1")]
    pub channel_id: i64,
    #[prost(int64, tag = "2")]
    pub other_user_id: i64,
}
/// 通知を行うチャンネル
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ChannelSubscription {
    #[prost(int64, tag = "1")]
    pub channel_id: i64,
    #[prost(bool, tag = "2")]
    pub notify: bool,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct CreateChannelRequest {
    #[prost(oneof = "create_channel_request::CreateChannelGroup", tags = "1, 2")]
    pub create_channel_group: ::core::option::Option<create_channel_request::CreateChannelGroup>,
}
/// Nested message and enum types in `CreateChannelRequest`.
pub mod create_channel_request {
    #[allow(clippy::derive_partial_eq_without_eq)]
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum CreateChannelGroup {
        #[prost(message, tag = "1")]
        DefaultChannelRequest(super::CreateDefaultChannelRequest),
        #[prost(message, tag = "2")]
        DmChannelRequest(super::CreateDmRequest),
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct CreateDefaultChannelRequest {
    #[prost(enumeration = "ChannelType", tag = "1")]
    pub channel_type: i32,
    #[prost(int64, tag = "2")]
    pub position: i64,
    #[prost(string, tag = "3")]
    pub name: ::prost::alloc::string::String,
    #[prost(bool, tag = "4")]
    pub private: bool,
    #[prost(bool, tag = "5")]
    pub nsfw: bool,
    #[prost(int64, repeated, tag = "6")]
    pub user_id: ::prost::alloc::vec::Vec<i64>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct CreateDmRequest {
    #[prost(int64, tag = "1")]
    pub other_user_id: i64,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MessageChannelRequest {
    #[prost(int64, tag = "1")]
    pub channel_id: i64,
}
/// KC3やHackathonのEventをいつやるかなどチャンネルに告知する機能
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ChannelEvent {
    #[prost(int64, tag = "1")]
    pub id: i64,
    #[prost(int64, tag = "2")]
    pub channel_id: i64,
    #[prost(string, tag = "3")]
    pub event_type: ::prost::alloc::string::String,
    #[prost(string, tag = "4")]
    pub title: ::prost::alloc::string::String,
    #[prost(string, tag = "5")]
    pub detail: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "6")]
    pub date: ::core::option::Option<super::super::google::protobuf::Timestamp>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct CreateChannelEventRequest {
    #[prost(int64, tag = "1")]
    pub channel_id: i64,
    #[prost(string, tag = "2")]
    pub event_type: ::prost::alloc::string::String,
    #[prost(string, tag = "3")]
    pub title: ::prost::alloc::string::String,
    #[prost(string, tag = "4")]
    pub detail: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "5")]
    pub date: ::core::option::Option<super::super::google::protobuf::Timestamp>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ChannelEventId {
    #[prost(int64, tag = "1")]
    pub channel_event_id: i64,
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum ChannelType {
    ChannelText = 0,
    ChannelVoice = 1,
}
impl ChannelType {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            ChannelType::ChannelText => "CHANNEL_TEXT",
            ChannelType::ChannelVoice => "CHANNEL_VOICE",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "CHANNEL_TEXT" => Some(Self::ChannelText),
            "CHANNEL_VOICE" => Some(Self::ChannelVoice),
            _ => None,
        }
    }
}
/// Generated client implementations.
pub mod channel_service_client {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::http::Uri;
    use tonic::codegen::*;
    #[derive(Debug, Clone)]
    pub struct ChannelServiceClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl ChannelServiceClient<tonic::transport::Channel> {
        /// Attempt to create a new client by connecting to a given endpoint.
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> ChannelServiceClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_origin(inner: T, origin: Uri) -> Self {
            let inner = tonic::client::Grpc::with_origin(inner, origin);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> ChannelServiceClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::BoxBody>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<http::Request<tonic::body::BoxBody>>>::Error:
                Into<StdError> + Send + Sync,
        {
            ChannelServiceClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with the given encoding.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.send_compressed(encoding);
            self
        }
        /// Enable decompressing responses.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.accept_compressed(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_decoding_message_size(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_encoding_message_size(limit);
            self
        }
        pub async fn create_channel(
            &mut self,
            request: impl tonic::IntoRequest<super::CreateChannelRequest>,
        ) -> std::result::Result<
            tonic::Response<super::super::super::google::protobuf::Empty>,
            tonic::Status,
        > {
            self.inner.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e.into()),
                )
            })?;
            let codec = tonic::codec::ProstCodec::default();
            let path =
                http::uri::PathAndQuery::from_static("/channel.v1.ChannelService/CreateChannel");
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new(
                "channel.v1.ChannelService",
                "CreateChannel",
            ));
            self.inner.unary(req, path, codec).await
        }
        pub async fn get_channel(
            &mut self,
            request: impl tonic::IntoRequest<super::ChannelId>,
        ) -> std::result::Result<tonic::Response<super::Channel>, tonic::Status> {
            self.inner.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e.into()),
                )
            })?;
            let codec = tonic::codec::ProstCodec::default();
            let path =
                http::uri::PathAndQuery::from_static("/channel.v1.ChannelService/GetChannel");
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("channel.v1.ChannelService", "GetChannel"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn update_channel(
            &mut self,
            request: impl tonic::IntoRequest<super::Channel>,
        ) -> std::result::Result<
            tonic::Response<super::super::super::google::protobuf::Empty>,
            tonic::Status,
        > {
            self.inner.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e.into()),
                )
            })?;
            let codec = tonic::codec::ProstCodec::default();
            let path =
                http::uri::PathAndQuery::from_static("/channel.v1.ChannelService/UpdateChannel");
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new(
                "channel.v1.ChannelService",
                "UpdateChannel",
            ));
            self.inner.unary(req, path, codec).await
        }
        /// Only Admin and Creator can delete
        pub async fn delete_channel(
            &mut self,
            request: impl tonic::IntoRequest<super::ChannelId>,
        ) -> std::result::Result<
            tonic::Response<super::super::super::google::protobuf::Empty>,
            tonic::Status,
        > {
            self.inner.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e.into()),
                )
            })?;
            let codec = tonic::codec::ProstCodec::default();
            let path =
                http::uri::PathAndQuery::from_static("/channel.v1.ChannelService/DeleteChannel");
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new(
                "channel.v1.ChannelService",
                "DeleteChannel",
            ));
            self.inner.unary(req, path, codec).await
        }
    }
}
/// Generated client implementations.
pub mod message_channel_service_client {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::http::Uri;
    use tonic::codegen::*;
    /// メッセージチャンネルの機能（メッセージチャンネルに入らないとメッセージが見れない）
    #[derive(Debug, Clone)]
    pub struct MessageChannelServiceClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl MessageChannelServiceClient<tonic::transport::Channel> {
        /// Attempt to create a new client by connecting to a given endpoint.
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> MessageChannelServiceClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_origin(inner: T, origin: Uri) -> Self {
            let inner = tonic::client::Grpc::with_origin(inner, origin);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> MessageChannelServiceClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::BoxBody>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<http::Request<tonic::body::BoxBody>>>::Error:
                Into<StdError> + Send + Sync,
        {
            MessageChannelServiceClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with the given encoding.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.send_compressed(encoding);
            self
        }
        /// Enable decompressing responses.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.accept_compressed(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_decoding_message_size(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_encoding_message_size(limit);
            self
        }
        pub async fn leave_message_channel(
            &mut self,
            request: impl tonic::IntoRequest<super::MessageChannelRequest>,
        ) -> std::result::Result<
            tonic::Response<super::super::super::google::protobuf::Empty>,
            tonic::Status,
        > {
            self.inner.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e.into()),
                )
            })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/channel.v1.MessageChannelService/LeaveMessageChannel",
            );
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new(
                "channel.v1.MessageChannelService",
                "LeaveMessageChannel",
            ));
            self.inner.unary(req, path, codec).await
        }
        pub async fn join_message_channel(
            &mut self,
            request: impl tonic::IntoRequest<super::MessageChannelRequest>,
        ) -> std::result::Result<
            tonic::Response<super::super::super::google::protobuf::Empty>,
            tonic::Status,
        > {
            self.inner.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e.into()),
                )
            })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/channel.v1.MessageChannelService/JoinMessageChannel",
            );
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new(
                "channel.v1.MessageChannelService",
                "JoinMessageChannel",
            ));
            self.inner.unary(req, path, codec).await
        }
    }
}
/// Generated client implementations.
pub mod channel_event_service_client {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::http::Uri;
    use tonic::codegen::*;
    /// KC3やHackathonのEventをいつやるかなどチャンネルに告知する機能
    #[derive(Debug, Clone)]
    pub struct ChannelEventServiceClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl ChannelEventServiceClient<tonic::transport::Channel> {
        /// Attempt to create a new client by connecting to a given endpoint.
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> ChannelEventServiceClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_origin(inner: T, origin: Uri) -> Self {
            let inner = tonic::client::Grpc::with_origin(inner, origin);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> ChannelEventServiceClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::BoxBody>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<http::Request<tonic::body::BoxBody>>>::Error:
                Into<StdError> + Send + Sync,
        {
            ChannelEventServiceClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with the given encoding.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.send_compressed(encoding);
            self
        }
        /// Enable decompressing responses.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.accept_compressed(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_decoding_message_size(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_encoding_message_size(limit);
            self
        }
        pub async fn create_channel_event(
            &mut self,
            request: impl tonic::IntoRequest<super::CreateChannelEventRequest>,
        ) -> std::result::Result<
            tonic::Response<super::super::super::google::protobuf::Empty>,
            tonic::Status,
        > {
            self.inner.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e.into()),
                )
            })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/channel.v1.ChannelEventService/CreateChannelEvent",
            );
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new(
                "channel.v1.ChannelEventService",
                "CreateChannelEvent",
            ));
            self.inner.unary(req, path, codec).await
        }
        pub async fn get_channel_event(
            &mut self,
            request: impl tonic::IntoRequest<super::ChannelEventId>,
        ) -> std::result::Result<tonic::Response<super::ChannelEvent>, tonic::Status> {
            self.inner.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e.into()),
                )
            })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/channel.v1.ChannelEventService/GetChannelEvent",
            );
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new(
                "channel.v1.ChannelEventService",
                "GetChannelEvent",
            ));
            self.inner.unary(req, path, codec).await
        }
        pub async fn update_channel_event(
            &mut self,
            request: impl tonic::IntoRequest<super::ChannelEvent>,
        ) -> std::result::Result<
            tonic::Response<super::super::super::google::protobuf::Empty>,
            tonic::Status,
        > {
            self.inner.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e.into()),
                )
            })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/channel.v1.ChannelEventService/UpdateChannelEvent",
            );
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new(
                "channel.v1.ChannelEventService",
                "UpdateChannelEvent",
            ));
            self.inner.unary(req, path, codec).await
        }
        /// Only Admin and Creator can delete
        pub async fn delete_channel_event(
            &mut self,
            request: impl tonic::IntoRequest<super::ChannelEventId>,
        ) -> std::result::Result<
            tonic::Response<super::super::super::google::protobuf::Empty>,
            tonic::Status,
        > {
            self.inner.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e.into()),
                )
            })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/channel.v1.ChannelEventService/DeleteChannelEvent",
            );
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new(
                "channel.v1.ChannelEventService",
                "DeleteChannelEvent",
            ));
            self.inner.unary(req, path, codec).await
        }
    }
}
/// Generated server implementations.
pub mod channel_service_server {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    /// Generated trait containing gRPC methods that should be implemented for use with ChannelServiceServer.
    #[async_trait]
    pub trait ChannelService: Send + Sync + 'static {
        async fn create_channel(
            &self,
            request: tonic::Request<super::CreateChannelRequest>,
        ) -> std::result::Result<
            tonic::Response<super::super::super::google::protobuf::Empty>,
            tonic::Status,
        >;
        async fn get_channel(
            &self,
            request: tonic::Request<super::ChannelId>,
        ) -> std::result::Result<tonic::Response<super::Channel>, tonic::Status>;
        async fn update_channel(
            &self,
            request: tonic::Request<super::Channel>,
        ) -> std::result::Result<
            tonic::Response<super::super::super::google::protobuf::Empty>,
            tonic::Status,
        >;
        /// Only Admin and Creator can delete
        async fn delete_channel(
            &self,
            request: tonic::Request<super::ChannelId>,
        ) -> std::result::Result<
            tonic::Response<super::super::super::google::protobuf::Empty>,
            tonic::Status,
        >;
    }
    #[derive(Debug)]
    pub struct ChannelServiceServer<T: ChannelService> {
        inner: _Inner<T>,
        accept_compression_encodings: EnabledCompressionEncodings,
        send_compression_encodings: EnabledCompressionEncodings,
        max_decoding_message_size: Option<usize>,
        max_encoding_message_size: Option<usize>,
    }
    struct _Inner<T>(Arc<T>);
    impl<T: ChannelService> ChannelServiceServer<T> {
        pub fn new(inner: T) -> Self {
            Self::from_arc(Arc::new(inner))
        }
        pub fn from_arc(inner: Arc<T>) -> Self {
            let inner = _Inner(inner);
            Self {
                inner,
                accept_compression_encodings: Default::default(),
                send_compression_encodings: Default::default(),
                max_decoding_message_size: None,
                max_encoding_message_size: None,
            }
        }
        pub fn with_interceptor<F>(inner: T, interceptor: F) -> InterceptedService<Self, F>
        where
            F: tonic::service::Interceptor,
        {
            InterceptedService::new(Self::new(inner), interceptor)
        }
        /// Enable decompressing requests with the given encoding.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.accept_compression_encodings.enable(encoding);
            self
        }
        /// Compress responses with the given encoding, if the client supports it.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.send_compression_encodings.enable(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.max_decoding_message_size = Some(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.max_encoding_message_size = Some(limit);
            self
        }
    }
    impl<T, B> tonic::codegen::Service<http::Request<B>> for ChannelServiceServer<T>
    where
        T: ChannelService,
        B: Body + Send + 'static,
        B::Error: Into<StdError> + Send + 'static,
    {
        type Response = http::Response<tonic::body::BoxBody>;
        type Error = std::convert::Infallible;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(
            &mut self,
            _cx: &mut Context<'_>,
        ) -> Poll<std::result::Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            let inner = self.inner.clone();
            match req.uri().path() {
                "/channel.v1.ChannelService/CreateChannel" => {
                    #[allow(non_camel_case_types)]
                    struct CreateChannelSvc<T: ChannelService>(pub Arc<T>);
                    impl<T: ChannelService> tonic::server::UnaryService<super::CreateChannelRequest>
                        for CreateChannelSvc<T>
                    {
                        type Response = super::super::super::google::protobuf::Empty;
                        type Future = BoxFuture<tonic::Response<Self::Response>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::CreateChannelRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as ChannelService>::create_channel(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = CreateChannelSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/channel.v1.ChannelService/GetChannel" => {
                    #[allow(non_camel_case_types)]
                    struct GetChannelSvc<T: ChannelService>(pub Arc<T>);
                    impl<T: ChannelService> tonic::server::UnaryService<super::ChannelId> for GetChannelSvc<T> {
                        type Response = super::Channel;
                        type Future = BoxFuture<tonic::Response<Self::Response>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ChannelId>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as ChannelService>::get_channel(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = GetChannelSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/channel.v1.ChannelService/UpdateChannel" => {
                    #[allow(non_camel_case_types)]
                    struct UpdateChannelSvc<T: ChannelService>(pub Arc<T>);
                    impl<T: ChannelService> tonic::server::UnaryService<super::Channel> for UpdateChannelSvc<T> {
                        type Response = super::super::super::google::protobuf::Empty;
                        type Future = BoxFuture<tonic::Response<Self::Response>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::Channel>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as ChannelService>::update_channel(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = UpdateChannelSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/channel.v1.ChannelService/DeleteChannel" => {
                    #[allow(non_camel_case_types)]
                    struct DeleteChannelSvc<T: ChannelService>(pub Arc<T>);
                    impl<T: ChannelService> tonic::server::UnaryService<super::ChannelId> for DeleteChannelSvc<T> {
                        type Response = super::super::super::google::protobuf::Empty;
                        type Future = BoxFuture<tonic::Response<Self::Response>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ChannelId>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as ChannelService>::delete_channel(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = DeleteChannelSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => Box::pin(async move {
                    Ok(http::Response::builder()
                        .status(200)
                        .header("grpc-status", "12")
                        .header("content-type", "application/grpc")
                        .body(empty_body())
                        .unwrap())
                }),
            }
        }
    }
    impl<T: ChannelService> Clone for ChannelServiceServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self {
                inner,
                accept_compression_encodings: self.accept_compression_encodings,
                send_compression_encodings: self.send_compression_encodings,
                max_decoding_message_size: self.max_decoding_message_size,
                max_encoding_message_size: self.max_encoding_message_size,
            }
        }
    }
    impl<T: ChannelService> Clone for _Inner<T> {
        fn clone(&self) -> Self {
            Self(Arc::clone(&self.0))
        }
    }
    impl<T: std::fmt::Debug> std::fmt::Debug for _Inner<T> {
        fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
            write!(f, "{:?}", self.0)
        }
    }
    impl<T: ChannelService> tonic::server::NamedService for ChannelServiceServer<T> {
        const NAME: &'static str = "channel.v1.ChannelService";
    }
}
/// Generated server implementations.
pub mod message_channel_service_server {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    /// Generated trait containing gRPC methods that should be implemented for use with MessageChannelServiceServer.
    #[async_trait]
    pub trait MessageChannelService: Send + Sync + 'static {
        async fn leave_message_channel(
            &self,
            request: tonic::Request<super::MessageChannelRequest>,
        ) -> std::result::Result<
            tonic::Response<super::super::super::google::protobuf::Empty>,
            tonic::Status,
        >;
        async fn join_message_channel(
            &self,
            request: tonic::Request<super::MessageChannelRequest>,
        ) -> std::result::Result<
            tonic::Response<super::super::super::google::protobuf::Empty>,
            tonic::Status,
        >;
    }
    /// メッセージチャンネルの機能（メッセージチャンネルに入らないとメッセージが見れない）
    #[derive(Debug)]
    pub struct MessageChannelServiceServer<T: MessageChannelService> {
        inner: _Inner<T>,
        accept_compression_encodings: EnabledCompressionEncodings,
        send_compression_encodings: EnabledCompressionEncodings,
        max_decoding_message_size: Option<usize>,
        max_encoding_message_size: Option<usize>,
    }
    struct _Inner<T>(Arc<T>);
    impl<T: MessageChannelService> MessageChannelServiceServer<T> {
        pub fn new(inner: T) -> Self {
            Self::from_arc(Arc::new(inner))
        }
        pub fn from_arc(inner: Arc<T>) -> Self {
            let inner = _Inner(inner);
            Self {
                inner,
                accept_compression_encodings: Default::default(),
                send_compression_encodings: Default::default(),
                max_decoding_message_size: None,
                max_encoding_message_size: None,
            }
        }
        pub fn with_interceptor<F>(inner: T, interceptor: F) -> InterceptedService<Self, F>
        where
            F: tonic::service::Interceptor,
        {
            InterceptedService::new(Self::new(inner), interceptor)
        }
        /// Enable decompressing requests with the given encoding.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.accept_compression_encodings.enable(encoding);
            self
        }
        /// Compress responses with the given encoding, if the client supports it.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.send_compression_encodings.enable(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.max_decoding_message_size = Some(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.max_encoding_message_size = Some(limit);
            self
        }
    }
    impl<T, B> tonic::codegen::Service<http::Request<B>> for MessageChannelServiceServer<T>
    where
        T: MessageChannelService,
        B: Body + Send + 'static,
        B::Error: Into<StdError> + Send + 'static,
    {
        type Response = http::Response<tonic::body::BoxBody>;
        type Error = std::convert::Infallible;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(
            &mut self,
            _cx: &mut Context<'_>,
        ) -> Poll<std::result::Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            let inner = self.inner.clone();
            match req.uri().path() {
                "/channel.v1.MessageChannelService/LeaveMessageChannel" => {
                    #[allow(non_camel_case_types)]
                    struct LeaveMessageChannelSvc<T: MessageChannelService>(pub Arc<T>);
                    impl<T: MessageChannelService>
                        tonic::server::UnaryService<super::MessageChannelRequest>
                        for LeaveMessageChannelSvc<T>
                    {
                        type Response = super::super::super::google::protobuf::Empty;
                        type Future = BoxFuture<tonic::Response<Self::Response>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MessageChannelRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as MessageChannelService>::leave_message_channel(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = LeaveMessageChannelSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/channel.v1.MessageChannelService/JoinMessageChannel" => {
                    #[allow(non_camel_case_types)]
                    struct JoinMessageChannelSvc<T: MessageChannelService>(pub Arc<T>);
                    impl<T: MessageChannelService>
                        tonic::server::UnaryService<super::MessageChannelRequest>
                        for JoinMessageChannelSvc<T>
                    {
                        type Response = super::super::super::google::protobuf::Empty;
                        type Future = BoxFuture<tonic::Response<Self::Response>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MessageChannelRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as MessageChannelService>::join_message_channel(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = JoinMessageChannelSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => Box::pin(async move {
                    Ok(http::Response::builder()
                        .status(200)
                        .header("grpc-status", "12")
                        .header("content-type", "application/grpc")
                        .body(empty_body())
                        .unwrap())
                }),
            }
        }
    }
    impl<T: MessageChannelService> Clone for MessageChannelServiceServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self {
                inner,
                accept_compression_encodings: self.accept_compression_encodings,
                send_compression_encodings: self.send_compression_encodings,
                max_decoding_message_size: self.max_decoding_message_size,
                max_encoding_message_size: self.max_encoding_message_size,
            }
        }
    }
    impl<T: MessageChannelService> Clone for _Inner<T> {
        fn clone(&self) -> Self {
            Self(Arc::clone(&self.0))
        }
    }
    impl<T: std::fmt::Debug> std::fmt::Debug for _Inner<T> {
        fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
            write!(f, "{:?}", self.0)
        }
    }
    impl<T: MessageChannelService> tonic::server::NamedService for MessageChannelServiceServer<T> {
        const NAME: &'static str = "channel.v1.MessageChannelService";
    }
}
/// Generated server implementations.
pub mod channel_event_service_server {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    /// Generated trait containing gRPC methods that should be implemented for use with ChannelEventServiceServer.
    #[async_trait]
    pub trait ChannelEventService: Send + Sync + 'static {
        async fn create_channel_event(
            &self,
            request: tonic::Request<super::CreateChannelEventRequest>,
        ) -> std::result::Result<
            tonic::Response<super::super::super::google::protobuf::Empty>,
            tonic::Status,
        >;
        async fn get_channel_event(
            &self,
            request: tonic::Request<super::ChannelEventId>,
        ) -> std::result::Result<tonic::Response<super::ChannelEvent>, tonic::Status>;
        async fn update_channel_event(
            &self,
            request: tonic::Request<super::ChannelEvent>,
        ) -> std::result::Result<
            tonic::Response<super::super::super::google::protobuf::Empty>,
            tonic::Status,
        >;
        /// Only Admin and Creator can delete
        async fn delete_channel_event(
            &self,
            request: tonic::Request<super::ChannelEventId>,
        ) -> std::result::Result<
            tonic::Response<super::super::super::google::protobuf::Empty>,
            tonic::Status,
        >;
    }
    /// KC3やHackathonのEventをいつやるかなどチャンネルに告知する機能
    #[derive(Debug)]
    pub struct ChannelEventServiceServer<T: ChannelEventService> {
        inner: _Inner<T>,
        accept_compression_encodings: EnabledCompressionEncodings,
        send_compression_encodings: EnabledCompressionEncodings,
        max_decoding_message_size: Option<usize>,
        max_encoding_message_size: Option<usize>,
    }
    struct _Inner<T>(Arc<T>);
    impl<T: ChannelEventService> ChannelEventServiceServer<T> {
        pub fn new(inner: T) -> Self {
            Self::from_arc(Arc::new(inner))
        }
        pub fn from_arc(inner: Arc<T>) -> Self {
            let inner = _Inner(inner);
            Self {
                inner,
                accept_compression_encodings: Default::default(),
                send_compression_encodings: Default::default(),
                max_decoding_message_size: None,
                max_encoding_message_size: None,
            }
        }
        pub fn with_interceptor<F>(inner: T, interceptor: F) -> InterceptedService<Self, F>
        where
            F: tonic::service::Interceptor,
        {
            InterceptedService::new(Self::new(inner), interceptor)
        }
        /// Enable decompressing requests with the given encoding.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.accept_compression_encodings.enable(encoding);
            self
        }
        /// Compress responses with the given encoding, if the client supports it.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.send_compression_encodings.enable(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.max_decoding_message_size = Some(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.max_encoding_message_size = Some(limit);
            self
        }
    }
    impl<T, B> tonic::codegen::Service<http::Request<B>> for ChannelEventServiceServer<T>
    where
        T: ChannelEventService,
        B: Body + Send + 'static,
        B::Error: Into<StdError> + Send + 'static,
    {
        type Response = http::Response<tonic::body::BoxBody>;
        type Error = std::convert::Infallible;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(
            &mut self,
            _cx: &mut Context<'_>,
        ) -> Poll<std::result::Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            let inner = self.inner.clone();
            match req.uri().path() {
                "/channel.v1.ChannelEventService/CreateChannelEvent" => {
                    #[allow(non_camel_case_types)]
                    struct CreateChannelEventSvc<T: ChannelEventService>(pub Arc<T>);
                    impl<T: ChannelEventService>
                        tonic::server::UnaryService<super::CreateChannelEventRequest>
                        for CreateChannelEventSvc<T>
                    {
                        type Response = super::super::super::google::protobuf::Empty;
                        type Future = BoxFuture<tonic::Response<Self::Response>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::CreateChannelEventRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as ChannelEventService>::create_channel_event(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = CreateChannelEventSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/channel.v1.ChannelEventService/GetChannelEvent" => {
                    #[allow(non_camel_case_types)]
                    struct GetChannelEventSvc<T: ChannelEventService>(pub Arc<T>);
                    impl<T: ChannelEventService> tonic::server::UnaryService<super::ChannelEventId>
                        for GetChannelEventSvc<T>
                    {
                        type Response = super::ChannelEvent;
                        type Future = BoxFuture<tonic::Response<Self::Response>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ChannelEventId>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as ChannelEventService>::get_channel_event(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = GetChannelEventSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/channel.v1.ChannelEventService/UpdateChannelEvent" => {
                    #[allow(non_camel_case_types)]
                    struct UpdateChannelEventSvc<T: ChannelEventService>(pub Arc<T>);
                    impl<T: ChannelEventService> tonic::server::UnaryService<super::ChannelEvent>
                        for UpdateChannelEventSvc<T>
                    {
                        type Response = super::super::super::google::protobuf::Empty;
                        type Future = BoxFuture<tonic::Response<Self::Response>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ChannelEvent>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as ChannelEventService>::update_channel_event(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = UpdateChannelEventSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/channel.v1.ChannelEventService/DeleteChannelEvent" => {
                    #[allow(non_camel_case_types)]
                    struct DeleteChannelEventSvc<T: ChannelEventService>(pub Arc<T>);
                    impl<T: ChannelEventService> tonic::server::UnaryService<super::ChannelEventId>
                        for DeleteChannelEventSvc<T>
                    {
                        type Response = super::super::super::google::protobuf::Empty;
                        type Future = BoxFuture<tonic::Response<Self::Response>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ChannelEventId>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as ChannelEventService>::delete_channel_event(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = DeleteChannelEventSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => Box::pin(async move {
                    Ok(http::Response::builder()
                        .status(200)
                        .header("grpc-status", "12")
                        .header("content-type", "application/grpc")
                        .body(empty_body())
                        .unwrap())
                }),
            }
        }
    }
    impl<T: ChannelEventService> Clone for ChannelEventServiceServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self {
                inner,
                accept_compression_encodings: self.accept_compression_encodings,
                send_compression_encodings: self.send_compression_encodings,
                max_decoding_message_size: self.max_decoding_message_size,
                max_encoding_message_size: self.max_encoding_message_size,
            }
        }
    }
    impl<T: ChannelEventService> Clone for _Inner<T> {
        fn clone(&self) -> Self {
            Self(Arc::clone(&self.0))
        }
    }
    impl<T: std::fmt::Debug> std::fmt::Debug for _Inner<T> {
        fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
            write!(f, "{:?}", self.0)
        }
    }
    impl<T: ChannelEventService> tonic::server::NamedService for ChannelEventServiceServer<T> {
        const NAME: &'static str = "channel.v1.ChannelEventService";
    }
}
