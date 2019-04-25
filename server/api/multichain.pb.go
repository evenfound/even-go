// Code generated by protoc-gen-go. DO NOT EDIT.
// source: multichain.proto

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

type BalancesRequest_Network int32

const (
	BalancesRequest_testnet BalancesRequest_Network = 0
	BalancesRequest_mainnet BalancesRequest_Network = 1
)

var BalancesRequest_Network_name = map[int32]string{
	0: "testnet",
	1: "mainnet",
}

var BalancesRequest_Network_value = map[string]int32{
	"testnet": 0,
	"mainnet": 1,
}

func (x BalancesRequest_Network) String() string {
	return proto.EnumName(BalancesRequest_Network_name, int32(x))
}

func (BalancesRequest_Network) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dd326252c7e9300e, []int{1, 0}
}

type Addrlist struct {
	Addresses            []*Addrlist_Address `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Addrlist) Reset()         { *m = Addrlist{} }
func (m *Addrlist) String() string { return proto.CompactTextString(m) }
func (*Addrlist) ProtoMessage()    {}
func (*Addrlist) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd326252c7e9300e, []int{0}
}

func (m *Addrlist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Addrlist.Unmarshal(m, b)
}
func (m *Addrlist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Addrlist.Marshal(b, m, deterministic)
}
func (m *Addrlist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addrlist.Merge(m, src)
}
func (m *Addrlist) XXX_Size() int {
	return xxx_messageInfo_Addrlist.Size(m)
}
func (m *Addrlist) XXX_DiscardUnknown() {
	xxx_messageInfo_Addrlist.DiscardUnknown(m)
}

var xxx_messageInfo_Addrlist proto.InternalMessageInfo

func (m *Addrlist) GetAddresses() []*Addrlist_Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type Addrlist_Address struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Addrlist_Address) Reset()         { *m = Addrlist_Address{} }
func (m *Addrlist_Address) String() string { return proto.CompactTextString(m) }
func (*Addrlist_Address) ProtoMessage()    {}
func (*Addrlist_Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd326252c7e9300e, []int{0, 0}
}

func (m *Addrlist_Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Addrlist_Address.Unmarshal(m, b)
}
func (m *Addrlist_Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Addrlist_Address.Marshal(b, m, deterministic)
}
func (m *Addrlist_Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addrlist_Address.Merge(m, src)
}
func (m *Addrlist_Address) XXX_Size() int {
	return xxx_messageInfo_Addrlist_Address.Size(m)
}
func (m *Addrlist_Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Addrlist_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Addrlist_Address proto.InternalMessageInfo

func (m *Addrlist_Address) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type BalancesRequest struct {
	Token                string                  `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Network              BalancesRequest_Network `protobuf:"varint,2,opt,name=network,proto3,enum=api.BalancesRequest_Network" json:"network,omitempty"`
	Addrlist             *Addrlist               `protobuf:"bytes,3,opt,name=addrlist,proto3" json:"addrlist,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *BalancesRequest) Reset()         { *m = BalancesRequest{} }
func (m *BalancesRequest) String() string { return proto.CompactTextString(m) }
func (*BalancesRequest) ProtoMessage()    {}
func (*BalancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd326252c7e9300e, []int{1}
}

func (m *BalancesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalancesRequest.Unmarshal(m, b)
}
func (m *BalancesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalancesRequest.Marshal(b, m, deterministic)
}
func (m *BalancesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalancesRequest.Merge(m, src)
}
func (m *BalancesRequest) XXX_Size() int {
	return xxx_messageInfo_BalancesRequest.Size(m)
}
func (m *BalancesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BalancesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BalancesRequest proto.InternalMessageInfo

func (m *BalancesRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *BalancesRequest) GetNetwork() BalancesRequest_Network {
	if m != nil {
		return m.Network
	}
	return BalancesRequest_testnet
}

func (m *BalancesRequest) GetAddrlist() *Addrlist {
	if m != nil {
		return m.Addrlist
	}
	return nil
}

type BalancesResponse struct {
	Balances             map[string]float64 `protobuf:"bytes,1,rep,name=balances,proto3" json:"balances,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BalancesResponse) Reset()         { *m = BalancesResponse{} }
func (m *BalancesResponse) String() string { return proto.CompactTextString(m) }
func (*BalancesResponse) ProtoMessage()    {}
func (*BalancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd326252c7e9300e, []int{2}
}

func (m *BalancesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalancesResponse.Unmarshal(m, b)
}
func (m *BalancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalancesResponse.Marshal(b, m, deterministic)
}
func (m *BalancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalancesResponse.Merge(m, src)
}
func (m *BalancesResponse) XXX_Size() int {
	return xxx_messageInfo_BalancesResponse.Size(m)
}
func (m *BalancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BalancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BalancesResponse proto.InternalMessageInfo

func (m *BalancesResponse) GetBalances() map[string]float64 {
	if m != nil {
		return m.Balances
	}
	return nil
}

func init() {
	proto.RegisterEnum("api.BalancesRequest_Network", BalancesRequest_Network_name, BalancesRequest_Network_value)
	proto.RegisterType((*Addrlist)(nil), "api.Addrlist")
	proto.RegisterType((*Addrlist_Address)(nil), "api.Addrlist.Address")
	proto.RegisterType((*BalancesRequest)(nil), "api.BalancesRequest")
	proto.RegisterType((*BalancesResponse)(nil), "api.BalancesResponse")
	proto.RegisterMapType((map[string]float64)(nil), "api.BalancesResponse.BalancesEntry")
}

func init() { proto.RegisterFile("multichain.proto", fileDescriptor_dd326252c7e9300e) }

var fileDescriptor_dd326252c7e9300e = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x41, 0x4b, 0xf3, 0x40,
	0x14, 0xec, 0x36, 0x7c, 0xdf, 0xb6, 0xaf, 0x56, 0xc3, 0xd2, 0x42, 0x28, 0x1e, 0x42, 0x7a, 0x89,
	0x97, 0x1c, 0x52, 0x10, 0x51, 0x44, 0x14, 0x14, 0x4f, 0x1e, 0xf6, 0x1f, 0x6c, 0x9b, 0x07, 0x5d,
	0x92, 0x6e, 0x62, 0x76, 0xa3, 0xf4, 0x5f, 0xf8, 0x4f, 0xfc, 0x8b, 0x92, 0x64, 0x93, 0xd2, 0xd2,
	0xdb, 0x9b, 0x79, 0x33, 0xec, 0xec, 0xf0, 0xc0, 0xdd, 0x55, 0x99, 0x91, 0x9b, 0xad, 0x90, 0x2a,
	0x2a, 0xca, 0xdc, 0xe4, 0xcc, 0x11, 0x85, 0x0c, 0x12, 0x18, 0x3d, 0x27, 0x49, 0x99, 0x49, 0x6d,
	0xd8, 0x0a, 0xc6, 0x22, 0x49, 0x4a, 0xd4, 0x1a, 0xb5, 0x47, 0x7c, 0x27, 0x9c, 0xc4, 0xf3, 0x48,
	0x14, 0x32, 0xea, 0x14, 0xcd, 0x80, 0x5a, 0xf3, 0x83, 0x6e, 0xb1, 0x04, 0x6a, 0x59, 0xe6, 0x01,
	0xb5, 0xbc, 0x47, 0x7c, 0x12, 0x8e, 0x79, 0x07, 0x83, 0x5f, 0x02, 0x57, 0x2f, 0x22, 0x13, 0x6a,
	0x83, 0x9a, 0xe3, 0x67, 0x85, 0xda, 0xb0, 0x19, 0xfc, 0x33, 0x79, 0x8a, 0xca, 0x6a, 0x5b, 0xc0,
	0x6e, 0x81, 0x2a, 0x34, 0xdf, 0x79, 0x99, 0x7a, 0x43, 0x9f, 0x84, 0x97, 0xf1, 0x75, 0x93, 0xe0,
	0xc4, 0x1c, 0x7d, 0xb4, 0x1a, 0xde, 0x89, 0xd9, 0x0d, 0x8c, 0x84, 0x4d, 0xe9, 0x39, 0x3e, 0x09,
	0x27, 0xf1, 0xf4, 0x28, 0x3a, 0xef, 0xd7, 0xc1, 0x12, 0xa8, 0xb5, 0xb3, 0x09, 0x50, 0x83, 0xda,
	0x28, 0x34, 0xee, 0xa0, 0x06, 0x3b, 0x21, 0x55, 0x0d, 0x48, 0xf0, 0x43, 0xc0, 0x3d, 0x3c, 0xaa,
	0x8b, 0x5c, 0x69, 0x64, 0x4f, 0x30, 0x5a, 0x5b, 0xce, 0xf6, 0xb3, 0x3c, 0x49, 0xd7, 0x0a, 0x7b,
	0xe2, 0x55, 0x99, 0x72, 0xcf, 0x7b, 0xd3, 0xe2, 0x01, 0xa6, 0x47, 0x2b, 0xe6, 0x82, 0x93, 0xe2,
	0xde, 0x56, 0x50, 0x8f, 0x75, 0x2d, 0x5f, 0x22, 0xab, 0xb0, 0xf9, 0x3e, 0xe1, 0x2d, 0xb8, 0x1f,
	0xde, 0x91, 0xf8, 0x1d, 0xa8, 0x35, 0xb3, 0x47, 0xb8, 0x78, 0x43, 0xb3, 0xd9, 0x76, 0x78, 0x76,
	0xae, 0xa4, 0xc5, 0xfc, 0x6c, 0xb8, 0x60, 0xb0, 0xfe, 0xdf, 0x1c, 0xc0, 0xea, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x94, 0xfa, 0x38, 0x45, 0x14, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BalanceClient is the client API for Balance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BalanceClient interface {
	FetchBalance(ctx context.Context, in *BalancesRequest, opts ...grpc.CallOption) (*BalancesResponse, error)
}

type balanceClient struct {
	cc *grpc.ClientConn
}

func NewBalanceClient(cc *grpc.ClientConn) BalanceClient {
	return &balanceClient{cc}
}

func (c *balanceClient) FetchBalance(ctx context.Context, in *BalancesRequest, opts ...grpc.CallOption) (*BalancesResponse, error) {
	out := new(BalancesResponse)
	err := c.cc.Invoke(ctx, "/api.Balance/FetchBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalanceServer is the server API for Balance service.
type BalanceServer interface {
	FetchBalance(context.Context, *BalancesRequest) (*BalancesResponse, error)
}

// UnimplementedBalanceServer can be embedded to have forward compatible implementations.
type UnimplementedBalanceServer struct {
}

func (*UnimplementedBalanceServer) FetchBalance(ctx context.Context, req *BalancesRequest) (*BalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchBalance not implemented")
}

func RegisterBalanceServer(s *grpc.Server, srv BalanceServer) {
	s.RegisterService(&_Balance_serviceDesc, srv)
}

func _Balance_FetchBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServer).FetchBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Balance/FetchBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServer).FetchBalance(ctx, req.(*BalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Balance_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Balance",
	HandlerType: (*BalanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchBalance",
			Handler:    _Balance_FetchBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "multichain.proto",
}