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
	// VPN name identifies requested vrf name
	VpnName string `protobuf:"bytes,2,opt,name=vpn_name,json=vpnName,proto3" json:"vpn_name,omitempty"`
	// Identifies if request sent for ipv4 prefixes in this case this field should
	// be set to true or ipv6, in this case this field should be set to false.
	Ipv4 bool `protobuf:"varint,3,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	// rt identifies requested vrf route target
	Rt                   *any.Any `protobuf:"bytes,4,opt,name=rt,proto3" json:"rt,omitempty"`
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

func (m *L3VpnRequest) GetVpnName() string {
	if m != nil {
		return m.VpnName
	}
	return ""
}

func (m *L3VpnRequest) GetIpv4() bool {
	if m != nil {
		return m.Ipv4
	}
	return false
}

func (m *L3VpnRequest) GetRt() *any.Any {
	if m != nil {
		return m.Rt
	}
	return nil
}

type VpnRTRequest struct {
	// id is a unique identificator of a gateway client.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// VPN name identifies requested vrf name
	VpnName              string   `protobuf:"bytes,2,opt,name=vpn_name,json=vpnName,proto3" json:"vpn_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VpnRTRequest) Reset()         { *m = VpnRTRequest{} }
func (m *VpnRTRequest) String() string { return proto.CompactTextString(m) }
func (*VpnRTRequest) ProtoMessage()    {}
func (*VpnRTRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4982045cb0164b85, []int{1}
}

func (m *VpnRTRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VpnRTRequest.Unmarshal(m, b)
}
func (m *VpnRTRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VpnRTRequest.Marshal(b, m, deterministic)
}
func (m *VpnRTRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VpnRTRequest.Merge(m, src)
}
func (m *VpnRTRequest) XXX_Size() int {
	return xxx_messageInfo_VpnRTRequest.Size(m)
}
func (m *VpnRTRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VpnRTRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VpnRTRequest proto.InternalMessageInfo

func (m *VpnRTRequest) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *VpnRTRequest) GetVpnName() string {
	if m != nil {
		return m.VpnName
	}
	return ""
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
	return fileDescriptor_4982045cb0164b85, []int{2}
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
	return fileDescriptor_4982045cb0164b85, []int{3}
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

type VpnRTResponse struct {
	// rt identifies requested vpn route target
	Rt                   *any.Any `protobuf:"bytes,1,opt,name=rt,proto3" json:"rt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VpnRTResponse) Reset()         { *m = VpnRTResponse{} }
func (m *VpnRTResponse) String() string { return proto.CompactTextString(m) }
func (*VpnRTResponse) ProtoMessage()    {}
func (*VpnRTResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4982045cb0164b85, []int{4}
}

func (m *VpnRTResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VpnRTResponse.Unmarshal(m, b)
}
func (m *VpnRTResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VpnRTResponse.Marshal(b, m, deterministic)
}
func (m *VpnRTResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VpnRTResponse.Merge(m, src)
}
func (m *VpnRTResponse) XXX_Size() int {
	return xxx_messageInfo_VpnRTResponse.Size(m)
}
func (m *VpnRTResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VpnRTResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VpnRTResponse proto.InternalMessageInfo

func (m *VpnRTResponse) GetRt() *any.Any {
	if m != nil {
		return m.Rt
	}
	return nil
}

func init() {
	proto.RegisterType((*L3VpnRequest)(nil), "apis.L3VpnRequest")
	proto.RegisterType((*VpnRTRequest)(nil), "apis.VpnRTRequest")
	proto.RegisterType((*Client)(nil), "apis.Client")
	proto.RegisterType((*SRv6L3Response)(nil), "apis.SRv6L3Response")
	proto.RegisterType((*VpnRTResponse)(nil), "apis.VpnRTResponse")
}

func init() { proto.RegisterFile("gateway-services.proto", fileDescriptor_4982045cb0164b85) }

var fileDescriptor_4982045cb0164b85 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xe5, 0xb4, 0x6c, 0xeb, 0x6b, 0x08, 0xc2, 0x4c, 0x55, 0x16, 0x2e, 0x51, 0xc4, 0x21,
	0x17, 0x32, 0x68, 0x58, 0x25, 0xc4, 0xa9, 0xa2, 0x13, 0x42, 0x1a, 0x53, 0xe5, 0xc1, 0xae, 0x53,
	0xba, 0xbc, 0x45, 0x96, 0x52, 0xdb, 0xc4, 0x6e, 0xa0, 0xff, 0x16, 0x7f, 0x21, 0x4a, 0x9c, 0x4e,
	0x2d, 0x15, 0x20, 0xd8, 0xed, 0xf9, 0xfd, 0xf8, 0xbe, 0xcf, 0xd7, 0x0f, 0x46, 0x45, 0x66, 0xf0,
	0x5b, 0xb6, 0x7e, 0xa9, 0xb1, 0xaa, 0xf9, 0x2d, 0xea, 0x44, 0x55, 0xd2, 0x48, 0xda, 0xcf, 0x14,
	0xd7, 0xc1, 0x49, 0x21, 0x65, 0x51, 0xe2, 0x69, 0x9b, 0x5b, 0xac, 0xee, 0x4e, 0x33, 0xb1, 0xb6,
	0x0d, 0xc1, 0xf3, 0x5f, 0x4b, 0xb8, 0x54, 0x66, 0x53, 0x7c, 0xa2, 0x2a, 0xbc, 0xe3, 0xdf, 0x35,
	0xcf, 0xbb, 0xc4, 0x60, 0x51, 0x28, 0x1b, 0x46, 0x1a, 0xdc, 0x8b, 0xf4, 0x5a, 0x09, 0x86, 0x5f,
	0x57, 0xa8, 0x0d, 0xf5, 0xc0, 0xe1, 0xb9, 0x4f, 0x42, 0x12, 0xbb, 0xcc, 0xe1, 0x39, 0x3d, 0x81,
	0xa3, 0x5a, 0x89, 0x1b, 0x91, 0x2d, 0xd1, 0x77, 0x42, 0x12, 0x0f, 0xd8, 0x61, 0xad, 0xc4, 0x65,
	0xb6, 0x44, 0x4a, 0xa1, 0xcf, 0x55, 0xfd, 0xc6, 0xef, 0x85, 0x24, 0x3e, 0x62, 0x6d, 0x4c, 0x5f,
	0x80, 0x53, 0x19, 0xbf, 0x1f, 0x92, 0x78, 0x38, 0x3e, 0x4e, 0x2c, 0x54, 0xb2, 0x81, 0x4a, 0xa6,
	0x62, 0xcd, 0x9c, 0xca, 0x44, 0x6f, 0xc1, 0x6d, 0x56, 0x7e, 0xfe, 0xf7, 0xa5, 0x91, 0x0f, 0x07,
	0xef, 0x4b, 0x8e, 0x62, 0x6f, 0x28, 0x3a, 0x07, 0xef, 0x8a, 0xd5, 0x93, 0x8b, 0x94, 0xa1, 0x56,
	0x52, 0x68, 0xa4, 0x29, 0x0c, 0x75, 0x55, 0x4f, 0x6e, 0xac, 0x7d, 0x9f, 0x84, 0xbd, 0x78, 0x38,
	0xa6, 0x49, 0xf3, 0x97, 0x89, 0x6d, 0x9d, 0xb7, 0x15, 0x06, 0x4d, 0x9b, 0x8d, 0xa3, 0x33, 0x78,
	0xdc, 0xb1, 0x75, 0x2a, 0xd6, 0x12, 0xf9, 0xb3, 0xa5, 0xf1, 0x8f, 0x1e, 0x78, 0x1f, 0xec, 0xf1,
	0xae, 0xec, 0xed, 0xe8, 0x6b, 0x38, 0xfc, 0x24, 0x05, 0x37, 0xb2, 0xa2, 0xae, 0x5d, 0x6a, 0xc9,
	0x83, 0xd1, 0x9e, 0xca, 0x79, 0x73, 0xad, 0x98, 0xd0, 0x33, 0x18, 0x58, 0xb0, 0xeb, 0xf9, 0x25,
	0xed, 0x48, 0xb7, 0xcf, 0x13, 0x1c, 0x6f, 0xd3, 0xdf, 0x23, 0xbe, 0x82, 0x47, 0x2d, 0xf3, 0x66,
	0x64, 0xfb, 0x73, 0x83, 0x67, 0x3b, 0xb9, 0x6e, 0x62, 0x06, 0x74, 0x9a, 0xe7, 0x1f, 0xe7, 0xf5,
	0xe4, 0x8b, 0xe0, 0xb7, 0x99, 0x61, 0x72, 0x65, 0x90, 0x8e, 0x6c, 0xeb, 0x7d, 0x5a, 0xdb, 0xfc,
	0xef, 0x80, 0x1b, 0x95, 0x19, 0x96, 0x0f, 0x55, 0x79, 0x07, 0xde, 0x34, 0xcf, 0x3b, 0x4b, 0xad,
	0xc2, 0xd3, 0x1d, 0x97, 0x7f, 0x1b, 0x9e, 0x61, 0xf9, 0x7f, 0xc3, 0x8b, 0x83, 0xf6, 0x9d, 0xfe,
	0x0c, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x56, 0xfc, 0xa9, 0x77, 0x03, 0x00, 0x00,
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
	SRv6L3VPN(ctx context.Context, in *L3VpnRequest, opts ...grpc.CallOption) (*SRv6L3Response, error)
	VpnRT(ctx context.Context, in *VpnRTRequest, opts ...grpc.CallOption) (*VpnRTResponse, error)
	AddIPv6UnicatRoute(ctx context.Context, in *IPv6UnicastRoute, opts ...grpc.CallOption) (*empty.Empty, error)
	DelIPv6UnicatRoute(ctx context.Context, in *IPv6UnicastRoute, opts ...grpc.CallOption) (*empty.Empty, error)
	AddSRv6L3Route(ctx context.Context, in *SRv6L3Route, opts ...grpc.CallOption) (*empty.Empty, error)
	DelSRv6L3Route(ctx context.Context, in *SRv6L3Route, opts ...grpc.CallOption) (*empty.Empty, error)
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

func (c *gatewayServiceClient) SRv6L3VPN(ctx context.Context, in *L3VpnRequest, opts ...grpc.CallOption) (*SRv6L3Response, error) {
	out := new(SRv6L3Response)
	err := c.cc.Invoke(ctx, "/apis.GatewayService/SRv6L3VPN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) VpnRT(ctx context.Context, in *VpnRTRequest, opts ...grpc.CallOption) (*VpnRTResponse, error) {
	out := new(VpnRTResponse)
	err := c.cc.Invoke(ctx, "/apis.GatewayService/VpnRT", in, out, opts...)
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

func (c *gatewayServiceClient) AddSRv6L3Route(ctx context.Context, in *SRv6L3Route, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/apis.GatewayService/AddSRv6L3Route", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) DelSRv6L3Route(ctx context.Context, in *SRv6L3Route, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/apis.GatewayService/DelSRv6L3Route", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServiceServer is the server API for GatewayService service.
type GatewayServiceServer interface {
	Monitor(GatewayService_MonitorServer) error
	SRv6L3VPN(context.Context, *L3VpnRequest) (*SRv6L3Response, error)
	VpnRT(context.Context, *VpnRTRequest) (*VpnRTResponse, error)
	AddIPv6UnicatRoute(context.Context, *IPv6UnicastRoute) (*empty.Empty, error)
	DelIPv6UnicatRoute(context.Context, *IPv6UnicastRoute) (*empty.Empty, error)
	AddSRv6L3Route(context.Context, *SRv6L3Route) (*empty.Empty, error)
	DelSRv6L3Route(context.Context, *SRv6L3Route) (*empty.Empty, error)
}

// UnimplementedGatewayServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayServiceServer struct {
}

func (*UnimplementedGatewayServiceServer) Monitor(srv GatewayService_MonitorServer) error {
	return status.Errorf(codes.Unimplemented, "method Monitor not implemented")
}
func (*UnimplementedGatewayServiceServer) SRv6L3VPN(ctx context.Context, req *L3VpnRequest) (*SRv6L3Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SRv6L3VPN not implemented")
}
func (*UnimplementedGatewayServiceServer) VpnRT(ctx context.Context, req *VpnRTRequest) (*VpnRTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VpnRT not implemented")
}
func (*UnimplementedGatewayServiceServer) AddIPv6UnicatRoute(ctx context.Context, req *IPv6UnicastRoute) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddIPv6UnicatRoute not implemented")
}
func (*UnimplementedGatewayServiceServer) DelIPv6UnicatRoute(ctx context.Context, req *IPv6UnicastRoute) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelIPv6UnicatRoute not implemented")
}
func (*UnimplementedGatewayServiceServer) AddSRv6L3Route(ctx context.Context, req *SRv6L3Route) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSRv6L3Route not implemented")
}
func (*UnimplementedGatewayServiceServer) DelSRv6L3Route(ctx context.Context, req *SRv6L3Route) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelSRv6L3Route not implemented")
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

func _GatewayService_VpnRT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VpnRTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).VpnRT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.GatewayService/VpnRT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).VpnRT(ctx, req.(*VpnRTRequest))
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

func _GatewayService_AddSRv6L3Route_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SRv6L3Route)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).AddSRv6L3Route(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.GatewayService/AddSRv6L3Route",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).AddSRv6L3Route(ctx, req.(*SRv6L3Route))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_DelSRv6L3Route_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SRv6L3Route)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).DelSRv6L3Route(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.GatewayService/DelSRv6L3Route",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).DelSRv6L3Route(ctx, req.(*SRv6L3Route))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apis.GatewayService",
	HandlerType: (*GatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SRv6L3VPN",
			Handler:    _GatewayService_SRv6L3VPN_Handler,
		},
		{
			MethodName: "VpnRT",
			Handler:    _GatewayService_VpnRT_Handler,
		},
		{
			MethodName: "AddIPv6UnicatRoute",
			Handler:    _GatewayService_AddIPv6UnicatRoute_Handler,
		},
		{
			MethodName: "DelIPv6UnicatRoute",
			Handler:    _GatewayService_DelIPv6UnicatRoute_Handler,
		},
		{
			MethodName: "AddSRv6L3Route",
			Handler:    _GatewayService_AddSRv6L3Route_Handler,
		},
		{
			MethodName: "DelSRv6L3Route",
			Handler:    _GatewayService_DelSRv6L3Route_Handler,
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
