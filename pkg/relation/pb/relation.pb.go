// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: relation.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Relation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProfileId   uint64 `protobuf:"varint,2,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	FollowingId uint64 `protobuf:"varint,3,opt,name=following_id,json=followingId,proto3" json:"following_id,omitempty"`
}

func (x *Relation) Reset() {
	*x = Relation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Relation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relation) ProtoMessage() {}

func (x *Relation) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relation.ProtoReflect.Descriptor instead.
func (*Relation) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{0}
}

func (x *Relation) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Relation) GetProfileId() uint64 {
	if x != nil {
		return x.ProfileId
	}
	return 0
}

func (x *Relation) GetFollowingId() uint64 {
	if x != nil {
		return x.FollowingId
	}
	return 0
}

type CreatRelationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileId   uint64 `protobuf:"varint,2,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	FollowingId uint64 `protobuf:"varint,3,opt,name=following_id,json=followingId,proto3" json:"following_id,omitempty"`
}

func (x *CreatRelationRequest) Reset() {
	*x = CreatRelationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatRelationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatRelationRequest) ProtoMessage() {}

func (x *CreatRelationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatRelationRequest.ProtoReflect.Descriptor instead.
func (*CreatRelationRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{1}
}

func (x *CreatRelationRequest) GetProfileId() uint64 {
	if x != nil {
		return x.ProfileId
	}
	return 0
}

func (x *CreatRelationRequest) GetFollowingId() uint64 {
	if x != nil {
		return x.FollowingId
	}
	return 0
}

type CreatRelationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relation *Relation `protobuf:"bytes,1,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *CreatRelationResponse) Reset() {
	*x = CreatRelationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatRelationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatRelationResponse) ProtoMessage() {}

func (x *CreatRelationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatRelationResponse.ProtoReflect.Descriptor instead.
func (*CreatRelationResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{2}
}

func (x *CreatRelationResponse) GetRelation() *Relation {
	if x != nil {
		return x.Relation
	}
	return nil
}

type GetProfileFollowingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileId uint64 `protobuf:"varint,1,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
}

func (x *GetProfileFollowingRequest) Reset() {
	*x = GetProfileFollowingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileFollowingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileFollowingRequest) ProtoMessage() {}

func (x *GetProfileFollowingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileFollowingRequest.ProtoReflect.Descriptor instead.
func (*GetProfileFollowingRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{3}
}

func (x *GetProfileFollowingRequest) GetProfileId() uint64 {
	if x != nil {
		return x.ProfileId
	}
	return 0
}

type GetProfileFollowingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileId []uint64 `protobuf:"varint,1,rep,packed,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
}

func (x *GetProfileFollowingResponse) Reset() {
	*x = GetProfileFollowingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileFollowingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileFollowingResponse) ProtoMessage() {}

func (x *GetProfileFollowingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileFollowingResponse.ProtoReflect.Descriptor instead.
func (*GetProfileFollowingResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{4}
}

func (x *GetProfileFollowingResponse) GetProfileId() []uint64 {
	if x != nil {
		return x.ProfileId
	}
	return nil
}

type GetProfileFollowerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileId uint64 `protobuf:"varint,1,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
}

func (x *GetProfileFollowerRequest) Reset() {
	*x = GetProfileFollowerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileFollowerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileFollowerRequest) ProtoMessage() {}

func (x *GetProfileFollowerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileFollowerRequest.ProtoReflect.Descriptor instead.
func (*GetProfileFollowerRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{5}
}

func (x *GetProfileFollowerRequest) GetProfileId() uint64 {
	if x != nil {
		return x.ProfileId
	}
	return 0
}

type GetProfileFollowerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileId []uint64 `protobuf:"varint,1,rep,packed,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
}

func (x *GetProfileFollowerResponse) Reset() {
	*x = GetProfileFollowerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileFollowerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileFollowerResponse) ProtoMessage() {}

func (x *GetProfileFollowerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileFollowerResponse.ProtoReflect.Descriptor instead.
func (*GetProfileFollowerResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{6}
}

func (x *GetProfileFollowerResponse) GetProfileId() []uint64 {
	if x != nil {
		return x.ProfileId
	}
	return nil
}

var File_relation_proto protoreflect.FileDescriptor

var file_relation_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x5c, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x22, 0x58, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x3b, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x64, 0x32, 0x8a, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x58, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12,
	0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relation_proto_rawDescOnce sync.Once
	file_relation_proto_rawDescData = file_relation_proto_rawDesc
)

func file_relation_proto_rawDescGZIP() []byte {
	file_relation_proto_rawDescOnce.Do(func() {
		file_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_relation_proto_rawDescData)
	})
	return file_relation_proto_rawDescData
}

var file_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_relation_proto_goTypes = []interface{}{
	(*Relation)(nil),                    // 0: pb.Relation
	(*CreatRelationRequest)(nil),        // 1: pb.CreatRelationRequest
	(*CreatRelationResponse)(nil),       // 2: pb.CreatRelationResponse
	(*GetProfileFollowingRequest)(nil),  // 3: pb.GetProfileFollowingRequest
	(*GetProfileFollowingResponse)(nil), // 4: pb.GetProfileFollowingResponse
	(*GetProfileFollowerRequest)(nil),   // 5: pb.GetProfileFollowerRequest
	(*GetProfileFollowerResponse)(nil),  // 6: pb.GetProfileFollowerResponse
}
var file_relation_proto_depIdxs = []int32{
	0, // 0: pb.CreatRelationResponse.relation:type_name -> pb.Relation
	1, // 1: pb.RelationService.CreatRelation:input_type -> pb.CreatRelationRequest
	3, // 2: pb.RelationService.GetProfileFollowing:input_type -> pb.GetProfileFollowingRequest
	5, // 3: pb.RelationService.GetProfileFollower:input_type -> pb.GetProfileFollowerRequest
	2, // 4: pb.RelationService.CreatRelation:output_type -> pb.CreatRelationResponse
	4, // 5: pb.RelationService.GetProfileFollowing:output_type -> pb.GetProfileFollowingResponse
	6, // 6: pb.RelationService.GetProfileFollower:output_type -> pb.GetProfileFollowerResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_relation_proto_init() }
func file_relation_proto_init() {
	if File_relation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Relation); i {
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
		file_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatRelationRequest); i {
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
		file_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatRelationResponse); i {
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
		file_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileFollowingRequest); i {
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
		file_relation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileFollowingResponse); i {
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
		file_relation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileFollowerRequest); i {
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
		file_relation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileFollowerResponse); i {
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
			RawDescriptor: file_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relation_proto_goTypes,
		DependencyIndexes: file_relation_proto_depIdxs,
		MessageInfos:      file_relation_proto_msgTypes,
	}.Build()
	File_relation_proto = out.File
	file_relation_proto_rawDesc = nil
	file_relation_proto_goTypes = nil
	file_relation_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RelationServiceClient is the client API for RelationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RelationServiceClient interface {
	CreatRelation(ctx context.Context, in *CreatRelationRequest, opts ...grpc.CallOption) (*CreatRelationResponse, error)
	GetProfileFollowing(ctx context.Context, in *GetProfileFollowingRequest, opts ...grpc.CallOption) (*GetProfileFollowingResponse, error)
	GetProfileFollower(ctx context.Context, in *GetProfileFollowerRequest, opts ...grpc.CallOption) (*GetProfileFollowerResponse, error)
}

type relationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationServiceClient(cc grpc.ClientConnInterface) RelationServiceClient {
	return &relationServiceClient{cc}
}

func (c *relationServiceClient) CreatRelation(ctx context.Context, in *CreatRelationRequest, opts ...grpc.CallOption) (*CreatRelationResponse, error) {
	out := new(CreatRelationResponse)
	err := c.cc.Invoke(ctx, "/pb.RelationService/CreatRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetProfileFollowing(ctx context.Context, in *GetProfileFollowingRequest, opts ...grpc.CallOption) (*GetProfileFollowingResponse, error) {
	out := new(GetProfileFollowingResponse)
	err := c.cc.Invoke(ctx, "/pb.RelationService/GetProfileFollowing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetProfileFollower(ctx context.Context, in *GetProfileFollowerRequest, opts ...grpc.CallOption) (*GetProfileFollowerResponse, error) {
	out := new(GetProfileFollowerResponse)
	err := c.cc.Invoke(ctx, "/pb.RelationService/GetProfileFollower", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationServiceServer is the server API for RelationService service.
type RelationServiceServer interface {
	CreatRelation(context.Context, *CreatRelationRequest) (*CreatRelationResponse, error)
	GetProfileFollowing(context.Context, *GetProfileFollowingRequest) (*GetProfileFollowingResponse, error)
	GetProfileFollower(context.Context, *GetProfileFollowerRequest) (*GetProfileFollowerResponse, error)
}

// UnimplementedRelationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRelationServiceServer struct {
}

func (*UnimplementedRelationServiceServer) CreatRelation(context.Context, *CreatRelationRequest) (*CreatRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatRelation not implemented")
}
func (*UnimplementedRelationServiceServer) GetProfileFollowing(context.Context, *GetProfileFollowingRequest) (*GetProfileFollowingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileFollowing not implemented")
}
func (*UnimplementedRelationServiceServer) GetProfileFollower(context.Context, *GetProfileFollowerRequest) (*GetProfileFollowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileFollower not implemented")
}

func RegisterRelationServiceServer(s *grpc.Server, srv RelationServiceServer) {
	s.RegisterService(&_RelationService_serviceDesc, srv)
}

func _RelationService_CreatRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).CreatRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RelationService/CreatRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).CreatRelation(ctx, req.(*CreatRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetProfileFollowing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileFollowingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetProfileFollowing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RelationService/GetProfileFollowing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetProfileFollowing(ctx, req.(*GetProfileFollowingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetProfileFollower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetProfileFollower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RelationService/GetProfileFollower",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetProfileFollower(ctx, req.(*GetProfileFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RelationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RelationService",
	HandlerType: (*RelationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatRelation",
			Handler:    _RelationService_CreatRelation_Handler,
		},
		{
			MethodName: "GetProfileFollowing",
			Handler:    _RelationService_GetProfileFollowing_Handler,
		},
		{
			MethodName: "GetProfileFollower",
			Handler:    _RelationService_GetProfileFollower_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relation.proto",
}