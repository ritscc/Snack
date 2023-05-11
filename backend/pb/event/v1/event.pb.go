// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: event/v1/event.proto

package event

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Event int32

const (
	// UserCreated ユーザーが追加された
	//
	//	Fields:
	//		user_id: int64
	Event_USER_CREATED Event = 0
	// UserUpdated ユーザー情報が更新された
	//
	//	Fields:
	//		user_id: int64
	Event_USER_UPDATED Event = 1
	// UserDeleted ユーザーが消去された
	//
	//	Fields:
	//		user_id: int64
	Event_USER_DELETED Event = 2
	// UserOnline ユーザーがオンラインになった
	//
	//	Fields:
	//		user_id: int64
	Event_USER_ONLINE Event = 5
	// UserOffline ユーザーがオフラインになった
	//
	//	Fields:
	//		user_id: int64
	Event_USER_OFFLINE Event = 6
	// UserGroupCreated ユーザーグループが作成された
	//
	//	Fields:
	//		group_id: int64
	Event_USER_GROUP_CREATED Event = 10
	// UserGroupUpdated ユーザーグループ設定が更新された
	//
	//	Fields:
	//		group_id: int64
	Event_USER_GROUP_UPDATED Event = 11
	// UserGroupDeleted ユーザーグループが消去された
	//
	//	Fields:
	//		group_id: int64
	Event_USER_GROUP_DELETED Event = 12
	// MemberAdded ユーザーグループにユーザーが追加された
	//
	//	Fields:
	//		user_id: int64
	//		group_id: int64
	Event_MEMBER_ADDED Event = 15
	// MemberUpdated ユーザーグループでのユーザー設定が更新された
	//
	//	Fields:
	//		user_id: int64
	//		group_id: int64
	Event_MEMBER_UPDATED Event = 16
	// MemberRemoved ユーザーグループからユーザーが消去された
	//
	//	Fields:
	//		user_id: int64
	//		group_id: int64
	Event_MEMBER_REMOVED Event = 17
	// ChannelCreated チャンネルが作成された
	//
	//	Fields:
	//		channel_id: int64
	Event_CHANNEL_CREATED Event = 20
	// ChannelUpdated チャンネル設定が更新された
	//
	//	Fields:
	//		channel_id: int64
	Event_CHANNEL_UPDATED Event = 21
	// ChannelDeleted チャンネルが消去された
	//
	//	Fields:
	//		channel_id: int64
	Event_CHANNEL_DELETED Event = 22
	// MessageChannelJoined メッセージチャンネルにユーザーが参加した
	//
	//	Fields:
	//		user_id: int64
	//		channel_id: int64
	Event_MESSAGE_CHANNEL_JOINED Event = 25
	// MessageChannelLeaved メッセージチャンネルからユーザーが退室した
	//
	//	Fields:
	//		user_id: int64
	//		channel_id: int64
	Event_MESSAGE_CHANNEL_LEAVED Event = 26
	// VoiceChannelJoined ボイスチャンネルにユーザーが参加した
	//
	//	Fields:
	//		user_id: int64
	//		channel_id: int64
	Event_VOICE_CHANNEL_JOINED Event = 30
	// VoiceChannelLeaved ボイスチャンネルからユーザーが退室した
	//
	//	Fields:
	//		user_id: int64
	//		channel_id: int64
	Event_VOICE_CHANNEL_LEAVED Event = 31
	// ChannelEventCreated チャンネルイベントが作成された
	//
	//	Fields:
	//		channel_event_id: int64
	Event_CHANNEL_EVENT_CREATED Event = 35
	// ChannelEventUpdated チャンネルイベントが更新された
	//
	//	Fields:
	//		channel_event_id: int64
	Event_CHANNEL_EVENT_UPDATED Event = 36
	// ChannelEventDeleted チャンネルイベントが消去された
	//
	//	Fields:
	//		channel_event_id: int64
	Event_CHANNEL_EVENT_DELETED Event = 37
	// ChannelEventStamped チャンネルイベントがスタンプされた
	//
	//	Fields:
	//		user_id: int64
	//		channel_event_id: int64
	//		stamp_id: int64
	Event_CHANNEL_EVENT_STAMPED Event = 40
	// ChannelEventUnstamped チャンネルイベントに対してのスタンプが消去された
	//
	//	Fields:
	//		user_id: int64
	//		channel_event_id: int64
	//		stamp_id: int64
	Event_CHANNEL_EVENT_UNSTAMPED Event = 41
	// MessageCreated メッセージが作成された
	//
	//	Fields:
	//		message_id: int64
	Event_MESSAGE_CREATED Event = 50
	// MessageUpdated メッセージが更新された
	//
	//	Fields:
	//		message_id: int64
	Event_MESSAGE_UPDATED Event = 51
	// MessageDeleted メッセージが消去された
	//
	//	Fields:
	//		message_id: int64
	Event_MESSAGE_DELETED Event = 52
	// MessageStamped メッセージがスタンプされた
	//
	//	Fields:
	//		user_id: int64
	//		message_id: int64
	//		stamp_id: int64
	Event_MESSAGE_STAMPED Event = 55
	// MessageUnstamped メッセージに対してのスタンプが消去された
	//
	//	Fields:
	//		user_id: int64
	//		message_id: int64
	//		stamp_id: int64
	Event_MESSAGE_UNSTAMPED Event = 56
	// MessagePinned メッセージがピン止めされた
	//
	//	Fields:
	//		user_id: int64
	//		message_id: int64
	Event_MESSAGE_PINNED Event = 57
	// MessageUnpinned メッセージのピン止めが消去された
	//
	//	Fields:
	//		user_id: int64
	//		message_id: int64
	Event_MESSAGE_UNPINNED Event = 58
	// StampCreated スタンプが作成された
	//
	//	Fields:
	//		stamp_id: int64
	Event_STAMP_CREATED Event = 60
	// StampUpdated スタンプ情報が更新された
	//
	//	Fields:
	//		stamp_id: int64
	Event_STAMP_UPDATED Event = 61
	// StampDeleted スタンプが消去された
	//
	//	Fields:
	//		stamp_id: int64
	Event_STAMP_DELETED Event = 62
)

// Enum value maps for Event.
var (
	Event_name = map[int32]string{
		0:  "USER_CREATED",
		1:  "USER_UPDATED",
		2:  "USER_DELETED",
		5:  "USER_ONLINE",
		6:  "USER_OFFLINE",
		10: "USER_GROUP_CREATED",
		11: "USER_GROUP_UPDATED",
		12: "USER_GROUP_DELETED",
		15: "MEMBER_ADDED",
		16: "MEMBER_UPDATED",
		17: "MEMBER_REMOVED",
		20: "CHANNEL_CREATED",
		21: "CHANNEL_UPDATED",
		22: "CHANNEL_DELETED",
		25: "MESSAGE_CHANNEL_JOINED",
		26: "MESSAGE_CHANNEL_LEAVED",
		30: "VOICE_CHANNEL_JOINED",
		31: "VOICE_CHANNEL_LEAVED",
		35: "CHANNEL_EVENT_CREATED",
		36: "CHANNEL_EVENT_UPDATED",
		37: "CHANNEL_EVENT_DELETED",
		40: "CHANNEL_EVENT_STAMPED",
		41: "CHANNEL_EVENT_UNSTAMPED",
		50: "MESSAGE_CREATED",
		51: "MESSAGE_UPDATED",
		52: "MESSAGE_DELETED",
		55: "MESSAGE_STAMPED",
		56: "MESSAGE_UNSTAMPED",
		57: "MESSAGE_PINNED",
		58: "MESSAGE_UNPINNED",
		60: "STAMP_CREATED",
		61: "STAMP_UPDATED",
		62: "STAMP_DELETED",
	}
	Event_value = map[string]int32{
		"USER_CREATED":            0,
		"USER_UPDATED":            1,
		"USER_DELETED":            2,
		"USER_ONLINE":             5,
		"USER_OFFLINE":            6,
		"USER_GROUP_CREATED":      10,
		"USER_GROUP_UPDATED":      11,
		"USER_GROUP_DELETED":      12,
		"MEMBER_ADDED":            15,
		"MEMBER_UPDATED":          16,
		"MEMBER_REMOVED":          17,
		"CHANNEL_CREATED":         20,
		"CHANNEL_UPDATED":         21,
		"CHANNEL_DELETED":         22,
		"MESSAGE_CHANNEL_JOINED":  25,
		"MESSAGE_CHANNEL_LEAVED":  26,
		"VOICE_CHANNEL_JOINED":    30,
		"VOICE_CHANNEL_LEAVED":    31,
		"CHANNEL_EVENT_CREATED":   35,
		"CHANNEL_EVENT_UPDATED":   36,
		"CHANNEL_EVENT_DELETED":   37,
		"CHANNEL_EVENT_STAMPED":   40,
		"CHANNEL_EVENT_UNSTAMPED": 41,
		"MESSAGE_CREATED":         50,
		"MESSAGE_UPDATED":         51,
		"MESSAGE_DELETED":         52,
		"MESSAGE_STAMPED":         55,
		"MESSAGE_UNSTAMPED":       56,
		"MESSAGE_PINNED":          57,
		"MESSAGE_UNPINNED":        58,
		"STAMP_CREATED":           60,
		"STAMP_UPDATED":           61,
		"STAMP_DELETED":           62,
	}
)

func (x Event) Enum() *Event {
	p := new(Event)
	*p = x
	return p
}

func (x Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event) Descriptor() protoreflect.EnumDescriptor {
	return file_event_v1_event_proto_enumTypes[0].Descriptor()
}

func (Event) Type() protoreflect.EnumType {
	return &file_event_v1_event_proto_enumTypes[0]
}

func (x Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event.Descriptor instead.
func (Event) EnumDescriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{0}
}

type EventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType      Event  `protobuf:"varint,1,opt,name=event_type,json=eventType,proto3,enum=event.v1.Event" json:"event_type,omitempty"`
	UserId         *int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
	GroupId        *int64 `protobuf:"varint,3,opt,name=group_id,json=groupId,proto3,oneof" json:"group_id,omitempty"`
	ChannelId      *int64 `protobuf:"varint,4,opt,name=channel_id,json=channelId,proto3,oneof" json:"channel_id,omitempty"`
	ChannelEventId *int64 `protobuf:"varint,5,opt,name=channel_event_id,json=channelEventId,proto3,oneof" json:"channel_event_id,omitempty"`
	MessageId      *int64 `protobuf:"varint,6,opt,name=message_id,json=messageId,proto3,oneof" json:"message_id,omitempty"`
	StampId        *int64 `protobuf:"varint,7,opt,name=stamp_id,json=stampId,proto3,oneof" json:"stamp_id,omitempty"`
}

func (x *EventResponse) Reset() {
	*x = EventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_v1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponse) ProtoMessage() {}

func (x *EventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_v1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponse.ProtoReflect.Descriptor instead.
func (*EventResponse) Descriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *EventResponse) GetEventType() Event {
	if x != nil {
		return x.EventType
	}
	return Event_USER_CREATED
}

func (x *EventResponse) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *EventResponse) GetGroupId() int64 {
	if x != nil && x.GroupId != nil {
		return *x.GroupId
	}
	return 0
}

func (x *EventResponse) GetChannelId() int64 {
	if x != nil && x.ChannelId != nil {
		return *x.ChannelId
	}
	return 0
}

func (x *EventResponse) GetChannelEventId() int64 {
	if x != nil && x.ChannelEventId != nil {
		return *x.ChannelEventId
	}
	return 0
}

func (x *EventResponse) GetMessageId() int64 {
	if x != nil && x.MessageId != nil {
		return *x.MessageId
	}
	return 0
}

func (x *EventResponse) GetStampId() int64 {
	if x != nil && x.StampId != nil {
		return *x.StampId
	}
	return 0
}

var File_event_v1_event_proto protoreflect.FileDescriptor

var file_event_v1_event_proto_rawDesc = []byte{
	0x0a, 0x14, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x02,
	0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1c, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a,
	0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x01, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a,
	0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x02, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x2d, 0x0a, 0x10, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x03, 0x52, 0x0e, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x22, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x04, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x48, 0x05, 0x52, 0x07, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x42, 0x13, 0x0a, 0x11,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x69, 0x64, 0x2a, 0xe4, 0x05,
	0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x05, 0x12, 0x10,
	0x0a, 0x0c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x06,
	0x12, 0x16, 0x0a, 0x12, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x0b,
	0x12, 0x16, 0x0a, 0x12, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x45, 0x4d, 0x42,
	0x45, 0x52, 0x5f, 0x41, 0x44, 0x44, 0x45, 0x44, 0x10, 0x0f, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x45,
	0x4d, 0x42, 0x45, 0x52, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x10, 0x12, 0x12,
	0x0a, 0x0e, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44,
	0x10, 0x11, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x14, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x48, 0x41, 0x4e, 0x4e,
	0x45, 0x4c, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x15, 0x12, 0x13, 0x0a, 0x0f,
	0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10,
	0x16, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x19, 0x12, 0x1a, 0x0a,
	0x16, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c,
	0x5f, 0x4c, 0x45, 0x41, 0x56, 0x45, 0x44, 0x10, 0x1a, 0x12, 0x18, 0x0a, 0x14, 0x56, 0x4f, 0x49,
	0x43, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x45,
	0x44, 0x10, 0x1e, 0x12, 0x18, 0x0a, 0x14, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x4c, 0x45, 0x41, 0x56, 0x45, 0x44, 0x10, 0x1f, 0x12, 0x19, 0x0a,
	0x15, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x23, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x48, 0x41, 0x4e,
	0x4e, 0x45, 0x4c, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x24, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x25, 0x12, 0x19,
	0x0a, 0x15, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x53, 0x54, 0x41, 0x4d, 0x50, 0x45, 0x44, 0x10, 0x28, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x48, 0x41,
	0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x54, 0x41,
	0x4d, 0x50, 0x45, 0x44, 0x10, 0x29, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47,
	0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x32, 0x12, 0x13, 0x0a, 0x0f, 0x4d,
	0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x33,
	0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x45, 0x44, 0x10, 0x34, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45,
	0x5f, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x45, 0x44, 0x10, 0x37, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x45, 0x44, 0x10,
	0x38, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x49, 0x4e,
	0x4e, 0x45, 0x44, 0x10, 0x39, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45,
	0x5f, 0x55, 0x4e, 0x50, 0x49, 0x4e, 0x4e, 0x45, 0x44, 0x10, 0x3a, 0x12, 0x11, 0x0a, 0x0d, 0x53,
	0x54, 0x41, 0x4d, 0x50, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x3c, 0x12, 0x11,
	0x0a, 0x0d, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x3d, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54,
	0x45, 0x44, 0x10, 0x3e, 0x32, 0x52, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x74, 0x73, 0x63, 0x63, 0x2f, 0x53, 0x6e,
	0x61, 0x63, 0x6b, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x62, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_v1_event_proto_rawDescOnce sync.Once
	file_event_v1_event_proto_rawDescData = file_event_v1_event_proto_rawDesc
)

func file_event_v1_event_proto_rawDescGZIP() []byte {
	file_event_v1_event_proto_rawDescOnce.Do(func() {
		file_event_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_v1_event_proto_rawDescData)
	})
	return file_event_v1_event_proto_rawDescData
}

var file_event_v1_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_event_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_event_v1_event_proto_goTypes = []interface{}{
	(Event)(0),            // 0: event.v1.Event
	(*EventResponse)(nil), // 1: event.v1.EventResponse
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
}
var file_event_v1_event_proto_depIdxs = []int32{
	0, // 0: event.v1.EventResponse.event_type:type_name -> event.v1.Event
	2, // 1: event.v1.EventService.EventStream:input_type -> google.protobuf.Empty
	1, // 2: event.v1.EventService.EventStream:output_type -> event.v1.EventResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_event_v1_event_proto_init() }
func file_event_v1_event_proto_init() {
	if File_event_v1_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_v1_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_event_v1_event_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_v1_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_event_v1_event_proto_goTypes,
		DependencyIndexes: file_event_v1_event_proto_depIdxs,
		EnumInfos:         file_event_v1_event_proto_enumTypes,
		MessageInfos:      file_event_v1_event_proto_msgTypes,
	}.Build()
	File_event_v1_event_proto = out.File
	file_event_v1_event_proto_rawDesc = nil
	file_event_v1_event_proto_goTypes = nil
	file_event_v1_event_proto_depIdxs = nil
}
