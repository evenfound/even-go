// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet.proto

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

type WalletInput struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WalletInput) Reset()         { *m = WalletInput{} }
func (m *WalletInput) String() string { return proto.CompactTextString(m) }
func (*WalletInput) ProtoMessage()    {}
func (*WalletInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{0}
}

func (m *WalletInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletInput.Unmarshal(m, b)
}
func (m *WalletInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletInput.Marshal(b, m, deterministic)
}
func (m *WalletInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletInput.Merge(m, src)
}
func (m *WalletInput) XXX_Size() int {
	return xxx_messageInfo_WalletInput.Size(m)
}
func (m *WalletInput) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletInput.DiscardUnknown(m)
}

var xxx_messageInfo_WalletInput proto.InternalMessageInfo

func (m *WalletInput) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WalletInput) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateWalletInput struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Mnemonic             string   `protobuf:"bytes,3,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateWalletInput) Reset()         { *m = CreateWalletInput{} }
func (m *CreateWalletInput) String() string { return proto.CompactTextString(m) }
func (*CreateWalletInput) ProtoMessage()    {}
func (*CreateWalletInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{1}
}

func (m *CreateWalletInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWalletInput.Unmarshal(m, b)
}
func (m *CreateWalletInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWalletInput.Marshal(b, m, deterministic)
}
func (m *CreateWalletInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWalletInput.Merge(m, src)
}
func (m *CreateWalletInput) XXX_Size() int {
	return xxx_messageInfo_CreateWalletInput.Size(m)
}
func (m *CreateWalletInput) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWalletInput.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWalletInput proto.InternalMessageInfo

func (m *CreateWalletInput) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateWalletInput) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateWalletInput) GetMnemonic() string {
	if m != nil {
		return m.Mnemonic
	}
	return ""
}

type WalletAccountInput struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Account              string   `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WalletAccountInput) Reset()         { *m = WalletAccountInput{} }
func (m *WalletAccountInput) String() string { return proto.CompactTextString(m) }
func (*WalletAccountInput) ProtoMessage()    {}
func (*WalletAccountInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{2}
}

func (m *WalletAccountInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletAccountInput.Unmarshal(m, b)
}
func (m *WalletAccountInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletAccountInput.Marshal(b, m, deterministic)
}
func (m *WalletAccountInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletAccountInput.Merge(m, src)
}
func (m *WalletAccountInput) XXX_Size() int {
	return xxx_messageInfo_WalletAccountInput.Size(m)
}
func (m *WalletAccountInput) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletAccountInput.DiscardUnknown(m)
}

var xxx_messageInfo_WalletAccountInput proto.InternalMessageInfo

func (m *WalletAccountInput) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WalletAccountInput) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *WalletAccountInput) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type WalletResult struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Result               string   `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WalletResult) Reset()         { *m = WalletResult{} }
func (m *WalletResult) String() string { return proto.CompactTextString(m) }
func (*WalletResult) ProtoMessage()    {}
func (*WalletResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{3}
}

func (m *WalletResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletResult.Unmarshal(m, b)
}
func (m *WalletResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletResult.Marshal(b, m, deterministic)
}
func (m *WalletResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletResult.Merge(m, src)
}
func (m *WalletResult) XXX_Size() int {
	return xxx_messageInfo_WalletResult.Size(m)
}
func (m *WalletResult) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletResult.DiscardUnknown(m)
}

var xxx_messageInfo_WalletResult proto.InternalMessageInfo

func (m *WalletResult) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *WalletResult) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*WalletInput)(nil), "api.WalletInput")
	proto.RegisterType((*CreateWalletInput)(nil), "api.CreateWalletInput")
	proto.RegisterType((*WalletAccountInput)(nil), "api.WalletAccountInput")
	proto.RegisterType((*WalletResult)(nil), "api.WalletResult")
}

func init() { proto.RegisterFile("wallet.proto", fileDescriptor_b88fd140af4deb6f) }

var fileDescriptor_b88fd140af4deb6f = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xed, 0x1f, 0x6a, 0x1d, 0x63, 0xb1, 0x03, 0xd6, 0x50, 0x2f, 0x92, 0x93, 0xa7, 0x1e,
	0x14, 0x2b, 0x28, 0x1e, 0xfc, 0x47, 0xd5, 0x42, 0x91, 0xa8, 0x78, 0x53, 0xb6, 0x71, 0xd4, 0xd0,
	0x64, 0x37, 0x6c, 0x37, 0xa6, 0x7e, 0x4e, 0xbf, 0x90, 0xb8, 0xdb, 0x94, 0x84, 0xf6, 0xd0, 0xe8,
	0x2d, 0xf3, 0x66, 0xdf, 0xef, 0xc1, 0x1b, 0x02, 0x56, 0xc2, 0x82, 0x80, 0x54, 0x27, 0x92, 0x42,
	0x09, 0xac, 0xb0, 0xc8, 0x77, 0x4e, 0x61, 0xfd, 0x49, 0x8b, 0x37, 0x3c, 0x8a, 0x15, 0x22, 0x54,
	0x39, 0x0b, 0xc9, 0x2e, 0xed, 0x96, 0xf6, 0xd6, 0x5c, 0xfd, 0x8d, 0x6d, 0xa8, 0x47, 0x6c, 0x3c,
	0x4e, 0x84, 0x7c, 0xb5, 0xcb, 0x5a, 0x9f, 0xcd, 0xce, 0x0b, 0x34, 0x2f, 0x24, 0x31, 0x45, 0xff,
	0x80, 0xfc, 0xee, 0x42, 0x4e, 0xa1, 0xe0, 0xbe, 0x67, 0x57, 0xcc, 0x2e, 0x9d, 0x9d, 0x67, 0x40,
	0x83, 0x3e, 0xf3, 0x3c, 0x11, 0xf3, 0x3f, 0x26, 0xd8, 0xb0, 0xca, 0x8c, 0x7f, 0x1a, 0x90, 0x8e,
	0x4e, 0x17, 0x2c, 0xc3, 0x77, 0x69, 0x1c, 0x07, 0x0a, 0x1b, 0x50, 0x16, 0x23, 0xcd, 0xad, 0xbb,
	0x65, 0x31, 0xc2, 0x16, 0xd4, 0xa4, 0xde, 0x4c, 0x99, 0xd3, 0x69, 0xff, 0xbb, 0x0a, 0x35, 0x63,
	0xc4, 0x23, 0x68, 0xf4, 0x88, 0x93, 0x9c, 0xb5, 0x80, 0x9b, 0x1d, 0x16, 0xf9, 0x9d, 0x4c, 0x25,
	0xed, 0x66, 0x46, 0x31, 0x49, 0xce, 0x0a, 0x9e, 0x80, 0x95, 0x2d, 0x0f, 0x5b, 0xfa, 0xd1, 0x5c,
	0x9f, 0x8b, 0xcd, 0x87, 0x60, 0x3d, 0xf2, 0x40, 0x78, 0xa3, 0x62, 0x99, 0xc7, 0xd0, 0x34, 0xca,
	0x80, 0x26, 0x69, 0xa7, 0xcb, 0x7a, 0xfb, 0xb0, 0x93, 0xbb, 0xc5, 0x65, 0x1c, 0x46, 0x77, 0xd2,
	0xff, 0x64, 0x8a, 0xfa, 0xf4, 0x85, 0xdb, 0x19, 0x4f, 0xf6, 0x5a, 0x8b, 0x61, 0xb7, 0xd0, 0x9e,
	0x87, 0xc5, 0xc3, 0xc0, 0xf7, 0x8a, 0xb3, 0xae, 0xc1, 0xce, 0x3d, 0xbd, 0xff, 0x10, 0xc9, 0x39,
	0x0b, 0x18, 0xf7, 0xa8, 0x20, 0xa9, 0x0b, 0x1b, 0x3d, 0x52, 0x69, 0x13, 0x6f, 0x62, 0xd9, 0x6a,
	0xae, 0x60, 0x2b, 0x17, 0xf1, 0x30, 0x19, 0x50, 0xe2, 0xd2, 0x7b, 0xb1, 0xf8, 0x61, 0x4d, 0xff,
	0x99, 0x07, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x68, 0x81, 0xe0, 0xa9, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WalletClient is the client API for Wallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WalletClient interface {
	// GenerateWallet returns result=seed on success.
	GenerateWallet(ctx context.Context, in *WalletInput, opts ...grpc.CallOption) (*WalletResult, error)
	// CreateWallet returns result=seed on success.
	CreateWallet(ctx context.Context, in *CreateWalletInput, opts ...grpc.CallOption) (*WalletResult, error)
	// UnlockWallet returns result=duration on success.
	UnlockWallet(ctx context.Context, in *WalletInput, opts ...grpc.CallOption) (*WalletResult, error)
	// WalletNextAccount returns result=address on success.
	WalletNextAccount(ctx context.Context, in *WalletInput, opts ...grpc.CallOption) (*WalletResult, error)
	// WalletAccountDumpPrivateKey returns result=privkey on success.
	WalletAccountDumpPrivateKey(ctx context.Context, in *WalletAccountInput, opts ...grpc.CallOption) (*WalletResult, error)
	// WalletAccountDumpPublicKey returns result=pubkey on success.
	WalletAccountDumpPublicKey(ctx context.Context, in *WalletAccountInput, opts ...grpc.CallOption) (*WalletResult, error)
	// WalletAccountShowBalance returns result=balance on success.
	WalletAccountShowBalance(ctx context.Context, in *WalletAccountInput, opts ...grpc.CallOption) (*WalletResult, error)
	// GetWalletInfo returns result=info on success.
	GetWalletInfo(ctx context.Context, in *WalletInput, opts ...grpc.CallOption) (*WalletResult, error)
	// WalletAccountTxNewReg creates initial transaction.
	WalletAccountTxNewReg(ctx context.Context, in *WalletAccountInput, opts ...grpc.CallOption) (*WalletResult, error)
}

type walletClient struct {
	cc *grpc.ClientConn
}

func NewWalletClient(cc *grpc.ClientConn) WalletClient {
	return &walletClient{cc}
}

func (c *walletClient) GenerateWallet(ctx context.Context, in *WalletInput, opts ...grpc.CallOption) (*WalletResult, error) {
	out := new(WalletResult)
	err := c.cc.Invoke(ctx, "/api.Wallet/GenerateWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) CreateWallet(ctx context.Context, in *CreateWalletInput, opts ...grpc.CallOption) (*WalletResult, error) {
	out := new(WalletResult)
	err := c.cc.Invoke(ctx, "/api.Wallet/CreateWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) UnlockWallet(ctx context.Context, in *WalletInput, opts ...grpc.CallOption) (*WalletResult, error) {
	out := new(WalletResult)
	err := c.cc.Invoke(ctx, "/api.Wallet/UnlockWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) WalletNextAccount(ctx context.Context, in *WalletInput, opts ...grpc.CallOption) (*WalletResult, error) {
	out := new(WalletResult)
	err := c.cc.Invoke(ctx, "/api.Wallet/WalletNextAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) WalletAccountDumpPrivateKey(ctx context.Context, in *WalletAccountInput, opts ...grpc.CallOption) (*WalletResult, error) {
	out := new(WalletResult)
	err := c.cc.Invoke(ctx, "/api.Wallet/WalletAccountDumpPrivateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) WalletAccountDumpPublicKey(ctx context.Context, in *WalletAccountInput, opts ...grpc.CallOption) (*WalletResult, error) {
	out := new(WalletResult)
	err := c.cc.Invoke(ctx, "/api.Wallet/WalletAccountDumpPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) WalletAccountShowBalance(ctx context.Context, in *WalletAccountInput, opts ...grpc.CallOption) (*WalletResult, error) {
	out := new(WalletResult)
	err := c.cc.Invoke(ctx, "/api.Wallet/WalletAccountShowBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetWalletInfo(ctx context.Context, in *WalletInput, opts ...grpc.CallOption) (*WalletResult, error) {
	out := new(WalletResult)
	err := c.cc.Invoke(ctx, "/api.Wallet/GetWalletInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) WalletAccountTxNewReg(ctx context.Context, in *WalletAccountInput, opts ...grpc.CallOption) (*WalletResult, error) {
	out := new(WalletResult)
	err := c.cc.Invoke(ctx, "/api.Wallet/WalletAccountTxNewReg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServer is the server API for Wallet service.
type WalletServer interface {
	// GenerateWallet returns result=seed on success.
	GenerateWallet(context.Context, *WalletInput) (*WalletResult, error)
	// CreateWallet returns result=seed on success.
	CreateWallet(context.Context, *CreateWalletInput) (*WalletResult, error)
	// UnlockWallet returns result=duration on success.
	UnlockWallet(context.Context, *WalletInput) (*WalletResult, error)
	// WalletNextAccount returns result=address on success.
	WalletNextAccount(context.Context, *WalletInput) (*WalletResult, error)
	// WalletAccountDumpPrivateKey returns result=privkey on success.
	WalletAccountDumpPrivateKey(context.Context, *WalletAccountInput) (*WalletResult, error)
	// WalletAccountDumpPublicKey returns result=pubkey on success.
	WalletAccountDumpPublicKey(context.Context, *WalletAccountInput) (*WalletResult, error)
	// WalletAccountShowBalance returns result=balance on success.
	WalletAccountShowBalance(context.Context, *WalletAccountInput) (*WalletResult, error)
	// GetWalletInfo returns result=info on success.
	GetWalletInfo(context.Context, *WalletInput) (*WalletResult, error)
	// WalletAccountTxNewReg creates initial transaction.
	WalletAccountTxNewReg(context.Context, *WalletAccountInput) (*WalletResult, error)
}

// UnimplementedWalletServer can be embedded to have forward compatible implementations.
type UnimplementedWalletServer struct {
}

func (*UnimplementedWalletServer) GenerateWallet(ctx context.Context, req *WalletInput) (*WalletResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateWallet not implemented")
}
func (*UnimplementedWalletServer) CreateWallet(ctx context.Context, req *CreateWalletInput) (*WalletResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (*UnimplementedWalletServer) UnlockWallet(ctx context.Context, req *WalletInput) (*WalletResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlockWallet not implemented")
}
func (*UnimplementedWalletServer) WalletNextAccount(ctx context.Context, req *WalletInput) (*WalletResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WalletNextAccount not implemented")
}
func (*UnimplementedWalletServer) WalletAccountDumpPrivateKey(ctx context.Context, req *WalletAccountInput) (*WalletResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WalletAccountDumpPrivateKey not implemented")
}
func (*UnimplementedWalletServer) WalletAccountDumpPublicKey(ctx context.Context, req *WalletAccountInput) (*WalletResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WalletAccountDumpPublicKey not implemented")
}
func (*UnimplementedWalletServer) WalletAccountShowBalance(ctx context.Context, req *WalletAccountInput) (*WalletResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WalletAccountShowBalance not implemented")
}
func (*UnimplementedWalletServer) GetWalletInfo(ctx context.Context, req *WalletInput) (*WalletResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWalletInfo not implemented")
}
func (*UnimplementedWalletServer) WalletAccountTxNewReg(ctx context.Context, req *WalletAccountInput) (*WalletResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WalletAccountTxNewReg not implemented")
}

func RegisterWalletServer(s *grpc.Server, srv WalletServer) {
	s.RegisterService(&_Wallet_serviceDesc, srv)
}

func _Wallet_GenerateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GenerateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Wallet/GenerateWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GenerateWallet(ctx, req.(*WalletInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Wallet/CreateWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).CreateWallet(ctx, req.(*CreateWalletInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_UnlockWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).UnlockWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Wallet/UnlockWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).UnlockWallet(ctx, req.(*WalletInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_WalletNextAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).WalletNextAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Wallet/WalletNextAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).WalletNextAccount(ctx, req.(*WalletInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_WalletAccountDumpPrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletAccountInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).WalletAccountDumpPrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Wallet/WalletAccountDumpPrivateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).WalletAccountDumpPrivateKey(ctx, req.(*WalletAccountInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_WalletAccountDumpPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletAccountInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).WalletAccountDumpPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Wallet/WalletAccountDumpPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).WalletAccountDumpPublicKey(ctx, req.(*WalletAccountInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_WalletAccountShowBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletAccountInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).WalletAccountShowBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Wallet/WalletAccountShowBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).WalletAccountShowBalance(ctx, req.(*WalletAccountInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetWalletInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetWalletInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Wallet/GetWalletInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetWalletInfo(ctx, req.(*WalletInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_WalletAccountTxNewReg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletAccountInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).WalletAccountTxNewReg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Wallet/WalletAccountTxNewReg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).WalletAccountTxNewReg(ctx, req.(*WalletAccountInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _Wallet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Wallet",
	HandlerType: (*WalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateWallet",
			Handler:    _Wallet_GenerateWallet_Handler,
		},
		{
			MethodName: "CreateWallet",
			Handler:    _Wallet_CreateWallet_Handler,
		},
		{
			MethodName: "UnlockWallet",
			Handler:    _Wallet_UnlockWallet_Handler,
		},
		{
			MethodName: "WalletNextAccount",
			Handler:    _Wallet_WalletNextAccount_Handler,
		},
		{
			MethodName: "WalletAccountDumpPrivateKey",
			Handler:    _Wallet_WalletAccountDumpPrivateKey_Handler,
		},
		{
			MethodName: "WalletAccountDumpPublicKey",
			Handler:    _Wallet_WalletAccountDumpPublicKey_Handler,
		},
		{
			MethodName: "WalletAccountShowBalance",
			Handler:    _Wallet_WalletAccountShowBalance_Handler,
		},
		{
			MethodName: "GetWalletInfo",
			Handler:    _Wallet_GetWalletInfo_Handler,
		},
		{
			MethodName: "WalletAccountTxNewReg",
			Handler:    _Wallet_WalletAccountTxNewReg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet.proto",
}
