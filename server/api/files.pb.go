// Code generated by protoc-gen-go. DO NOT EDIT.
// source: files.proto

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

type FileStatResponse_Type int32

const (
	FileStatResponse_Directory FileStatResponse_Type = 0
	FileStatResponse_File      FileStatResponse_Type = 1
)

var FileStatResponse_Type_name = map[int32]string{
	0: "Directory",
	1: "File",
}

var FileStatResponse_Type_value = map[string]int32{
	"Directory": 0,
	"File":      1,
}

func (x FileStatResponse_Type) String() string {
	return proto.EnumName(FileStatResponse_Type_name, int32(x))
}

func (FileStatResponse_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{6, 0}
}

type FileRequest struct {
	Cid                  string   `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	Output               string   `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileRequest) Reset()         { *m = FileRequest{} }
func (m *FileRequest) String() string { return proto.CompactTextString(m) }
func (*FileRequest) ProtoMessage()    {}
func (*FileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{0}
}

func (m *FileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileRequest.Unmarshal(m, b)
}
func (m *FileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileRequest.Marshal(b, m, deterministic)
}
func (m *FileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileRequest.Merge(m, src)
}
func (m *FileRequest) XXX_Size() int {
	return xxx_messageInfo_FileRequest.Size(m)
}
func (m *FileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileRequest proto.InternalMessageInfo

func (m *FileRequest) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *FileRequest) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

type FileResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileResponse) Reset()         { *m = FileResponse{} }
func (m *FileResponse) String() string { return proto.CompactTextString(m) }
func (*FileResponse) ProtoMessage()    {}
func (*FileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{1}
}

func (m *FileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileResponse.Unmarshal(m, b)
}
func (m *FileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileResponse.Marshal(b, m, deterministic)
}
func (m *FileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileResponse.Merge(m, src)
}
func (m *FileResponse) XXX_Size() int {
	return xxx_messageInfo_FileResponse.Size(m)
}
func (m *FileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FileResponse proto.InternalMessageInfo

func (m *FileResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type FileCreateRequest struct {
	Fname                string   `protobuf:"bytes,2,opt,name=fname,proto3" json:"fname,omitempty"`
	Source               string   `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileCreateRequest) Reset()         { *m = FileCreateRequest{} }
func (m *FileCreateRequest) String() string { return proto.CompactTextString(m) }
func (*FileCreateRequest) ProtoMessage()    {}
func (*FileCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{2}
}

func (m *FileCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileCreateRequest.Unmarshal(m, b)
}
func (m *FileCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileCreateRequest.Marshal(b, m, deterministic)
}
func (m *FileCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileCreateRequest.Merge(m, src)
}
func (m *FileCreateRequest) XXX_Size() int {
	return xxx_messageInfo_FileCreateRequest.Size(m)
}
func (m *FileCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileCreateRequest proto.InternalMessageInfo

func (m *FileCreateRequest) GetFname() string {
	if m != nil {
		return m.Fname
	}
	return ""
}

func (m *FileCreateRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type FileMkdirRequest struct {
	DirName              string   `protobuf:"bytes,1,opt,name=dirName,proto3" json:"dirName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileMkdirRequest) Reset()         { *m = FileMkdirRequest{} }
func (m *FileMkdirRequest) String() string { return proto.CompactTextString(m) }
func (*FileMkdirRequest) ProtoMessage()    {}
func (*FileMkdirRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{3}
}

func (m *FileMkdirRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileMkdirRequest.Unmarshal(m, b)
}
func (m *FileMkdirRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileMkdirRequest.Marshal(b, m, deterministic)
}
func (m *FileMkdirRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileMkdirRequest.Merge(m, src)
}
func (m *FileMkdirRequest) XXX_Size() int {
	return xxx_messageInfo_FileMkdirRequest.Size(m)
}
func (m *FileMkdirRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileMkdirRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileMkdirRequest proto.InternalMessageInfo

func (m *FileMkdirRequest) GetDirName() string {
	if m != nil {
		return m.DirName
	}
	return ""
}

type FileDataResponse struct {
	Cid                  string   `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileDataResponse) Reset()         { *m = FileDataResponse{} }
func (m *FileDataResponse) String() string { return proto.CompactTextString(m) }
func (*FileDataResponse) ProtoMessage()    {}
func (*FileDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{4}
}

func (m *FileDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileDataResponse.Unmarshal(m, b)
}
func (m *FileDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileDataResponse.Marshal(b, m, deterministic)
}
func (m *FileDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileDataResponse.Merge(m, src)
}
func (m *FileDataResponse) XXX_Size() int {
	return xxx_messageInfo_FileDataResponse.Size(m)
}
func (m *FileDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FileDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FileDataResponse proto.InternalMessageInfo

func (m *FileDataResponse) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

type FileStateRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileStateRequest) Reset()         { *m = FileStateRequest{} }
func (m *FileStateRequest) String() string { return proto.CompactTextString(m) }
func (*FileStateRequest) ProtoMessage()    {}
func (*FileStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{5}
}

func (m *FileStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileStateRequest.Unmarshal(m, b)
}
func (m *FileStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileStateRequest.Marshal(b, m, deterministic)
}
func (m *FileStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileStateRequest.Merge(m, src)
}
func (m *FileStateRequest) XXX_Size() int {
	return xxx_messageInfo_FileStateRequest.Size(m)
}
func (m *FileStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileStateRequest proto.InternalMessageInfo

func (m *FileStateRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type FileStatResponse struct {
	Cid                  string                `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	BlockSize            int64                 `protobuf:"varint,2,opt,name=blockSize,proto3" json:"blockSize,omitempty"`
	CumulativeSize       int64                 `protobuf:"varint,3,opt,name=cumulativeSize,proto3" json:"cumulativeSize,omitempty"`
	DataSize             int64                 `protobuf:"varint,4,opt,name=dataSize,proto3" json:"dataSize,omitempty"`
	NumLinks             int64                 `protobuf:"varint,5,opt,name=numLinks,proto3" json:"numLinks,omitempty"`
	LinksSize            int64                 `protobuf:"varint,6,opt,name=linksSize,proto3" json:"linksSize,omitempty"`
	Type                 FileStatResponse_Type `protobuf:"varint,7,opt,name=type,proto3,enum=api.FileStatResponse_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FileStatResponse) Reset()         { *m = FileStatResponse{} }
func (m *FileStatResponse) String() string { return proto.CompactTextString(m) }
func (*FileStatResponse) ProtoMessage()    {}
func (*FileStatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{6}
}

func (m *FileStatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileStatResponse.Unmarshal(m, b)
}
func (m *FileStatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileStatResponse.Marshal(b, m, deterministic)
}
func (m *FileStatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileStatResponse.Merge(m, src)
}
func (m *FileStatResponse) XXX_Size() int {
	return xxx_messageInfo_FileStatResponse.Size(m)
}
func (m *FileStatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FileStatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FileStatResponse proto.InternalMessageInfo

func (m *FileStatResponse) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *FileStatResponse) GetBlockSize() int64 {
	if m != nil {
		return m.BlockSize
	}
	return 0
}

func (m *FileStatResponse) GetCumulativeSize() int64 {
	if m != nil {
		return m.CumulativeSize
	}
	return 0
}

func (m *FileStatResponse) GetDataSize() int64 {
	if m != nil {
		return m.DataSize
	}
	return 0
}

func (m *FileStatResponse) GetNumLinks() int64 {
	if m != nil {
		return m.NumLinks
	}
	return 0
}

func (m *FileStatResponse) GetLinksSize() int64 {
	if m != nil {
		return m.LinksSize
	}
	return 0
}

func (m *FileStatResponse) GetType() FileStatResponse_Type {
	if m != nil {
		return m.Type
	}
	return FileStatResponse_Directory
}

func init() {
	proto.RegisterEnum("api.FileStatResponse_Type", FileStatResponse_Type_name, FileStatResponse_Type_value)
	proto.RegisterType((*FileRequest)(nil), "api.FileRequest")
	proto.RegisterType((*FileResponse)(nil), "api.FileResponse")
	proto.RegisterType((*FileCreateRequest)(nil), "api.FileCreateRequest")
	proto.RegisterType((*FileMkdirRequest)(nil), "api.FileMkdirRequest")
	proto.RegisterType((*FileDataResponse)(nil), "api.FileDataResponse")
	proto.RegisterType((*FileStateRequest)(nil), "api.FileStateRequest")
	proto.RegisterType((*FileStatResponse)(nil), "api.FileStatResponse")
}

func init() { proto.RegisterFile("files.proto", fileDescriptor_cac8f32ecfdd343c) }

var fileDescriptor_cac8f32ecfdd343c = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x4d, 0x9a, 0x6c, 0xba, 0x3b, 0xa5, 0x55, 0x3a, 0xa2, 0x55, 0x14, 0x21, 0x51, 0x59, 0x68,
	0xd5, 0x03, 0xca, 0xa1, 0x88, 0xf6, 0x0c, 0x54, 0xc0, 0x01, 0x38, 0x64, 0xf9, 0x03, 0x6e, 0xe2,
	0xaa, 0xd6, 0x66, 0x13, 0xe3, 0x8f, 0x4a, 0xe1, 0x5f, 0xf0, 0x43, 0xf9, 0x0f, 0xc8, 0x76, 0x36,
	0x1b, 0xca, 0x72, 0xf3, 0x7b, 0xf3, 0x66, 0x9e, 0xc7, 0x79, 0x81, 0xa3, 0x7b, 0xde, 0x30, 0x55,
	0x08, 0xd9, 0xe9, 0x0e, 0x23, 0x2a, 0x38, 0xb9, 0x81, 0xa3, 0x8f, 0xbc, 0x61, 0x25, 0xfb, 0x61,
	0x98, 0xd2, 0x98, 0x42, 0x54, 0xf1, 0x3a, 0x0b, 0x2f, 0xc2, 0xcb, 0x45, 0x69, 0x8f, 0x78, 0x0e,
	0x49, 0x67, 0xb4, 0x30, 0x3a, 0x3b, 0x70, 0xe4, 0x80, 0xc8, 0x12, 0x9e, 0xf9, 0x46, 0x25, 0xba,
	0x56, 0x31, 0xab, 0x53, 0x9a, 0x6a, 0xa3, 0x5c, 0xf3, 0xbc, 0x1c, 0x10, 0x79, 0x07, 0xa7, 0x56,
	0xf7, 0x41, 0x32, 0xaa, 0x47, 0x9b, 0xe7, 0x30, 0xbb, 0x6f, 0xe9, 0x86, 0x0d, 0x33, 0x3d, 0x70,
	0x23, 0x3a, 0x23, 0x2b, 0x96, 0x45, 0xde, 0xca, 0x23, 0xf2, 0x1a, 0x52, 0x3b, 0xe2, 0xeb, 0xba,
	0xe6, 0x72, 0x3b, 0x21, 0x83, 0xc3, 0x9a, 0xcb, 0x6f, 0x76, 0x86, 0xbf, 0xec, 0x16, 0x92, 0x57,
	0x5e, 0x7d, 0x4b, 0x35, 0x1d, 0x2f, 0xf7, 0xcf, 0x5a, 0x64, 0xe9, 0x55, 0x2b, 0x3d, 0xb9, 0x15,
	0x42, 0x2c, 0xa8, 0x7e, 0x18, 0x64, 0xee, 0x4c, 0x7e, 0x1d, 0xec, 0x84, 0xff, 0x1f, 0x87, 0x2f,
	0x60, 0x71, 0xd7, 0x74, 0xd5, 0x7a, 0xc5, 0x7f, 0xfa, 0xa5, 0xa2, 0x72, 0x47, 0xe0, 0x12, 0x4e,
	0x2a, 0xb3, 0x31, 0x0d, 0xd5, 0xfc, 0x91, 0x39, 0x49, 0xe4, 0x24, 0x4f, 0x58, 0xcc, 0x61, 0x5e,
	0x53, 0x4d, 0x9d, 0x22, 0x76, 0x8a, 0x11, 0xdb, 0x5a, 0x6b, 0x36, 0x5f, 0x78, 0xbb, 0x56, 0xd9,
	0xcc, 0xd7, 0xb6, 0xd8, 0xba, 0x37, 0xf6, 0xe0, 0x1a, 0x13, 0xef, 0x3e, 0x12, 0x58, 0x40, 0xac,
	0x7b, 0xc1, 0xb2, 0xc3, 0x8b, 0xf0, 0xf2, 0xe4, 0x2a, 0x2f, 0xa8, 0xe0, 0xc5, 0xd3, 0x95, 0x8a,
	0xef, 0xbd, 0x60, 0xa5, 0xd3, 0x91, 0x97, 0x10, 0x5b, 0x84, 0xc7, 0xb0, 0xb8, 0xe5, 0x92, 0x55,
	0xba, 0x93, 0x7d, 0x1a, 0xe0, 0x1c, 0x62, 0xdb, 0x95, 0x86, 0x57, 0xbf, 0x43, 0x1f, 0x9a, 0x15,
	0x93, 0x8f, 0xbc, 0x62, 0x78, 0x0d, 0xc7, 0x9f, 0x98, 0xb6, 0xcc, 0xfb, 0xfe, 0x33, 0x55, 0x0f,
	0x98, 0x8e, 0x1e, 0xc3, 0xd3, 0xe6, 0xa7, 0x13, 0xc6, 0x3b, 0x92, 0x00, 0xdf, 0x42, 0xe2, 0x63,
	0x81, 0xe7, 0x63, 0xf9, 0xaf, 0x9c, 0xec, 0x6f, 0xbb, 0x81, 0x99, 0x8b, 0x02, 0x9e, 0x8d, 0xd5,
	0x69, 0x34, 0xf2, 0x1d, 0x3d, 0xcd, 0x00, 0x09, 0xf0, 0x1a, 0x62, 0xbb, 0xf3, 0xa4, 0x6f, 0xfa,
	0xf9, 0xf3, 0xb3, 0xbd, 0x2f, 0x43, 0x82, 0xbb, 0xc4, 0xfd, 0x2f, 0x6f, 0xfe, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xb4, 0x19, 0x98, 0x3c, 0x3e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileServiceClient interface {
	GetFileByHash(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileResponse, error)
	Create(ctx context.Context, in *FileCreateRequest, opts ...grpc.CallOption) (*FileResponse, error)
	Mkdir(ctx context.Context, in *FileMkdirRequest, opts ...grpc.CallOption) (*FileDataResponse, error)
	Stat(ctx context.Context, in *FileStateRequest, opts ...grpc.CallOption) (*FileStatResponse, error)
}

type fileServiceClient struct {
	cc *grpc.ClientConn
}

func NewFileServiceClient(cc *grpc.ClientConn) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) GetFileByHash(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileResponse, error) {
	out := new(FileResponse)
	err := c.cc.Invoke(ctx, "/api.FileService/GetFileByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) Create(ctx context.Context, in *FileCreateRequest, opts ...grpc.CallOption) (*FileResponse, error) {
	out := new(FileResponse)
	err := c.cc.Invoke(ctx, "/api.FileService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) Mkdir(ctx context.Context, in *FileMkdirRequest, opts ...grpc.CallOption) (*FileDataResponse, error) {
	out := new(FileDataResponse)
	err := c.cc.Invoke(ctx, "/api.FileService/Mkdir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) Stat(ctx context.Context, in *FileStateRequest, opts ...grpc.CallOption) (*FileStatResponse, error) {
	out := new(FileStatResponse)
	err := c.cc.Invoke(ctx, "/api.FileService/Stat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
type FileServiceServer interface {
	GetFileByHash(context.Context, *FileRequest) (*FileResponse, error)
	Create(context.Context, *FileCreateRequest) (*FileResponse, error)
	Mkdir(context.Context, *FileMkdirRequest) (*FileDataResponse, error)
	Stat(context.Context, *FileStateRequest) (*FileStatResponse, error)
}

// UnimplementedFileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (*UnimplementedFileServiceServer) GetFileByHash(ctx context.Context, req *FileRequest) (*FileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileByHash not implemented")
}
func (*UnimplementedFileServiceServer) Create(ctx context.Context, req *FileCreateRequest) (*FileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedFileServiceServer) Mkdir(ctx context.Context, req *FileMkdirRequest) (*FileDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mkdir not implemented")
}
func (*UnimplementedFileServiceServer) Stat(ctx context.Context, req *FileStateRequest) (*FileStatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}

func RegisterFileServiceServer(s *grpc.Server, srv FileServiceServer) {
	s.RegisterService(&_FileService_serviceDesc, srv)
}

func _FileService_GetFileByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetFileByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FileService/GetFileByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetFileByHash(ctx, req.(*FileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FileService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).Create(ctx, req.(*FileCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_Mkdir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileMkdirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).Mkdir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FileService/Mkdir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).Mkdir(ctx, req.(*FileMkdirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FileService/Stat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).Stat(ctx, req.(*FileStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFileByHash",
			Handler:    _FileService_GetFileByHash_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _FileService_Create_Handler,
		},
		{
			MethodName: "Mkdir",
			Handler:    _FileService_Mkdir_Handler,
		},
		{
			MethodName: "Stat",
			Handler:    _FileService_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "files.proto",
}