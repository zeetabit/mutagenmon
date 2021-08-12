// Code generated by protoc-gen-go. DO NOT EDIT.
// source: synchronization/core/archive.proto

package core

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

// Archive is a wrapper around Entry that allows identification of non-existent
// roots when serializing. In-memory, a nil-Entry (that arrives without any
// error) represents an absence of content on the filesystem. Unfortunately,
// there is no way to represent that as an encoded message (an empty byte slice
// would successfully decode to an empty directory entry). By adding a level of
// indirection that allows for an unset root entry, we can encode Entry messages
// in a way that allows us to represent absence.
type Archive struct {
	Root                 *Entry   `protobuf:"bytes,1,opt,name=root,proto3" json:"root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Archive) Reset()         { *m = Archive{} }
func (m *Archive) String() string { return proto.CompactTextString(m) }
func (*Archive) ProtoMessage()    {}
func (*Archive) Descriptor() ([]byte, []int) {
	return fileDescriptor_b604a3e22a5b0a8f, []int{0}
}

func (m *Archive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Archive.Unmarshal(m, b)
}
func (m *Archive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Archive.Marshal(b, m, deterministic)
}
func (m *Archive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Archive.Merge(m, src)
}
func (m *Archive) XXX_Size() int {
	return xxx_messageInfo_Archive.Size(m)
}
func (m *Archive) XXX_DiscardUnknown() {
	xxx_messageInfo_Archive.DiscardUnknown(m)
}

var xxx_messageInfo_Archive proto.InternalMessageInfo

func (m *Archive) GetRoot() *Entry {
	if m != nil {
		return m.Root
	}
	return nil
}

func init() {
	proto.RegisterType((*Archive)(nil), "core.Archive")
}

func init() { proto.RegisterFile("synchronization/core/archive.proto", fileDescriptor_b604a3e22a5b0a8f) }

var fileDescriptor_b604a3e22a5b0a8f = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0xae, 0xcc, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcb, 0xac, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x4f, 0xce, 0x2f, 0x4a,
	0xd5, 0x4f, 0x2c, 0x4a, 0xce, 0xc8, 0x2c, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x01, 0x89, 0x49, 0x29, 0x60, 0x55, 0x99, 0x9a, 0x57, 0x52, 0x54, 0x09, 0x51, 0xa7, 0xa4, 0xc5,
	0xc5, 0xee, 0x08, 0xd1, 0x28, 0x24, 0xcf, 0xc5, 0x52, 0x94, 0x9f, 0x5f, 0x22, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x6d, 0xc4, 0xad, 0x07, 0x52, 0xab, 0xe7, 0x0a, 0x52, 0x1b, 0x04, 0x96, 0x70, 0xb2,
	0x88, 0x32, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2d, 0x2d,
	0x49, 0x4c, 0x4f, 0xcd, 0xd3, 0xcd, 0xcc, 0x87, 0x31, 0xf5, 0x0b, 0xb2, 0xd3, 0xf5, 0xb1, 0xd9,
	0x98, 0xc4, 0x06, 0xb6, 0xcc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xba, 0x6b, 0xbe, 0xd7, 0xba,
	0x00, 0x00, 0x00,
}
