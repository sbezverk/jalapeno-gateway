// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

package apis

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
type GatewayError int32

const (
	GatewayError_OK GatewayError = 0
	// Returned when a gateway encountered a generic error
	GatewayError_EIO GatewayError = 1
	// Returned when request key does not exist
	GatewayError_ENOENT GatewayError = 2
	// Returned when an operation triggered by the client's requested timed out
	// and was canceled
	GatewayError_ETIMEDOUT GatewayError = 3
	// Returned when a gateway cannot reach a DB host
	GatewayError_EHOSTDOWN GatewayError = 4
)

var GatewayError_name = map[int32]string{
	0: "OK",
	1: "EIO",
	2: "ENOENT",
	3: "ETIMEDOUT",
	4: "EHOSTDOWN",
}

var GatewayError_value = map[string]int32{
	"OK":        0,
	"EIO":       1,
	"ENOENT":    2,
	"ETIMEDOUT": 3,
	"EHOSTDOWN": 4,
}

func (x GatewayError) String() string {
	return proto.EnumName(GatewayError_name, int32(x))
}

func (GatewayError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{0}
}

type Prefix struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	MaskLength           uint32   `protobuf:"varint,2,opt,name=mask_length,json=maskLength,proto3" json:"mask_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Prefix) Reset()         { *m = Prefix{} }
func (m *Prefix) String() string { return proto.CompactTextString(m) }
func (*Prefix) ProtoMessage()    {}
func (*Prefix) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{0}
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

func init() {
	proto.RegisterEnum("apis.GatewayError", GatewayError_name, GatewayError_value)
	proto.RegisterType((*Prefix)(nil), "apis.Prefix")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor_db1b6b0986796150) }

var fileDescriptor_db1b6b0986796150 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xce, 0x4f, 0x0b, 0x82, 0x30,
	0x18, 0x80, 0xf1, 0xa6, 0x32, 0xe9, 0x4d, 0x61, 0xbc, 0x27, 0x6f, 0x49, 0x27, 0xe9, 0xd0, 0xa5,
	0x8f, 0x90, 0xa3, 0xa4, 0x72, 0x61, 0x8b, 0x8e, 0x31, 0x71, 0x95, 0xf4, 0x47, 0xd9, 0x84, 0xea,
	0xdb, 0x47, 0x42, 0xc7, 0xe7, 0x77, 0x7a, 0x00, 0x4a, 0x65, 0xf5, 0xac, 0x35, 0x4d, 0xd7, 0xa0,
	0xa7, 0xda, 0xda, 0x4e, 0x16, 0x40, 0x77, 0x46, 0x9f, 0xeb, 0x37, 0x46, 0xe0, 0xab, 0xaa, 0x32,
	0xda, 0xda, 0x88, 0xc4, 0x24, 0x09, 0x8a, 0x7f, 0xe2, 0x18, 0x46, 0x0f, 0x65, 0x6f, 0xa7, 0xbb,
	0x7e, 0x5e, 0xba, 0x6b, 0xe4, 0xc4, 0x24, 0x09, 0x0b, 0xf8, 0xd1, 0xa6, 0x97, 0x69, 0x06, 0xc1,
	0x52, 0x75, 0xfa, 0xa5, 0x3e, 0xdc, 0x98, 0xc6, 0x20, 0x05, 0x47, 0xac, 0xd9, 0x00, 0x7d, 0x70,
	0x79, 0x26, 0x18, 0x41, 0x00, 0xca, 0x73, 0xc1, 0x73, 0xc9, 0x1c, 0x0c, 0x61, 0xc8, 0x65, 0xb6,
	0xe5, 0xa9, 0x38, 0x48, 0xe6, 0xf6, 0xb9, 0x12, 0x7b, 0x99, 0x8a, 0x63, 0xce, 0xbc, 0x92, 0xf6,
	0x73, 0xf3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xb5, 0xd4, 0x58, 0xaa, 0x00, 0x00, 0x00,
}
