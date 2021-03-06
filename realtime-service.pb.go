// Code generated by protoc-gen-go. DO NOT EDIT.
// source: realtime-service.proto

package realtime

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

type RouteRequest struct {
	Route                string   `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteRequest) Reset()         { *m = RouteRequest{} }
func (m *RouteRequest) String() string { return proto.CompactTextString(m) }
func (*RouteRequest) ProtoMessage()    {}
func (*RouteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e1541a3a8064139, []int{0}
}

func (m *RouteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteRequest.Unmarshal(m, b)
}
func (m *RouteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteRequest.Marshal(b, m, deterministic)
}
func (m *RouteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteRequest.Merge(m, src)
}
func (m *RouteRequest) XXX_Size() int {
	return xxx_messageInfo_RouteRequest.Size(m)
}
func (m *RouteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RouteRequest proto.InternalMessageInfo

func (m *RouteRequest) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

// A response message containing a greeting
type RouteResponse struct {
	Json                 string   `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
	FromCache            bool     `protobuf:"varint,2,opt,name=fromCache,proto3" json:"fromCache,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteResponse) Reset()         { *m = RouteResponse{} }
func (m *RouteResponse) String() string { return proto.CompactTextString(m) }
func (*RouteResponse) ProtoMessage()    {}
func (*RouteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e1541a3a8064139, []int{1}
}

func (m *RouteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteResponse.Unmarshal(m, b)
}
func (m *RouteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteResponse.Marshal(b, m, deterministic)
}
func (m *RouteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteResponse.Merge(m, src)
}
func (m *RouteResponse) XXX_Size() int {
	return xxx_messageInfo_RouteResponse.Size(m)
}
func (m *RouteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RouteResponse proto.InternalMessageInfo

func (m *RouteResponse) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

func (m *RouteResponse) GetFromCache() bool {
	if m != nil {
		return m.FromCache
	}
	return false
}

func init() {
	proto.RegisterType((*RouteRequest)(nil), "realtime.RouteRequest")
	proto.RegisterType((*RouteResponse)(nil), "realtime.RouteResponse")
}

func init() { proto.RegisterFile("realtime-service.proto", fileDescriptor_3e1541a3a8064139) }

var fileDescriptor_3e1541a3a8064139 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x4a, 0x4d, 0xcc,
	0x29, 0xc9, 0xcc, 0x4d, 0xd5, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x80, 0x89, 0x2b, 0xa9, 0x70, 0xf1, 0x04, 0xe5, 0x97, 0x96, 0xa4, 0x06,
	0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x70, 0xb1, 0x16, 0x81, 0xf8, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x92, 0x23, 0x17, 0x2f, 0x54, 0x55, 0x71, 0x41, 0x7e, 0x5e,
	0x71, 0xaa, 0x90, 0x10, 0x17, 0x4b, 0x56, 0x71, 0x7e, 0x1e, 0x54, 0x15, 0x98, 0x2d, 0x24, 0xc3,
	0xc5, 0x99, 0x56, 0x94, 0x9f, 0xeb, 0x9c, 0x98, 0x9c, 0x91, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1,
	0x11, 0x84, 0x10, 0x30, 0x9a, 0xc7, 0xc8, 0xc5, 0x1f, 0x04, 0xb5, 0x35, 0x18, 0xe2, 0x18, 0x21,
	0x4f, 0x2e, 0x21, 0xf7, 0xd4, 0x12, 0x9f, 0xc4, 0x92, 0xd4, 0xe2, 0x92, 0xe0, 0x92, 0xa2, 0xd4,
	0xc4, 0xdc, 0xcc, 0xbc, 0x74, 0x21, 0x31, 0x3d, 0x98, 0xeb, 0xf4, 0x90, 0x9d, 0x26, 0x25, 0x8e,
	0x21, 0x0e, 0x71, 0x8c, 0x12, 0x83, 0x01, 0xa3, 0x90, 0x03, 0x17, 0x27, 0xdc, 0x28, 0xb2, 0x4c,
	0x48, 0x62, 0x03, 0x07, 0x8d, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xce, 0xe1, 0x6a, 0x7a, 0x34,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RealtimeServiceClient is the client API for RealtimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RealtimeServiceClient interface {
	//
	//GetLatestStreaming
	//returns latest values, and future values for provided namespace
	GetLatestStreaming(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (RealtimeService_GetLatestStreamingClient, error)
	//
	//GetLatest
	//returns latest values for provided namespace
	GetLatest(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (RealtimeService_GetLatestClient, error)
}

type realtimeServiceClient struct {
	cc *grpc.ClientConn
}

func NewRealtimeServiceClient(cc *grpc.ClientConn) RealtimeServiceClient {
	return &realtimeServiceClient{cc}
}

func (c *realtimeServiceClient) GetLatestStreaming(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (RealtimeService_GetLatestStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RealtimeService_serviceDesc.Streams[0], "/realtime.RealtimeService/GetLatestStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &realtimeServiceGetLatestStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RealtimeService_GetLatestStreamingClient interface {
	Recv() (*RouteResponse, error)
	grpc.ClientStream
}

type realtimeServiceGetLatestStreamingClient struct {
	grpc.ClientStream
}

func (x *realtimeServiceGetLatestStreamingClient) Recv() (*RouteResponse, error) {
	m := new(RouteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *realtimeServiceClient) GetLatest(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (RealtimeService_GetLatestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RealtimeService_serviceDesc.Streams[1], "/realtime.RealtimeService/GetLatest", opts...)
	if err != nil {
		return nil, err
	}
	x := &realtimeServiceGetLatestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RealtimeService_GetLatestClient interface {
	Recv() (*RouteResponse, error)
	grpc.ClientStream
}

type realtimeServiceGetLatestClient struct {
	grpc.ClientStream
}

func (x *realtimeServiceGetLatestClient) Recv() (*RouteResponse, error) {
	m := new(RouteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RealtimeServiceServer is the server API for RealtimeService service.
type RealtimeServiceServer interface {
	//
	//GetLatestStreaming
	//returns latest values, and future values for provided namespace
	GetLatestStreaming(*RouteRequest, RealtimeService_GetLatestStreamingServer) error
	//
	//GetLatest
	//returns latest values for provided namespace
	GetLatest(*RouteRequest, RealtimeService_GetLatestServer) error
}

// UnimplementedRealtimeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRealtimeServiceServer struct {
}

func (*UnimplementedRealtimeServiceServer) GetLatestStreaming(req *RouteRequest, srv RealtimeService_GetLatestStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLatestStreaming not implemented")
}
func (*UnimplementedRealtimeServiceServer) GetLatest(req *RouteRequest, srv RealtimeService_GetLatestServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLatest not implemented")
}

func RegisterRealtimeServiceServer(s *grpc.Server, srv RealtimeServiceServer) {
	s.RegisterService(&_RealtimeService_serviceDesc, srv)
}

func _RealtimeService_GetLatestStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RouteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RealtimeServiceServer).GetLatestStreaming(m, &realtimeServiceGetLatestStreamingServer{stream})
}

type RealtimeService_GetLatestStreamingServer interface {
	Send(*RouteResponse) error
	grpc.ServerStream
}

type realtimeServiceGetLatestStreamingServer struct {
	grpc.ServerStream
}

func (x *realtimeServiceGetLatestStreamingServer) Send(m *RouteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RealtimeService_GetLatest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RouteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RealtimeServiceServer).GetLatest(m, &realtimeServiceGetLatestServer{stream})
}

type RealtimeService_GetLatestServer interface {
	Send(*RouteResponse) error
	grpc.ServerStream
}

type realtimeServiceGetLatestServer struct {
	grpc.ServerStream
}

func (x *realtimeServiceGetLatestServer) Send(m *RouteResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _RealtimeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "realtime.RealtimeService",
	HandlerType: (*RealtimeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLatestStreaming",
			Handler:       _RealtimeService_GetLatestStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetLatest",
			Handler:       _RealtimeService_GetLatest_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "realtime-service.proto",
}
