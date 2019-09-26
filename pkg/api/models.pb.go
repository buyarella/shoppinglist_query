// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models/v1/models.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Currency int32

const (
	Currency_EURO      Currency = 0
	Currency_POUND     Currency = 1
	Currency_US_DOLLAR Currency = 2
)

var Currency_name = map[int32]string{
	0: "EURO",
	1: "POUND",
	2: "US_DOLLAR",
}

var Currency_value = map[string]int32{
	"EURO":      0,
	"POUND":     1,
	"US_DOLLAR": 2,
}

func (x Currency) String() string {
	return proto.EnumName(Currency_name, int32(x))
}

func (Currency) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c211c9176b87007, []int{0}
}

type Unit int32

const (
	Unit_PIECES     Unit = 0
	Unit_KILO_GRAMM Unit = 1
	Unit_GRAMM      Unit = 2
	Unit_METER      Unit = 3
)

var Unit_name = map[int32]string{
	0: "PIECES",
	1: "KILO_GRAMM",
	2: "GRAMM",
	3: "METER",
}

var Unit_value = map[string]int32{
	"PIECES":     0,
	"KILO_GRAMM": 1,
	"GRAMM":      2,
	"METER":      3,
}

func (x Unit) String() string {
	return proto.EnumName(Unit_name, int32(x))
}

func (Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c211c9176b87007, []int{1}
}

type ShoppingList struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string               `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Entries              []*ShoppingListEntry `protobuf:"bytes,3,rep,name=entries,proto3" json:"entries,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Creator              *User                `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ShoppingList) Reset()         { *m = ShoppingList{} }
func (m *ShoppingList) String() string { return proto.CompactTextString(m) }
func (*ShoppingList) ProtoMessage()    {}
func (*ShoppingList) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c211c9176b87007, []int{0}
}

func (m *ShoppingList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShoppingList.Unmarshal(m, b)
}
func (m *ShoppingList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShoppingList.Marshal(b, m, deterministic)
}
func (m *ShoppingList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShoppingList.Merge(m, src)
}
func (m *ShoppingList) XXX_Size() int {
	return xxx_messageInfo_ShoppingList.Size(m)
}
func (m *ShoppingList) XXX_DiscardUnknown() {
	xxx_messageInfo_ShoppingList.DiscardUnknown(m)
}

var xxx_messageInfo_ShoppingList proto.InternalMessageInfo

func (m *ShoppingList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ShoppingList) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ShoppingList) GetEntries() []*ShoppingListEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *ShoppingList) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ShoppingList) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *ShoppingList) GetCreator() *User {
	if m != nil {
		return m.Creator
	}
	return nil
}

type ShoppingListEntry struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	BoughtTime           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=bought_time,json=boughtTime,proto3" json:"bought_time,omitempty"`
	RemovedTime          *timestamp.Timestamp `protobuf:"bytes,5,opt,name=removed_time,json=removedTime,proto3" json:"removed_time,omitempty"`
	Item                 *Item                `protobuf:"bytes,7,opt,name=item,proto3" json:"item,omitempty"`
	Amount               *Amount              `protobuf:"bytes,8,opt,name=amount,proto3" json:"amount,omitempty"`
	Requestor            *User                `protobuf:"bytes,9,opt,name=requestor,proto3" json:"requestor,omitempty"`
	Offering             *Offering            `protobuf:"bytes,10,opt,name=offering,proto3" json:"offering,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ShoppingListEntry) Reset()         { *m = ShoppingListEntry{} }
func (m *ShoppingListEntry) String() string { return proto.CompactTextString(m) }
func (*ShoppingListEntry) ProtoMessage()    {}
func (*ShoppingListEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c211c9176b87007, []int{1}
}

func (m *ShoppingListEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShoppingListEntry.Unmarshal(m, b)
}
func (m *ShoppingListEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShoppingListEntry.Marshal(b, m, deterministic)
}
func (m *ShoppingListEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShoppingListEntry.Merge(m, src)
}
func (m *ShoppingListEntry) XXX_Size() int {
	return xxx_messageInfo_ShoppingListEntry.Size(m)
}
func (m *ShoppingListEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ShoppingListEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ShoppingListEntry proto.InternalMessageInfo

func (m *ShoppingListEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ShoppingListEntry) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ShoppingListEntry) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *ShoppingListEntry) GetBoughtTime() *timestamp.Timestamp {
	if m != nil {
		return m.BoughtTime
	}
	return nil
}

func (m *ShoppingListEntry) GetRemovedTime() *timestamp.Timestamp {
	if m != nil {
		return m.RemovedTime
	}
	return nil
}

func (m *ShoppingListEntry) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *ShoppingListEntry) GetAmount() *Amount {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *ShoppingListEntry) GetRequestor() *User {
	if m != nil {
		return m.Requestor
	}
	return nil
}

func (m *ShoppingListEntry) GetOffering() *Offering {
	if m != nil {
		return m.Offering
	}
	return nil
}

type Item struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string   `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	ImageUrl             string   `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c211c9176b87007, []int{2}
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

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Item) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

type Amount struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Unit                 Unit     `protobuf:"varint,2,opt,name=unit,proto3,enum=buyarella.Unit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Amount) Reset()         { *m = Amount{} }
func (m *Amount) String() string { return proto.CompactTextString(m) }
func (*Amount) ProtoMessage()    {}
func (*Amount) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c211c9176b87007, []int{3}
}

func (m *Amount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Amount.Unmarshal(m, b)
}
func (m *Amount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Amount.Marshal(b, m, deterministic)
}
func (m *Amount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Amount.Merge(m, src)
}
func (m *Amount) XXX_Size() int {
	return xxx_messageInfo_Amount.Size(m)
}
func (m *Amount) XXX_DiscardUnknown() {
	xxx_messageInfo_Amount.DiscardUnknown(m)
}

var xxx_messageInfo_Amount proto.InternalMessageInfo

func (m *Amount) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Amount) GetUnit() Unit {
	if m != nil {
		return m.Unit
	}
	return Unit_PIECES
}

type Offering struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Item                 *Item    `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	Vendor               *Vendor  `protobuf:"bytes,3,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Price                float32  `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Currency             Currency `protobuf:"varint,5,opt,name=currency,proto3,enum=buyarella.Currency" json:"currency,omitempty"`
	Per                  *Amount  `protobuf:"bytes,6,opt,name=per,proto3" json:"per,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Offering) Reset()         { *m = Offering{} }
func (m *Offering) String() string { return proto.CompactTextString(m) }
func (*Offering) ProtoMessage()    {}
func (*Offering) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c211c9176b87007, []int{4}
}

func (m *Offering) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Offering.Unmarshal(m, b)
}
func (m *Offering) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Offering.Marshal(b, m, deterministic)
}
func (m *Offering) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Offering.Merge(m, src)
}
func (m *Offering) XXX_Size() int {
	return xxx_messageInfo_Offering.Size(m)
}
func (m *Offering) XXX_DiscardUnknown() {
	xxx_messageInfo_Offering.DiscardUnknown(m)
}

var xxx_messageInfo_Offering proto.InternalMessageInfo

func (m *Offering) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Offering) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *Offering) GetVendor() *Vendor {
	if m != nil {
		return m.Vendor
	}
	return nil
}

func (m *Offering) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Offering) GetCurrency() Currency {
	if m != nil {
		return m.Currency
	}
	return Currency_EURO
}

func (m *Offering) GetPer() *Amount {
	if m != nil {
		return m.Per
	}
	return nil
}

type Vendor struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string               `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Vendor) Reset()         { *m = Vendor{} }
func (m *Vendor) String() string { return proto.CompactTextString(m) }
func (*Vendor) ProtoMessage()    {}
func (*Vendor) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c211c9176b87007, []int{5}
}

func (m *Vendor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vendor.Unmarshal(m, b)
}
func (m *Vendor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vendor.Marshal(b, m, deterministic)
}
func (m *Vendor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vendor.Merge(m, src)
}
func (m *Vendor) XXX_Size() int {
	return xxx_messageInfo_Vendor.Size(m)
}
func (m *Vendor) XXX_DiscardUnknown() {
	xxx_messageInfo_Vendor.DiscardUnknown(m)
}

var xxx_messageInfo_Vendor proto.InternalMessageInfo

func (m *Vendor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vendor) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Vendor) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Vendor) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Fistname             string   `protobuf:"bytes,3,opt,name=fistname,proto3" json:"fistname,omitempty"`
	Lastname             string   `protobuf:"bytes,4,opt,name=lastname,proto3" json:"lastname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c211c9176b87007, []int{6}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetFistname() string {
	if m != nil {
		return m.Fistname
	}
	return ""
}

func (m *User) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func init() {
	proto.RegisterEnum("buyarella.Currency", Currency_name, Currency_value)
	proto.RegisterEnum("buyarella.Unit", Unit_name, Unit_value)
	proto.RegisterType((*ShoppingList)(nil), "buyarella.ShoppingList")
	proto.RegisterType((*ShoppingListEntry)(nil), "buyarella.ShoppingListEntry")
	proto.RegisterType((*Item)(nil), "buyarella.Item")
	proto.RegisterType((*Amount)(nil), "buyarella.Amount")
	proto.RegisterType((*Offering)(nil), "buyarella.Offering")
	proto.RegisterType((*Vendor)(nil), "buyarella.Vendor")
	proto.RegisterType((*User)(nil), "buyarella.User")
}

func init() { proto.RegisterFile("models/v1/models.proto", fileDescriptor_0c211c9176b87007) }

var fileDescriptor_0c211c9176b87007 = []byte{
	// 643 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xdd, 0x4e, 0xdb, 0x30,
	0x14, 0x26, 0x3f, 0x6d, 0x93, 0x53, 0xc6, 0x8a, 0x87, 0xa6, 0x88, 0x4d, 0x1a, 0x0b, 0x37, 0x80,
	0xb4, 0x54, 0x63, 0xd2, 0x34, 0x09, 0xed, 0x82, 0x41, 0x35, 0xa1, 0x15, 0x8a, 0x0c, 0xdd, 0x05,
	0x37, 0x55, 0xda, 0xba, 0xc1, 0x5a, 0x12, 0x67, 0x8e, 0x83, 0xd4, 0xd7, 0xd9, 0x5b, 0xec, 0x21,
	0xf6, 0x02, 0x7b, 0x9a, 0xc9, 0x76, 0xd2, 0x15, 0xa8, 0xc4, 0xdf, 0x9d, 0x8f, 0xbf, 0xef, 0xd8,
	0xdf, 0xf9, 0xce, 0x39, 0xf0, 0x32, 0x61, 0x63, 0x12, 0xe7, 0xed, 0xab, 0xf7, 0x6d, 0x7d, 0x0a,
	0x32, 0xce, 0x04, 0x43, 0xee, 0xb0, 0x98, 0x86, 0x9c, 0xc4, 0x71, 0xb8, 0xfe, 0x26, 0x62, 0x2c,
	0x8a, 0x49, 0x5b, 0x01, 0xc3, 0x62, 0xd2, 0x16, 0x34, 0x21, 0xb9, 0x08, 0x93, 0x4c, 0x73, 0xfd,
	0x5f, 0x26, 0x2c, 0x9f, 0x5d, 0xb2, 0x2c, 0xa3, 0x69, 0xd4, 0xa5, 0xb9, 0x40, 0x08, 0xec, 0x34,
	0x4c, 0x88, 0x67, 0x6c, 0x18, 0x5b, 0x2e, 0x56, 0x67, 0xf4, 0x16, 0x96, 0xc7, 0x34, 0xcf, 0xe2,
	0x70, 0x3a, 0x50, 0x98, 0xa9, 0xb0, 0x66, 0x79, 0x77, 0x22, 0x29, 0x1f, 0xa1, 0x41, 0x52, 0xc1,
	0x29, 0xc9, 0x3d, 0x6b, 0xc3, 0xda, 0x6a, 0xee, 0xbe, 0x0e, 0x66, 0x2a, 0x82, 0xf9, 0x0f, 0x3a,
	0xa9, 0xe0, 0x53, 0x5c, 0x91, 0xd1, 0x1e, 0x34, 0x47, 0x9c, 0x84, 0x82, 0x0c, 0xa4, 0x32, 0xcf,
	0xde, 0x30, 0xb6, 0x9a, 0xbb, 0xeb, 0x81, 0x96, 0x1d, 0x54, 0xb2, 0x83, 0xf3, 0x4a, 0x36, 0x06,
	0x4d, 0x97, 0x17, 0x32, 0xb9, 0xc8, 0xc6, 0xb3, 0xe4, 0xda, 0xdd, 0xc9, 0x9a, 0xae, 0x92, 0xb7,
	0xa1, 0xa1, 0x9e, 0x62, 0xdc, 0xab, 0xab, 0xc4, 0xe7, 0x73, 0x8a, 0xfb, 0x39, 0xe1, 0xb8, 0xc2,
	0xfd, 0x3f, 0x16, 0xac, 0xde, 0xaa, 0x61, 0xa1, 0x53, 0x37, 0xca, 0x31, 0x9f, 0x52, 0x8e, 0xf5,
	0xa0, 0x72, 0xf6, 0xa0, 0x39, 0x64, 0x45, 0x74, 0x29, 0xee, 0x6d, 0xa4, 0xa6, 0xab, 0xe4, 0xcf,
	0xb0, 0xcc, 0x49, 0xc2, 0xae, 0xc8, 0xf8, 0xbe, 0x4e, 0x36, 0x4b, 0xbe, 0x4a, 0xdf, 0x04, 0x9b,
	0x0a, 0x92, 0x78, 0x8d, 0x5b, 0x3e, 0x1e, 0x09, 0x92, 0x60, 0x05, 0xa2, 0x6d, 0xa8, 0x87, 0x09,
	0x2b, 0x52, 0xe1, 0x39, 0x8a, 0xb6, 0x3a, 0x47, 0xdb, 0x57, 0x00, 0x2e, 0x09, 0xe8, 0x1d, 0xb8,
	0x9c, 0xfc, 0x2c, 0x48, 0x2e, 0x9b, 0xe3, 0x2e, 0x6e, 0xce, 0x7f, 0x06, 0x6a, 0x83, 0xc3, 0x26,
	0x13, 0xc2, 0x69, 0x1a, 0x79, 0xa0, 0xd8, 0x2f, 0xe6, 0xd8, 0xbd, 0x12, 0xc2, 0x33, 0x92, 0x7f,
	0x01, 0xb6, 0x14, 0xf6, 0xd8, 0x59, 0x7f, 0x05, 0x2e, 0x4d, 0xc2, 0x88, 0x0c, 0x0a, 0x1e, 0xab,
	0x2e, 0xb9, 0xd8, 0x51, 0x17, 0x7d, 0x1e, 0xfb, 0x07, 0x50, 0xd7, 0xd5, 0xa0, 0x35, 0xa8, 0x8d,
	0x54, 0xbd, 0xf2, 0xf9, 0x1a, 0xd6, 0x81, 0xf4, 0xaa, 0x48, 0xa9, 0x50, 0xef, 0xae, 0x5c, 0x2f,
	0x2b, 0xa5, 0x02, 0x2b, 0xd0, 0xff, 0x6b, 0x80, 0x53, 0xe9, 0x5e, 0xa8, 0xb2, 0x72, 0xdc, 0xbc,
	0xc3, 0xf1, 0x2b, 0x92, 0x8e, 0x19, 0x2f, 0x47, 0x69, 0xde, 0xf1, 0xef, 0x0a, 0xc0, 0x25, 0x41,
	0x6a, 0xcd, 0x38, 0x1d, 0xe9, 0xb9, 0x31, 0xb1, 0x0e, 0xa4, 0xb1, 0xa3, 0x82, 0x73, 0x92, 0x8e,
	0xa6, 0x6a, 0x24, 0x56, 0xae, 0x19, 0x7b, 0x50, 0x42, 0x78, 0x46, 0x42, 0x9b, 0x60, 0x65, 0xa4,
	0xda, 0xa7, 0x05, 0x0d, 0x96, 0xa8, 0xff, 0xdb, 0x80, 0xba, 0xfe, 0xfe, 0xb1, 0x0d, 0xb8, 0xb1,
	0x65, 0xd6, 0x53, 0xb6, 0xcc, 0x7e, 0xc8, 0x96, 0xf9, 0x97, 0x60, 0xcb, 0xe9, 0x5b, 0x28, 0x7c,
	0x0d, 0x6a, 0x24, 0x09, 0x69, 0x5c, 0x2a, 0xd6, 0x01, 0x5a, 0x07, 0x67, 0x42, 0x73, 0xa1, 0xd8,
	0xe5, 0xac, 0x54, 0xb1, 0xc4, 0xe2, 0xb0, 0xc4, 0x6c, 0x8d, 0x55, 0xf1, 0x4e, 0x00, 0x4e, 0x65,
	0x30, 0x72, 0xc0, 0xee, 0xf4, 0x71, 0xaf, 0xb5, 0x84, 0x5c, 0xa8, 0x9d, 0xf6, 0xfa, 0x27, 0x87,
	0x2d, 0x03, 0x3d, 0x03, 0xb7, 0x7f, 0x36, 0x38, 0xec, 0x75, 0xbb, 0xfb, 0xb8, 0x65, 0xee, 0x7c,
	0x02, 0x5b, 0x0e, 0x10, 0x02, 0xa8, 0x9f, 0x1e, 0x75, 0x0e, 0x3a, 0x67, 0xad, 0x25, 0xb4, 0x02,
	0xf0, 0xed, 0xa8, 0xdb, 0x1b, 0x7c, 0xc5, 0xfb, 0xc7, 0xc7, 0x2d, 0x43, 0x66, 0xeb, 0xa3, 0x29,
	0x8f, 0xc7, 0x9d, 0xf3, 0x0e, 0x6e, 0x59, 0x5f, 0xdc, 0x8b, 0x46, 0xf6, 0x23, 0x6a, 0x87, 0x19,
	0x1d, 0xd6, 0x55, 0xf9, 0x1f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x64, 0xa5, 0xa1, 0x5a,
	0x06, 0x00, 0x00,
}