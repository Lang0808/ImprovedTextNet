// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: BlogService.proto

package GrpcBlogService

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PrivacyMode int32

const (
	PrivacyMode_PrivacyMode_PUBLIC  PrivacyMode = 0
	PrivacyMode_PrivacyMode_FRIEND  PrivacyMode = 1
	PrivacyMode_PrivacyMode_PRIVATE PrivacyMode = 2
)

// Enum value maps for PrivacyMode.
var (
	PrivacyMode_name = map[int32]string{
		0: "PrivacyMode_PUBLIC",
		1: "PrivacyMode_FRIEND",
		2: "PrivacyMode_PRIVATE",
	}
	PrivacyMode_value = map[string]int32{
		"PrivacyMode_PUBLIC":  0,
		"PrivacyMode_FRIEND":  1,
		"PrivacyMode_PRIVATE": 2,
	}
)

func (x PrivacyMode) Enum() *PrivacyMode {
	p := new(PrivacyMode)
	*p = x
	return p
}

func (x PrivacyMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PrivacyMode) Descriptor() protoreflect.EnumDescriptor {
	return file_BlogService_proto_enumTypes[0].Descriptor()
}

func (PrivacyMode) Type() protoreflect.EnumType {
	return &file_BlogService_proto_enumTypes[0]
}

func (x PrivacyMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PrivacyMode.Descriptor instead.
func (PrivacyMode) EnumDescriptor() ([]byte, []int) {
	return file_BlogService_proto_rawDescGZIP(), []int{0}
}

type CreateBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string      `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Privacy PrivacyMode `protobuf:"varint,2,opt,name=Privacy,proto3,enum=PrivacyMode" json:"Privacy,omitempty"`
	Author  int32       `protobuf:"varint,3,opt,name=Author,proto3" json:"Author,omitempty"`
	Content string      `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *CreateBlogRequest) Reset() {
	*x = CreateBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BlogService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogRequest) ProtoMessage() {}

func (x *CreateBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_BlogService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogRequest.ProtoReflect.Descriptor instead.
func (*CreateBlogRequest) Descriptor() ([]byte, []int) {
	return file_BlogService_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBlogRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateBlogRequest) GetPrivacy() PrivacyMode {
	if x != nil {
		return x.Privacy
	}
	return PrivacyMode_PrivacyMode_PUBLIC
}

func (x *CreateBlogRequest) GetAuthor() int32 {
	if x != nil {
		return x.Author
	}
	return 0
}

func (x *CreateBlogRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreateBlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateBlogResponse) Reset() {
	*x = CreateBlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BlogService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogResponse) ProtoMessage() {}

func (x *CreateBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_BlogService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogResponse.ProtoReflect.Descriptor instead.
func (*CreateBlogResponse) Descriptor() ([]byte, []int) {
	return file_BlogService_proto_rawDescGZIP(), []int{1}
}

type MultiGetBlogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogIds []int32 `protobuf:"varint,1,rep,packed,name=BlogIds,proto3" json:"BlogIds,omitempty"`
}

func (x *MultiGetBlogsRequest) Reset() {
	*x = MultiGetBlogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BlogService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiGetBlogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiGetBlogsRequest) ProtoMessage() {}

func (x *MultiGetBlogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_BlogService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiGetBlogsRequest.ProtoReflect.Descriptor instead.
func (*MultiGetBlogsRequest) Descriptor() ([]byte, []int) {
	return file_BlogService_proto_rawDescGZIP(), []int{2}
}

func (x *MultiGetBlogsRequest) GetBlogIds() []int32 {
	if x != nil {
		return x.BlogIds
	}
	return nil
}

type Blog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string      `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Privacy     PrivacyMode `protobuf:"varint,2,opt,name=Privacy,proto3,enum=PrivacyMode" json:"Privacy,omitempty"`
	Author      int32       `protobuf:"varint,3,opt,name=Author,proto3" json:"Author,omitempty"`
	Content     string      `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	CreatedTime int64       `protobuf:"varint,5,opt,name=CreatedTime,proto3" json:"CreatedTime,omitempty"`
	UpdatedTime int64       `protobuf:"varint,6,opt,name=UpdatedTime,proto3" json:"UpdatedTime,omitempty"`
}

func (x *Blog) Reset() {
	*x = Blog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BlogService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_BlogService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blog.ProtoReflect.Descriptor instead.
func (*Blog) Descriptor() ([]byte, []int) {
	return file_BlogService_proto_rawDescGZIP(), []int{3}
}

func (x *Blog) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Blog) GetPrivacy() PrivacyMode {
	if x != nil {
		return x.Privacy
	}
	return PrivacyMode_PrivacyMode_PUBLIC
}

func (x *Blog) GetAuthor() int32 {
	if x != nil {
		return x.Author
	}
	return 0
}

func (x *Blog) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Blog) GetCreatedTime() int64 {
	if x != nil {
		return x.CreatedTime
	}
	return 0
}

func (x *Blog) GetUpdatedTime() int64 {
	if x != nil {
		return x.UpdatedTime
	}
	return 0
}

type MultiGetBlogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blogs []*Blog `protobuf:"bytes,1,rep,name=Blogs,proto3" json:"Blogs,omitempty"`
}

func (x *MultiGetBlogsResponse) Reset() {
	*x = MultiGetBlogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BlogService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiGetBlogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiGetBlogsResponse) ProtoMessage() {}

func (x *MultiGetBlogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_BlogService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiGetBlogsResponse.ProtoReflect.Descriptor instead.
func (*MultiGetBlogsResponse) Descriptor() ([]byte, []int) {
	return file_BlogService_proto_rawDescGZIP(), []int{4}
}

func (x *MultiGetBlogsResponse) GetBlogs() []*Blog {
	if x != nil {
		return x.Blogs
	}
	return nil
}

type GetBlogsOfUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32 `protobuf:"varint,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetBlogsOfUserRequest) Reset() {
	*x = GetBlogsOfUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BlogService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlogsOfUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogsOfUserRequest) ProtoMessage() {}

func (x *GetBlogsOfUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_BlogService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlogsOfUserRequest.ProtoReflect.Descriptor instead.
func (*GetBlogsOfUserRequest) Descriptor() ([]byte, []int) {
	return file_BlogService_proto_rawDescGZIP(), []int{5}
}

func (x *GetBlogsOfUserRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetBlogsOfUserRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetBlogsOfUserRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetBlogsOfUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blogs []*Blog `protobuf:"bytes,1,rep,name=Blogs,proto3" json:"Blogs,omitempty"`
}

func (x *GetBlogsOfUserResponse) Reset() {
	*x = GetBlogsOfUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BlogService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlogsOfUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogsOfUserResponse) ProtoMessage() {}

func (x *GetBlogsOfUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_BlogService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlogsOfUserResponse.ProtoReflect.Descriptor instead.
func (*GetBlogsOfUserResponse) Descriptor() ([]byte, []int) {
	return file_BlogService_proto_rawDescGZIP(), []int{6}
}

func (x *GetBlogsOfUserResponse) GetBlogs() []*Blog {
	if x != nil {
		return x.Blogs
	}
	return nil
}

type GetUserBlogsWithRelationshipsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int32   `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Offset       int32   `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit        int32   `protobuf:"varint,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Relationship []int32 `protobuf:"varint,4,rep,packed,name=Relationship,proto3" json:"Relationship,omitempty"`
}

func (x *GetUserBlogsWithRelationshipsRequest) Reset() {
	*x = GetUserBlogsWithRelationshipsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BlogService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserBlogsWithRelationshipsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserBlogsWithRelationshipsRequest) ProtoMessage() {}

func (x *GetUserBlogsWithRelationshipsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_BlogService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserBlogsWithRelationshipsRequest.ProtoReflect.Descriptor instead.
func (*GetUserBlogsWithRelationshipsRequest) Descriptor() ([]byte, []int) {
	return file_BlogService_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserBlogsWithRelationshipsRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetUserBlogsWithRelationshipsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetUserBlogsWithRelationshipsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetUserBlogsWithRelationshipsRequest) GetRelationship() []int32 {
	if x != nil {
		return x.Relationship
	}
	return nil
}

var File_BlogService_proto protoreflect.FileDescriptor

var file_BlogService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x26, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0c, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x07,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x30, 0x0a, 0x14, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x67, 0x49,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x64,
	0x73, 0x22, 0xba, 0x01, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x26, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x34,
	0x0a, 0x15, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x05, 0x42,
	0x6c, 0x6f, 0x67, 0x73, 0x22, 0x5d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73,
	0x4f, 0x66, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x35, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x4f,
	0x66, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a,
	0x05, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42,
	0x6c, 0x6f, 0x67, 0x52, 0x05, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x24, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x57, 0x69, 0x74, 0x68, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x0c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2a, 0x56, 0x0a,
	0x0b, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x12,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x50, 0x55, 0x42, 0x4c,
	0x49, 0x43, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4d,
	0x6f, 0x64, 0x65, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x50, 0x52, 0x49, 0x56,
	0x41, 0x54, 0x45, 0x10, 0x02, 0x32, 0xa9, 0x02, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x6c, 0x6f, 0x67, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x0d, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x15,
	0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x4f, 0x66, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x16, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x4f, 0x66, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f,
	0x67, 0x73, 0x4f, 0x66, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5f, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x67, 0x73,
	0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x73, 0x12, 0x25, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x67, 0x73,
	0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x67, 0x73, 0x4f, 0x66, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x47, 0x72, 0x70, 0x63, 0x42, 0x6c, 0x6f, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BlogService_proto_rawDescOnce sync.Once
	file_BlogService_proto_rawDescData = file_BlogService_proto_rawDesc
)

func file_BlogService_proto_rawDescGZIP() []byte {
	file_BlogService_proto_rawDescOnce.Do(func() {
		file_BlogService_proto_rawDescData = protoimpl.X.CompressGZIP(file_BlogService_proto_rawDescData)
	})
	return file_BlogService_proto_rawDescData
}

var file_BlogService_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_BlogService_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_BlogService_proto_goTypes = []interface{}{
	(PrivacyMode)(0),                             // 0: PrivacyMode
	(*CreateBlogRequest)(nil),                    // 1: CreateBlogRequest
	(*CreateBlogResponse)(nil),                   // 2: CreateBlogResponse
	(*MultiGetBlogsRequest)(nil),                 // 3: MultiGetBlogsRequest
	(*Blog)(nil),                                 // 4: Blog
	(*MultiGetBlogsResponse)(nil),                // 5: MultiGetBlogsResponse
	(*GetBlogsOfUserRequest)(nil),                // 6: GetBlogsOfUserRequest
	(*GetBlogsOfUserResponse)(nil),               // 7: GetBlogsOfUserResponse
	(*GetUserBlogsWithRelationshipsRequest)(nil), // 8: GetUserBlogsWithRelationshipsRequest
}
var file_BlogService_proto_depIdxs = []int32{
	0, // 0: CreateBlogRequest.Privacy:type_name -> PrivacyMode
	0, // 1: Blog.Privacy:type_name -> PrivacyMode
	4, // 2: MultiGetBlogsResponse.Blogs:type_name -> Blog
	4, // 3: GetBlogsOfUserResponse.Blogs:type_name -> Blog
	1, // 4: BlogService.CreatedBlog:input_type -> CreateBlogRequest
	3, // 5: BlogService.MultiGetBlogs:input_type -> MultiGetBlogsRequest
	6, // 6: BlogService.GetBlogsOfUser:input_type -> GetBlogsOfUserRequest
	8, // 7: BlogService.GetUserBlogsWithRelationships:input_type -> GetUserBlogsWithRelationshipsRequest
	2, // 8: BlogService.CreatedBlog:output_type -> CreateBlogResponse
	5, // 9: BlogService.MultiGetBlogs:output_type -> MultiGetBlogsResponse
	7, // 10: BlogService.GetBlogsOfUser:output_type -> GetBlogsOfUserResponse
	7, // 11: BlogService.GetUserBlogsWithRelationships:output_type -> GetBlogsOfUserResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_BlogService_proto_init() }
func file_BlogService_proto_init() {
	if File_BlogService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BlogService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlogRequest); i {
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
		file_BlogService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlogResponse); i {
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
		file_BlogService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiGetBlogsRequest); i {
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
		file_BlogService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blog); i {
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
		file_BlogService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiGetBlogsResponse); i {
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
		file_BlogService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlogsOfUserRequest); i {
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
		file_BlogService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlogsOfUserResponse); i {
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
		file_BlogService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserBlogsWithRelationshipsRequest); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_BlogService_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_BlogService_proto_goTypes,
		DependencyIndexes: file_BlogService_proto_depIdxs,
		EnumInfos:         file_BlogService_proto_enumTypes,
		MessageInfos:      file_BlogService_proto_msgTypes,
	}.Build()
	File_BlogService_proto = out.File
	file_BlogService_proto_rawDesc = nil
	file_BlogService_proto_goTypes = nil
	file_BlogService_proto_depIdxs = nil
}
