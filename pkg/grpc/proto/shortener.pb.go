// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: shortener.proto

package shortener

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

type URLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"` // URL
}

func (x *URLRequest) Reset() {
	*x = URLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLRequest) ProtoMessage() {}

func (x *URLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLRequest.ProtoReflect.Descriptor instead.
func (*URLRequest) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *URLRequest) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type URLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	URL  string `protobuf:"bytes,2,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *URLResponse) Reset() {
	*x = URLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLResponse) ProtoMessage() {}

func (x *URLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLResponse.ProtoReflect.Descriptor instead.
func (*URLResponse) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *URLResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *URLResponse) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type BunchLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
	Id  int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BunchLink) Reset() {
	*x = BunchLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BunchLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BunchLink) ProtoMessage() {}

func (x *BunchLink) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BunchLink.ProtoReflect.Descriptor instead.
func (*BunchLink) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *BunchLink) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *BunchLink) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type BunchSaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Links []*BunchLink `protobuf:"bytes,1,rep,name=links,proto3" json:"links,omitempty"`
}

func (x *BunchSaveRequest) Reset() {
	*x = BunchSaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BunchSaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BunchSaveRequest) ProtoMessage() {}

func (x *BunchSaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BunchSaveRequest.ProtoReflect.Descriptor instead.
func (*BunchSaveRequest) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *BunchSaveRequest) GetLinks() []*BunchLink {
	if x != nil {
		return x.Links
	}
	return nil
}

type BunchSaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Links []*BunchLink `protobuf:"bytes,1,rep,name=links,proto3" json:"links,omitempty"`
	Code  int32        `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *BunchSaveResponse) Reset() {
	*x = BunchSaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BunchSaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BunchSaveResponse) ProtoMessage() {}

func (x *BunchSaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BunchSaveResponse.ProtoReflect.Descriptor instead.
func (*BunchSaveResponse) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{4}
}

func (x *BunchSaveResponse) GetLinks() []*BunchLink {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *BunchSaveResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{5}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{6}
}

type GetUserShortsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Links []*BunchLink `protobuf:"bytes,1,rep,name=links,proto3" json:"links,omitempty"`
	Code  int32        `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetUserShortsResponse) Reset() {
	*x = GetUserShortsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserShortsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserShortsResponse) ProtoMessage() {}

func (x *GetUserShortsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserShortsResponse.ProtoReflect.Descriptor instead.
func (*GetUserShortsResponse) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserShortsResponse) GetLinks() []*BunchLink {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *GetUserShortsResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type CodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CodeResponse) Reset() {
	*x = CodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeResponse) ProtoMessage() {}

func (x *CodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeResponse.ProtoReflect.Descriptor instead.
func (*CodeResponse) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{8}
}

func (x *CodeResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []int32 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteRequest) GetId() []int32 {
	if x != nil {
		return x.Id
	}
	return nil
}

type StatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URLs  int32 `protobuf:"varint,1,opt,name=URLs,proto3" json:"URLs,omitempty"`
	Users int32 `protobuf:"varint,2,opt,name=Users,proto3" json:"Users,omitempty"`
	Code  int32 `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *StatsResponse) Reset() {
	*x = StatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsResponse) ProtoMessage() {}

func (x *StatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsResponse.ProtoReflect.Descriptor instead.
func (*StatsResponse) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{10}
}

func (x *StatsResponse) GetURLs() int32 {
	if x != nil {
		return x.URLs
	}
	return 0
}

func (x *StatsResponse) GetUsers() int32 {
	if x != nil {
		return x.Users
	}
	return 0
}

func (x *StatsResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_shortener_proto protoreflect.FileDescriptor

var file_shortener_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x22, 0x1e, 0x0a, 0x0a,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52,
	0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x22, 0x33, 0x0a, 0x0b,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52,
	0x4c, 0x22, 0x2d, 0x0a, 0x09, 0x42, 0x75, 0x6e, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x10,
	0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x3e, 0x0a, 0x10, 0x42, 0x75, 0x6e, 0x63, 0x68, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e,
	0x42, 0x75, 0x6e, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73,
	0x22, 0x53, 0x0a, 0x11, 0x42, 0x75, 0x6e, 0x63, 0x68, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x2e, 0x42, 0x75, 0x6e, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x57, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x42, 0x75, 0x6e, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e,
	0x6b, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x22, 0x0a, 0x0c, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x0d, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x55, 0x52, 0x4c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x55, 0x52, 0x4c, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0xfb, 0x03, 0x0a, 0x09, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x41, 0x50, 0x49, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x42, 0x75, 0x6e, 0x63, 0x68,
	0x53, 0x61, 0x76, 0x65, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x1b, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x6e, 0x63, 0x68, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x42, 0x75, 0x6e, 0x63, 0x68, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12,
	0x15, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x73, 0x12,
	0x10, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x20, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x2e, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x18, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x10, 0x2e, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6c, 0x79, 0x61, 0x6b, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x6f, 0x6b, 0x6f, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x75, 0x73, 0x74, 0x68, 0x61, 0x76,
	0x65, 0x2d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2d, 0x74, 0x70, 0x6c, 0x2f,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_shortener_proto_rawDescOnce sync.Once
	file_shortener_proto_rawDescData = file_shortener_proto_rawDesc
)

func file_shortener_proto_rawDescGZIP() []byte {
	file_shortener_proto_rawDescOnce.Do(func() {
		file_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortener_proto_rawDescData)
	})
	return file_shortener_proto_rawDescData
}

var file_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_shortener_proto_goTypes = []interface{}{
	(*URLRequest)(nil),            // 0: shortener.URLRequest
	(*URLResponse)(nil),           // 1: shortener.URLResponse
	(*BunchLink)(nil),             // 2: shortener.BunchLink
	(*BunchSaveRequest)(nil),      // 3: shortener.BunchSaveRequest
	(*BunchSaveResponse)(nil),     // 4: shortener.BunchSaveResponse
	(*PingRequest)(nil),           // 5: shortener.PingRequest
	(*Empty)(nil),                 // 6: shortener.Empty
	(*GetUserShortsResponse)(nil), // 7: shortener.GetUserShortsResponse
	(*CodeResponse)(nil),          // 8: shortener.CodeResponse
	(*DeleteRequest)(nil),         // 9: shortener.DeleteRequest
	(*StatsResponse)(nil),         // 10: shortener.StatsResponse
}
var file_shortener_proto_depIdxs = []int32{
	2,  // 0: shortener.BunchSaveRequest.links:type_name -> shortener.BunchLink
	2,  // 1: shortener.BunchSaveResponse.links:type_name -> shortener.BunchLink
	2,  // 2: shortener.GetUserShortsResponse.links:type_name -> shortener.BunchLink
	0,  // 3: shortener.Shortener.CreateShort:input_type -> shortener.URLRequest
	0,  // 4: shortener.Shortener.APICreateShort:input_type -> shortener.URLRequest
	3,  // 5: shortener.Shortener.BunchSaveJSON:input_type -> shortener.BunchSaveRequest
	0,  // 6: shortener.Shortener.GetShort:input_type -> shortener.URLRequest
	6,  // 7: shortener.Shortener.GetUserShorts:input_type -> shortener.Empty
	6,  // 8: shortener.Shortener.Ping:input_type -> shortener.Empty
	9,  // 9: shortener.Shortener.Delete:input_type -> shortener.DeleteRequest
	6,  // 10: shortener.Shortener.Stats:input_type -> shortener.Empty
	1,  // 11: shortener.Shortener.CreateShort:output_type -> shortener.URLResponse
	1,  // 12: shortener.Shortener.APICreateShort:output_type -> shortener.URLResponse
	4,  // 13: shortener.Shortener.BunchSaveJSON:output_type -> shortener.BunchSaveResponse
	1,  // 14: shortener.Shortener.GetShort:output_type -> shortener.URLResponse
	7,  // 15: shortener.Shortener.GetUserShorts:output_type -> shortener.GetUserShortsResponse
	8,  // 16: shortener.Shortener.Ping:output_type -> shortener.CodeResponse
	8,  // 17: shortener.Shortener.Delete:output_type -> shortener.CodeResponse
	10, // 18: shortener.Shortener.Stats:output_type -> shortener.StatsResponse
	11, // [11:19] is the sub-list for method output_type
	3,  // [3:11] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_shortener_proto_init() }
func file_shortener_proto_init() {
	if File_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URLRequest); i {
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
		file_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URLResponse); i {
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
		file_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BunchLink); i {
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
		file_shortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BunchSaveRequest); i {
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
		file_shortener_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BunchSaveResponse); i {
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
		file_shortener_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_shortener_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_shortener_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserShortsResponse); i {
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
		file_shortener_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeResponse); i {
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
		file_shortener_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_shortener_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsResponse); i {
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
			RawDescriptor: file_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shortener_proto_goTypes,
		DependencyIndexes: file_shortener_proto_depIdxs,
		MessageInfos:      file_shortener_proto_msgTypes,
	}.Build()
	File_shortener_proto = out.File
	file_shortener_proto_rawDesc = nil
	file_shortener_proto_goTypes = nil
	file_shortener_proto_depIdxs = nil
}