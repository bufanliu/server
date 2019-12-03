// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client/client.proto

package client

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ClientInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientInfo) Reset()         { *m = ClientInfo{} }
func (m *ClientInfo) String() string { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()    {}
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{0}
}

func (m *ClientInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientInfo.Unmarshal(m, b)
}
func (m *ClientInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientInfo.Marshal(b, m, deterministic)
}
func (m *ClientInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientInfo.Merge(m, src)
}
func (m *ClientInfo) XXX_Size() int {
	return xxx_messageInfo_ClientInfo.Size(m)
}
func (m *ClientInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClientInfo proto.InternalMessageInfo

func (m *ClientInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ClientInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MC_ClientLogon struct {
	ClientId             int64    `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientName           string   `protobuf:"bytes,2,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MC_ClientLogon) Reset()         { *m = MC_ClientLogon{} }
func (m *MC_ClientLogon) String() string { return proto.CompactTextString(m) }
func (*MC_ClientLogon) ProtoMessage()    {}
func (*MC_ClientLogon) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{1}
}

func (m *MC_ClientLogon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MC_ClientLogon.Unmarshal(m, b)
}
func (m *MC_ClientLogon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MC_ClientLogon.Marshal(b, m, deterministic)
}
func (m *MC_ClientLogon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MC_ClientLogon.Merge(m, src)
}
func (m *MC_ClientLogon) XXX_Size() int {
	return xxx_messageInfo_MC_ClientLogon.Size(m)
}
func (m *MC_ClientLogon) XXX_DiscardUnknown() {
	xxx_messageInfo_MC_ClientLogon.DiscardUnknown(m)
}

var xxx_messageInfo_MC_ClientLogon proto.InternalMessageInfo

func (m *MC_ClientLogon) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *MC_ClientLogon) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

type MS_ClientLogon struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MS_ClientLogon) Reset()         { *m = MS_ClientLogon{} }
func (m *MS_ClientLogon) String() string { return proto.CompactTextString(m) }
func (*MS_ClientLogon) ProtoMessage()    {}
func (*MS_ClientLogon) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{2}
}

func (m *MS_ClientLogon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MS_ClientLogon.Unmarshal(m, b)
}
func (m *MS_ClientLogon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MS_ClientLogon.Marshal(b, m, deterministic)
}
func (m *MS_ClientLogon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MS_ClientLogon.Merge(m, src)
}
func (m *MS_ClientLogon) XXX_Size() int {
	return xxx_messageInfo_MS_ClientLogon.Size(m)
}
func (m *MS_ClientLogon) XXX_DiscardUnknown() {
	xxx_messageInfo_MS_ClientLogon.DiscardUnknown(m)
}

var xxx_messageInfo_MS_ClientLogon proto.InternalMessageInfo

type MC_HeartBeat struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MC_HeartBeat) Reset()         { *m = MC_HeartBeat{} }
func (m *MC_HeartBeat) String() string { return proto.CompactTextString(m) }
func (*MC_HeartBeat) ProtoMessage()    {}
func (*MC_HeartBeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{3}
}

func (m *MC_HeartBeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MC_HeartBeat.Unmarshal(m, b)
}
func (m *MC_HeartBeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MC_HeartBeat.Marshal(b, m, deterministic)
}
func (m *MC_HeartBeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MC_HeartBeat.Merge(m, src)
}
func (m *MC_HeartBeat) XXX_Size() int {
	return xxx_messageInfo_MC_HeartBeat.Size(m)
}
func (m *MC_HeartBeat) XXX_DiscardUnknown() {
	xxx_messageInfo_MC_HeartBeat.DiscardUnknown(m)
}

var xxx_messageInfo_MC_HeartBeat proto.InternalMessageInfo

type MS_HeartBeat struct {
	Timestamp            uint32   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MS_HeartBeat) Reset()         { *m = MS_HeartBeat{} }
func (m *MS_HeartBeat) String() string { return proto.CompactTextString(m) }
func (*MS_HeartBeat) ProtoMessage()    {}
func (*MS_HeartBeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{4}
}

func (m *MS_HeartBeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MS_HeartBeat.Unmarshal(m, b)
}
func (m *MS_HeartBeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MS_HeartBeat.Marshal(b, m, deterministic)
}
func (m *MS_HeartBeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MS_HeartBeat.Merge(m, src)
}
func (m *MS_HeartBeat) XXX_Size() int {
	return xxx_messageInfo_MS_HeartBeat.Size(m)
}
func (m *MS_HeartBeat) XXX_DiscardUnknown() {
	xxx_messageInfo_MS_HeartBeat.DiscardUnknown(m)
}

var xxx_messageInfo_MS_HeartBeat proto.InternalMessageInfo

func (m *MS_HeartBeat) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type MC_ClientConnected struct {
	ClientId             int64    `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MC_ClientConnected) Reset()         { *m = MC_ClientConnected{} }
func (m *MC_ClientConnected) String() string { return proto.CompactTextString(m) }
func (*MC_ClientConnected) ProtoMessage()    {}
func (*MC_ClientConnected) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{5}
}

func (m *MC_ClientConnected) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MC_ClientConnected.Unmarshal(m, b)
}
func (m *MC_ClientConnected) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MC_ClientConnected.Marshal(b, m, deterministic)
}
func (m *MC_ClientConnected) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MC_ClientConnected.Merge(m, src)
}
func (m *MC_ClientConnected) XXX_Size() int {
	return xxx_messageInfo_MC_ClientConnected.Size(m)
}
func (m *MC_ClientConnected) XXX_DiscardUnknown() {
	xxx_messageInfo_MC_ClientConnected.DiscardUnknown(m)
}

var xxx_messageInfo_MC_ClientConnected proto.InternalMessageInfo

func (m *MC_ClientConnected) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *MC_ClientConnected) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MC_ClientDisconnect struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MC_ClientDisconnect) Reset()         { *m = MC_ClientDisconnect{} }
func (m *MC_ClientDisconnect) String() string { return proto.CompactTextString(m) }
func (*MC_ClientDisconnect) ProtoMessage()    {}
func (*MC_ClientDisconnect) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{6}
}

func (m *MC_ClientDisconnect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MC_ClientDisconnect.Unmarshal(m, b)
}
func (m *MC_ClientDisconnect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MC_ClientDisconnect.Marshal(b, m, deterministic)
}
func (m *MC_ClientDisconnect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MC_ClientDisconnect.Merge(m, src)
}
func (m *MC_ClientDisconnect) XXX_Size() int {
	return xxx_messageInfo_MC_ClientDisconnect.Size(m)
}
func (m *MC_ClientDisconnect) XXX_DiscardUnknown() {
	xxx_messageInfo_MC_ClientDisconnect.DiscardUnknown(m)
}

var xxx_messageInfo_MC_ClientDisconnect proto.InternalMessageInfo

////////////////////////////////////////////////
// player
type Hero struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TypeId               int32    `protobuf:"varint,2,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hero) Reset()         { *m = Hero{} }
func (m *Hero) String() string { return proto.CompactTextString(m) }
func (*Hero) ProtoMessage()    {}
func (*Hero) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{7}
}

func (m *Hero) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hero.Unmarshal(m, b)
}
func (m *Hero) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hero.Marshal(b, m, deterministic)
}
func (m *Hero) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hero.Merge(m, src)
}
func (m *Hero) XXX_Size() int {
	return xxx_messageInfo_Hero.Size(m)
}
func (m *Hero) XXX_DiscardUnknown() {
	xxx_messageInfo_Hero.DiscardUnknown(m)
}

var xxx_messageInfo_Hero proto.InternalMessageInfo

func (m *Hero) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Hero) GetTypeId() int32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

type MC_AddHero struct {
	TypeId               int32    `protobuf:"varint,1,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MC_AddHero) Reset()         { *m = MC_AddHero{} }
func (m *MC_AddHero) String() string { return proto.CompactTextString(m) }
func (*MC_AddHero) ProtoMessage()    {}
func (*MC_AddHero) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{8}
}

func (m *MC_AddHero) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MC_AddHero.Unmarshal(m, b)
}
func (m *MC_AddHero) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MC_AddHero.Marshal(b, m, deterministic)
}
func (m *MC_AddHero) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MC_AddHero.Merge(m, src)
}
func (m *MC_AddHero) XXX_Size() int {
	return xxx_messageInfo_MC_AddHero.Size(m)
}
func (m *MC_AddHero) XXX_DiscardUnknown() {
	xxx_messageInfo_MC_AddHero.DiscardUnknown(m)
}

var xxx_messageInfo_MC_AddHero proto.InternalMessageInfo

func (m *MC_AddHero) GetTypeId() int32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

type MS_HeroList struct {
	Heros                []*Hero  `protobuf:"bytes,1,rep,name=heros,proto3" json:"heros,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MS_HeroList) Reset()         { *m = MS_HeroList{} }
func (m *MS_HeroList) String() string { return proto.CompactTextString(m) }
func (*MS_HeroList) ProtoMessage()    {}
func (*MS_HeroList) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{9}
}

func (m *MS_HeroList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MS_HeroList.Unmarshal(m, b)
}
func (m *MS_HeroList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MS_HeroList.Marshal(b, m, deterministic)
}
func (m *MS_HeroList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MS_HeroList.Merge(m, src)
}
func (m *MS_HeroList) XXX_Size() int {
	return xxx_messageInfo_MS_HeroList.Size(m)
}
func (m *MS_HeroList) XXX_DiscardUnknown() {
	xxx_messageInfo_MS_HeroList.DiscardUnknown(m)
}

var xxx_messageInfo_MS_HeroList proto.InternalMessageInfo

func (m *MS_HeroList) GetHeros() []*Hero {
	if m != nil {
		return m.Heros
	}
	return nil
}

type Item struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TypeId               int32    `protobuf:"varint,2,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{10}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetTypeId() int32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

type MC_AddItem struct {
	TypeId               int32    `protobuf:"varint,1,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MC_AddItem) Reset()         { *m = MC_AddItem{} }
func (m *MC_AddItem) String() string { return proto.CompactTextString(m) }
func (*MC_AddItem) ProtoMessage()    {}
func (*MC_AddItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{11}
}

func (m *MC_AddItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MC_AddItem.Unmarshal(m, b)
}
func (m *MC_AddItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MC_AddItem.Marshal(b, m, deterministic)
}
func (m *MC_AddItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MC_AddItem.Merge(m, src)
}
func (m *MC_AddItem) XXX_Size() int {
	return xxx_messageInfo_MC_AddItem.Size(m)
}
func (m *MC_AddItem) XXX_DiscardUnknown() {
	xxx_messageInfo_MC_AddItem.DiscardUnknown(m)
}

var xxx_messageInfo_MC_AddItem proto.InternalMessageInfo

func (m *MC_AddItem) GetTypeId() int32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

type MS_ItemList struct {
	Items                []*Item  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MS_ItemList) Reset()         { *m = MS_ItemList{} }
func (m *MS_ItemList) String() string { return proto.CompactTextString(m) }
func (*MS_ItemList) ProtoMessage()    {}
func (*MS_ItemList) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{12}
}

func (m *MS_ItemList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MS_ItemList.Unmarshal(m, b)
}
func (m *MS_ItemList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MS_ItemList.Marshal(b, m, deterministic)
}
func (m *MS_ItemList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MS_ItemList.Merge(m, src)
}
func (m *MS_ItemList) XXX_Size() int {
	return xxx_messageInfo_MS_ItemList.Size(m)
}
func (m *MS_ItemList) XXX_DiscardUnknown() {
	xxx_messageInfo_MS_ItemList.DiscardUnknown(m)
}

var xxx_messageInfo_MS_ItemList proto.InternalMessageInfo

func (m *MS_ItemList) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type MC_ChangeExp struct {
	AddExp               int64    `protobuf:"varint,1,opt,name=add_exp,json=addExp,proto3" json:"add_exp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MC_ChangeExp) Reset()         { *m = MC_ChangeExp{} }
func (m *MC_ChangeExp) String() string { return proto.CompactTextString(m) }
func (*MC_ChangeExp) ProtoMessage()    {}
func (*MC_ChangeExp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{13}
}

func (m *MC_ChangeExp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MC_ChangeExp.Unmarshal(m, b)
}
func (m *MC_ChangeExp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MC_ChangeExp.Marshal(b, m, deterministic)
}
func (m *MC_ChangeExp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MC_ChangeExp.Merge(m, src)
}
func (m *MC_ChangeExp) XXX_Size() int {
	return xxx_messageInfo_MC_ChangeExp.Size(m)
}
func (m *MC_ChangeExp) XXX_DiscardUnknown() {
	xxx_messageInfo_MC_ChangeExp.DiscardUnknown(m)
}

var xxx_messageInfo_MC_ChangeExp proto.InternalMessageInfo

func (m *MC_ChangeExp) GetAddExp() int64 {
	if m != nil {
		return m.AddExp
	}
	return 0
}

type MS_ExpUpdate struct {
	Exp                  int64    `protobuf:"varint,1,opt,name=exp,proto3" json:"exp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MS_ExpUpdate) Reset()         { *m = MS_ExpUpdate{} }
func (m *MS_ExpUpdate) String() string { return proto.CompactTextString(m) }
func (*MS_ExpUpdate) ProtoMessage()    {}
func (*MS_ExpUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{14}
}

func (m *MS_ExpUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MS_ExpUpdate.Unmarshal(m, b)
}
func (m *MS_ExpUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MS_ExpUpdate.Marshal(b, m, deterministic)
}
func (m *MS_ExpUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MS_ExpUpdate.Merge(m, src)
}
func (m *MS_ExpUpdate) XXX_Size() int {
	return xxx_messageInfo_MS_ExpUpdate.Size(m)
}
func (m *MS_ExpUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_MS_ExpUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_MS_ExpUpdate proto.InternalMessageInfo

func (m *MS_ExpUpdate) GetExp() int64 {
	if m != nil {
		return m.Exp
	}
	return 0
}

type MC_ChangeLevel struct {
	AddLevel             int32    `protobuf:"varint,1,opt,name=add_level,json=addLevel,proto3" json:"add_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MC_ChangeLevel) Reset()         { *m = MC_ChangeLevel{} }
func (m *MC_ChangeLevel) String() string { return proto.CompactTextString(m) }
func (*MC_ChangeLevel) ProtoMessage()    {}
func (*MC_ChangeLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{15}
}

func (m *MC_ChangeLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MC_ChangeLevel.Unmarshal(m, b)
}
func (m *MC_ChangeLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MC_ChangeLevel.Marshal(b, m, deterministic)
}
func (m *MC_ChangeLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MC_ChangeLevel.Merge(m, src)
}
func (m *MC_ChangeLevel) XXX_Size() int {
	return xxx_messageInfo_MC_ChangeLevel.Size(m)
}
func (m *MC_ChangeLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_MC_ChangeLevel.DiscardUnknown(m)
}

var xxx_messageInfo_MC_ChangeLevel proto.InternalMessageInfo

func (m *MC_ChangeLevel) GetAddLevel() int32 {
	if m != nil {
		return m.AddLevel
	}
	return 0
}

type MS_LevelUpdate struct {
	Level                int32    `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MS_LevelUpdate) Reset()         { *m = MS_LevelUpdate{} }
func (m *MS_LevelUpdate) String() string { return proto.CompactTextString(m) }
func (*MS_LevelUpdate) ProtoMessage()    {}
func (*MS_LevelUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{16}
}

func (m *MS_LevelUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MS_LevelUpdate.Unmarshal(m, b)
}
func (m *MS_LevelUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MS_LevelUpdate.Marshal(b, m, deterministic)
}
func (m *MS_LevelUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MS_LevelUpdate.Merge(m, src)
}
func (m *MS_LevelUpdate) XXX_Size() int {
	return xxx_messageInfo_MS_LevelUpdate.Size(m)
}
func (m *MS_LevelUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_MS_LevelUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_MS_LevelUpdate proto.InternalMessageInfo

func (m *MS_LevelUpdate) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func init() {
	proto.RegisterType((*ClientInfo)(nil), "yokai_client.ClientInfo")
	proto.RegisterType((*MC_ClientLogon)(nil), "yokai_client.MC_ClientLogon")
	proto.RegisterType((*MS_ClientLogon)(nil), "yokai_client.MS_ClientLogon")
	proto.RegisterType((*MC_HeartBeat)(nil), "yokai_client.MC_HeartBeat")
	proto.RegisterType((*MS_HeartBeat)(nil), "yokai_client.MS_HeartBeat")
	proto.RegisterType((*MC_ClientConnected)(nil), "yokai_client.MC_ClientConnected")
	proto.RegisterType((*MC_ClientDisconnect)(nil), "yokai_client.MC_ClientDisconnect")
	proto.RegisterType((*Hero)(nil), "yokai_client.Hero")
	proto.RegisterType((*MC_AddHero)(nil), "yokai_client.MC_AddHero")
	proto.RegisterType((*MS_HeroList)(nil), "yokai_client.MS_HeroList")
	proto.RegisterType((*Item)(nil), "yokai_client.Item")
	proto.RegisterType((*MC_AddItem)(nil), "yokai_client.MC_AddItem")
	proto.RegisterType((*MS_ItemList)(nil), "yokai_client.MS_ItemList")
	proto.RegisterType((*MC_ChangeExp)(nil), "yokai_client.MC_ChangeExp")
	proto.RegisterType((*MS_ExpUpdate)(nil), "yokai_client.MS_ExpUpdate")
	proto.RegisterType((*MC_ChangeLevel)(nil), "yokai_client.MC_ChangeLevel")
	proto.RegisterType((*MS_LevelUpdate)(nil), "yokai_client.MS_LevelUpdate")
}

func init() { proto.RegisterFile("client/client.proto", fileDescriptor_4d3551c163a1d198) }

var fileDescriptor_4d3551c163a1d198 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0x76, 0x2d, 0xeb, 0x6b, 0x89, 0x26, 0x0f, 0x44, 0x25, 0x90, 0x88, 0x2c, 0x01,
	0x39, 0x8c, 0x04, 0xc1, 0x81, 0x33, 0x0b, 0x91, 0x16, 0xa9, 0xd9, 0x21, 0x13, 0x17, 0x2e, 0x91,
	0x17, 0x3f, 0x5a, 0x8b, 0x26, 0xb6, 0x12, 0x33, 0x65, 0xff, 0x3d, 0xb2, 0x1d, 0xda, 0x20, 0x7e,
	0x49, 0x9c, 0xf2, 0xde, 0xf3, 0xf7, 0x7d, 0xfc, 0x75, 0x9e, 0x0d, 0xe7, 0xd5, 0x5e, 0x60, 0xa3,
	0x63, 0xf7, 0x89, 0x54, 0x2b, 0xb5, 0x24, 0xab, 0x7b, 0xf9, 0x95, 0x89, 0xd2, 0xd5, 0xe8, 0x1b,
	0x80, 0xc4, 0x46, 0x59, 0xf3, 0x45, 0x12, 0x1f, 0x26, 0x82, 0xaf, 0xbd, 0xc0, 0x0b, 0xa7, 0xc5,
	0x44, 0x70, 0x42, 0xe0, 0xa4, 0x61, 0x35, 0xae, 0x27, 0x81, 0x17, 0x2e, 0x0a, 0x1b, 0xd3, 0x6b,
	0xf0, 0xf3, 0xa4, 0x74, 0x4d, 0x1b, 0xb9, 0x95, 0x0d, 0x79, 0x0a, 0x0b, 0x47, 0x2b, 0x0f, 0xcd,
	0xa7, 0xae, 0x90, 0x71, 0xf2, 0x1c, 0x96, 0xc3, 0xe2, 0x88, 0x04, 0xae, 0x74, 0x6d, 0x78, 0x67,
	0xe0, 0xe7, 0x37, 0x63, 0x1e, 0xf5, 0x61, 0x95, 0x27, 0xe5, 0x15, 0xb2, 0x56, 0x5f, 0x22, 0xd3,
	0xf4, 0x02, 0x56, 0xf9, 0xcd, 0x31, 0x27, 0xcf, 0x60, 0xa1, 0x45, 0x8d, 0x9d, 0x66, 0xb5, 0xb2,
	0xfb, 0x3d, 0x2c, 0x8e, 0x05, 0x9a, 0x02, 0x39, 0xf8, 0x4b, 0x64, 0xd3, 0x60, 0xa5, 0x91, 0xff,
	0xdd, 0xe3, 0xef, 0x8e, 0xf9, 0x18, 0xce, 0x0f, 0x98, 0x8f, 0xa2, 0xab, 0x1c, 0x89, 0xc6, 0x70,
	0x72, 0x85, 0xed, 0xaf, 0x7f, 0xea, 0x09, 0x3c, 0xd0, 0xf7, 0x0a, 0x0d, 0xdd, 0x50, 0x66, 0xc5,
	0xdc, 0xa4, 0x19, 0xa7, 0x2f, 0x00, 0xf2, 0xa4, 0xfc, 0xc0, 0xb9, 0x6d, 0x1b, 0xc9, 0xbc, 0x9f,
	0x64, 0xef, 0x61, 0x69, 0xcf, 0xd8, 0xca, 0x8d, 0xe8, 0x34, 0x09, 0x61, 0xb6, 0xc3, 0x56, 0x76,
	0x6b, 0x2f, 0x98, 0x86, 0xcb, 0xb7, 0x24, 0x1a, 0x0f, 0x2d, 0x32, 0xb2, 0xc2, 0x09, 0x8c, 0xa1,
	0x4c, 0x63, 0xfd, 0x1f, 0x86, 0x6c, 0xdb, 0x3f, 0x0c, 0x19, 0xcd, 0x0f, 0x43, 0x42, 0x63, 0xfd,
	0x07, 0x43, 0x46, 0x56, 0x38, 0x01, 0x7d, 0x65, 0xa7, 0x97, 0xec, 0x58, 0xb3, 0xc5, 0xb4, 0x57,
	0x66, 0x07, 0xc6, 0x79, 0x89, 0xbd, 0x1a, 0xdc, 0xcd, 0x19, 0xe7, 0x69, 0xaf, 0x68, 0x60, 0xc7,
	0x9a, 0xf6, 0xea, 0x93, 0xe2, 0x4c, 0x23, 0x39, 0x83, 0xe9, 0x51, 0x64, 0x42, 0xfa, 0xda, 0x5d,
	0x35, 0x8b, 0xda, 0xe0, 0x1d, 0xee, 0xcd, 0x18, 0x0d, 0x6c, 0x6f, 0x92, 0xc1, 0xf0, 0x29, 0xe3,
	0xdc, 0x2e, 0xd2, 0x97, 0xf6, 0x26, 0xd9, 0x78, 0x40, 0x3e, 0x82, 0xd9, 0x58, 0xea, 0x92, 0xcb,
	0xe8, 0xf3, 0xc5, 0x56, 0xe8, 0xdd, 0xb7, 0xdb, 0xa8, 0x92, 0x75, 0x6c, 0x0f, 0x22, 0xa4, 0xfb,
	0x96, 0x1d, 0xb6, 0x77, 0xd8, 0xc6, 0xf6, 0xa9, 0x0c, 0xef, 0xe6, 0x76, 0x6e, 0xb3, 0x77, 0xdf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xae, 0x10, 0x89, 0x6f, 0x4f, 0x03, 0x00, 0x00,
}
