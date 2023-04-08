// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shipping/shipping.proto

package shipping

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

type CreateShippingRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateShippingRequest) Reset()         { *m = CreateShippingRequest{} }
func (m *CreateShippingRequest) String() string { return proto.CompactTextString(m) }
func (*CreateShippingRequest) ProtoMessage()    {}
func (*CreateShippingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_218d76102d381bfb, []int{0}
}

func (m *CreateShippingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateShippingRequest.Unmarshal(m, b)
}
func (m *CreateShippingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateShippingRequest.Marshal(b, m, deterministic)
}
func (m *CreateShippingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateShippingRequest.Merge(m, src)
}
func (m *CreateShippingRequest) XXX_Size() int {
	return xxx_messageInfo_CreateShippingRequest.Size(m)
}
func (m *CreateShippingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateShippingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateShippingRequest proto.InternalMessageInfo

func (m *CreateShippingRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type CreateShippingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateShippingResponse) Reset()         { *m = CreateShippingResponse{} }
func (m *CreateShippingResponse) String() string { return proto.CompactTextString(m) }
func (*CreateShippingResponse) ProtoMessage()    {}
func (*CreateShippingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_218d76102d381bfb, []int{1}
}

func (m *CreateShippingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateShippingResponse.Unmarshal(m, b)
}
func (m *CreateShippingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateShippingResponse.Marshal(b, m, deterministic)
}
func (m *CreateShippingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateShippingResponse.Merge(m, src)
}
func (m *CreateShippingResponse) XXX_Size() int {
	return xxx_messageInfo_CreateShippingResponse.Size(m)
}
func (m *CreateShippingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateShippingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateShippingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateShippingRequest)(nil), "CreateShippingRequest")
	proto.RegisterType((*CreateShippingResponse)(nil), "CreateShippingResponse")
}

func init() { proto.RegisterFile("shipping/shipping.proto", fileDescriptor_218d76102d381bfb) }

var fileDescriptor_218d76102d381bfb = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0xce, 0xc8, 0x2c,
	0x28, 0xc8, 0xcc, 0x4b, 0xd7, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x95, 0x0c, 0xb9,
	0x44, 0x9d, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0x83, 0xa1, 0xe2, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5,
	0x25, 0x42, 0x12, 0x5c, 0xec, 0x89, 0x29, 0x29, 0x45, 0xa9, 0xc5, 0xc5, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x30, 0xae, 0x92, 0x04, 0x97, 0x18, 0xba, 0x96, 0xe2, 0x82, 0xfc, 0xbc, 0xe2,
	0x54, 0x23, 0x77, 0x2e, 0x0e, 0x98, 0x98, 0x90, 0x35, 0x17, 0x1b, 0x44, 0x95, 0x90, 0x98, 0x1e,
	0x56, 0x1b, 0xa4, 0xc4, 0xf5, 0xb0, 0x1b, 0xa3, 0xc4, 0xe0, 0x64, 0x14, 0x65, 0x90, 0x9e, 0x59,
	0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x56, 0x94, 0x5a, 0x9c, 0x98, 0x97, 0x5f,
	0xa6, 0x9f, 0x9b, 0x99, 0x5c, 0x94, 0x5f, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x5a, 0xac, 0x0b,
	0xf6, 0x00, 0xdc, 0x3f, 0x49, 0x6c, 0x60, 0xbe, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xee, 0x8c,
	0x79, 0xa7, 0xeb, 0x00, 0x00, 0x00,
}
