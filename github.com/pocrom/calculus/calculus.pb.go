// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculus.proto

package calculus

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

type Data struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26a829c8b298ac1, []int{0}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Data) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type Result struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26a829c8b298ac1, []int{1}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Data)(nil), "calculus.Data")
	proto.RegisterType((*Result)(nil), "calculus.Result")
}

func init() {
	proto.RegisterFile("calculus.proto", fileDescriptor_c26a829c8b298ac1)
}

var fileDescriptor_c26a829c8b298ac1 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4e, 0xcc, 0x49,
	0x2e, 0xcd, 0x29, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x94,
	0xb8, 0x58, 0x5c, 0x12, 0x4b, 0x12, 0x85, 0x78, 0xb8, 0x18, 0x13, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x58, 0x83, 0x18, 0xc1, 0xbc, 0x24, 0x09, 0x26, 0x08, 0x2f, 0x49, 0x49, 0x81, 0x8b, 0x2d, 0x28,
	0xb5, 0xb8, 0x34, 0xa7, 0x44, 0x48, 0x8c, 0x8b, 0xad, 0x08, 0xcc, 0x82, 0x2a, 0x85, 0xf2, 0x8c,
	0x12, 0xb9, 0x38, 0x9d, 0xc1, 0x26, 0x26, 0x96, 0xa4, 0x0a, 0x69, 0x72, 0x31, 0x3b, 0xa6, 0xa4,
	0x08, 0xf1, 0xe9, 0xc1, 0x2d, 0x05, 0xd9, 0x20, 0x25, 0x80, 0xe0, 0x43, 0x4c, 0x53, 0x62, 0x00,
	0x29, 0x0d, 0x2e, 0x4d, 0x22, 0x46, 0xa9, 0x93, 0x4c, 0x94, 0x54, 0x7a, 0x66, 0x49, 0x46, 0x69,
	0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x41, 0x7e, 0x72, 0x51, 0x7e, 0xae, 0x3e, 0x4c, 0x59, 0x12,
	0x1b, 0xd8, 0x5f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xc1, 0xa6, 0xe9, 0xe9, 0x00,
	0x00, 0x00,
}