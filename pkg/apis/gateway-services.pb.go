// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway-services.proto

package apis

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// L3VpnRequest defines a structure of a request for the prefixes
// which belons to a specific VPN, identified by the route distingusiher.
// Further selection criteria includes: family ipv6/ipv6, a list of route
// targets or a specific vpn prefix.
type L3VpnRequest struct {
	// id is a unique identificator of a gateway client.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// RD identifies the vpn route distinguisher, it is a mandatory parameter
	Rd *any.Any `protobuf:"bytes,2,opt,name=rd,proto3" json:"rd,omitempty"`
	// Identifies if request sent for ipv4 prefixes in this case this field should
	// be set to true or ipv6, in this case this field should be set to false.
	Ipv4 bool `protobuf:"varint,3,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	// Route Targets is optional filtering parameter.
	Rt []*any.Any `protobuf:"bytes,4,rep,name=rt,proto3" json:"rt,omitempty"`
	// Prefix is optional filtering parameter.
	VpnPrefix            *Prefix  `protobuf:"bytes,5,opt,name=vpn_prefix,json=vpnPrefix,proto3" json:"vpn_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L3VpnRequest) Reset()         { *m = L3VpnRequest{} }
func (m *L3VpnRequest) String() string { return proto.CompactTextString(m) }
func (*L3VpnRequest) ProtoMessage()    {}
func (*L3VpnRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4982045cb0164b85, []int{0}
}

func (m *L3VpnRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L3VpnRequest.Unmarshal(m, b)
}
func (m *L3VpnRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L3VpnRequest.Marshal(b, m, deterministic)
}
func (m *L3VpnRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L3VpnRequest.Merge(m, src)
}
func (m *L3VpnRequest) XXX_Size() int {
	return xxx_messageInfo_L3VpnRequest.Size(m)
}
func (m *L3VpnRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_L3VpnRequest.DiscardUnknown(m)
}

var xxx_messageInfo_L3VpnRequest proto.InternalMessageInfo

func (m *L3VpnRequest) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *L3VpnRequest) GetRd() *any.Any {
	if m != nil {
		return m.Rd
	}
	return nil
}

func (m *L3VpnRequest) GetIpv4() bool {
	if m != nil {
		return m.Ipv4
	}
	return false
}

func (m *L3VpnRequest) GetRt() []*any.Any {
	if m != nil {
		return m.Rt
	}
	return nil
}

func (m *L3VpnRequest) GetVpnPrefix() *Prefix {
	if m != nil {
		return m.VpnPrefix
	}
	return nil
}

type Client struct {
	// id is a unique identificator of a gateway client.
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_4982045cb0164b85, []int{1}
}

func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (m *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(m, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type SRv6L3Response struct {
	Srv6Prefix           []*SRv6L3Prefix `protobuf:"bytes,1,rep,name=srv6_prefix,json=srv6Prefix,proto3" json:"srv6_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SRv6L3Response) Reset()         { *m = SRv6L3Response{} }
func (m *SRv6L3Response) String() string { return proto.CompactTextString(m) }
func (*SRv6L3Response) ProtoMessage()    {}
func (*SRv6L3Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_4982045cb0164b85, []int{2}
}

func (m *SRv6L3Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SRv6L3Response.Unmarshal(m, b)
}
func (m *SRv6L3Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SRv6L3Response.Marshal(b, m, deterministic)
}
func (m *SRv6L3Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SRv6L3Response.Merge(m, src)
}
func (m *SRv6L3Response) XXX_Size() int {
	return xxx_messageInfo_SRv6L3Response.Size(m)
}
func (m *SRv6L3Response) XXX_DiscardUnknown() {
	xxx_messageInfo_SRv6L3Response.DiscardUnknown(m)
}

var xxx_messageInfo_SRv6L3Response proto.InternalMessageInfo

func (m *SRv6L3Response) GetSrv6Prefix() []*SRv6L3Prefix {
	if m != nil {
		return m.Srv6Prefix
	}
	return nil
}

type MPLSL3Response struct {
	MplsPrefix           []*MPLSL3Prefix `protobuf:"bytes,1,rep,name=mpls_prefix,json=mplsPrefix,proto3" json:"mpls_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MPLSL3Response) Reset()         { *m = MPLSL3Response{} }
func (m *MPLSL3Response) String() string { return proto.CompactTextString(m) }
func (*MPLSL3Response) ProtoMessage()    {}
func (*MPLSL3Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_4982045cb0164b85, []int{3}
}

func (m *MPLSL3Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MPLSL3Response.Unmarshal(m, b)
}
func (m *MPLSL3Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MPLSL3Response.Marshal(b, m, deterministic)
}
func (m *MPLSL3Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MPLSL3Response.Merge(m, src)
}
func (m *MPLSL3Response) XXX_Size() int {
	return xxx_messageInfo_MPLSL3Response.Size(m)
}
func (m *MPLSL3Response) XXX_DiscardUnknown() {
	xxx_messageInfo_MPLSL3Response.DiscardUnknown(m)
}

var xxx_messageInfo_MPLSL3Response proto.InternalMessageInfo

func (m *MPLSL3Response) GetMplsPrefix() []*MPLSL3Prefix {
	if m != nil {
		return m.MplsPrefix
	}
	return nil
}

func init() {
	proto.RegisterType((*L3VpnRequest)(nil), "apis.L3VpnRequest")
	proto.RegisterType((*Client)(nil), "apis.Client")
	proto.RegisterType((*SRv6L3Response)(nil), "apis.SRv6L3Response")
	proto.RegisterType((*MPLSL3Response)(nil), "apis.MPLSL3Response")
}

func init() { proto.RegisterFile("gateway-services.proto", fileDescriptor_4982045cb0164b85) }

var fileDescriptor_4982045cb0164b85 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x51, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x25, 0xd9, 0xba, 0x9a, 0xbb, 0x25, 0xc2, 0xb0, 0x94, 0x18, 0x5f, 0x42, 0xf1, 0x21, 0x20,
	0x66, 0x71, 0xe3, 0xf6, 0xbd, 0xd8, 0x22, 0x42, 0x2b, 0x61, 0x8a, 0x7d, 0x95, 0xa4, 0x99, 0x86,
	0x81, 0x74, 0x66, 0xcc, 0x4c, 0xa3, 0xfd, 0x4d, 0xfe, 0x06, 0xff, 0x9b, 0x24, 0x37, 0x29, 0xb5,
	0xf5, 0xe3, 0x61, 0xdf, 0xce, 0xdc, 0x39, 0xf7, 0x9c, 0x73, 0xef, 0x85, 0x51, 0x91, 0x1a, 0xf6,
	0x2d, 0x3d, 0xbc, 0xd1, 0xac, 0xaa, 0xf9, 0x86, 0xe9, 0x48, 0x55, 0xd2, 0x48, 0x32, 0x48, 0x15,
	0xd7, 0xfe, 0x8b, 0x42, 0xca, 0xa2, 0x64, 0x77, 0x6d, 0x2d, 0xdb, 0x6f, 0xef, 0x52, 0x71, 0x40,
	0x82, 0xff, 0xf2, 0xfc, 0x8b, 0xed, 0x94, 0xe9, 0x3f, 0x21, 0x4b, 0x35, 0xeb, 0xf0, 0x73, 0x55,
	0xb1, 0x2d, 0xff, 0xae, 0x79, 0xde, 0x7f, 0xee, 0x54, 0xd9, 0xd9, 0xf8, 0x4e, 0x56, 0x28, 0x84,
	0xe3, 0x1f, 0x16, 0x0c, 0x17, 0xf1, 0x5a, 0x09, 0xca, 0xbe, 0xee, 0x99, 0x36, 0xc4, 0x05, 0x9b,
	0xe7, 0x9e, 0x15, 0x58, 0xe1, 0x90, 0xda, 0x3c, 0x27, 0xaf, 0xc0, 0xae, 0x72, 0xcf, 0x0e, 0xac,
	0xf0, 0xe6, 0xfe, 0x36, 0x42, 0xfb, 0xa8, 0xb7, 0x8f, 0xa6, 0xe2, 0x40, 0xed, 0x2a, 0x27, 0x04,
	0x06, 0x5c, 0xd5, 0xef, 0xbc, 0xab, 0xc0, 0x0a, 0x9f, 0xd1, 0x16, 0xb7, 0x9d, 0xc6, 0x1b, 0x04,
	0x57, 0xff, 0xe8, 0x34, 0xe4, 0x35, 0x40, 0xad, 0xc4, 0x17, 0x8c, 0xeb, 0x3d, 0x69, 0x7d, 0x86,
	0x51, 0xb3, 0x87, 0x28, 0x69, 0x6b, 0xd4, 0xa9, 0x95, 0x40, 0x38, 0xf6, 0xe0, 0xfa, 0x7d, 0xc9,
	0x99, 0xb8, 0x88, 0x39, 0x9e, 0x83, 0xbb, 0xa2, 0xf5, 0x64, 0x11, 0x53, 0xa6, 0x95, 0x14, 0x9a,
	0x91, 0x18, 0x6e, 0x74, 0x55, 0x4f, 0x7a, 0x65, 0xab, 0xcd, 0x41, 0x50, 0x19, 0xa9, 0x9d, 0x3e,
	0x34, 0xb4, 0xce, 0x60, 0x0e, 0xee, 0x32, 0x59, 0xac, 0x7e, 0x97, 0x69, 0x36, 0xf7, 0x47, 0x19,
	0xa4, 0xf6, 0x32, 0x0d, 0x0d, 0xf1, 0xfd, 0x4f, 0x1b, 0xdc, 0x0f, 0x78, 0xe2, 0x15, 0x5e, 0x98,
	0xbc, 0x85, 0xa7, 0x4b, 0x29, 0xb8, 0x91, 0x15, 0xe9, 0xc6, 0xc3, 0x49, 0xfc, 0xd1, 0xc5, 0x6a,
	0xe6, 0xcd, 0x4d, 0x43, 0x8b, 0x3c, 0x80, 0x83, 0x0e, 0xeb, 0xe4, 0x13, 0xe9, 0x2c, 0x4f, 0x6f,
	0xe5, 0xdf, 0x9e, 0xc6, 0x38, 0x26, 0x7e, 0x00, 0x07, 0xe7, 0xfb, 0x4f, 0xdb, 0xd9, 0xbe, 0x66,
	0x40, 0xa6, 0x79, 0xfe, 0x31, 0xa9, 0x27, 0x9f, 0x05, 0xdf, 0xa4, 0x86, 0xca, 0xbd, 0x61, 0x64,
	0x84, 0xdc, 0x63, 0x59, 0x63, 0xfd, 0x6f, 0xa9, 0x1b, 0x95, 0x19, 0x2b, 0x1f, 0xa9, 0x92, 0x5d,
	0xb7, 0xef, 0xf8, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x9e, 0x3b, 0xac, 0x28, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GatewayServiceClient is the client API for GatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayServiceClient interface {
	Monitor(ctx context.Context, opts ...grpc.CallOption) (GatewayService_MonitorClient, error)
	MPLSL3VPN(ctx context.Context, in *L3VpnRequest, opts ...grpc.CallOption) (*MPLSL3Response, error)
	SRv6L3VPN(ctx context.Context, in *L3VpnRequest, opts ...grpc.CallOption) (*SRv6L3Response, error)
	AddIPv6UnicatRoute(ctx context.Context, in *IPv6UnicastRoute, opts ...grpc.CallOption) (*empty.Empty, error)
	DelIPv6UnicatRoute(ctx context.Context, in *IPv6UnicastRoute, opts ...grpc.CallOption) (*empty.Empty, error)
}

type gatewayServiceClient struct {
	cc *grpc.ClientConn
}

func NewGatewayServiceClient(cc *grpc.ClientConn) GatewayServiceClient {
	return &gatewayServiceClient{cc}
}

func (c *gatewayServiceClient) Monitor(ctx context.Context, opts ...grpc.CallOption) (GatewayService_MonitorClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GatewayService_serviceDesc.Streams[0], "/apis.GatewayService/Monitor", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewayServiceMonitorClient{stream}
	return x, nil
}

type GatewayService_MonitorClient interface {
	Send(*Client) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type gatewayServiceMonitorClient struct {
	grpc.ClientStream
}

func (x *gatewayServiceMonitorClient) Send(m *Client) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gatewayServiceMonitorClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gatewayServiceClient) MPLSL3VPN(ctx context.Context, in *L3VpnRequest, opts ...grpc.CallOption) (*MPLSL3Response, error) {
	out := new(MPLSL3Response)
	err := c.cc.Invoke(ctx, "/apis.GatewayService/MPLSL3VPN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) SRv6L3VPN(ctx context.Context, in *L3VpnRequest, opts ...grpc.CallOption) (*SRv6L3Response, error) {
	out := new(SRv6L3Response)
	err := c.cc.Invoke(ctx, "/apis.GatewayService/SRv6L3VPN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) AddIPv6UnicatRoute(ctx context.Context, in *IPv6UnicastRoute, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/apis.GatewayService/AddIPv6UnicatRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) DelIPv6UnicatRoute(ctx context.Context, in *IPv6UnicastRoute, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/apis.GatewayService/DelIPv6UnicatRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServiceServer is the server API for GatewayService service.
type GatewayServiceServer interface {
	Monitor(GatewayService_MonitorServer) error
	MPLSL3VPN(context.Context, *L3VpnRequest) (*MPLSL3Response, error)
	SRv6L3VPN(context.Context, *L3VpnRequest) (*SRv6L3Response, error)
	AddIPv6UnicatRoute(context.Context, *IPv6UnicastRoute) (*empty.Empty, error)
	DelIPv6UnicatRoute(context.Context, *IPv6UnicastRoute) (*empty.Empty, error)
}

// UnimplementedGatewayServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayServiceServer struct {
}

func (*UnimplementedGatewayServiceServer) Monitor(srv GatewayService_MonitorServer) error {
	return status.Errorf(codes.Unimplemented, "method Monitor not implemented")
}
func (*UnimplementedGatewayServiceServer) MPLSL3VPN(ctx context.Context, req *L3VpnRequest) (*MPLSL3Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MPLSL3VPN not implemented")
}
func (*UnimplementedGatewayServiceServer) SRv6L3VPN(ctx context.Context, req *L3VpnRequest) (*SRv6L3Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SRv6L3VPN not implemented")
}
func (*UnimplementedGatewayServiceServer) AddIPv6UnicatRoute(ctx context.Context, req *IPv6UnicastRoute) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddIPv6UnicatRoute not implemented")
}
func (*UnimplementedGatewayServiceServer) DelIPv6UnicatRoute(ctx context.Context, req *IPv6UnicastRoute) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelIPv6UnicatRoute not implemented")
}

func RegisterGatewayServiceServer(s *grpc.Server, srv GatewayServiceServer) {
	s.RegisterService(&_GatewayService_serviceDesc, srv)
}

func _GatewayService_Monitor_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GatewayServiceServer).Monitor(&gatewayServiceMonitorServer{stream})
}

type GatewayService_MonitorServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*Client, error)
	grpc.ServerStream
}

type gatewayServiceMonitorServer struct {
	grpc.ServerStream
}

func (x *gatewayServiceMonitorServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gatewayServiceMonitorServer) Recv() (*Client, error) {
	m := new(Client)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GatewayService_MPLSL3VPN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(L3VpnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).MPLSL3VPN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.GatewayService/MPLSL3VPN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).MPLSL3VPN(ctx, req.(*L3VpnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_SRv6L3VPN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(L3VpnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).SRv6L3VPN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.GatewayService/SRv6L3VPN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).SRv6L3VPN(ctx, req.(*L3VpnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_AddIPv6UnicatRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPv6UnicastRoute)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).AddIPv6UnicatRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.GatewayService/AddIPv6UnicatRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).AddIPv6UnicatRoute(ctx, req.(*IPv6UnicastRoute))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_DelIPv6UnicatRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPv6UnicastRoute)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).DelIPv6UnicatRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.GatewayService/DelIPv6UnicatRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).DelIPv6UnicatRoute(ctx, req.(*IPv6UnicastRoute))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apis.GatewayService",
	HandlerType: (*GatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MPLSL3VPN",
			Handler:    _GatewayService_MPLSL3VPN_Handler,
		},
		{
			MethodName: "SRv6L3VPN",
			Handler:    _GatewayService_SRv6L3VPN_Handler,
		},
		{
			MethodName: "AddIPv6UnicatRoute",
			Handler:    _GatewayService_AddIPv6UnicatRoute_Handler,
		},
		{
			MethodName: "DelIPv6UnicatRoute",
			Handler:    _GatewayService_DelIPv6UnicatRoute_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Monitor",
			Handler:       _GatewayService_Monitor_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "gateway-services.proto",
}
