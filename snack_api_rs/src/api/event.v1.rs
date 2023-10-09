#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct EventResponse {
    #[prost(enumeration = "Event", tag = "1")]
    pub event_type: i32,
    #[prost(int64, optional, tag = "2")]
    pub user_id: ::core::option::Option<i64>,
    #[prost(int64, optional, tag = "3")]
    pub group_id: ::core::option::Option<i64>,
    #[prost(int64, optional, tag = "4")]
    pub channel_id: ::core::option::Option<i64>,
    #[prost(int64, optional, tag = "5")]
    pub channel_event_id: ::core::option::Option<i64>,
    #[prost(int64, optional, tag = "6")]
    pub message_id: ::core::option::Option<i64>,
    #[prost(int64, optional, tag = "7")]
    pub stamp_id: ::core::option::Option<i64>,
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum Event {
    /// UserCreated ユーザーが追加された
    /// 	Fields:
    /// 		user_id: int64
    UserCreated = 0,
    /// UserUpdated ユーザー情報が更新された
    /// 	Fields:
    /// 		user_id: int64
    UserUpdated = 1,
    /// UserDeleted ユーザーが消去された
    /// 	Fields:
    /// 		user_id: int64
    UserDeleted = 2,
    /// UserOnline ユーザーがオンラインになった
    /// 	Fields:
    /// 		user_id: int64
    UserOnline = 5,
    /// UserOffline ユーザーがオフラインになった
    /// 	Fields:
    /// 		user_id: int64
    UserOffline = 6,
    /// UserGroupCreated ユーザーグループが作成された
    /// 	Fields:
    /// 		group_id: int64
    UserGroupCreated = 10,
    /// UserGroupUpdated ユーザーグループ設定が更新された
    /// 	Fields:
    /// 		group_id: int64
    UserGroupUpdated = 11,
    /// UserGroupDeleted ユーザーグループが消去された
    /// 	Fields:
    /// 		group_id: int64
    UserGroupDeleted = 12,
    /// MemberAdded ユーザーグループにユーザーが追加された
    /// 	Fields:
    /// 		user_id: int64
    /// 		group_id: int64
    MemberAdded = 15,
    /// MemberUpdated ユーザーグループでのユーザー設定が更新された
    /// 	Fields:
    /// 		user_id: int64
    /// 		group_id: int64
    MemberUpdated = 16,
    /// MemberRemoved ユーザーグループからユーザーが消去された
    /// 	Fields:
    /// 		user_id: int64
    /// 		group_id: int64
    MemberRemoved = 17,
    /// ChannelCreated チャンネルが作成された
    /// 	Fields:
    /// 		channel_id: int64
    ChannelCreated = 20,
    /// ChannelUpdated チャンネル設定が更新された
    /// 	Fields:
    /// 		channel_id: int64
    ChannelUpdated = 21,
    /// ChannelDeleted チャンネルが消去された
    /// 	Fields:
    /// 		channel_id: int64
    ChannelDeleted = 22,
    /// MessageChannelJoined メッセージチャンネルにユーザーが参加した
    /// 	Fields:
    /// 		user_id: int64
    /// 		channel_id: int64
    MessageChannelJoined = 25,
    /// MessageChannelLeaved メッセージチャンネルからユーザーが退室した
    /// 	Fields:
    /// 		user_id: int64
    /// 		channel_id: int64
    MessageChannelLeaved = 26,
    /// VoiceChannelJoined ボイスチャンネルにユーザーが参加した
    /// 	Fields:
    /// 		user_id: int64
    /// 		channel_id: int64
    VoiceChannelJoined = 30,
    /// VoiceChannelLeaved ボイスチャンネルからユーザーが退室した
    /// 	Fields:
    /// 		user_id: int64
    /// 		channel_id: int64
    VoiceChannelLeaved = 31,
    /// ChannelEventCreated チャンネルイベントが作成された
    /// 	Fields:
    /// 		channel_event_id: int64
    ChannelEventCreated = 35,
    /// ChannelEventUpdated チャンネルイベントが更新された
    /// 	Fields:
    /// 		channel_event_id: int64
    ChannelEventUpdated = 36,
    /// ChannelEventDeleted チャンネルイベントが消去された
    /// 	Fields:
    /// 		channel_event_id: int64
    ChannelEventDeleted = 37,
    /// ChannelEventStamped チャンネルイベントがスタンプされた
    /// 	Fields:
    /// 		user_id: int64
    /// 		channel_event_id: int64
    /// 		stamp_id: int64
    ChannelEventStamped = 40,
    /// ChannelEventUnstamped チャンネルイベントに対してのスタンプが消去された
    /// 	Fields:
    /// 		user_id: int64
    /// 		channel_event_id: int64
    /// 		stamp_id: int64
    ChannelEventUnstamped = 41,
    /// MessageCreated メッセージが作成された
    /// 	Fields:
    /// 		message_id: int64
    MessageCreated = 50,
    /// MessageUpdated メッセージが更新された
    /// 	Fields:
    /// 		message_id: int64
    MessageUpdated = 51,
    /// MessageDeleted メッセージが消去された
    /// 	Fields:
    /// 		message_id: int64
    MessageDeleted = 52,
    /// MessageStamped メッセージがスタンプされた
    /// 	Fields:
    /// 		user_id: int64
    /// 		message_id: int64
    /// 		stamp_id: int64
    MessageStamped = 55,
    /// MessageUnstamped メッセージに対してのスタンプが消去された
    /// 	Fields:
    /// 		user_id: int64
    /// 		message_id: int64
    /// 		stamp_id: int64
    MessageUnstamped = 56,
    /// MessagePinned メッセージがピン止めされた
    /// 	Fields:
    /// 		user_id: int64
    /// 		message_id: int64
    MessagePinned = 57,
    /// MessageUnpinned メッセージのピン止めが消去された
    /// 	Fields:
    /// 		user_id: int64
    /// 		message_id: int64
    MessageUnpinned = 58,
    /// StampCreated スタンプが作成された
    /// 	Fields:
    /// 		stamp_id: int64
    StampCreated = 60,
    /// StampUpdated スタンプ情報が更新された
    /// 	Fields:
    /// 		stamp_id: int64
    StampUpdated = 61,
    /// StampDeleted スタンプが消去された
    /// 	Fields:
    /// 		stamp_id: int64
    StampDeleted = 62,
}
impl Event {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            Event::UserCreated => "USER_CREATED",
            Event::UserUpdated => "USER_UPDATED",
            Event::UserDeleted => "USER_DELETED",
            Event::UserOnline => "USER_ONLINE",
            Event::UserOffline => "USER_OFFLINE",
            Event::UserGroupCreated => "USER_GROUP_CREATED",
            Event::UserGroupUpdated => "USER_GROUP_UPDATED",
            Event::UserGroupDeleted => "USER_GROUP_DELETED",
            Event::MemberAdded => "MEMBER_ADDED",
            Event::MemberUpdated => "MEMBER_UPDATED",
            Event::MemberRemoved => "MEMBER_REMOVED",
            Event::ChannelCreated => "CHANNEL_CREATED",
            Event::ChannelUpdated => "CHANNEL_UPDATED",
            Event::ChannelDeleted => "CHANNEL_DELETED",
            Event::MessageChannelJoined => "MESSAGE_CHANNEL_JOINED",
            Event::MessageChannelLeaved => "MESSAGE_CHANNEL_LEAVED",
            Event::VoiceChannelJoined => "VOICE_CHANNEL_JOINED",
            Event::VoiceChannelLeaved => "VOICE_CHANNEL_LEAVED",
            Event::ChannelEventCreated => "CHANNEL_EVENT_CREATED",
            Event::ChannelEventUpdated => "CHANNEL_EVENT_UPDATED",
            Event::ChannelEventDeleted => "CHANNEL_EVENT_DELETED",
            Event::ChannelEventStamped => "CHANNEL_EVENT_STAMPED",
            Event::ChannelEventUnstamped => "CHANNEL_EVENT_UNSTAMPED",
            Event::MessageCreated => "MESSAGE_CREATED",
            Event::MessageUpdated => "MESSAGE_UPDATED",
            Event::MessageDeleted => "MESSAGE_DELETED",
            Event::MessageStamped => "MESSAGE_STAMPED",
            Event::MessageUnstamped => "MESSAGE_UNSTAMPED",
            Event::MessagePinned => "MESSAGE_PINNED",
            Event::MessageUnpinned => "MESSAGE_UNPINNED",
            Event::StampCreated => "STAMP_CREATED",
            Event::StampUpdated => "STAMP_UPDATED",
            Event::StampDeleted => "STAMP_DELETED",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "USER_CREATED" => Some(Self::UserCreated),
            "USER_UPDATED" => Some(Self::UserUpdated),
            "USER_DELETED" => Some(Self::UserDeleted),
            "USER_ONLINE" => Some(Self::UserOnline),
            "USER_OFFLINE" => Some(Self::UserOffline),
            "USER_GROUP_CREATED" => Some(Self::UserGroupCreated),
            "USER_GROUP_UPDATED" => Some(Self::UserGroupUpdated),
            "USER_GROUP_DELETED" => Some(Self::UserGroupDeleted),
            "MEMBER_ADDED" => Some(Self::MemberAdded),
            "MEMBER_UPDATED" => Some(Self::MemberUpdated),
            "MEMBER_REMOVED" => Some(Self::MemberRemoved),
            "CHANNEL_CREATED" => Some(Self::ChannelCreated),
            "CHANNEL_UPDATED" => Some(Self::ChannelUpdated),
            "CHANNEL_DELETED" => Some(Self::ChannelDeleted),
            "MESSAGE_CHANNEL_JOINED" => Some(Self::MessageChannelJoined),
            "MESSAGE_CHANNEL_LEAVED" => Some(Self::MessageChannelLeaved),
            "VOICE_CHANNEL_JOINED" => Some(Self::VoiceChannelJoined),
            "VOICE_CHANNEL_LEAVED" => Some(Self::VoiceChannelLeaved),
            "CHANNEL_EVENT_CREATED" => Some(Self::ChannelEventCreated),
            "CHANNEL_EVENT_UPDATED" => Some(Self::ChannelEventUpdated),
            "CHANNEL_EVENT_DELETED" => Some(Self::ChannelEventDeleted),
            "CHANNEL_EVENT_STAMPED" => Some(Self::ChannelEventStamped),
            "CHANNEL_EVENT_UNSTAMPED" => Some(Self::ChannelEventUnstamped),
            "MESSAGE_CREATED" => Some(Self::MessageCreated),
            "MESSAGE_UPDATED" => Some(Self::MessageUpdated),
            "MESSAGE_DELETED" => Some(Self::MessageDeleted),
            "MESSAGE_STAMPED" => Some(Self::MessageStamped),
            "MESSAGE_UNSTAMPED" => Some(Self::MessageUnstamped),
            "MESSAGE_PINNED" => Some(Self::MessagePinned),
            "MESSAGE_UNPINNED" => Some(Self::MessageUnpinned),
            "STAMP_CREATED" => Some(Self::StampCreated),
            "STAMP_UPDATED" => Some(Self::StampUpdated),
            "STAMP_DELETED" => Some(Self::StampDeleted),
            _ => None,
        }
    }
}
/// Generated client implementations.
pub mod event_service_client {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::http::Uri;
    use tonic::codegen::*;
    #[derive(Debug, Clone)]
    pub struct EventServiceClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl EventServiceClient<tonic::transport::Channel> {
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
    impl<T> EventServiceClient<T>
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
        ) -> EventServiceClient<InterceptedService<T, F>>
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
            EventServiceClient::new(InterceptedService::new(inner, interceptor))
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
        pub async fn event_stream(
            &mut self,
            request: impl tonic::IntoRequest<super::super::super::google::protobuf::Empty>,
        ) -> std::result::Result<
            tonic::Response<tonic::codec::Streaming<super::EventResponse>>,
            tonic::Status,
        > {
            self.inner.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e.into()),
                )
            })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/event.v1.EventService/EventStream");
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("event.v1.EventService", "EventStream"));
            self.inner.server_streaming(req, path, codec).await
        }
    }
}
/// Generated server implementations.
pub mod event_service_server {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    /// Generated trait containing gRPC methods that should be implemented for use with EventServiceServer.
    #[async_trait]
    pub trait EventService: Send + Sync + 'static {
        /// Server streaming response type for the EventStream method.
        type EventStreamStream: tonic::codegen::tokio_stream::Stream<
                Item = std::result::Result<super::EventResponse, tonic::Status>,
            > + Send
            + 'static;
        async fn event_stream(
            &self,
            request: tonic::Request<super::super::super::google::protobuf::Empty>,
        ) -> std::result::Result<tonic::Response<Self::EventStreamStream>, tonic::Status>;
    }
    #[derive(Debug)]
    pub struct EventServiceServer<T: EventService> {
        inner: _Inner<T>,
        accept_compression_encodings: EnabledCompressionEncodings,
        send_compression_encodings: EnabledCompressionEncodings,
        max_decoding_message_size: Option<usize>,
        max_encoding_message_size: Option<usize>,
    }
    struct _Inner<T>(Arc<T>);
    impl<T: EventService> EventServiceServer<T> {
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
    impl<T, B> tonic::codegen::Service<http::Request<B>> for EventServiceServer<T>
    where
        T: EventService,
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
                "/event.v1.EventService/EventStream" => {
                    #[allow(non_camel_case_types)]
                    struct EventStreamSvc<T: EventService>(pub Arc<T>);
                    impl<T: EventService>
                        tonic::server::ServerStreamingService<
                            super::super::super::google::protobuf::Empty,
                        > for EventStreamSvc<T>
                    {
                        type Response = super::EventResponse;
                        type ResponseStream = T::EventStreamStream;
                        type Future =
                            BoxFuture<tonic::Response<Self::ResponseStream>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::super::super::google::protobuf::Empty>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as EventService>::event_stream(&inner, request).await
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
                        let method = EventStreamSvc(inner);
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
                        let res = grpc.server_streaming(method, req).await;
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
    impl<T: EventService> Clone for EventServiceServer<T> {
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
    impl<T: EventService> Clone for _Inner<T> {
        fn clone(&self) -> Self {
            Self(Arc::clone(&self.0))
        }
    }
    impl<T: std::fmt::Debug> std::fmt::Debug for _Inner<T> {
        fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
            write!(f, "{:?}", self.0)
        }
    }
    impl<T: EventService> tonic::server::NamedService for EventServiceServer<T> {
        const NAME: &'static str = "event.v1.EventService";
    }
}
