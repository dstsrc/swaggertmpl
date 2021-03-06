// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/pb/service.proto

package example

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type StringMessage struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringMessage) Reset()         { *m = StringMessage{} }
func (m *StringMessage) String() string { return proto.CompactTextString(m) }
func (*StringMessage) ProtoMessage()    {}
func (*StringMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6d296d44b7b6a15, []int{0}
}

func (m *StringMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMessage.Unmarshal(m, b)
}
func (m *StringMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMessage.Marshal(b, m, deterministic)
}
func (m *StringMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMessage.Merge(m, src)
}
func (m *StringMessage) XXX_Size() int {
	return xxx_messageInfo_StringMessage.Size(m)
}
func (m *StringMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StringMessage proto.InternalMessageInfo

func (m *StringMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*StringMessage)(nil), "example.StringMessage")
}

func init() {
	proto.RegisterFile("pkg/pb/service.proto", fileDescriptor_d6d296d44b7b6a15)
}

var fileDescriptor_d6d296d44b7b6a15 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x29, 0xc8, 0x4e, 0xd7,
	0x2f, 0x48, 0xd2, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x95, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf,
	0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf,
	0x2b, 0x86, 0x28, 0x53, 0x52, 0xe5, 0xe2, 0x0d, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0xf7, 0x4d, 0x2d,
	0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x8c, 0x0e, 0x33, 0x72, 0x71, 0x47, 0xe6, 0x97, 0x16, 0x05,
	0x43, 0xec, 0x10, 0x0a, 0xe6, 0x62, 0x71, 0x4d, 0xce, 0xc8, 0x17, 0x12, 0xd3, 0x83, 0x5a, 0xa3,
	0x87, 0x62, 0x8a, 0x14, 0x0e, 0x71, 0x25, 0xe9, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0x89, 0x2a, 0x09,
	0xe8, 0x97, 0x19, 0xea, 0x43, 0x95, 0xe8, 0xa7, 0x02, 0x4d, 0xb2, 0x62, 0xd4, 0x12, 0x8a, 0xe7,
	0xe2, 0x02, 0x19, 0x0a, 0xd4, 0x91, 0x9a, 0x98, 0x4b, 0xb2, 0xd1, 0xb2, 0x60, 0xa3, 0xc5, 0x95,
	0x84, 0x90, 0x8d, 0x2e, 0x06, 0x9b, 0x05, 0x34, 0x5c, 0x83, 0xd1, 0x80, 0x31, 0x89, 0x0d, 0xec,
	0x67, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x54, 0xdb, 0xbd, 0x32, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// YourServiceClient is the client API for YourService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type YourServiceClient interface {
	Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
	EchoStream(ctx context.Context, opts ...grpc.CallOption) (YourService_EchoStreamClient, error)
}

type yourServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYourServiceClient(cc grpc.ClientConnInterface) YourServiceClient {
	return &yourServiceClient{cc}
}

func (c *yourServiceClient) Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, "/example.YourService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yourServiceClient) EchoStream(ctx context.Context, opts ...grpc.CallOption) (YourService_EchoStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_YourService_serviceDesc.Streams[0], "/example.YourService/EchoStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &yourServiceEchoStreamClient{stream}
	return x, nil
}

type YourService_EchoStreamClient interface {
	Send(*StringMessage) error
	Recv() (*StringMessage, error)
	grpc.ClientStream
}

type yourServiceEchoStreamClient struct {
	grpc.ClientStream
}

func (x *yourServiceEchoStreamClient) Send(m *StringMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *yourServiceEchoStreamClient) Recv() (*StringMessage, error) {
	m := new(StringMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// YourServiceServer is the server API for YourService service.
type YourServiceServer interface {
	Echo(context.Context, *StringMessage) (*StringMessage, error)
	EchoStream(YourService_EchoStreamServer) error
}

// UnimplementedYourServiceServer can be embedded to have forward compatible implementations.
type UnimplementedYourServiceServer struct {
}

func (*UnimplementedYourServiceServer) Echo(ctx context.Context, req *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (*UnimplementedYourServiceServer) EchoStream(srv YourService_EchoStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EchoStream not implemented")
}

func RegisterYourServiceServer(s *grpc.Server, srv YourServiceServer) {
	s.RegisterService(&_YourService_serviceDesc, srv)
}

func _YourService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YourServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.YourService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YourServiceServer).Echo(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _YourService_EchoStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(YourServiceServer).EchoStream(&yourServiceEchoStreamServer{stream})
}

type YourService_EchoStreamServer interface {
	Send(*StringMessage) error
	Recv() (*StringMessage, error)
	grpc.ServerStream
}

type yourServiceEchoStreamServer struct {
	grpc.ServerStream
}

func (x *yourServiceEchoStreamServer) Send(m *StringMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *yourServiceEchoStreamServer) Recv() (*StringMessage, error) {
	m := new(StringMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _YourService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.YourService",
	HandlerType: (*YourServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _YourService_Echo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EchoStream",
			Handler:       _YourService_EchoStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/pb/service.proto",
}
