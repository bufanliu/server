// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: server/rank/rank.proto

package rank

import (
	global "github.com/east-eden/server/proto/global"
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

// 通过id查询排行
type QueryRankByObjIdRq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankId int32 `protobuf:"varint,1,opt,name=RankId,proto3" json:"RankId,omitempty"`
	ObjId  int64 `protobuf:"varint,2,opt,name=ObjId,proto3" json:"ObjId,omitempty"`
}

func (x *QueryRankByObjIdRq) Reset() {
	*x = QueryRankByObjIdRq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_rank_rank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRankByObjIdRq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRankByObjIdRq) ProtoMessage() {}

func (x *QueryRankByObjIdRq) ProtoReflect() protoreflect.Message {
	mi := &file_server_rank_rank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRankByObjIdRq.ProtoReflect.Descriptor instead.
func (*QueryRankByObjIdRq) Descriptor() ([]byte, []int) {
	return file_server_rank_rank_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRankByObjIdRq) GetRankId() int32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

func (x *QueryRankByObjIdRq) GetObjId() int64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

type QueryRankByObjIdRs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankId    int32                `protobuf:"varint,1,opt,name=RankId,proto3" json:"RankId,omitempty"`
	ObjId     int64                `protobuf:"varint,2,opt,name=ObjId,proto3" json:"ObjId,omitempty"`
	RankIndex int32                `protobuf:"varint,3,opt,name=RankIndex,proto3" json:"RankIndex,omitempty"` // 排行榜中位置：从0开始
	Metadata  *global.RankMetadata `protobuf:"bytes,4,opt,name=Metadata,proto3" json:"Metadata,omitempty"`
}

func (x *QueryRankByObjIdRs) Reset() {
	*x = QueryRankByObjIdRs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_rank_rank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRankByObjIdRs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRankByObjIdRs) ProtoMessage() {}

func (x *QueryRankByObjIdRs) ProtoReflect() protoreflect.Message {
	mi := &file_server_rank_rank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRankByObjIdRs.ProtoReflect.Descriptor instead.
func (*QueryRankByObjIdRs) Descriptor() ([]byte, []int) {
	return file_server_rank_rank_proto_rawDescGZIP(), []int{1}
}

func (x *QueryRankByObjIdRs) GetRankId() int32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

func (x *QueryRankByObjIdRs) GetObjId() int64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

func (x *QueryRankByObjIdRs) GetRankIndex() int32 {
	if x != nil {
		return x.RankIndex
	}
	return 0
}

func (x *QueryRankByObjIdRs) GetMetadata() *global.RankMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type QueryRankByRangeRq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankId int32 `protobuf:"varint,1,opt,name=RankId,proto3" json:"RankId,omitempty"`
	Start  int64 `protobuf:"varint,2,opt,name=Start,proto3" json:"Start,omitempty"`
	End    int64 `protobuf:"varint,3,opt,name=End,proto3" json:"End,omitempty"` // End == -1时代表查询所有数据
}

func (x *QueryRankByRangeRq) Reset() {
	*x = QueryRankByRangeRq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_rank_rank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRankByRangeRq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRankByRangeRq) ProtoMessage() {}

func (x *QueryRankByRangeRq) ProtoReflect() protoreflect.Message {
	mi := &file_server_rank_rank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRankByRangeRq.ProtoReflect.Descriptor instead.
func (*QueryRankByRangeRq) Descriptor() ([]byte, []int) {
	return file_server_rank_rank_proto_rawDescGZIP(), []int{2}
}

func (x *QueryRankByRangeRq) GetRankId() int32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

func (x *QueryRankByRangeRq) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *QueryRankByRangeRq) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

type QueryRankByRangeRs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankId    int32                  `protobuf:"varint,1,opt,name=RankId,proto3" json:"RankId,omitempty"`
	Start     int64                  `protobuf:"varint,2,opt,name=Start,proto3" json:"Start,omitempty"`
	End       int64                  `protobuf:"varint,3,opt,name=End,proto3" json:"End,omitempty"`
	Metadatas []*global.RankMetadata `protobuf:"bytes,4,rep,name=Metadatas,proto3" json:"Metadatas,omitempty"` // 按照start到end顺序
}

func (x *QueryRankByRangeRs) Reset() {
	*x = QueryRankByRangeRs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_rank_rank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRankByRangeRs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRankByRangeRs) ProtoMessage() {}

func (x *QueryRankByRangeRs) ProtoReflect() protoreflect.Message {
	mi := &file_server_rank_rank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRankByRangeRs.ProtoReflect.Descriptor instead.
func (*QueryRankByRangeRs) Descriptor() ([]byte, []int) {
	return file_server_rank_rank_proto_rawDescGZIP(), []int{3}
}

func (x *QueryRankByRangeRs) GetRankId() int32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

func (x *QueryRankByRangeRs) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *QueryRankByRangeRs) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *QueryRankByRangeRs) GetMetadatas() []*global.RankMetadata {
	if x != nil {
		return x.Metadatas
	}
	return nil
}

// 设置排行榜分数
type SetRankScoreRq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankId   int32                `protobuf:"varint,1,opt,name=RankId,proto3" json:"RankId,omitempty"`    // 排行榜id
	Metadata *global.RankMetadata `protobuf:"bytes,2,opt,name=Metadata,proto3" json:"Metadata,omitempty"` // 元数据
}

func (x *SetRankScoreRq) Reset() {
	*x = SetRankScoreRq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_rank_rank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRankScoreRq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRankScoreRq) ProtoMessage() {}

func (x *SetRankScoreRq) ProtoReflect() protoreflect.Message {
	mi := &file_server_rank_rank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRankScoreRq.ProtoReflect.Descriptor instead.
func (*SetRankScoreRq) Descriptor() ([]byte, []int) {
	return file_server_rank_rank_proto_rawDescGZIP(), []int{4}
}

func (x *SetRankScoreRq) GetRankId() int32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

func (x *SetRankScoreRq) GetMetadata() *global.RankMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type SetRankScoreRs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetRankScoreRs) Reset() {
	*x = SetRankScoreRs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_rank_rank_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRankScoreRs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRankScoreRs) ProtoMessage() {}

func (x *SetRankScoreRs) ProtoReflect() protoreflect.Message {
	mi := &file_server_rank_rank_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRankScoreRs.ProtoReflect.Descriptor instead.
func (*SetRankScoreRs) Descriptor() ([]byte, []int) {
	return file_server_rank_rank_proto_rawDescGZIP(), []int{5}
}

// 踢掉其他节点排行榜缓存
type KickRankDataRq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankId     int32 `protobuf:"varint,1,opt,name=RankId,proto3" json:"RankId,omitempty"`         // 排行榜id
	RankNodeId int32 `protobuf:"varint,2,opt,name=RankNodeId,proto3" json:"RankNodeId,omitempty"` // 所在rank服务节点id
}

func (x *KickRankDataRq) Reset() {
	*x = KickRankDataRq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_rank_rank_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickRankDataRq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickRankDataRq) ProtoMessage() {}

func (x *KickRankDataRq) ProtoReflect() protoreflect.Message {
	mi := &file_server_rank_rank_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickRankDataRq.ProtoReflect.Descriptor instead.
func (*KickRankDataRq) Descriptor() ([]byte, []int) {
	return file_server_rank_rank_proto_rawDescGZIP(), []int{6}
}

func (x *KickRankDataRq) GetRankId() int32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

func (x *KickRankDataRq) GetRankNodeId() int32 {
	if x != nil {
		return x.RankNodeId
	}
	return 0
}

type KickRankDataRs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankId int32  `protobuf:"varint,1,opt,name=RankId,proto3" json:"RankId,omitempty"` // 排行榜id
	Error  string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *KickRankDataRs) Reset() {
	*x = KickRankDataRs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_rank_rank_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickRankDataRs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickRankDataRs) ProtoMessage() {}

func (x *KickRankDataRs) ProtoReflect() protoreflect.Message {
	mi := &file_server_rank_rank_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickRankDataRs.ProtoReflect.Descriptor instead.
func (*KickRankDataRs) Descriptor() ([]byte, []int) {
	return file_server_rank_rank_proto_rawDescGZIP(), []int{7}
}

func (x *KickRankDataRs) GetRankId() int32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

func (x *KickRankDataRs) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_server_rank_rank_proto protoreflect.FileDescriptor

var file_server_rank_rank_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x2f, 0x72, 0x61,
	0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x1a, 0x13,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x6b,
	0x42, 0x79, 0x4f, 0x62, 0x6a, 0x49, 0x64, 0x52, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x6e,
	0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x61, 0x6e, 0x6b, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x62, 0x6a, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x4f, 0x62, 0x6a, 0x49, 0x64, 0x22, 0x91, 0x01, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x61, 0x6e, 0x6b, 0x42, 0x79, 0x4f, 0x62, 0x6a, 0x49, 0x64, 0x52, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x52, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x62, 0x6a, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4f, 0x62, 0x6a, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x52, 0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2f, 0x0a, 0x08, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x54, 0x0a, 0x12, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x42, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x45, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x45, 0x6e,
	0x64, 0x22, 0x87, 0x01, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x42,
	0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x6e, 0x6b,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x6e, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x45, 0x6e, 0x64, 0x12, 0x31, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x73, 0x22, 0x59, 0x0a, 0x0e, 0x53,
	0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52,
	0x61, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x61, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x52, 0x61, 0x6e,
	0x6b, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x73, 0x22, 0x48, 0x0a, 0x0e, 0x4b, 0x69, 0x63, 0x6b,
	0x52, 0x61, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61,
	0x6e, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x61, 0x6e, 0x6b,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x61, 0x6e, 0x6b, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x52, 0x61, 0x6e, 0x6b, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x22, 0x3e, 0x0a, 0x0e, 0x4b, 0x69, 0x63, 0x6b, 0x52, 0x61, 0x6e, 0x6b, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x32, 0x9d, 0x02, 0x0a, 0x0b, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x48, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x42,
	0x79, 0x4f, 0x62, 0x6a, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x42, 0x79, 0x4f, 0x62, 0x6a, 0x49, 0x64, 0x52, 0x71,
	0x1a, 0x18, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x61, 0x6e,
	0x6b, 0x42, 0x79, 0x4f, 0x62, 0x6a, 0x49, 0x64, 0x52, 0x73, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x10,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x42, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x18, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x61, 0x6e,
	0x6b, 0x42, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x71, 0x1a, 0x18, 0x2e, 0x72, 0x61, 0x6e,
	0x6b, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x42, 0x79, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x73, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x52, 0x61, 0x6e,
	0x6b, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x53, 0x65,
	0x74, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x71, 0x1a, 0x14, 0x2e, 0x72,
	0x61, 0x6e, 0x6b, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x52, 0x73, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0c, 0x4b, 0x69, 0x63, 0x6b, 0x52, 0x61, 0x6e, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x4b, 0x69, 0x63, 0x6b,
	0x52, 0x61, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x71, 0x1a, 0x14, 0x2e, 0x72, 0x61, 0x6e,
	0x6b, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x52, 0x61, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x73,
	0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x61, 0x73, 0x74, 0x2d, 0x65, 0x64, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x72,
	0x61, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_rank_rank_proto_rawDescOnce sync.Once
	file_server_rank_rank_proto_rawDescData = file_server_rank_rank_proto_rawDesc
)

func file_server_rank_rank_proto_rawDescGZIP() []byte {
	file_server_rank_rank_proto_rawDescOnce.Do(func() {
		file_server_rank_rank_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_rank_rank_proto_rawDescData)
	})
	return file_server_rank_rank_proto_rawDescData
}

var file_server_rank_rank_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_server_rank_rank_proto_goTypes = []interface{}{
	(*QueryRankByObjIdRq)(nil),  // 0: rank.QueryRankByObjIdRq
	(*QueryRankByObjIdRs)(nil),  // 1: rank.QueryRankByObjIdRs
	(*QueryRankByRangeRq)(nil),  // 2: rank.QueryRankByRangeRq
	(*QueryRankByRangeRs)(nil),  // 3: rank.QueryRankByRangeRs
	(*SetRankScoreRq)(nil),      // 4: rank.SetRankScoreRq
	(*SetRankScoreRs)(nil),      // 5: rank.SetRankScoreRs
	(*KickRankDataRq)(nil),      // 6: rank.KickRankDataRq
	(*KickRankDataRs)(nil),      // 7: rank.KickRankDataRs
	(*global.RankMetadata)(nil), // 8: proto.RankMetadata
}
var file_server_rank_rank_proto_depIdxs = []int32{
	8, // 0: rank.QueryRankByObjIdRs.Metadata:type_name -> proto.RankMetadata
	8, // 1: rank.QueryRankByRangeRs.Metadatas:type_name -> proto.RankMetadata
	8, // 2: rank.SetRankScoreRq.Metadata:type_name -> proto.RankMetadata
	0, // 3: rank.RankService.QueryRankByObjId:input_type -> rank.QueryRankByObjIdRq
	2, // 4: rank.RankService.QueryRankByRange:input_type -> rank.QueryRankByRangeRq
	4, // 5: rank.RankService.SetRankScore:input_type -> rank.SetRankScoreRq
	6, // 6: rank.RankService.KickRankData:input_type -> rank.KickRankDataRq
	1, // 7: rank.RankService.QueryRankByObjId:output_type -> rank.QueryRankByObjIdRs
	3, // 8: rank.RankService.QueryRankByRange:output_type -> rank.QueryRankByRangeRs
	5, // 9: rank.RankService.SetRankScore:output_type -> rank.SetRankScoreRs
	7, // 10: rank.RankService.KickRankData:output_type -> rank.KickRankDataRs
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_server_rank_rank_proto_init() }
func file_server_rank_rank_proto_init() {
	if File_server_rank_rank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_rank_rank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRankByObjIdRq); i {
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
		file_server_rank_rank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRankByObjIdRs); i {
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
		file_server_rank_rank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRankByRangeRq); i {
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
		file_server_rank_rank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRankByRangeRs); i {
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
		file_server_rank_rank_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRankScoreRq); i {
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
		file_server_rank_rank_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRankScoreRs); i {
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
		file_server_rank_rank_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickRankDataRq); i {
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
		file_server_rank_rank_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickRankDataRs); i {
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
			RawDescriptor: file_server_rank_rank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_rank_rank_proto_goTypes,
		DependencyIndexes: file_server_rank_rank_proto_depIdxs,
		MessageInfos:      file_server_rank_rank_proto_msgTypes,
	}.Build()
	File_server_rank_rank_proto = out.File
	file_server_rank_rank_proto_rawDesc = nil
	file_server_rank_rank_proto_goTypes = nil
	file_server_rank_rank_proto_depIdxs = nil
}
