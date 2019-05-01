// Code generated by protoc-gen-go. DO NOT EDIT.
// source: assignments.proto

package proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Assignment struct {
	// Common Types
	AsID  string `protobuf:"bytes,1,opt,name=AsID,json=asID,proto3" json:"AsID,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=Label,json=label,proto3" json:"Label,omitempty"`
	// Specific types
	CpID     string           `protobuf:"bytes,20,opt,name=CpID,json=cpID,proto3" json:"CpID,omitempty"`
	DvID     string           `protobuf:"bytes,21,opt,name=DvID,json=dvID,proto3" json:"DvID,omitempty"`
	Level    int32            `protobuf:"varint,22,opt,name=Level,json=level,proto3" json:"Level,omitempty"`
	Settings []*CommonSetting `protobuf:"bytes,24,rep,name=Settings,json=settings,proto3" json:"Settings,omitempty"`
	// More Generic Types
	Description string `protobuf:"bytes,40,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
	// Datastore Key
	KeyID                int64    `protobuf:"zigzag64,100,opt,name=KeyID,json=keyID,proto3" json:"KeyID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Assignment) Reset()         { *m = Assignment{} }
func (m *Assignment) String() string { return proto.CompactTextString(m) }
func (*Assignment) ProtoMessage()    {}
func (*Assignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_887509349ea6962f, []int{0}
}

func (m *Assignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Assignment.Unmarshal(m, b)
}
func (m *Assignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Assignment.Marshal(b, m, deterministic)
}
func (m *Assignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Assignment.Merge(m, src)
}
func (m *Assignment) XXX_Size() int {
	return xxx_messageInfo_Assignment.Size(m)
}
func (m *Assignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Assignment.DiscardUnknown(m)
}

var xxx_messageInfo_Assignment proto.InternalMessageInfo

func (m *Assignment) GetAsID() string {
	if m != nil {
		return m.AsID
	}
	return ""
}

func (m *Assignment) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Assignment) GetCpID() string {
	if m != nil {
		return m.CpID
	}
	return ""
}

func (m *Assignment) GetDvID() string {
	if m != nil {
		return m.DvID
	}
	return ""
}

func (m *Assignment) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Assignment) GetSettings() []*CommonSetting {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *Assignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Assignment) GetKeyID() int64 {
	if m != nil {
		return m.KeyID
	}
	return 0
}

func init() {
	proto.RegisterType((*Assignment)(nil), "proto.Assignment")
}

func init() { proto.RegisterFile("assignments.proto", fileDescriptor_887509349ea6962f) }

var fileDescriptor_887509349ea6962f = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0xc1, 0x4e, 0xc4, 0x20,
	0x14, 0x45, 0x45, 0x07, 0x33, 0x52, 0x37, 0x92, 0x6a, 0x5e, 0x5c, 0x11, 0x57, 0xac, 0x26, 0x46,
	0xbf, 0x60, 0x32, 0x6f, 0x43, 0x74, 0x61, 0xea, 0x17, 0x74, 0x5a, 0x32, 0x69, 0xa4, 0x40, 0xe6,
	0x91, 0x26, 0xfe, 0xac, 0xdf, 0x62, 0x80, 0x46, 0x57, 0xe4, 0x1e, 0x0e, 0xdc, 0x2b, 0xee, 0x7a,
	0xa2, 0xe9, 0xe4, 0x67, 0xeb, 0x13, 0xed, 0xe2, 0x39, 0xa4, 0x20, 0x79, 0x39, 0x1e, 0x6f, 0x87,
	0x30, 0xcf, 0xc1, 0x57, 0xf8, 0xf4, 0xc3, 0x84, 0xd8, 0xff, 0xa9, 0x52, 0x8a, 0xcd, 0x9e, 0x0c,
	0x02, 0x53, 0x4c, 0xdf, 0x74, 0x9b, 0x9e, 0x0c, 0xca, 0x56, 0xf0, 0xf7, 0xfe, 0x68, 0x1d, 0x5c,
	0x16, 0xc8, 0x5d, 0x0e, 0xd9, 0x3c, 0x44, 0x83, 0xd0, 0x56, 0x73, 0x88, 0x06, 0x33, 0xc3, 0xc5,
	0x20, 0xdc, 0x57, 0x36, 0x2e, 0xeb, 0x6b, 0xbb, 0x58, 0x07, 0x0f, 0x8a, 0x69, 0xde, 0x71, 0x97,
	0x83, 0x7c, 0x16, 0xdb, 0x4f, 0x9b, 0xd2, 0xe4, 0x4f, 0x04, 0xa0, 0xae, 0x74, 0xf3, 0xd2, 0xd6,
	0x41, 0xbb, 0x43, 0x59, 0xb7, 0x5e, 0x76, 0x5b, 0x5a, 0x2d, 0xa9, 0x44, 0x83, 0x96, 0x86, 0xf3,
	0x14, 0xd3, 0x14, 0x3c, 0xe8, 0x52, 0xd1, 0x8c, 0xff, 0x28, 0x37, 0xbd, 0xd9, 0x6f, 0x83, 0x30,
	0x2a, 0xa6, 0x65, 0xc7, 0xbf, 0x72, 0xf8, 0xb8, 0x38, 0x5e, 0x97, 0x8f, 0x5f, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xc8, 0xb9, 0x35, 0x37, 0x13, 0x01, 0x00, 0x00,
}