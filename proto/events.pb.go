// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events.proto

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

type Event struct {
	EvID                 string   `protobuf:"bytes,1,opt,name=EvID,json=evID,proto3" json:"EvID,omitempty"`
	NtID                 string   `protobuf:"bytes,2,opt,name=NtID,json=ntID,proto3" json:"NtID,omitempty"`
	CpID                 string   `protobuf:"bytes,3,opt,name=CpID,json=cpID,proto3" json:"CpID,omitempty"`
	DvID                 string   `protobuf:"bytes,4,opt,name=DvID,json=dvID,proto3" json:"DvID,omitempty"`
	Visibility           string   `protobuf:"bytes,5,opt,name=Visibility,json=visibility,proto3" json:"Visibility,omitempty"`
	EvType               string   `protobuf:"bytes,6,opt,name=EvType,json=evType,proto3" json:"EvType,omitempty"`
	EvSubType            string   `protobuf:"bytes,7,opt,name=EvSubType,json=evSubType,proto3" json:"EvSubType,omitempty"`
	EvDescription        string   `protobuf:"bytes,8,opt,name=EvDescription,json=evDescription,proto3" json:"EvDescription,omitempty"`
	KeyID                int64    `protobuf:"zigzag64,100,opt,name=KeyID,json=keyID,proto3" json:"KeyID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetEvID() string {
	if m != nil {
		return m.EvID
	}
	return ""
}

func (m *Event) GetNtID() string {
	if m != nil {
		return m.NtID
	}
	return ""
}

func (m *Event) GetCpID() string {
	if m != nil {
		return m.CpID
	}
	return ""
}

func (m *Event) GetDvID() string {
	if m != nil {
		return m.DvID
	}
	return ""
}

func (m *Event) GetVisibility() string {
	if m != nil {
		return m.Visibility
	}
	return ""
}

func (m *Event) GetEvType() string {
	if m != nil {
		return m.EvType
	}
	return ""
}

func (m *Event) GetEvSubType() string {
	if m != nil {
		return m.EvSubType
	}
	return ""
}

func (m *Event) GetEvDescription() string {
	if m != nil {
		return m.EvDescription
	}
	return ""
}

func (m *Event) GetKeyID() int64 {
	if m != nil {
		return m.KeyID
	}
	return 0
}

func init() {
	proto.RegisterType((*Event)(nil), "proto.Event")
}

func init() { proto.RegisterFile("events.proto", fileDescriptor_8f22242cb04491f9) }

var fileDescriptor_8f22242cb04491f9 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcf, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x80, 0x61, 0xa2, 0x4d, 0xb4, 0x83, 0xbd, 0x0c, 0x22, 0x39, 0x88, 0x14, 0xf1, 0xd0, 0x93,
	0x17, 0x1f, 0xc1, 0xe4, 0x10, 0x04, 0x0f, 0x55, 0x7c, 0x80, 0xb6, 0x73, 0x08, 0x4a, 0x1b, 0xda,
	0x18, 0xe8, 0x23, 0xef, 0x5b, 0x2c, 0x99, 0xb0, 0xcb, 0x9e, 0xc2, 0x7c, 0x7f, 0xe6, 0x30, 0x70,
	0x47, 0x89, 0xe6, 0xb8, 0xbd, 0x86, 0x75, 0x89, 0x0b, 0x4a, 0x7e, 0x9e, 0x0f, 0x02, 0xa4, 0xcd,
	0x8e, 0x08, 0x95, 0x4d, 0xce, 0x68, 0xd1, 0x8a, 0xae, 0xee, 0x2b, 0x4a, 0xce, 0x64, 0xfb, 0x8c,
	0xce, 0xe8, 0xab, 0x62, 0x73, 0x2c, 0xf6, 0x1e, 0x9c, 0xd1, 0xd7, 0xc5, 0xc6, 0x50, 0xcc, 0xe4,
	0xdd, 0xaa, 0xd8, 0x94, 0x77, 0x9f, 0x00, 0x7e, 0xfc, 0xe6, 0x07, 0xff, 0xe7, 0xe3, 0xae, 0x25,
	0x17, 0x48, 0x67, 0xc1, 0x07, 0x50, 0x36, 0x7d, 0xef, 0x81, 0xb4, 0xe2, 0xa6, 0x88, 0x27, 0x7c,
	0x84, 0xda, 0xa6, 0xaf, 0xff, 0x81, 0xd3, 0x0d, 0xa7, 0x9a, 0x4e, 0x80, 0x2f, 0xd0, 0xd8, 0x64,
	0x68, 0x1b, 0x57, 0x1f, 0xa2, 0x5f, 0x66, 0x7d, 0xcb, 0x3f, 0x1a, 0xba, 0x44, 0xbc, 0x07, 0xf9,
	0x41, 0xbb, 0x33, 0x7a, 0x6a, 0x45, 0x87, 0xbd, 0xfc, 0xcd, 0xc3, 0xa0, 0xf8, 0xe4, 0xb7, 0x63,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x06, 0x91, 0x22, 0x8e, 0x09, 0x01, 0x00, 0x00,
}
