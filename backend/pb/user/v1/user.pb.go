// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: user/v1/user.proto

package user

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

type Status int32

const (
	Status_ONLINE  Status = 0
	Status_OFFLINE Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "ONLINE",
		1: "OFFLINE",
	}
	Status_value = map[string]int32{
		"ONLINE":  0,
		"OFFLINE": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_user_v1_user_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_user_v1_user_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // unique
	Username  string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`            // unique
	Nick      string   `protobuf:"bytes,3,opt,name=nick,proto3" json:"nick,omitempty"`                    // nickname
	RealName  string   `protobuf:"bytes,4,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	Avatar    string   `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"` // icon
	Role      string   `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
	Locale    string   `protobuf:"bytes,7,opt,name=locale,proto3" json:"locale,omitempty"`                           // Japan or Korea
	RitsEmail string   `protobuf:"bytes,8,opt,name=rits_email,json=ritsEmail,proto3" json:"rits_email,omitempty"`    // ed.ritsumei.ac.jp
	OwnEmail  *string  `protobuf:"bytes,9,opt,name=own_email,json=ownEmail,proto3,oneof" json:"own_email,omitempty"` // gmail.com など
	Comment   string   `protobuf:"bytes,10,opt,name=comment,proto3" json:"comment,omitempty"`
	Status    Status   `protobuf:"varint,11,opt,name=status,proto3,enum=user.v1.Status" json:"status,omitempty"` // ONLINE or OFFLINE
	Services  *Service `protobuf:"bytes,20,opt,name=services,proto3,oneof" json:"services,omitempty"`            // Twitter, Github, Discord
	IsDeleted bool     `protobuf:"varint,30,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetNick() string {
	if x != nil {
		return x.Nick
	}
	return ""
}

func (x *User) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *User) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *User) GetRitsEmail() string {
	if x != nil {
		return x.RitsEmail
	}
	return ""
}

func (x *User) GetOwnEmail() string {
	if x != nil && x.OwnEmail != nil {
		return *x.OwnEmail
	}
	return ""
}

func (x *User) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *User) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_ONLINE
}

func (x *User) GetServices() *Service {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *User) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Twitter *string `protobuf:"bytes,1,opt,name=twitter,proto3,oneof" json:"twitter,omitempty"`
	Github  *string `protobuf:"bytes,2,opt,name=github,proto3,oneof" json:"github,omitempty"`
	Discord *string `protobuf:"bytes,3,opt,name=discord,proto3,oneof" json:"discord,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *Service) GetTwitter() string {
	if x != nil && x.Twitter != nil {
		return *x.Twitter
	}
	return ""
}

func (x *Service) GetGithub() string {
	if x != nil && x.Github != nil {
		return *x.Github
	}
	return ""
}

func (x *Service) GetDiscord() string {
	if x != nil && x.Discord != nil {
		return *x.Discord
	}
	return ""
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_user_v1_user_proto protoreflect.FileDescriptor

var file_user_v1_user_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x03, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x69, 0x63, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x69, 0x63, 0x6b, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x69, 0x74, 0x73, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x69, 0x74, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x09,
	0x6f, 0x77, 0x6e, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x08, 0x6f, 0x77, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x31, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x01, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6f, 0x77, 0x6e, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x87,
	0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x74, 0x77,
	0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x74,
	0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x72, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65,
	0x72, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x61, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x29, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x2a, 0x21, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x32, 0xfd, 0x01, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x33, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x74, 0x73, 0x63, 0x63, 0x2f, 0x53,
	0x6e, 0x61, 0x63, 0x6b, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x62, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_v1_user_proto_rawDescOnce sync.Once
	file_user_v1_user_proto_rawDescData = file_user_v1_user_proto_rawDesc
)

func file_user_v1_user_proto_rawDescGZIP() []byte {
	file_user_v1_user_proto_rawDescOnce.Do(func() {
		file_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_v1_user_proto_rawDescData)
	})
	return file_user_v1_user_proto_rawDescData
}

var file_user_v1_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_v1_user_proto_goTypes = []interface{}{
	(Status)(0),               // 0: user.v1.Status
	(*User)(nil),              // 1: user.v1.User
	(*Service)(nil),           // 2: user.v1.Service
	(*CreateUserRequest)(nil), // 3: user.v1.CreateUserRequest
	(*GetUserRequest)(nil),    // 4: user.v1.GetUserRequest
	(*emptypb.Empty)(nil),     // 5: google.protobuf.Empty
}
var file_user_v1_user_proto_depIdxs = []int32{
	0, // 0: user.v1.User.status:type_name -> user.v1.Status
	2, // 1: user.v1.User.services:type_name -> user.v1.Service
	3, // 2: user.v1.UserService.CreateUser:input_type -> user.v1.CreateUserRequest
	4, // 3: user.v1.UserService.GetUser:input_type -> user.v1.GetUserRequest
	1, // 4: user.v1.UserService.UpdateUser:input_type -> user.v1.User
	5, // 5: user.v1.UserService.DeleteUser:input_type -> google.protobuf.Empty
	5, // 6: user.v1.UserService.CreateUser:output_type -> google.protobuf.Empty
	1, // 7: user.v1.UserService.GetUser:output_type -> user.v1.User
	5, // 8: user.v1.UserService.UpdateUser:output_type -> google.protobuf.Empty
	5, // 9: user.v1.UserService.DeleteUser:output_type -> google.protobuf.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_v1_user_proto_init() }
func file_user_v1_user_proto_init() {
	if File_user_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_user_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_user_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_user_v1_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
	file_user_v1_user_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_user_v1_user_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_v1_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_v1_user_proto_goTypes,
		DependencyIndexes: file_user_v1_user_proto_depIdxs,
		EnumInfos:         file_user_v1_user_proto_enumTypes,
		MessageInfos:      file_user_v1_user_proto_msgTypes,
	}.Build()
	File_user_v1_user_proto = out.File
	file_user_v1_user_proto_rawDesc = nil
	file_user_v1_user_proto_goTypes = nil
	file_user_v1_user_proto_depIdxs = nil
}
