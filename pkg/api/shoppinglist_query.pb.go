// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shoppinglist_query/v1/shoppinglist_query.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetAllShoppingListsRequest struct {
	Filter               *ShoppingListFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	PageSize             int32               `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            string              `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetAllShoppingListsRequest) Reset()         { *m = GetAllShoppingListsRequest{} }
func (m *GetAllShoppingListsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllShoppingListsRequest) ProtoMessage()    {}
func (*GetAllShoppingListsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_594f819429a1a196, []int{0}
}

func (m *GetAllShoppingListsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllShoppingListsRequest.Unmarshal(m, b)
}
func (m *GetAllShoppingListsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllShoppingListsRequest.Marshal(b, m, deterministic)
}
func (m *GetAllShoppingListsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllShoppingListsRequest.Merge(m, src)
}
func (m *GetAllShoppingListsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllShoppingListsRequest.Size(m)
}
func (m *GetAllShoppingListsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllShoppingListsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllShoppingListsRequest proto.InternalMessageInfo

func (m *GetAllShoppingListsRequest) GetFilter() *ShoppingListFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *GetAllShoppingListsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetAllShoppingListsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type GetAllShoppingListsResponse struct {
	ItemsToBuy           []*ShoppingList `protobuf:"bytes,1,rep,name=items_to_buy,json=itemsToBuy,proto3" json:"items_to_buy,omitempty"`
	NextPageToken        string          `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetAllShoppingListsResponse) Reset()         { *m = GetAllShoppingListsResponse{} }
func (m *GetAllShoppingListsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllShoppingListsResponse) ProtoMessage()    {}
func (*GetAllShoppingListsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_594f819429a1a196, []int{1}
}

func (m *GetAllShoppingListsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllShoppingListsResponse.Unmarshal(m, b)
}
func (m *GetAllShoppingListsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllShoppingListsResponse.Marshal(b, m, deterministic)
}
func (m *GetAllShoppingListsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllShoppingListsResponse.Merge(m, src)
}
func (m *GetAllShoppingListsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllShoppingListsResponse.Size(m)
}
func (m *GetAllShoppingListsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllShoppingListsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllShoppingListsResponse proto.InternalMessageInfo

func (m *GetAllShoppingListsResponse) GetItemsToBuy() []*ShoppingList {
	if m != nil {
		return m.ItemsToBuy
	}
	return nil
}

func (m *GetAllShoppingListsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetShoppingListRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShoppingListRequest) Reset()         { *m = GetShoppingListRequest{} }
func (m *GetShoppingListRequest) String() string { return proto.CompactTextString(m) }
func (*GetShoppingListRequest) ProtoMessage()    {}
func (*GetShoppingListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_594f819429a1a196, []int{2}
}

func (m *GetShoppingListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShoppingListRequest.Unmarshal(m, b)
}
func (m *GetShoppingListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShoppingListRequest.Marshal(b, m, deterministic)
}
func (m *GetShoppingListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShoppingListRequest.Merge(m, src)
}
func (m *GetShoppingListRequest) XXX_Size() int {
	return xxx_messageInfo_GetShoppingListRequest.Size(m)
}
func (m *GetShoppingListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShoppingListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShoppingListRequest proto.InternalMessageInfo

func (m *GetShoppingListRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetActiveShoppingListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetActiveShoppingListRequest) Reset()         { *m = GetActiveShoppingListRequest{} }
func (m *GetActiveShoppingListRequest) String() string { return proto.CompactTextString(m) }
func (*GetActiveShoppingListRequest) ProtoMessage()    {}
func (*GetActiveShoppingListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_594f819429a1a196, []int{3}
}

func (m *GetActiveShoppingListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetActiveShoppingListRequest.Unmarshal(m, b)
}
func (m *GetActiveShoppingListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetActiveShoppingListRequest.Marshal(b, m, deterministic)
}
func (m *GetActiveShoppingListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetActiveShoppingListRequest.Merge(m, src)
}
func (m *GetActiveShoppingListRequest) XXX_Size() int {
	return xxx_messageInfo_GetActiveShoppingListRequest.Size(m)
}
func (m *GetActiveShoppingListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetActiveShoppingListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetActiveShoppingListRequest proto.InternalMessageInfo

type ShoppingListFilter struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShoppingListFilter) Reset()         { *m = ShoppingListFilter{} }
func (m *ShoppingListFilter) String() string { return proto.CompactTextString(m) }
func (*ShoppingListFilter) ProtoMessage()    {}
func (*ShoppingListFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_594f819429a1a196, []int{4}
}

func (m *ShoppingListFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShoppingListFilter.Unmarshal(m, b)
}
func (m *ShoppingListFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShoppingListFilter.Marshal(b, m, deterministic)
}
func (m *ShoppingListFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShoppingListFilter.Merge(m, src)
}
func (m *ShoppingListFilter) XXX_Size() int {
	return xxx_messageInfo_ShoppingListFilter.Size(m)
}
func (m *ShoppingListFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ShoppingListFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ShoppingListFilter proto.InternalMessageInfo

type GetAllItemsRequest struct {
	Filter               *ItemsFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	PageSize             int32        `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            string       `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetAllItemsRequest) Reset()         { *m = GetAllItemsRequest{} }
func (m *GetAllItemsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllItemsRequest) ProtoMessage()    {}
func (*GetAllItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_594f819429a1a196, []int{5}
}

func (m *GetAllItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllItemsRequest.Unmarshal(m, b)
}
func (m *GetAllItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllItemsRequest.Marshal(b, m, deterministic)
}
func (m *GetAllItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllItemsRequest.Merge(m, src)
}
func (m *GetAllItemsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllItemsRequest.Size(m)
}
func (m *GetAllItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllItemsRequest proto.InternalMessageInfo

func (m *GetAllItemsRequest) GetFilter() *ItemsFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *GetAllItemsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetAllItemsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ItemsFilter struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemsFilter) Reset()         { *m = ItemsFilter{} }
func (m *ItemsFilter) String() string { return proto.CompactTextString(m) }
func (*ItemsFilter) ProtoMessage()    {}
func (*ItemsFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_594f819429a1a196, []int{6}
}

func (m *ItemsFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemsFilter.Unmarshal(m, b)
}
func (m *ItemsFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemsFilter.Marshal(b, m, deterministic)
}
func (m *ItemsFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemsFilter.Merge(m, src)
}
func (m *ItemsFilter) XXX_Size() int {
	return xxx_messageInfo_ItemsFilter.Size(m)
}
func (m *ItemsFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemsFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ItemsFilter proto.InternalMessageInfo

type GetAllItemsResponse struct {
	Items                []*Item  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllItemsResponse) Reset()         { *m = GetAllItemsResponse{} }
func (m *GetAllItemsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllItemsResponse) ProtoMessage()    {}
func (*GetAllItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_594f819429a1a196, []int{7}
}

func (m *GetAllItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllItemsResponse.Unmarshal(m, b)
}
func (m *GetAllItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllItemsResponse.Marshal(b, m, deterministic)
}
func (m *GetAllItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllItemsResponse.Merge(m, src)
}
func (m *GetAllItemsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllItemsResponse.Size(m)
}
func (m *GetAllItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllItemsResponse proto.InternalMessageInfo

func (m *GetAllItemsResponse) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *GetAllItemsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAllShoppingListsRequest)(nil), "buyarella.shoppinglist.query.GetAllShoppingListsRequest")
	proto.RegisterType((*GetAllShoppingListsResponse)(nil), "buyarella.shoppinglist.query.GetAllShoppingListsResponse")
	proto.RegisterType((*GetShoppingListRequest)(nil), "buyarella.shoppinglist.query.GetShoppingListRequest")
	proto.RegisterType((*GetActiveShoppingListRequest)(nil), "buyarella.shoppinglist.query.GetActiveShoppingListRequest")
	proto.RegisterType((*ShoppingListFilter)(nil), "buyarella.shoppinglist.query.ShoppingListFilter")
	proto.RegisterType((*GetAllItemsRequest)(nil), "buyarella.shoppinglist.query.GetAllItemsRequest")
	proto.RegisterType((*ItemsFilter)(nil), "buyarella.shoppinglist.query.ItemsFilter")
	proto.RegisterType((*GetAllItemsResponse)(nil), "buyarella.shoppinglist.query.GetAllItemsResponse")
}

func init() {
	proto.RegisterFile("shoppinglist_query/v1/shoppinglist_query.proto", fileDescriptor_594f819429a1a196)
}

var fileDescriptor_594f819429a1a196 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xed, 0x36, 0xb4, 0x90, 0x09, 0x55, 0xa4, 0x29, 0x94, 0xc8, 0x2d, 0x28, 0x5a, 0x09, 0x14,
	0x24, 0xe4, 0xb4, 0x81, 0x03, 0xe5, 0xd6, 0x1e, 0x28, 0x48, 0x1c, 0xc0, 0xed, 0x89, 0x8b, 0x65,
	0xd3, 0x21, 0xac, 0xea, 0x78, 0x5d, 0xef, 0xba, 0xc2, 0x3d, 0xf1, 0x01, 0xdc, 0xf8, 0x03, 0x3e,
	0x90, 0x6f, 0x40, 0x5e, 0x3b, 0x68, 0xdd, 0x3a, 0x6e, 0x40, 0xdc, 0x56, 0xb3, 0xf3, 0xe6, 0xbd,
	0x7d, 0xf3, 0xb4, 0xe0, 0xaa, 0x2f, 0x32, 0x49, 0x44, 0x3c, 0x8d, 0x84, 0xd2, 0xfe, 0x79, 0x46,
	0x69, 0x3e, 0xbe, 0xd8, 0x1b, 0x5f, 0xaf, 0xba, 0x49, 0x2a, 0xb5, 0xc4, 0x9d, 0x30, 0xcb, 0x83,
	0x94, 0xa2, 0x28, 0xa8, 0x21, 0x5d, 0xd3, 0xe3, 0x6c, 0xcd, 0xe4, 0x29, 0x45, 0xaa, 0x98, 0x50,
	0x9e, 0x4a, 0x14, 0xff, 0xc9, 0xc0, 0x39, 0x22, 0x7d, 0x10, 0x45, 0xc7, 0x15, 0xe8, 0x9d, 0x50,
	0x5a, 0x79, 0x74, 0x9e, 0x91, 0xd2, 0xf8, 0x06, 0xd6, 0x3f, 0x8b, 0x48, 0x53, 0x3a, 0x60, 0x43,
	0x36, 0xea, 0x4d, 0x76, 0xdd, 0x36, 0x16, 0xd7, 0x9e, 0xf1, 0xda, 0xe0, 0xbc, 0x0a, 0x8f, 0xdb,
	0xd0, 0x4d, 0x82, 0x29, 0xf9, 0x4a, 0x5c, 0xd2, 0x60, 0x75, 0xc8, 0x46, 0x6b, 0xde, 0x9d, 0xa2,
	0x70, 0x2c, 0x2e, 0x09, 0x1f, 0x02, 0x98, 0x4b, 0x2d, 0xcf, 0x28, 0x1e, 0x74, 0x86, 0x6c, 0xd4,
	0xf5, 0x4c, 0xfb, 0x49, 0x51, 0xe0, 0xdf, 0x18, 0x6c, 0x37, 0x8a, 0x54, 0x89, 0x8c, 0x15, 0xe1,
	0x3e, 0xdc, 0x15, 0x9a, 0x66, 0xca, 0xd7, 0xd2, 0x0f, 0xb3, 0x7c, 0xc0, 0x86, 0x9d, 0x51, 0x6f,
	0xf2, 0xc0, 0xd2, 0x6a, 0xe3, 0x3c, 0x30, 0xcd, 0x27, 0xf2, 0x30, 0xcb, 0xf1, 0x09, 0xf4, 0x63,
	0xfa, 0xaa, 0x7d, 0x8b, 0x7e, 0xd5, 0xd0, 0x6f, 0x14, 0xe5, 0xf7, 0x7f, 0x24, 0x3c, 0x83, 0xad,
	0x23, 0xd2, 0xb5, 0x31, 0x95, 0x45, 0x08, 0xb7, 0xe2, 0x60, 0x46, 0xc6, 0xa0, 0xae, 0x67, 0xce,
	0xfc, 0x11, 0xec, 0x14, 0x7a, 0x3f, 0x69, 0x71, 0x41, 0x0d, 0x18, 0x7e, 0x0f, 0xf0, 0xba, 0x55,
	0xfc, 0x07, 0x03, 0x2c, 0x9f, 0xf9, 0xb6, 0x10, 0x38, 0x27, 0x38, 0xb8, 0xb2, 0x83, 0xa7, 0xed,
	0x3b, 0x30, 0xd8, 0xff, 0x68, 0xfe, 0x06, 0xf4, 0xac, 0x91, 0xfc, 0x14, 0x36, 0x6b, 0x1a, 0xab,
	0x15, 0x3c, 0x86, 0x35, 0xe3, 0x6a, 0xe5, 0x7d, 0xdf, 0xd2, 0x58, 0x34, 0x7a, 0xe5, 0xed, 0xb2,
	0x76, 0x4f, 0x7e, 0x75, 0x60, 0xd3, 0x76, 0xe8, 0x43, 0x46, 0xa9, 0x20, 0x85, 0xdf, 0xd9, 0x9c,
	0xbe, 0x96, 0x04, 0x7c, 0xd9, 0xee, 0xc9, 0xe2, 0x84, 0x3b, 0xfb, 0xff, 0x80, 0x2c, 0xdf, 0xcc,
	0x57, 0x30, 0x84, 0xfe, 0x95, 0x54, 0xe0, 0x8b, 0x1b, 0xe7, 0x35, 0x04, 0xc2, 0x59, 0x94, 0x55,
	0xbe, 0x82, 0x31, 0xdc, 0x6f, 0xcc, 0x12, 0xbe, 0xba, 0x59, 0xf9, 0xa2, 0x00, 0xb6, 0xf1, 0x69,
	0xe8, 0x59, 0x0b, 0xc6, 0xdd, 0x65, 0xfc, 0xb1, 0xf3, 0xea, 0xec, 0xfd, 0x05, 0x62, 0xee, 0xe4,
	0x61, 0xf7, 0xe3, 0xed, 0xe4, 0x6c, 0x3a, 0x0e, 0x12, 0x11, 0xae, 0x9b, 0x9f, 0xe9, 0xf9, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x3f, 0x61, 0x26, 0x01, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShoppingListQueriesClient is the client API for ShoppingListQueries service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShoppingListQueriesClient interface {
	GetAllShoppingLists(ctx context.Context, in *GetAllShoppingListsRequest, opts ...grpc.CallOption) (*GetAllShoppingListsResponse, error)
	GetShoppingList(ctx context.Context, in *GetShoppingListRequest, opts ...grpc.CallOption) (*ShoppingList, error)
	GetActiveShoppingList(ctx context.Context, in *GetActiveShoppingListRequest, opts ...grpc.CallOption) (*ShoppingList, error)
	GetAllItems(ctx context.Context, in *GetAllItemsRequest, opts ...grpc.CallOption) (*GetAllItemsResponse, error)
}

type shoppingListQueriesClient struct {
	cc *grpc.ClientConn
}

func NewShoppingListQueriesClient(cc *grpc.ClientConn) ShoppingListQueriesClient {
	return &shoppingListQueriesClient{cc}
}

func (c *shoppingListQueriesClient) GetAllShoppingLists(ctx context.Context, in *GetAllShoppingListsRequest, opts ...grpc.CallOption) (*GetAllShoppingListsResponse, error) {
	out := new(GetAllShoppingListsResponse)
	err := c.cc.Invoke(ctx, "/buyarella.shoppinglist.query.ShoppingListQueries/GetAllShoppingLists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingListQueriesClient) GetShoppingList(ctx context.Context, in *GetShoppingListRequest, opts ...grpc.CallOption) (*ShoppingList, error) {
	out := new(ShoppingList)
	err := c.cc.Invoke(ctx, "/buyarella.shoppinglist.query.ShoppingListQueries/GetShoppingList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingListQueriesClient) GetActiveShoppingList(ctx context.Context, in *GetActiveShoppingListRequest, opts ...grpc.CallOption) (*ShoppingList, error) {
	out := new(ShoppingList)
	err := c.cc.Invoke(ctx, "/buyarella.shoppinglist.query.ShoppingListQueries/GetActiveShoppingList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingListQueriesClient) GetAllItems(ctx context.Context, in *GetAllItemsRequest, opts ...grpc.CallOption) (*GetAllItemsResponse, error) {
	out := new(GetAllItemsResponse)
	err := c.cc.Invoke(ctx, "/buyarella.shoppinglist.query.ShoppingListQueries/GetAllItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShoppingListQueriesServer is the server API for ShoppingListQueries service.
type ShoppingListQueriesServer interface {
	GetAllShoppingLists(context.Context, *GetAllShoppingListsRequest) (*GetAllShoppingListsResponse, error)
	GetShoppingList(context.Context, *GetShoppingListRequest) (*ShoppingList, error)
	GetActiveShoppingList(context.Context, *GetActiveShoppingListRequest) (*ShoppingList, error)
	GetAllItems(context.Context, *GetAllItemsRequest) (*GetAllItemsResponse, error)
}

// UnimplementedShoppingListQueriesServer can be embedded to have forward compatible implementations.
type UnimplementedShoppingListQueriesServer struct {
}

func (*UnimplementedShoppingListQueriesServer) GetAllShoppingLists(ctx context.Context, req *GetAllShoppingListsRequest) (*GetAllShoppingListsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllShoppingLists not implemented")
}
func (*UnimplementedShoppingListQueriesServer) GetShoppingList(ctx context.Context, req *GetShoppingListRequest) (*ShoppingList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShoppingList not implemented")
}
func (*UnimplementedShoppingListQueriesServer) GetActiveShoppingList(ctx context.Context, req *GetActiveShoppingListRequest) (*ShoppingList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveShoppingList not implemented")
}
func (*UnimplementedShoppingListQueriesServer) GetAllItems(ctx context.Context, req *GetAllItemsRequest) (*GetAllItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllItems not implemented")
}

func RegisterShoppingListQueriesServer(s *grpc.Server, srv ShoppingListQueriesServer) {
	s.RegisterService(&_ShoppingListQueries_serviceDesc, srv)
}

func _ShoppingListQueries_GetAllShoppingLists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllShoppingListsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingListQueriesServer).GetAllShoppingLists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buyarella.shoppinglist.query.ShoppingListQueries/GetAllShoppingLists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingListQueriesServer).GetAllShoppingLists(ctx, req.(*GetAllShoppingListsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingListQueries_GetShoppingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShoppingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingListQueriesServer).GetShoppingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buyarella.shoppinglist.query.ShoppingListQueries/GetShoppingList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingListQueriesServer).GetShoppingList(ctx, req.(*GetShoppingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingListQueries_GetActiveShoppingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActiveShoppingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingListQueriesServer).GetActiveShoppingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buyarella.shoppinglist.query.ShoppingListQueries/GetActiveShoppingList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingListQueriesServer).GetActiveShoppingList(ctx, req.(*GetActiveShoppingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingListQueries_GetAllItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingListQueriesServer).GetAllItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buyarella.shoppinglist.query.ShoppingListQueries/GetAllItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingListQueriesServer).GetAllItems(ctx, req.(*GetAllItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShoppingListQueries_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buyarella.shoppinglist.query.ShoppingListQueries",
	HandlerType: (*ShoppingListQueriesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllShoppingLists",
			Handler:    _ShoppingListQueries_GetAllShoppingLists_Handler,
		},
		{
			MethodName: "GetShoppingList",
			Handler:    _ShoppingListQueries_GetShoppingList_Handler,
		},
		{
			MethodName: "GetActiveShoppingList",
			Handler:    _ShoppingListQueries_GetActiveShoppingList_Handler,
		},
		{
			MethodName: "GetAllItems",
			Handler:    _ShoppingListQueries_GetAllItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shoppinglist_query/v1/shoppinglist_query.proto",
}
