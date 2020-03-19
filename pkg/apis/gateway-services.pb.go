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

// Errors defines error codes which a gateway can return to a client
type GatewayErrors int32

const (
	GatewayErrors_OK GatewayErrors = 0
	// Returned when a gateway encountered a generic error
	GatewayErrors_EIO GatewayErrors = 1
	// Returned when request key does not exist
	GatewayErrors_ENOENT GatewayErrors = 2
	// Returned when an operation triggered by the client's requested timed out
	// and was canceled
	GatewayErrors_ETIMEDOUT GatewayErrors = 3
	// Returned when a gateway cannot reach a DB host
	GatewayErrors_EHOSTDOWN GatewayErrors = 4
)

var GatewayErrors_name = map[int32]string{
	0: "OK",
	1: "EIO",
	2: "ENOENT",
	3: "ETIMEDOUT",
	4: "EHOSTDOWN",
}

var GatewayErrors_value = map[string]int32{
	"OK":        0,
	"EIO":       1,
	"ENOENT":    2,
	"ETIMEDOUT": 3,
	"EHOSTDOWN": 4,
}

func (x GatewayErrors) String() string {
	return proto.EnumName(GatewayErrors_name, int32(x))
}

func (GatewayErrors) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4982045cb0164b85, []int{0}
}

type RouteDistinguisherTwoOctetAS struct {
	Admin                uint32   `protobuf:"varint,1,opt,name=admin,proto3" json:"admin,omitempty"`
	Assigned             uint32   `protobuf:"varint,2,opt,name=assigned,proto3" json:"assigned,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteDistinguisherTwoOctetAS) Reset()         { *m = RouteDistinguisherTwoOctetAS{} }
func (m *RouteDistinguisherTwoOctetAS) String() string { return proto.CompactTextString(m) }
func (*RouteDistinguisherTwoOctetAS) ProtoMessage()    {}
func (*RouteDistinguisherTwoOctetAS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4982045cb0164b85, []int{0}
}

func (m *RouteDistinguisherTwoOctetAS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteDistinguisherTwoOctetAS.Unmarshal(m, b)
}
func (m *RouteDistinguisherTwoOctetAS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteDistinguisherTwoOctetAS.Marshal(b, m, deterministic)
}
func (m *RouteDistinguisherTwoOctetAS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteDistinguisherTwoOctetAS.Merge(m, src)
}
func (m *RouteDistinguisherTwoOctetAS) XXX_Size() int {
	return xxx_messageInfo_RouteDistinguisherTwoOctetAS.Size(m)
}
func (m *RouteDistinguisherTwoOctetAS) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteDistinguisherTwoOctetAS.DiscardUnknown(m)
}

var xxx_messageInfo_RouteDistinguisherTwoOctetAS proto.InternalMessageInfo

func (m *RouteDistinguisherTwoOctetAS) GetAdmin() uint32 {
	if m != nil {
		return m.Admin
	}
	return 0
}

func (m *RouteDistinguisherTwoOctetAS) GetAssigned() uint32 {
	if m != nil {
		return m.Assigned
	}
	return 0
}

type RouteDistinguisherIPAddress struct {
	Admin                string   `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	Assigned             uint32   `protobuf:"varint,2,opt,name=assigned,proto3" json:"assigned,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteDistinguisherIPAddress) Reset()         { *m = RouteDistinguisherIPAddress{} }
func (m *RouteDistinguisherIPAddress) String() string { return proto.CompactTextString(m) }
func (*RouteDistinguisherIPAddress) ProtoMessage()    {}
func (*RouteDistinguisherIPAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_4982045cb0164b85, []int{1}
}

func (m *RouteDistinguisherIPAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteDistinguisherIPAddress.Unmarshal(m, b)
}
func (m *RouteDistinguisherIPAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteDistinguisherIPAddress.Marshal(b, m, deterministic)
}
func (m *RouteDistinguisherIPAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteDistinguisherIPAddress.Merge(m, src)
}
func (m *RouteDistinguisherIPAddress) XXX_Size() int {
	return xxx_messageInfo_RouteDistinguisherIPAddress.Size(m)
}
func (m *RouteDistinguisherIPAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteDistinguisherIPAddress.DiscardUnknown(m)
}

var xxx_messageInfo_RouteDistinguisherIPAddress proto.InternalMessageInfo

func (m *RouteDistinguisherIPAddress) GetAdmin() string {
	if m != nil {
		return m.Admin
	}
	return ""
}

func (m *RouteDistinguisherIPAddress) GetAssigned() uint32 {
	if m != nil {
		return m.Assigned
	}
	return 0
}

type RouteDistinguisherFourOctetAS struct {
	Admin                uint32   `protobuf:"varint,1,opt,name=admin,proto3" json:"admin,omitempty"`
	Assigned             uint32   `protobuf:"varint,2,opt,name=assigned,proto3" json:"assigned,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteDistinguisherFourOctetAS) Reset()         { *m = RouteDistinguisherFourOctetAS{} }
func (m *RouteDistinguisherFourOctetAS) String() string { return proto.CompactTextString(m) }
func (*RouteDistinguisherFourOctetAS) ProtoMessage()    {}
func (*RouteDistinguisherFourOctetAS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4982045cb0164b85, []int{2}
}

func (m *RouteDistinguisherFourOctetAS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteDistinguisherFourOctetAS.Unmarshal(m, b)
}
func (m *RouteDistinguisherFourOctetAS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteDistinguisherFourOctetAS.Marshal(b, m, deterministic)
}
func (m *RouteDistinguisherFourOctetAS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteDistinguisherFourOctetAS.Merge(m, src)
}
func (m *RouteDistinguisherFourOctetAS) XXX_Size() int {
	return xxx_messageInfo_RouteDistinguisherFourOctetAS.Size(m)
}
func (m *RouteDistinguisherFourOctetAS) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteDistinguisherFourOctetAS.DiscardUnknown(m)
}

var xxx_messageInfo_RouteDistinguisherFourOctetAS proto.InternalMessageInfo

func (m *RouteDistinguisherFourOctetAS) GetAdmin() uint32 {
	if m != nil {
		return m.Admin
	}
	return 0
}

func (m *RouteDistinguisherFourOctetAS) GetAssigned() uint32 {
	if m != nil {
		return m.Assigned
	}
	return 0
}

type TwoOctetAsSpecificExtended struct {
	IsTransitive         bool     `protobuf:"varint,1,opt,name=is_transitive,json=isTransitive,proto3" json:"is_transitive,omitempty"`
	SubType              uint32   `protobuf:"varint,2,opt,name=sub_type,json=subType,proto3" json:"sub_type,omitempty"`
	As                   uint32   `protobuf:"varint,3,opt,name=as,proto3" json:"as,omitempty"`
	LocalAdmin           uint32   `protobuf:"varint,4,opt,name=local_admin,json=localAdmin,proto3" json:"local_admin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TwoOctetAsSpecificExtended) Reset()         { *m = TwoOctetAsSpecificExtended{} }
func (m *TwoOctetAsSpecificExtended) String() string { return proto.CompactTextString(m) }
func (*TwoOctetAsSpecificExtended) ProtoMessage()    {}
func (*TwoOctetAsSpecificExtended) Descriptor() ([]byte, []int) {
	return fileDescriptor_4982045cb0164b85, []int{3}
}

func (m *TwoOctetAsSpecificExtended) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TwoOctetAsSpecificExtended.Unmarshal(m, b)
}
func (m *TwoOctetAsSpecificExtended) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TwoOctetAsSpecificExtended.Marshal(b, m, deterministic)
}
func (m *TwoOctetAsSpecificExtended) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TwoOctetAsSpecificExtended.Merge(m, src)
}
func (m *TwoOctetAsSpecificExtended) XXX_Size() int {
	return xxx_messageInfo_TwoOctetAsSpecificExtended.Size(m)
}
func (m *TwoOctetAsSpecificExtended) XXX_DiscardUnknown() {
	xxx_messageInfo_TwoOctetAsSpecificExtended.DiscardUnknown(m)
}

var xxx_messageInfo_TwoOctetAsSpecificExtended proto.InternalMessageInfo

func (m *TwoOctetAsSpecificExtended) GetIsTransitive() bool {
	if m != nil {
		return m.IsTransitive
	}
	return false
}

func (m *TwoOctetAsSpecificExtended) GetSubType() uint32 {
	if m != nil {
		return m.SubType
	}
	return 0
}

func (m *TwoOctetAsSpecificExtended) GetAs() uint32 {
	if m != nil {
		return m.As
	}
	return 0
}

func (m *TwoOctetAsSpecificExtended) GetLocalAdmin() uint32 {
	if m != nil {
		return m.LocalAdmin
	}
	return 0
}

type IPv4AddressSpecificExtended struct {
	IsTransitive         bool     `protobuf:"varint,1,opt,name=is_transitive,json=isTransitive,proto3" json:"is_transitive,omitempty"`
	SubType              uint32   `protobuf:"varint,2,opt,name=sub_type,json=subType,proto3" json:"sub_type,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	LocalAdmin           uint32   `protobuf:"varint,4,opt,name=local_admin,json=localAdmin,proto3" json:"local_admin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPv4AddressSpecificExtended) Reset()         { *m = IPv4AddressSpecificExtended{} }
func (m *IPv4AddressSpecificExtended) String() string { return proto.CompactTextString(m) }
func (*IPv4AddressSpecificExtended) ProtoMessage()    {}
func (*IPv4AddressSpecificExtended) Descriptor() ([]byte, []int) {
	return fileDescriptor_4982045cb0164b85, []int{4}
}

func (m *IPv4AddressSpecificExtended) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPv4AddressSpecificExtended.Unmarshal(m, b)
}
func (m *IPv4AddressSpecificExtended) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPv4AddressSpecificExtended.Marshal(b, m, deterministic)
}
func (m *IPv4AddressSpecificExtended) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPv4AddressSpecificExtended.Merge(m, src)
}
func (m *IPv4AddressSpecificExtended) XXX_Size() int {
	return xxx_messageInfo_IPv4AddressSpecificExtended.Size(m)
}
func (m *IPv4AddressSpecificExtended) XXX_DiscardUnknown() {
	xxx_messageInfo_IPv4AddressSpecificExtended.DiscardUnknown(m)
}

var xxx_messageInfo_IPv4AddressSpecificExtended proto.InternalMessageInfo

func (m *IPv4AddressSpecificExtended) GetIsTransitive() bool {
	if m != nil {
		return m.IsTransitive
	}
	return false
}

func (m *IPv4AddressSpecificExtended) GetSubType() uint32 {
	if m != nil {
		return m.SubType
	}
	return 0
}

func (m *IPv4AddressSpecificExtended) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IPv4AddressSpecificExtended) GetLocalAdmin() uint32 {
	if m != nil {
		return m.LocalAdmin
	}
	return 0
}

type FourOctetAsSpecificExtended struct {
	IsTransitive         bool     `protobuf:"varint,1,opt,name=is_transitive,json=isTransitive,proto3" json:"is_transitive,omitempty"`
	SubType              uint32   `protobuf:"varint,2,opt,name=sub_type,json=subType,proto3" json:"sub_type,omitempty"`
	As                   uint32   `protobuf:"varint,3,opt,name=as,proto3" json:"as,omitempty"`
	LocalAdmin           uint32   `protobuf:"varint,4,opt,name=local_admin,json=localAdmin,proto3" json:"local_admin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FourOctetAsSpecificExtended) Reset()         { *m = FourOctetAsSpecificExtended{} }
func (m *FourOctetAsSpecificExtended) String() string { return proto.CompactTextString(m) }
func (*FourOctetAsSpecificExtended) ProtoMessage()    {}
func (*FourOctetAsSpecificExtended) Descriptor() ([]byte, []int) {
	return fileDescriptor_4982045cb0164b85, []int{5}
}

func (m *FourOctetAsSpecificExtended) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FourOctetAsSpecificExtended.Unmarshal(m, b)
}
func (m *FourOctetAsSpecificExtended) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FourOctetAsSpecificExtended.Marshal(b, m, deterministic)
}
func (m *FourOctetAsSpecificExtended) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FourOctetAsSpecificExtended.Merge(m, src)
}
func (m *FourOctetAsSpecificExtended) XXX_Size() int {
	return xxx_messageInfo_FourOctetAsSpecificExtended.Size(m)
}
func (m *FourOctetAsSpecificExtended) XXX_DiscardUnknown() {
	xxx_messageInfo_FourOctetAsSpecificExtended.DiscardUnknown(m)
}

var xxx_messageInfo_FourOctetAsSpecificExtended proto.InternalMessageInfo

func (m *FourOctetAsSpecificExtended) GetIsTransitive() bool {
	if m != nil {
		return m.IsTransitive
	}
	return false
}

func (m *FourOctetAsSpecificExtended) GetSubType() uint32 {
	if m != nil {
		return m.SubType
	}
	return 0
}

func (m *FourOctetAsSpecificExtended) GetAs() uint32 {
	if m != nil {
		return m.As
	}
	return 0
}

func (m *FourOctetAsSpecificExtended) GetLocalAdmin() uint32 {
	if m != nil {
		return m.LocalAdmin
	}
	return 0
}

type Prefix struct {
	Address    []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	MaskLength uint32 `protobuf:"varint,2,opt,name=mask_length,json=maskLength,proto3" json:"mask_length,omitempty"`
	// vpn label assigned to the prefix by the originating provider edge router
	VpnLabel uint32 `protobuf:"varint,3,opt,name=vpn_label,json=vpnLabel,proto3" json:"vpn_label,omitempty"`
	// sid label is an identifier used for a provider edge router identification
	SidLabel uint32 `protobuf:"varint,4,opt,name=sid_label,json=sidLabel,proto3" json:"sid_label,omitempty"`
	// Route Distinguisher must be one of
	// RouteDistinguisherTwoOctetAS,
	// RouteDistinguisherIPAddressAS,
	// or RouteDistinguisherFourOctetAS.
	// Mandatory parameter
	Rd *any.Any `protobuf:"bytes,5,opt,name=rd,proto3" json:"rd,omitempty"`
	// List of the a specific prefix's Route Targets. Each must be one of
	// TwoOctetAsSpecificExtended,
	// IPv4AddressSpecificExtended,
	// or FourOctetAsSpecificExtended.
	Rt                   []*any.Any `protobuf:"bytes,6,rep,name=rt,proto3" json:"rt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Prefix) Reset()         { *m = Prefix{} }
func (m *Prefix) String() string { return proto.CompactTextString(m) }
func (*Prefix) ProtoMessage()    {}
func (*Prefix) Descriptor() ([]byte, []int) {
	return fileDescriptor_4982045cb0164b85, []int{6}
}

func (m *Prefix) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prefix.Unmarshal(m, b)
}
func (m *Prefix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prefix.Marshal(b, m, deterministic)
}
func (m *Prefix) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prefix.Merge(m, src)
}
func (m *Prefix) XXX_Size() int {
	return xxx_messageInfo_Prefix.Size(m)
}
func (m *Prefix) XXX_DiscardUnknown() {
	xxx_messageInfo_Prefix.DiscardUnknown(m)
}

var xxx_messageInfo_Prefix proto.InternalMessageInfo

func (m *Prefix) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Prefix) GetMaskLength() uint32 {
	if m != nil {
		return m.MaskLength
	}
	return 0
}

func (m *Prefix) GetVpnLabel() uint32 {
	if m != nil {
		return m.VpnLabel
	}
	return 0
}

func (m *Prefix) GetSidLabel() uint32 {
	if m != nil {
		return m.SidLabel
	}
	return 0
}

func (m *Prefix) GetRd() *any.Any {
	if m != nil {
		return m.Rd
	}
	return nil
}

func (m *Prefix) GetRt() []*any.Any {
	if m != nil {
		return m.Rt
	}
	return nil
}

// RequestVPN call used to request L3 VPN entries, identified by one Route
// Distinguisher which can be one of listed below types, and one or more Route
// Targets.
type L3VPNRequest struct {
	// Route Distinguisher must be one of
	// RouteDistinguisherTwoOctetAS,
	// RouteDistinguisherIPAddressAS,
	// or RouteDistinguisherFourOctetAS.
	// Mandatory parameter
	Rd *any.Any `protobuf:"bytes,1,opt,name=rd,proto3" json:"rd,omitempty"`
	// Identifies if request sent for ipv4 prefixes in this case this field should
	// be set to true or ipv6, in this case this field should be set to false
	Ipv4 bool `protobuf:"varint,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	// List of the Route Targets. Each must be one of
	// TwoOctetAsSpecificExtended,
	// IPv4AddressSpecificExtended,
	// or FourOctetAsSpecificExtended.
	// Optional parameter
	Rt []*any.Any `protobuf:"bytes,3,rep,name=rt,proto3" json:"rt,omitempty"`
	// vpn_prefix is L3 VPN prefix which vpn label is requested for.
	// Optional parameter
	VpnPrefix            *Prefix  `protobuf:"bytes,4,opt,name=vpn_prefix,json=vpnPrefix,proto3" json:"vpn_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L3VPNRequest) Reset()         { *m = L3VPNRequest{} }
func (m *L3VPNRequest) String() string { return proto.CompactTextString(m) }
func (*L3VPNRequest) ProtoMessage()    {}
func (*L3VPNRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4982045cb0164b85, []int{7}
}

func (m *L3VPNRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L3VPNRequest.Unmarshal(m, b)
}
func (m *L3VPNRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L3VPNRequest.Marshal(b, m, deterministic)
}
func (m *L3VPNRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L3VPNRequest.Merge(m, src)
}
func (m *L3VPNRequest) XXX_Size() int {
	return xxx_messageInfo_L3VPNRequest.Size(m)
}
func (m *L3VPNRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_L3VPNRequest.DiscardUnknown(m)
}

var xxx_messageInfo_L3VPNRequest proto.InternalMessageInfo

func (m *L3VPNRequest) GetRd() *any.Any {
	if m != nil {
		return m.Rd
	}
	return nil
}

func (m *L3VPNRequest) GetIpv4() bool {
	if m != nil {
		return m.Ipv4
	}
	return false
}

func (m *L3VPNRequest) GetRt() []*any.Any {
	if m != nil {
		return m.Rt
	}
	return nil
}

func (m *L3VPNRequest) GetVpnPrefix() *Prefix {
	if m != nil {
		return m.VpnPrefix
	}
	return nil
}

type L3VPNResponse struct {
	// Identifies if a response carries ipv4 prefixes in this case this field is
	// set to true or ipv6, in this case this field is set to false
	Ipv4                 bool      `protobuf:"varint,1,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	VpnPrefix            []*Prefix `protobuf:"bytes,2,rep,name=vpn_prefix,json=vpnPrefix,proto3" json:"vpn_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *L3VPNResponse) Reset()         { *m = L3VPNResponse{} }
func (m *L3VPNResponse) String() string { return proto.CompactTextString(m) }
func (*L3VPNResponse) ProtoMessage()    {}
func (*L3VPNResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4982045cb0164b85, []int{8}
}

func (m *L3VPNResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L3VPNResponse.Unmarshal(m, b)
}
func (m *L3VPNResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L3VPNResponse.Marshal(b, m, deterministic)
}
func (m *L3VPNResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L3VPNResponse.Merge(m, src)
}
func (m *L3VPNResponse) XXX_Size() int {
	return xxx_messageInfo_L3VPNResponse.Size(m)
}
func (m *L3VPNResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_L3VPNResponse.DiscardUnknown(m)
}

var xxx_messageInfo_L3VPNResponse proto.InternalMessageInfo

func (m *L3VPNResponse) GetIpv4() bool {
	if m != nil {
		return m.Ipv4
	}
	return false
}

func (m *L3VPNResponse) GetVpnPrefix() []*Prefix {
	if m != nil {
		return m.VpnPrefix
	}
	return nil
}

// VPNv4Prefix defines a collection of VPNv4 prefixes, used in AdvBGPVPNv4 to
// advertise VPNv4 prefixes and in WdBGPVPNv4 to withdraw them.
type VPNv4Prefix struct {
	Prefix               []*Prefix `protobuf:"bytes,1,rep,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *VPNv4Prefix) Reset()         { *m = VPNv4Prefix{} }
func (m *VPNv4Prefix) String() string { return proto.CompactTextString(m) }
func (*VPNv4Prefix) ProtoMessage()    {}
func (*VPNv4Prefix) Descriptor() ([]byte, []int) {
	return fileDescriptor_4982045cb0164b85, []int{9}
}

func (m *VPNv4Prefix) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VPNv4Prefix.Unmarshal(m, b)
}
func (m *VPNv4Prefix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VPNv4Prefix.Marshal(b, m, deterministic)
}
func (m *VPNv4Prefix) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VPNv4Prefix.Merge(m, src)
}
func (m *VPNv4Prefix) XXX_Size() int {
	return xxx_messageInfo_VPNv4Prefix.Size(m)
}
func (m *VPNv4Prefix) XXX_DiscardUnknown() {
	xxx_messageInfo_VPNv4Prefix.DiscardUnknown(m)
}

var xxx_messageInfo_VPNv4Prefix proto.InternalMessageInfo

func (m *VPNv4Prefix) GetPrefix() []*Prefix {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func init() {
	proto.RegisterEnum("apis.GatewayErrors", GatewayErrors_name, GatewayErrors_value)
	proto.RegisterType((*RouteDistinguisherTwoOctetAS)(nil), "apis.RouteDistinguisherTwoOctetAS")
	proto.RegisterType((*RouteDistinguisherIPAddress)(nil), "apis.RouteDistinguisherIPAddress")
	proto.RegisterType((*RouteDistinguisherFourOctetAS)(nil), "apis.RouteDistinguisherFourOctetAS")
	proto.RegisterType((*TwoOctetAsSpecificExtended)(nil), "apis.TwoOctetAsSpecificExtended")
	proto.RegisterType((*IPv4AddressSpecificExtended)(nil), "apis.IPv4AddressSpecificExtended")
	proto.RegisterType((*FourOctetAsSpecificExtended)(nil), "apis.FourOctetAsSpecificExtended")
	proto.RegisterType((*Prefix)(nil), "apis.Prefix")
	proto.RegisterType((*L3VPNRequest)(nil), "apis.L3VPNRequest")
	proto.RegisterType((*L3VPNResponse)(nil), "apis.L3VPNResponse")
	proto.RegisterType((*VPNv4Prefix)(nil), "apis.VPNv4Prefix")
}

func init() { proto.RegisterFile("gateway-services.proto", fileDescriptor_4982045cb0164b85) }

var fileDescriptor_4982045cb0164b85 = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x51, 0x6f, 0xd3, 0x3c,
	0x14, 0xfd, 0x9c, 0x76, 0x5d, 0x7b, 0xdb, 0x4e, 0xfd, 0xcc, 0x34, 0x75, 0x2d, 0x88, 0xa9, 0xec,
	0x61, 0x02, 0xd1, 0xa1, 0x6d, 0x12, 0xbc, 0x16, 0x2d, 0x8c, 0xc2, 0x68, 0x42, 0x1a, 0xb6, 0xc7,
	0xca, 0x6d, 0xbc, 0xce, 0xa2, 0x4b, 0x82, 0xed, 0x64, 0xeb, 0x8f, 0x80, 0x67, 0x24, 0x7e, 0x07,
	0x3f, 0x82, 0x7f, 0x85, 0x62, 0x7b, 0xa5, 0x65, 0xb0, 0xa1, 0x89, 0x07, 0xde, 0x72, 0xcf, 0xb1,
	0xcf, 0xb9, 0xe7, 0x5a, 0xb9, 0xb0, 0x36, 0x26, 0x92, 0x9e, 0x93, 0xe9, 0x63, 0x41, 0x79, 0xca,
	0x46, 0x54, 0xb4, 0x63, 0x1e, 0xc9, 0x08, 0xe7, 0x49, 0xcc, 0x44, 0x63, 0x7d, 0x1c, 0x45, 0xe3,
	0x09, 0xdd, 0x56, 0xd8, 0x30, 0x39, 0xd9, 0x26, 0xe1, 0x54, 0x1f, 0x68, 0x34, 0x7f, 0xa6, 0xe8,
	0x59, 0x2c, 0x0d, 0xd9, 0x72, 0xe1, 0xae, 0x17, 0x25, 0x92, 0xee, 0x33, 0x21, 0x59, 0x38, 0x4e,
	0x98, 0x38, 0xa5, 0xdc, 0x3f, 0x8f, 0x9c, 0x91, 0xa4, 0xb2, 0xd3, 0xc7, 0xab, 0xb0, 0x44, 0x82,
	0x33, 0x16, 0xd6, 0xd1, 0x06, 0xda, 0xaa, 0x7a, 0xba, 0xc0, 0x0d, 0x28, 0x12, 0x21, 0xd8, 0x38,
	0xa4, 0x41, 0xdd, 0x52, 0xc4, 0xac, 0x6e, 0x39, 0xd0, 0xbc, 0xaa, 0xd8, 0x75, 0x3b, 0x41, 0xc0,
	0xa9, 0x10, 0x8b, 0x82, 0xa5, 0x3f, 0x11, 0x7c, 0x0b, 0xf7, 0xae, 0x0a, 0xbe, 0x88, 0x12, 0x7e,
	0xfb, 0x1e, 0x3f, 0x22, 0x68, 0xcc, 0x42, 0x8a, 0x7e, 0x4c, 0x47, 0xec, 0x84, 0x8d, 0xec, 0x0b,
	0x49, 0xc3, 0x80, 0x06, 0xf8, 0x01, 0x54, 0x99, 0x18, 0x48, 0x4e, 0x42, 0xc1, 0x24, 0x4b, 0xa9,
	0x12, 0x2e, 0x7a, 0x15, 0x26, 0xfc, 0x19, 0x86, 0xd7, 0xa1, 0x28, 0x92, 0xe1, 0x40, 0x4e, 0x63,
	0x6a, 0xf4, 0x97, 0x45, 0x32, 0xf4, 0xa7, 0x31, 0xc5, 0x2b, 0x60, 0x11, 0x51, 0xcf, 0x29, 0xd0,
	0x22, 0x02, 0xdf, 0x87, 0xf2, 0x24, 0x1a, 0x91, 0xc9, 0x40, 0xb7, 0x99, 0x57, 0x04, 0x28, 0xa8,
	0x93, 0x21, 0xad, 0xcf, 0x08, 0x9a, 0x5d, 0x37, 0xdd, 0x33, 0x43, 0xfa, 0xeb, 0x0d, 0xd5, 0x61,
	0x99, 0x68, 0x69, 0xd5, 0x55, 0xc9, 0xbb, 0x2c, 0x6f, 0x6e, 0xed, 0x13, 0x82, 0xe6, 0x8f, 0x61,
	0xff, 0x03, 0xb3, 0xfa, 0x86, 0xa0, 0xe0, 0x72, 0x7a, 0xc2, 0x2e, 0xe6, 0x63, 0x65, 0xae, 0x95,
	0x85, 0x58, 0x67, 0x44, 0xbc, 0x1f, 0x4c, 0x68, 0x38, 0x96, 0xa7, 0xc6, 0x13, 0x32, 0xe8, 0x50,
	0x21, 0xb8, 0x09, 0xa5, 0x34, 0x0e, 0x07, 0x13, 0x32, 0xa4, 0x13, 0xe3, 0x5e, 0x4c, 0xe3, 0xf0,
	0x30, 0xab, 0x33, 0x52, 0xb0, 0xc0, 0x90, 0xba, 0x83, 0xa2, 0x60, 0x81, 0x26, 0x37, 0xc1, 0xe2,
	0x41, 0x7d, 0x69, 0x03, 0x6d, 0x95, 0x77, 0x56, 0xdb, 0xfa, 0xdf, 0x6a, 0x5f, 0xfe, 0x5b, 0xed,
	0x4e, 0x38, 0xf5, 0x2c, 0x1e, 0xa8, 0x53, 0xb2, 0x5e, 0xd8, 0xc8, 0x5d, 0x73, 0x4a, 0xb6, 0xbe,
	0x20, 0xa8, 0x1c, 0xee, 0x1e, 0xb9, 0x3d, 0x8f, 0x7e, 0x48, 0xa8, 0x90, 0x46, 0x1c, 0xdd, 0x20,
	0x8e, 0x21, 0xcf, 0xe2, 0x74, 0x4f, 0xc5, 0x2a, 0x7a, 0xea, 0xdb, 0x18, 0xe6, 0xae, 0x37, 0xc4,
	0x8f, 0x00, 0xb2, 0xd8, 0xb1, 0x9a, 0x9f, 0x8a, 0x56, 0xde, 0xa9, 0xb4, 0xb3, 0x0d, 0xd2, 0xd6,
	0x33, 0xf5, 0xb2, 0xb1, 0xe8, 0xcf, 0x96, 0x0b, 0x55, 0xd3, 0x9c, 0x88, 0xa3, 0x50, 0xd0, 0x99,
	0x2f, 0x9a, 0xf3, 0x5d, 0x54, 0xb4, 0x94, 0xff, 0x6f, 0x15, 0x77, 0xa1, 0x7c, 0xe4, 0xf6, 0xd2,
	0x3d, 0xf3, 0x7e, 0x9b, 0x50, 0x30, 0xf7, 0xd0, 0x2f, 0xee, 0x19, 0xee, 0xe1, 0x2b, 0xa8, 0x1e,
	0xe8, 0xd5, 0x67, 0x73, 0x1e, 0x71, 0x81, 0x0b, 0x60, 0x39, 0xaf, 0x6b, 0xff, 0xe1, 0x65, 0xc8,
	0xd9, 0x5d, 0xa7, 0x86, 0x30, 0x40, 0xc1, 0xee, 0x39, 0x76, 0xcf, 0xaf, 0x59, 0xb8, 0x0a, 0x25,
	0xdb, 0xef, 0xbe, 0xb1, 0xf7, 0x9d, 0x77, 0x7e, 0x2d, 0xa7, 0xca, 0x97, 0x4e, 0xdf, 0xdf, 0x77,
	0x8e, 0x7b, 0xb5, 0xfc, 0xce, 0x57, 0x04, 0x2b, 0x46, 0xac, 0xaf, 0xd7, 0x28, 0x7e, 0x02, 0x4b,
	0x2a, 0x25, 0xc6, 0xda, 0x7d, 0xfe, 0x3d, 0x1a, 0x77, 0x16, 0x30, 0x33, 0x86, 0x67, 0x50, 0xee,
	0x04, 0xe9, 0xf3, 0x03, 0x57, 0x65, 0xc1, 0xff, 0xeb, 0x33, 0x73, 0xc1, 0x1a, 0x6b, 0x57, 0x1e,
	0xc0, 0xce, 0x76, 0x2e, 0x7e, 0x0a, 0x70, 0x1c, 0xdc, 0xe2, 0xe2, 0xb0, 0xa0, 0xea, 0xdd, 0xef,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xba, 0x8f, 0x73, 0x05, 0x06, 0x00, 0x00,
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
	// API to request L3 VPN label and list of prefixes for VRF specified by RD.
	// Optionally RT and/or Prefix can be specified as additional selection
	// creterias.
	L3VPN(ctx context.Context, in *L3VPNRequest, opts ...grpc.CallOption) (*L3VPNResponse, error)
	// API to advertise VPNv4 Prefix(s), no response other than Success or Failure
	// is generated
	AdvBGPVPNv4(ctx context.Context, in *VPNv4Prefix, opts ...grpc.CallOption) (*empty.Empty, error)
	// API to withdraw VPNv4 Prefix(s), no response other than Success or Failure
	// is generated
	WdBGPVPNv4(ctx context.Context, in *VPNv4Prefix, opts ...grpc.CallOption) (*empty.Empty, error)
}

type gatewayServiceClient struct {
	cc *grpc.ClientConn
}

func NewGatewayServiceClient(cc *grpc.ClientConn) GatewayServiceClient {
	return &gatewayServiceClient{cc}
}

func (c *gatewayServiceClient) L3VPN(ctx context.Context, in *L3VPNRequest, opts ...grpc.CallOption) (*L3VPNResponse, error) {
	out := new(L3VPNResponse)
	err := c.cc.Invoke(ctx, "/apis.GatewayService/L3VPN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) AdvBGPVPNv4(ctx context.Context, in *VPNv4Prefix, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/apis.GatewayService/AdvBGPVPNv4", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) WdBGPVPNv4(ctx context.Context, in *VPNv4Prefix, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/apis.GatewayService/WdBGPVPNv4", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServiceServer is the server API for GatewayService service.
type GatewayServiceServer interface {
	// API to request L3 VPN label and list of prefixes for VRF specified by RD.
	// Optionally RT and/or Prefix can be specified as additional selection
	// creterias.
	L3VPN(context.Context, *L3VPNRequest) (*L3VPNResponse, error)
	// API to advertise VPNv4 Prefix(s), no response other than Success or Failure
	// is generated
	AdvBGPVPNv4(context.Context, *VPNv4Prefix) (*empty.Empty, error)
	// API to withdraw VPNv4 Prefix(s), no response other than Success or Failure
	// is generated
	WdBGPVPNv4(context.Context, *VPNv4Prefix) (*empty.Empty, error)
}

// UnimplementedGatewayServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayServiceServer struct {
}

func (*UnimplementedGatewayServiceServer) L3VPN(ctx context.Context, req *L3VPNRequest) (*L3VPNResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method L3VPN not implemented")
}
func (*UnimplementedGatewayServiceServer) AdvBGPVPNv4(ctx context.Context, req *VPNv4Prefix) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdvBGPVPNv4 not implemented")
}
func (*UnimplementedGatewayServiceServer) WdBGPVPNv4(ctx context.Context, req *VPNv4Prefix) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WdBGPVPNv4 not implemented")
}

func RegisterGatewayServiceServer(s *grpc.Server, srv GatewayServiceServer) {
	s.RegisterService(&_GatewayService_serviceDesc, srv)
}

func _GatewayService_L3VPN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(L3VPNRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).L3VPN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.GatewayService/L3VPN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).L3VPN(ctx, req.(*L3VPNRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_AdvBGPVPNv4_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VPNv4Prefix)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).AdvBGPVPNv4(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.GatewayService/AdvBGPVPNv4",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).AdvBGPVPNv4(ctx, req.(*VPNv4Prefix))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_WdBGPVPNv4_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VPNv4Prefix)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).WdBGPVPNv4(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.GatewayService/WdBGPVPNv4",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).WdBGPVPNv4(ctx, req.(*VPNv4Prefix))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apis.GatewayService",
	HandlerType: (*GatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "L3VPN",
			Handler:    _GatewayService_L3VPN_Handler,
		},
		{
			MethodName: "AdvBGPVPNv4",
			Handler:    _GatewayService_AdvBGPVPNv4_Handler,
		},
		{
			MethodName: "WdBGPVPNv4",
			Handler:    _GatewayService_WdBGPVPNv4_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway-services.proto",
}
