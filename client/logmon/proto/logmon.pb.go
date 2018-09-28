// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client/logmon/proto/logmon.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StartRequest struct {
	LogDir               string   `protobuf:"bytes,1,opt,name=log_dir,json=logDir,proto3" json:"log_dir,omitempty"`
	StdoutFileName       string   `protobuf:"bytes,2,opt,name=stdout_file_name,json=stdoutFileName,proto3" json:"stdout_file_name,omitempty"`
	StderrFileName       string   `protobuf:"bytes,3,opt,name=stderr_file_name,json=stderrFileName,proto3" json:"stderr_file_name,omitempty"`
	MaxFiles             uint32   `protobuf:"varint,4,opt,name=max_files,json=maxFiles,proto3" json:"max_files,omitempty"`
	MaxFileSizeMb        uint32   `protobuf:"varint,5,opt,name=max_file_size_mb,json=maxFileSizeMb,proto3" json:"max_file_size_mb,omitempty"`
	StdoutFifo           string   `protobuf:"bytes,6,opt,name=stdout_fifo,json=stdoutFifo,proto3" json:"stdout_fifo,omitempty"`
	StderrFifo           string   `protobuf:"bytes,7,opt,name=stderr_fifo,json=stderrFifo,proto3" json:"stderr_fifo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartRequest) Reset()         { *m = StartRequest{} }
func (m *StartRequest) String() string { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()    {}
func (*StartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_logmon_6dbff459851a9ae9, []int{0}
}
func (m *StartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartRequest.Unmarshal(m, b)
}
func (m *StartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartRequest.Marshal(b, m, deterministic)
}
func (dst *StartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartRequest.Merge(dst, src)
}
func (m *StartRequest) XXX_Size() int {
	return xxx_messageInfo_StartRequest.Size(m)
}
func (m *StartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartRequest proto.InternalMessageInfo

func (m *StartRequest) GetLogDir() string {
	if m != nil {
		return m.LogDir
	}
	return ""
}

func (m *StartRequest) GetStdoutFileName() string {
	if m != nil {
		return m.StdoutFileName
	}
	return ""
}

func (m *StartRequest) GetStderrFileName() string {
	if m != nil {
		return m.StderrFileName
	}
	return ""
}

func (m *StartRequest) GetMaxFiles() uint32 {
	if m != nil {
		return m.MaxFiles
	}
	return 0
}

func (m *StartRequest) GetMaxFileSizeMb() uint32 {
	if m != nil {
		return m.MaxFileSizeMb
	}
	return 0
}

func (m *StartRequest) GetStdoutFifo() string {
	if m != nil {
		return m.StdoutFifo
	}
	return ""
}

func (m *StartRequest) GetStderrFifo() string {
	if m != nil {
		return m.StderrFifo
	}
	return ""
}

type StartResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartResponse) Reset()         { *m = StartResponse{} }
func (m *StartResponse) String() string { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()    {}
func (*StartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_logmon_6dbff459851a9ae9, []int{1}
}
func (m *StartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartResponse.Unmarshal(m, b)
}
func (m *StartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartResponse.Marshal(b, m, deterministic)
}
func (dst *StartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartResponse.Merge(dst, src)
}
func (m *StartResponse) XXX_Size() int {
	return xxx_messageInfo_StartResponse.Size(m)
}
func (m *StartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartResponse proto.InternalMessageInfo

type StopRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopRequest) Reset()         { *m = StopRequest{} }
func (m *StopRequest) String() string { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()    {}
func (*StopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_logmon_6dbff459851a9ae9, []int{2}
}
func (m *StopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopRequest.Unmarshal(m, b)
}
func (m *StopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopRequest.Marshal(b, m, deterministic)
}
func (dst *StopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopRequest.Merge(dst, src)
}
func (m *StopRequest) XXX_Size() int {
	return xxx_messageInfo_StopRequest.Size(m)
}
func (m *StopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopRequest proto.InternalMessageInfo

type StopResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopResponse) Reset()         { *m = StopResponse{} }
func (m *StopResponse) String() string { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()    {}
func (*StopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_logmon_6dbff459851a9ae9, []int{3}
}
func (m *StopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopResponse.Unmarshal(m, b)
}
func (m *StopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopResponse.Marshal(b, m, deterministic)
}
func (dst *StopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopResponse.Merge(dst, src)
}
func (m *StopResponse) XXX_Size() int {
	return xxx_messageInfo_StopResponse.Size(m)
}
func (m *StopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StartRequest)(nil), "hashicorp.nomad.client.logmon.proto.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "hashicorp.nomad.client.logmon.proto.StartResponse")
	proto.RegisterType((*StopRequest)(nil), "hashicorp.nomad.client.logmon.proto.StopRequest")
	proto.RegisterType((*StopResponse)(nil), "hashicorp.nomad.client.logmon.proto.StopResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogMonClient is the client API for LogMon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogMonClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
}

type logMonClient struct {
	cc *grpc.ClientConn
}

func NewLogMonClient(cc *grpc.ClientConn) LogMonClient {
	return &logMonClient{cc}
}

func (c *logMonClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.client.logmon.proto.LogMon/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logMonClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.client.logmon.proto.LogMon/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogMonServer is the server API for LogMon service.
type LogMonServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Stop(context.Context, *StopRequest) (*StopResponse, error)
}

func RegisterLogMonServer(s *grpc.Server, srv LogMonServer) {
	s.RegisterService(&_LogMon_serviceDesc, srv)
}

func _LogMon_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogMonServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.client.logmon.proto.LogMon/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogMonServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogMon_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogMonServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.client.logmon.proto.LogMon/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogMonServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogMon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.nomad.client.logmon.proto.LogMon",
	HandlerType: (*LogMonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _LogMon_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _LogMon_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client/logmon/proto/logmon.proto",
}

func init() {
	proto.RegisterFile("client/logmon/proto/logmon.proto", fileDescriptor_logmon_6dbff459851a9ae9)
}

var fileDescriptor_logmon_6dbff459851a9ae9 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x31, 0x6f, 0xc2, 0x30,
	0x10, 0x85, 0x1b, 0x0a, 0xa1, 0x1c, 0x0d, 0x45, 0x5e, 0x1a, 0xd1, 0xa1, 0x28, 0x1d, 0xca, 0x14,
	0x0a, 0xfd, 0x07, 0x55, 0xd5, 0xa9, 0x74, 0x80, 0xad, 0x4b, 0x64, 0xc0, 0x09, 0x96, 0x62, 0x5f,
	0x6a, 0x1b, 0x09, 0xb1, 0xf6, 0xd7, 0xf6, 0x5f, 0x54, 0x71, 0x4c, 0x94, 0x11, 0x26, 0xeb, 0xee,
	0x7d, 0x4f, 0xf7, 0xfc, 0x60, 0xbc, 0xc9, 0x39, 0x93, 0x66, 0x9a, 0x63, 0x26, 0x50, 0x4e, 0x0b,
	0x85, 0x06, 0xdd, 0x10, 0xdb, 0x81, 0x3c, 0xed, 0xa8, 0xde, 0xf1, 0x0d, 0xaa, 0x22, 0x96, 0x28,
	0xe8, 0x36, 0xae, 0x1c, 0x71, 0x13, 0x8a, 0x7e, 0x5b, 0x70, 0xbb, 0x32, 0x54, 0x99, 0x25, 0xfb,
	0xd9, 0x33, 0x6d, 0xc8, 0x3d, 0x74, 0x73, 0xcc, 0x92, 0x2d, 0x57, 0xa1, 0x37, 0xf6, 0x26, 0xbd,
	0xa5, 0x9f, 0x63, 0xf6, 0xce, 0x15, 0x99, 0xc0, 0x50, 0x9b, 0x2d, 0xee, 0x4d, 0x92, 0xf2, 0x9c,
	0x25, 0x92, 0x0a, 0x16, 0xb6, 0x2c, 0x31, 0xa8, 0xf6, 0x1f, 0x3c, 0x67, 0x5f, 0x54, 0x30, 0x47,
	0x32, 0xa5, 0x1a, 0xe4, 0x75, 0x4d, 0x32, 0xa5, 0x6a, 0xf2, 0x01, 0x7a, 0x82, 0x1e, 0x2c, 0xa6,
	0xc3, 0xf6, 0xd8, 0x9b, 0x04, 0xcb, 0x1b, 0x41, 0x0f, 0xa5, 0xae, 0xc9, 0x33, 0x0c, 0x4f, 0x62,
	0xa2, 0xf9, 0x91, 0x25, 0x62, 0x1d, 0x76, 0x2c, 0x13, 0x38, 0x66, 0xc5, 0x8f, 0x6c, 0xb1, 0x26,
	0x8f, 0xd0, 0xaf, 0x93, 0xa5, 0x18, 0xfa, 0xf6, 0x14, 0x9c, 0x42, 0xa5, 0xe8, 0x80, 0x2a, 0x50,
	0x8a, 0x61, 0xb7, 0x06, 0x6c, 0x96, 0x14, 0xa3, 0x3b, 0x08, 0x5c, 0x09, 0xba, 0x40, 0xa9, 0x59,
	0x14, 0x40, 0x7f, 0x65, 0xb0, 0x70, 0xa5, 0x44, 0x83, 0xb2, 0xa4, 0x72, 0xac, 0xe4, 0xf9, 0x9f,
	0x07, 0xfe, 0x27, 0x66, 0x0b, 0x94, 0xa4, 0x80, 0x8e, 0xb5, 0x92, 0x59, 0x7c, 0x46, 0xdf, 0x71,
	0xb3, 0xeb, 0xd1, 0xfc, 0x12, 0x8b, 0x4b, 0x76, 0x45, 0x04, 0xb4, 0xcb, 0x30, 0xe4, 0xe5, 0x4c,
	0x77, 0xfd, 0x8d, 0xd1, 0xec, 0x02, 0xc7, 0xe9, 0xdc, 0x5b, 0xf7, 0xbb, 0x63, 0xf7, 0x6b, 0xdf,
	0x3e, 0xaf, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xb4, 0x87, 0xfc, 0x7a, 0x02, 0x00, 0x00,
}
