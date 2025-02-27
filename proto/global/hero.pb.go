// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: hero.proto

package global

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

type C2S_DelHero struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *C2S_DelHero) Reset() {
	*x = C2S_DelHero{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_DelHero) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_DelHero) ProtoMessage() {}

func (x *C2S_DelHero) ProtoReflect() protoreflect.Message {
	mi := &file_hero_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_DelHero.ProtoReflect.Descriptor instead.
func (*C2S_DelHero) Descriptor() ([]byte, []int) {
	return file_hero_proto_rawDescGZIP(), []int{0}
}

func (x *C2S_DelHero) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type S2C_DelHero struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *S2C_DelHero) Reset() {
	*x = S2C_DelHero{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_DelHero) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_DelHero) ProtoMessage() {}

func (x *S2C_DelHero) ProtoReflect() protoreflect.Message {
	mi := &file_hero_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_DelHero.ProtoReflect.Descriptor instead.
func (*S2C_DelHero) Descriptor() ([]byte, []int) {
	return file_hero_proto_rawDescGZIP(), []int{1}
}

func (x *S2C_DelHero) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 升级
type C2S_HeroLevelup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId          int64 `protobuf:"varint,1,opt,name=HeroId,proto3" json:"HeroId,omitempty"`                   // 英雄id
	StuffItemTypeId int32 `protobuf:"varint,2,opt,name=StuffItemTypeId,proto3" json:"StuffItemTypeId,omitempty"` // 经验物品typeid
	UseNum          int32 `protobuf:"varint,3,opt,name=UseNum,proto3" json:"UseNum,omitempty"`                   // 使用物品个数
}

func (x *C2S_HeroLevelup) Reset() {
	*x = C2S_HeroLevelup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_HeroLevelup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_HeroLevelup) ProtoMessage() {}

func (x *C2S_HeroLevelup) ProtoReflect() protoreflect.Message {
	mi := &file_hero_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_HeroLevelup.ProtoReflect.Descriptor instead.
func (*C2S_HeroLevelup) Descriptor() ([]byte, []int) {
	return file_hero_proto_rawDescGZIP(), []int{2}
}

func (x *C2S_HeroLevelup) GetHeroId() int64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *C2S_HeroLevelup) GetStuffItemTypeId() int32 {
	if x != nil {
		return x.StuffItemTypeId
	}
	return 0
}

func (x *C2S_HeroLevelup) GetUseNum() int32 {
	if x != nil {
		return x.UseNum
	}
	return 0
}

type S2C_HeroLevelup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId   int64 `protobuf:"varint,1,opt,name=HeroId,proto3" json:"HeroId,omitempty"`
	CurLevel int32 `protobuf:"varint,2,opt,name=CurLevel,proto3" json:"CurLevel,omitempty"`
	CurExp   int32 `protobuf:"varint,3,opt,name=CurExp,proto3" json:"CurExp,omitempty"`
}

func (x *S2C_HeroLevelup) Reset() {
	*x = S2C_HeroLevelup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_HeroLevelup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_HeroLevelup) ProtoMessage() {}

func (x *S2C_HeroLevelup) ProtoReflect() protoreflect.Message {
	mi := &file_hero_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_HeroLevelup.ProtoReflect.Descriptor instead.
func (*S2C_HeroLevelup) Descriptor() ([]byte, []int) {
	return file_hero_proto_rawDescGZIP(), []int{3}
}

func (x *S2C_HeroLevelup) GetHeroId() int64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *S2C_HeroLevelup) GetCurLevel() int32 {
	if x != nil {
		return x.CurLevel
	}
	return 0
}

func (x *S2C_HeroLevelup) GetCurExp() int32 {
	if x != nil {
		return x.CurExp
	}
	return 0
}

// 突破
type C2S_HeroPromote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId int64 `protobuf:"varint,1,opt,name=HeroId,proto3" json:"HeroId,omitempty"`
}

func (x *C2S_HeroPromote) Reset() {
	*x = C2S_HeroPromote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_HeroPromote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_HeroPromote) ProtoMessage() {}

func (x *C2S_HeroPromote) ProtoReflect() protoreflect.Message {
	mi := &file_hero_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_HeroPromote.ProtoReflect.Descriptor instead.
func (*C2S_HeroPromote) Descriptor() ([]byte, []int) {
	return file_hero_proto_rawDescGZIP(), []int{4}
}

func (x *C2S_HeroPromote) GetHeroId() int64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

type S2C_HeroPromote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId          int64 `protobuf:"varint,1,opt,name=HeroId,proto3" json:"HeroId,omitempty"`
	CurPromoteLevel int32 `protobuf:"varint,2,opt,name=CurPromoteLevel,proto3" json:"CurPromoteLevel,omitempty"`
}

func (x *S2C_HeroPromote) Reset() {
	*x = S2C_HeroPromote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_HeroPromote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_HeroPromote) ProtoMessage() {}

func (x *S2C_HeroPromote) ProtoReflect() protoreflect.Message {
	mi := &file_hero_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_HeroPromote.ProtoReflect.Descriptor instead.
func (*S2C_HeroPromote) Descriptor() ([]byte, []int) {
	return file_hero_proto_rawDescGZIP(), []int{5}
}

func (x *S2C_HeroPromote) GetHeroId() int64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *S2C_HeroPromote) GetCurPromoteLevel() int32 {
	if x != nil {
		return x.CurPromoteLevel
	}
	return 0
}

// 升星
type C2S_HeroStarup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId int64 `protobuf:"varint,1,opt,name=HeroId,proto3" json:"HeroId,omitempty"`
}

func (x *C2S_HeroStarup) Reset() {
	*x = C2S_HeroStarup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_HeroStarup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_HeroStarup) ProtoMessage() {}

func (x *C2S_HeroStarup) ProtoReflect() protoreflect.Message {
	mi := &file_hero_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_HeroStarup.ProtoReflect.Descriptor instead.
func (*C2S_HeroStarup) Descriptor() ([]byte, []int) {
	return file_hero_proto_rawDescGZIP(), []int{6}
}

func (x *C2S_HeroStarup) GetHeroId() int64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

// 选择天赋
type C2S_HeroTalentChoose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId   int64 `protobuf:"varint,1,opt,name=HeroId,proto3" json:"HeroId,omitempty"`     // 英雄id
	TalentId int32 `protobuf:"varint,2,opt,name=TalentId,proto3" json:"TalentId,omitempty"` // 天赋id
}

func (x *C2S_HeroTalentChoose) Reset() {
	*x = C2S_HeroTalentChoose{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_HeroTalentChoose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_HeroTalentChoose) ProtoMessage() {}

func (x *C2S_HeroTalentChoose) ProtoReflect() protoreflect.Message {
	mi := &file_hero_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_HeroTalentChoose.ProtoReflect.Descriptor instead.
func (*C2S_HeroTalentChoose) Descriptor() ([]byte, []int) {
	return file_hero_proto_rawDescGZIP(), []int{7}
}

func (x *C2S_HeroTalentChoose) GetHeroId() int64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *C2S_HeroTalentChoose) GetTalentId() int32 {
	if x != nil {
		return x.TalentId
	}
	return 0
}

type S2C_HeroTalentChoose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId     int64     `protobuf:"varint,1,opt,name=HeroId,proto3" json:"HeroId,omitempty"`        // 英雄id
	TalentList []*Talent `protobuf:"bytes,2,rep,name=TalentList,proto3" json:"TalentList,omitempty"` // 天赋id
}

func (x *S2C_HeroTalentChoose) Reset() {
	*x = S2C_HeroTalentChoose{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_HeroTalentChoose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_HeroTalentChoose) ProtoMessage() {}

func (x *S2C_HeroTalentChoose) ProtoReflect() protoreflect.Message {
	mi := &file_hero_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_HeroTalentChoose.ProtoReflect.Descriptor instead.
func (*S2C_HeroTalentChoose) Descriptor() ([]byte, []int) {
	return file_hero_proto_rawDescGZIP(), []int{8}
}

func (x *S2C_HeroTalentChoose) GetHeroId() int64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *S2C_HeroTalentChoose) GetTalentList() []*Talent {
	if x != nil {
		return x.TalentList
	}
	return nil
}

type S2C_HeroInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Hero `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *S2C_HeroInfo) Reset() {
	*x = S2C_HeroInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_HeroInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_HeroInfo) ProtoMessage() {}

func (x *S2C_HeroInfo) ProtoReflect() protoreflect.Message {
	mi := &file_hero_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_HeroInfo.ProtoReflect.Descriptor instead.
func (*S2C_HeroInfo) Descriptor() ([]byte, []int) {
	return file_hero_proto_rawDescGZIP(), []int{9}
}

func (x *S2C_HeroInfo) GetInfo() *Hero {
	if x != nil {
		return x.Info
	}
	return nil
}

////////////////////////////////////////////////
// Att
type S2C_HeroAtts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId   int64   `protobuf:"varint,1,opt,name=HeroId,proto3" json:"HeroId,omitempty"`
	AttValue []int32 `protobuf:"varint,2,rep,packed,name=AttValue,proto3" json:"AttValue,omitempty"`
}

func (x *S2C_HeroAtts) Reset() {
	*x = S2C_HeroAtts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_HeroAtts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_HeroAtts) ProtoMessage() {}

func (x *S2C_HeroAtts) ProtoReflect() protoreflect.Message {
	mi := &file_hero_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_HeroAtts.ProtoReflect.Descriptor instead.
func (*S2C_HeroAtts) Descriptor() ([]byte, []int) {
	return file_hero_proto_rawDescGZIP(), []int{10}
}

func (x *S2C_HeroAtts) GetHeroId() int64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *S2C_HeroAtts) GetAttValue() []int32 {
	if x != nil {
		return x.AttValue
	}
	return nil
}

type S2C_HeroAttUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId int64  `protobuf:"varint,1,opt,name=HeroId,proto3" json:"HeroId,omitempty"`
	Atts   []*Att `protobuf:"bytes,2,rep,name=Atts,proto3" json:"Atts,omitempty"` // 更新的属性
}

func (x *S2C_HeroAttUpdate) Reset() {
	*x = S2C_HeroAttUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_HeroAttUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_HeroAttUpdate) ProtoMessage() {}

func (x *S2C_HeroAttUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_hero_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_HeroAttUpdate.ProtoReflect.Descriptor instead.
func (*S2C_HeroAttUpdate) Descriptor() ([]byte, []int) {
	return file_hero_proto_rawDescGZIP(), []int{11}
}

func (x *S2C_HeroAttUpdate) GetHeroId() int64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *S2C_HeroAttUpdate) GetAtts() []*Att {
	if x != nil {
		return x.Atts
	}
	return nil
}

var File_hero_proto protoreflect.FileDescriptor

var file_hero_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x09, 0x61, 0x74, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x0b,
	0x43, 0x32, 0x53, 0x5f, 0x44, 0x65, 0x6c, 0x48, 0x65, 0x72, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x1d, 0x0a, 0x0b, 0x53,
	0x32, 0x43, 0x5f, 0x44, 0x65, 0x6c, 0x48, 0x65, 0x72, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x6b, 0x0a, 0x0f, 0x43, 0x32,
	0x53, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x75, 0x70, 0x12, 0x16, 0x0a,
	0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48,
	0x65, 0x72, 0x6f, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x74, 0x75, 0x66, 0x66, 0x49, 0x74,
	0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x53, 0x74, 0x75, 0x66, 0x66, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x22, 0x5d, 0x0a, 0x0f, 0x53, 0x32, 0x43, 0x5f, 0x48,
	0x65, 0x72, 0x6f, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65,
	0x72, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48, 0x65, 0x72, 0x6f,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x43, 0x75, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x43, 0x75, 0x72, 0x45, 0x78, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x43, 0x75, 0x72, 0x45, 0x78, 0x70, 0x22, 0x29, 0x0a, 0x0f, 0x43, 0x32, 0x53, 0x5f, 0x48, 0x65,
	0x72, 0x6f, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x72,
	0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49,
	0x64, 0x22, 0x53, 0x0a, 0x0f, 0x53, 0x32, 0x43, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f,
	0x43, 0x75, 0x72, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x43, 0x75, 0x72, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x28, 0x0a, 0x0e, 0x43, 0x32, 0x53, 0x5f, 0x48, 0x65,
	0x72, 0x6f, 0x53, 0x74, 0x61, 0x72, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x72, 0x6f,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64,
	0x22, 0x4a, 0x0a, 0x14, 0x43, 0x32, 0x53, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x54, 0x61, 0x6c, 0x65,
	0x6e, 0x74, 0x43, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x72, 0x6f,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x14,
	0x53, 0x32, 0x43, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x43, 0x68,
	0x6f, 0x6f, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x0a,
	0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x52,
	0x0a, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2f, 0x0a, 0x0c, 0x53,
	0x32, 0x43, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x42, 0x0a, 0x0c,
	0x53, 0x32, 0x43, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x41, 0x74, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48, 0x65,
	0x72, 0x6f, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x74, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x41, 0x74, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x4b, 0x0a, 0x11, 0x53, 0x32, 0x43, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x41, 0x74, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x04, 0x41, 0x74, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x52, 0x04, 0x41, 0x74, 0x74, 0x73, 0x42, 0x32, 0x5a,
	0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x61, 0x73, 0x74,
	0x2d, 0x65, 0x64, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0xaa, 0x02, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hero_proto_rawDescOnce sync.Once
	file_hero_proto_rawDescData = file_hero_proto_rawDesc
)

func file_hero_proto_rawDescGZIP() []byte {
	file_hero_proto_rawDescOnce.Do(func() {
		file_hero_proto_rawDescData = protoimpl.X.CompressGZIP(file_hero_proto_rawDescData)
	})
	return file_hero_proto_rawDescData
}

var file_hero_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_hero_proto_goTypes = []interface{}{
	(*C2S_DelHero)(nil),          // 0: proto.C2S_DelHero
	(*S2C_DelHero)(nil),          // 1: proto.S2C_DelHero
	(*C2S_HeroLevelup)(nil),      // 2: proto.C2S_HeroLevelup
	(*S2C_HeroLevelup)(nil),      // 3: proto.S2C_HeroLevelup
	(*C2S_HeroPromote)(nil),      // 4: proto.C2S_HeroPromote
	(*S2C_HeroPromote)(nil),      // 5: proto.S2C_HeroPromote
	(*C2S_HeroStarup)(nil),       // 6: proto.C2S_HeroStarup
	(*C2S_HeroTalentChoose)(nil), // 7: proto.C2S_HeroTalentChoose
	(*S2C_HeroTalentChoose)(nil), // 8: proto.S2C_HeroTalentChoose
	(*S2C_HeroInfo)(nil),         // 9: proto.S2C_HeroInfo
	(*S2C_HeroAtts)(nil),         // 10: proto.S2C_HeroAtts
	(*S2C_HeroAttUpdate)(nil),    // 11: proto.S2C_HeroAttUpdate
	(*Talent)(nil),               // 12: proto.Talent
	(*Hero)(nil),                 // 13: proto.Hero
	(*Att)(nil),                  // 14: proto.Att
}
var file_hero_proto_depIdxs = []int32{
	12, // 0: proto.S2C_HeroTalentChoose.TalentList:type_name -> proto.Talent
	13, // 1: proto.S2C_HeroInfo.info:type_name -> proto.Hero
	14, // 2: proto.S2C_HeroAttUpdate.Atts:type_name -> proto.Att
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_hero_proto_init() }
func file_hero_proto_init() {
	if File_hero_proto != nil {
		return
	}
	file_define_proto_init()
	file_att_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_hero_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_DelHero); i {
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
		file_hero_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_DelHero); i {
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
		file_hero_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_HeroLevelup); i {
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
		file_hero_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_HeroLevelup); i {
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
		file_hero_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_HeroPromote); i {
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
		file_hero_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_HeroPromote); i {
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
		file_hero_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_HeroStarup); i {
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
		file_hero_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_HeroTalentChoose); i {
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
		file_hero_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_HeroTalentChoose); i {
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
		file_hero_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_HeroInfo); i {
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
		file_hero_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_HeroAtts); i {
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
		file_hero_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_HeroAttUpdate); i {
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
			RawDescriptor: file_hero_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hero_proto_goTypes,
		DependencyIndexes: file_hero_proto_depIdxs,
		MessageInfos:      file_hero_proto_msgTypes,
	}.Build()
	File_hero_proto = out.File
	file_hero_proto_rawDesc = nil
	file_hero_proto_goTypes = nil
	file_hero_proto_depIdxs = nil
}
