// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/tunneling/tunneling.proto

package tunneling

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	selection "github.com/mutagen-io/mutagen/pkg/selection"
	tunneling "github.com/mutagen-io/mutagen/pkg/tunneling"
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

type CreationSpecification struct {
	// Configuration is the base tunnel configuration. It is the result of
	// merging the global configuration (unless disabled), any manually
	// specified configuration file, and any command line configuration
	// parameters.
	Configuration *tunneling.Configuration `protobuf:"bytes,1,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// Name is the name for the tunnel object.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Labels are the labels for the tunnel object.
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Paused indicates whether or not to create the tunnel pre-paused.
	Paused               bool     `protobuf:"varint,4,opt,name=paused,proto3" json:"paused,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreationSpecification) Reset()         { *m = CreationSpecification{} }
func (m *CreationSpecification) String() string { return proto.CompactTextString(m) }
func (*CreationSpecification) ProtoMessage()    {}
func (*CreationSpecification) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e41fdc941b0f342, []int{0}
}

func (m *CreationSpecification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreationSpecification.Unmarshal(m, b)
}
func (m *CreationSpecification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreationSpecification.Marshal(b, m, deterministic)
}
func (m *CreationSpecification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreationSpecification.Merge(m, src)
}
func (m *CreationSpecification) XXX_Size() int {
	return xxx_messageInfo_CreationSpecification.Size(m)
}
func (m *CreationSpecification) XXX_DiscardUnknown() {
	xxx_messageInfo_CreationSpecification.DiscardUnknown(m)
}

var xxx_messageInfo_CreationSpecification proto.InternalMessageInfo

func (m *CreationSpecification) GetConfiguration() *tunneling.Configuration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *CreationSpecification) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreationSpecification) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *CreationSpecification) GetPaused() bool {
	if m != nil {
		return m.Paused
	}
	return false
}

type CreateRequest struct {
	Specification        *CreationSpecification `protobuf:"bytes,1,opt,name=specification,proto3" json:"specification,omitempty"`
	Response             string                 `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e41fdc941b0f342, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetSpecification() *CreationSpecification {
	if m != nil {
		return m.Specification
	}
	return nil
}

func (m *CreateRequest) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type CreateResponse struct {
	HostCredentials      *tunneling.TunnelHostCredentials `protobuf:"bytes,1,opt,name=hostCredentials,proto3" json:"hostCredentials,omitempty"`
	Message              string                           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Prompt               string                           `protobuf:"bytes,3,opt,name=prompt,proto3" json:"prompt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e41fdc941b0f342, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetHostCredentials() *tunneling.TunnelHostCredentials {
	if m != nil {
		return m.HostCredentials
	}
	return nil
}

func (m *CreateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateResponse) GetPrompt() string {
	if m != nil {
		return m.Prompt
	}
	return ""
}

type ListRequest struct {
	Selection            *selection.Selection `protobuf:"bytes,1,opt,name=selection,proto3" json:"selection,omitempty"`
	PreviousStateIndex   uint64               `protobuf:"varint,2,opt,name=previousStateIndex,proto3" json:"previousStateIndex,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e41fdc941b0f342, []int{3}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetSelection() *selection.Selection {
	if m != nil {
		return m.Selection
	}
	return nil
}

func (m *ListRequest) GetPreviousStateIndex() uint64 {
	if m != nil {
		return m.PreviousStateIndex
	}
	return 0
}

type ListResponse struct {
	StateIndex           uint64             `protobuf:"varint,1,opt,name=stateIndex,proto3" json:"stateIndex,omitempty"`
	TunnelStates         []*tunneling.State `protobuf:"bytes,2,rep,name=tunnelStates,proto3" json:"tunnelStates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e41fdc941b0f342, []int{4}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetStateIndex() uint64 {
	if m != nil {
		return m.StateIndex
	}
	return 0
}

func (m *ListResponse) GetTunnelStates() []*tunneling.State {
	if m != nil {
		return m.TunnelStates
	}
	return nil
}

type PauseRequest struct {
	Selection            *selection.Selection `protobuf:"bytes,1,opt,name=selection,proto3" json:"selection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PauseRequest) Reset()         { *m = PauseRequest{} }
func (m *PauseRequest) String() string { return proto.CompactTextString(m) }
func (*PauseRequest) ProtoMessage()    {}
func (*PauseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e41fdc941b0f342, []int{5}
}

func (m *PauseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PauseRequest.Unmarshal(m, b)
}
func (m *PauseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PauseRequest.Marshal(b, m, deterministic)
}
func (m *PauseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PauseRequest.Merge(m, src)
}
func (m *PauseRequest) XXX_Size() int {
	return xxx_messageInfo_PauseRequest.Size(m)
}
func (m *PauseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PauseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PauseRequest proto.InternalMessageInfo

func (m *PauseRequest) GetSelection() *selection.Selection {
	if m != nil {
		return m.Selection
	}
	return nil
}

type PauseResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PauseResponse) Reset()         { *m = PauseResponse{} }
func (m *PauseResponse) String() string { return proto.CompactTextString(m) }
func (*PauseResponse) ProtoMessage()    {}
func (*PauseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e41fdc941b0f342, []int{6}
}

func (m *PauseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PauseResponse.Unmarshal(m, b)
}
func (m *PauseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PauseResponse.Marshal(b, m, deterministic)
}
func (m *PauseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PauseResponse.Merge(m, src)
}
func (m *PauseResponse) XXX_Size() int {
	return xxx_messageInfo_PauseResponse.Size(m)
}
func (m *PauseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PauseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PauseResponse proto.InternalMessageInfo

func (m *PauseResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ResumeRequest struct {
	Selection            *selection.Selection `protobuf:"bytes,1,opt,name=selection,proto3" json:"selection,omitempty"`
	Response             string               `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ResumeRequest) Reset()         { *m = ResumeRequest{} }
func (m *ResumeRequest) String() string { return proto.CompactTextString(m) }
func (*ResumeRequest) ProtoMessage()    {}
func (*ResumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e41fdc941b0f342, []int{7}
}

func (m *ResumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeRequest.Unmarshal(m, b)
}
func (m *ResumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeRequest.Marshal(b, m, deterministic)
}
func (m *ResumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeRequest.Merge(m, src)
}
func (m *ResumeRequest) XXX_Size() int {
	return xxx_messageInfo_ResumeRequest.Size(m)
}
func (m *ResumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeRequest proto.InternalMessageInfo

func (m *ResumeRequest) GetSelection() *selection.Selection {
	if m != nil {
		return m.Selection
	}
	return nil
}

func (m *ResumeRequest) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type ResumeResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Prompt               string   `protobuf:"bytes,2,opt,name=prompt,proto3" json:"prompt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResumeResponse) Reset()         { *m = ResumeResponse{} }
func (m *ResumeResponse) String() string { return proto.CompactTextString(m) }
func (*ResumeResponse) ProtoMessage()    {}
func (*ResumeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e41fdc941b0f342, []int{8}
}

func (m *ResumeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeResponse.Unmarshal(m, b)
}
func (m *ResumeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeResponse.Marshal(b, m, deterministic)
}
func (m *ResumeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeResponse.Merge(m, src)
}
func (m *ResumeResponse) XXX_Size() int {
	return xxx_messageInfo_ResumeResponse.Size(m)
}
func (m *ResumeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeResponse proto.InternalMessageInfo

func (m *ResumeResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ResumeResponse) GetPrompt() string {
	if m != nil {
		return m.Prompt
	}
	return ""
}

type TerminateRequest struct {
	Selection            *selection.Selection `protobuf:"bytes,1,opt,name=selection,proto3" json:"selection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TerminateRequest) Reset()         { *m = TerminateRequest{} }
func (m *TerminateRequest) String() string { return proto.CompactTextString(m) }
func (*TerminateRequest) ProtoMessage()    {}
func (*TerminateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e41fdc941b0f342, []int{9}
}

func (m *TerminateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminateRequest.Unmarshal(m, b)
}
func (m *TerminateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminateRequest.Marshal(b, m, deterministic)
}
func (m *TerminateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminateRequest.Merge(m, src)
}
func (m *TerminateRequest) XXX_Size() int {
	return xxx_messageInfo_TerminateRequest.Size(m)
}
func (m *TerminateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TerminateRequest proto.InternalMessageInfo

func (m *TerminateRequest) GetSelection() *selection.Selection {
	if m != nil {
		return m.Selection
	}
	return nil
}

type TerminateResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminateResponse) Reset()         { *m = TerminateResponse{} }
func (m *TerminateResponse) String() string { return proto.CompactTextString(m) }
func (*TerminateResponse) ProtoMessage()    {}
func (*TerminateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e41fdc941b0f342, []int{10}
}

func (m *TerminateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminateResponse.Unmarshal(m, b)
}
func (m *TerminateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminateResponse.Marshal(b, m, deterministic)
}
func (m *TerminateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminateResponse.Merge(m, src)
}
func (m *TerminateResponse) XXX_Size() int {
	return xxx_messageInfo_TerminateResponse.Size(m)
}
func (m *TerminateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TerminateResponse proto.InternalMessageInfo

func (m *TerminateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CreationSpecification)(nil), "tunneling.CreationSpecification")
	proto.RegisterMapType((map[string]string)(nil), "tunneling.CreationSpecification.LabelsEntry")
	proto.RegisterType((*CreateRequest)(nil), "tunneling.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "tunneling.CreateResponse")
	proto.RegisterType((*ListRequest)(nil), "tunneling.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "tunneling.ListResponse")
	proto.RegisterType((*PauseRequest)(nil), "tunneling.PauseRequest")
	proto.RegisterType((*PauseResponse)(nil), "tunneling.PauseResponse")
	proto.RegisterType((*ResumeRequest)(nil), "tunneling.ResumeRequest")
	proto.RegisterType((*ResumeResponse)(nil), "tunneling.ResumeResponse")
	proto.RegisterType((*TerminateRequest)(nil), "tunneling.TerminateRequest")
	proto.RegisterType((*TerminateResponse)(nil), "tunneling.TerminateResponse")
}

func init() { proto.RegisterFile("service/tunneling/tunneling.proto", fileDescriptor_4e41fdc941b0f342) }

var fileDescriptor_4e41fdc941b0f342 = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x49, 0x1a, 0x9a, 0x49, 0x53, 0xca, 0xaa, 0x2d, 0xae, 0xf9, 0x50, 0xf0, 0x29, 0x48,
	0x34, 0x41, 0x29, 0x48, 0x94, 0x03, 0x42, 0x0d, 0x54, 0x80, 0x7a, 0x40, 0xdb, 0x9e, 0xb8, 0x20,
	0xd7, 0x99, 0xba, 0xab, 0xda, 0x6b, 0xd7, 0xbb, 0xae, 0xe8, 0x9d, 0x33, 0xbf, 0x91, 0x9f, 0x82,
	0xb2, 0xf6, 0xc6, 0xbb, 0x69, 0x44, 0x25, 0x7a, 0x9b, 0xef, 0xf7, 0x76, 0xe6, 0x59, 0x86, 0xe7,
	0x02, 0xf3, 0x2b, 0x16, 0xe2, 0x48, 0x16, 0x9c, 0x63, 0xcc, 0x78, 0x54, 0x5b, 0xc3, 0x2c, 0x4f,
	0x65, 0x4a, 0x3a, 0xf3, 0x80, 0xb7, 0x23, 0x30, 0xc6, 0x50, 0xb2, 0x94, 0x8f, 0xe6, 0x56, 0x59,
	0xe5, 0x3d, 0xad, 0x07, 0x84, 0x29, 0x3f, 0x63, 0x51, 0x91, 0x07, 0x46, 0x7a, 0xab, 0x4e, 0x0b,
	0x19, 0x48, 0xac, 0xc2, 0xdb, 0x8b, 0xb0, 0x65, 0xdc, 0xff, 0xd5, 0x80, 0xad, 0x49, 0x8e, 0x6a,
	0xc2, 0x71, 0x86, 0x21, 0x3b, 0x63, 0xa1, 0x72, 0xc8, 0x7b, 0xe8, 0x59, 0xf3, 0x5d, 0xa7, 0xef,
	0x0c, 0xba, 0x63, 0x77, 0x58, 0xd3, 0x9e, 0x98, 0x79, 0x6a, 0x97, 0x13, 0x02, 0x2d, 0x1e, 0x24,
	0xe8, 0x36, 0xfa, 0xce, 0xa0, 0x43, 0x95, 0x4d, 0x3e, 0x42, 0x3b, 0x0e, 0x4e, 0x31, 0x16, 0x6e,
	0xb3, 0xdf, 0x1c, 0x74, 0xc7, 0x2f, 0xcd, 0x61, 0xcb, 0x58, 0x0c, 0x8f, 0x54, 0xf9, 0x27, 0x2e,
	0xf3, 0x6b, 0x5a, 0xf5, 0x92, 0x6d, 0x68, 0x67, 0x41, 0x21, 0x70, 0xea, 0xb6, 0xfa, 0xce, 0x60,
	0x95, 0x56, 0x9e, 0xb7, 0x0f, 0x5d, 0xa3, 0x9c, 0x6c, 0x40, 0xf3, 0x02, 0xaf, 0x15, 0xed, 0x0e,
	0x9d, 0x99, 0x64, 0x13, 0x56, 0xae, 0x82, 0xb8, 0xd0, 0x9c, 0x4a, 0xe7, 0x5d, 0xe3, 0xad, 0xe3,
	0x0b, 0xe8, 0x29, 0x7c, 0xa4, 0x78, 0x59, 0xa0, 0x90, 0xe4, 0x10, 0x7a, 0xc2, 0x24, 0x52, 0xbd,
	0xbe, 0x7f, 0x1b, 0x61, 0x6a, 0xb7, 0x11, 0x0f, 0x56, 0x73, 0x14, 0x59, 0xca, 0x85, 0x46, 0x9d,
	0xfb, 0xfe, 0x6f, 0x07, 0xd6, 0x35, 0x6a, 0x19, 0x22, 0x5f, 0xe1, 0xc1, 0x79, 0x2a, 0xe4, 0x24,
	0xc7, 0x29, 0x72, 0xc9, 0x82, 0x58, 0x2c, 0x01, 0x3e, 0x51, 0xd6, 0x67, 0xbb, 0x8e, 0x2e, 0x36,
	0x12, 0x17, 0xee, 0x27, 0x28, 0x44, 0x10, 0x69, 0x64, 0xed, 0xaa, 0x05, 0xe6, 0x69, 0x92, 0x49,
	0xb7, 0xa9, 0x12, 0x95, 0xe7, 0x5f, 0x42, 0xf7, 0x88, 0x09, 0xa9, 0x77, 0x30, 0x86, 0xce, 0x5c,
	0x7c, 0x15, 0x8d, 0xcd, 0x61, 0x2d, 0xc7, 0x63, 0x6d, 0xd1, 0xba, 0x8c, 0x0c, 0x81, 0x64, 0x39,
	0x5e, 0xb1, 0xb4, 0x10, 0xc7, 0x33, 0xf9, 0x7d, 0xe1, 0x53, 0xfc, 0xa9, 0xf0, 0x5b, 0x74, 0x49,
	0xc6, 0x9f, 0xc2, 0x5a, 0x09, 0x59, 0x2d, 0xe0, 0x19, 0x80, 0xa8, 0xfb, 0x1c, 0xd5, 0x67, 0x44,
	0xc8, 0x6b, 0x58, 0x2b, 0x17, 0xa1, 0x66, 0x08, 0xb7, 0xa1, 0x74, 0xb4, 0x61, 0x6c, 0x47, 0x25,
	0xa8, 0x55, 0xe5, 0x1f, 0xc0, 0xda, 0xb7, 0x99, 0x46, 0xee, 0xf0, 0x32, 0xff, 0x05, 0xf4, 0xaa,
	0x19, 0x15, 0x55, 0x63, 0xbf, 0x8e, 0xb5, 0x5f, 0xff, 0x07, 0xf4, 0x28, 0x8a, 0x22, 0xb9, 0x0b,
	0xde, 0x3f, 0x95, 0x73, 0x00, 0xeb, 0x1a, 0xe0, 0x36, 0x32, 0xc6, 0xb1, 0x1b, 0xd6, 0xb1, 0x0f,
	0x61, 0xe3, 0x04, 0xf3, 0x84, 0x71, 0x43, 0xf5, 0xff, 0xb3, 0x97, 0x5d, 0x78, 0x68, 0xcc, 0xb9,
	0x8d, 0xce, 0xf8, 0x4f, 0x03, 0x3a, 0x27, 0xfa, 0x58, 0x64, 0x02, 0xed, 0xf2, 0x0b, 0x20, 0xee,
	0xe2, 0x97, 0xa5, 0x49, 0x79, 0x3b, 0x4b, 0x32, 0xd5, 0x1e, 0xee, 0x0d, 0x9c, 0x57, 0x0e, 0xd9,
	0x87, 0xd6, 0x4c, 0x43, 0x64, 0xdb, 0x28, 0x34, 0x74, 0xec, 0x3d, 0xba, 0x11, 0xd7, 0xed, 0xe4,
	0x03, 0xac, 0xa8, 0xa3, 0x12, 0xb3, 0xc6, 0x94, 0x8a, 0xe7, 0xde, 0x4c, 0x58, 0xe0, 0x13, 0x68,
	0x97, 0xa7, 0xb0, 0x5e, 0x60, 0x9d, 0xdf, 0x7a, 0x81, 0x7d, 0xb7, 0x6a, 0xc8, 0x11, 0x74, 0xe6,
	0x3b, 0x24, 0x8f, 0xcd, 0x4f, 0x7d, 0xe1, 0x42, 0xde, 0x93, 0xe5, 0x49, 0x73, 0xda, 0xc1, 0x9b,
	0xef, 0x7b, 0x11, 0x93, 0xe7, 0xc5, 0xe9, 0x30, 0x4c, 0x93, 0x51, 0x52, 0xc8, 0x20, 0x42, 0xbe,
	0xcb, 0x52, 0x6d, 0x8e, 0xb2, 0x8b, 0x68, 0x74, 0xe3, 0x77, 0x74, 0xda, 0x56, 0x7f, 0x84, 0xbd,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x49, 0xe7, 0x86, 0xaa, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TunnelingClient is the client API for Tunneling service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TunnelingClient interface {
	Create(ctx context.Context, opts ...grpc.CallOption) (Tunneling_CreateClient, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Pause(ctx context.Context, opts ...grpc.CallOption) (Tunneling_PauseClient, error)
	Resume(ctx context.Context, opts ...grpc.CallOption) (Tunneling_ResumeClient, error)
	Terminate(ctx context.Context, opts ...grpc.CallOption) (Tunneling_TerminateClient, error)
}

type tunnelingClient struct {
	cc grpc.ClientConnInterface
}

func NewTunnelingClient(cc grpc.ClientConnInterface) TunnelingClient {
	return &tunnelingClient{cc}
}

func (c *tunnelingClient) Create(ctx context.Context, opts ...grpc.CallOption) (Tunneling_CreateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Tunneling_serviceDesc.Streams[0], "/tunneling.Tunneling/Create", opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelingCreateClient{stream}
	return x, nil
}

type Tunneling_CreateClient interface {
	Send(*CreateRequest) error
	Recv() (*CreateResponse, error)
	grpc.ClientStream
}

type tunnelingCreateClient struct {
	grpc.ClientStream
}

func (x *tunnelingCreateClient) Send(m *CreateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelingCreateClient) Recv() (*CreateResponse, error) {
	m := new(CreateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tunnelingClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/tunneling.Tunneling/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunnelingClient) Pause(ctx context.Context, opts ...grpc.CallOption) (Tunneling_PauseClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Tunneling_serviceDesc.Streams[1], "/tunneling.Tunneling/Pause", opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelingPauseClient{stream}
	return x, nil
}

type Tunneling_PauseClient interface {
	Send(*PauseRequest) error
	Recv() (*PauseResponse, error)
	grpc.ClientStream
}

type tunnelingPauseClient struct {
	grpc.ClientStream
}

func (x *tunnelingPauseClient) Send(m *PauseRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelingPauseClient) Recv() (*PauseResponse, error) {
	m := new(PauseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tunnelingClient) Resume(ctx context.Context, opts ...grpc.CallOption) (Tunneling_ResumeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Tunneling_serviceDesc.Streams[2], "/tunneling.Tunneling/Resume", opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelingResumeClient{stream}
	return x, nil
}

type Tunneling_ResumeClient interface {
	Send(*ResumeRequest) error
	Recv() (*ResumeResponse, error)
	grpc.ClientStream
}

type tunnelingResumeClient struct {
	grpc.ClientStream
}

func (x *tunnelingResumeClient) Send(m *ResumeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelingResumeClient) Recv() (*ResumeResponse, error) {
	m := new(ResumeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tunnelingClient) Terminate(ctx context.Context, opts ...grpc.CallOption) (Tunneling_TerminateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Tunneling_serviceDesc.Streams[3], "/tunneling.Tunneling/Terminate", opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelingTerminateClient{stream}
	return x, nil
}

type Tunneling_TerminateClient interface {
	Send(*TerminateRequest) error
	Recv() (*TerminateResponse, error)
	grpc.ClientStream
}

type tunnelingTerminateClient struct {
	grpc.ClientStream
}

func (x *tunnelingTerminateClient) Send(m *TerminateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelingTerminateClient) Recv() (*TerminateResponse, error) {
	m := new(TerminateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TunnelingServer is the server API for Tunneling service.
type TunnelingServer interface {
	Create(Tunneling_CreateServer) error
	List(context.Context, *ListRequest) (*ListResponse, error)
	Pause(Tunneling_PauseServer) error
	Resume(Tunneling_ResumeServer) error
	Terminate(Tunneling_TerminateServer) error
}

// UnimplementedTunnelingServer can be embedded to have forward compatible implementations.
type UnimplementedTunnelingServer struct {
}

func (*UnimplementedTunnelingServer) Create(srv Tunneling_CreateServer) error {
	return status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedTunnelingServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedTunnelingServer) Pause(srv Tunneling_PauseServer) error {
	return status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (*UnimplementedTunnelingServer) Resume(srv Tunneling_ResumeServer) error {
	return status.Errorf(codes.Unimplemented, "method Resume not implemented")
}
func (*UnimplementedTunnelingServer) Terminate(srv Tunneling_TerminateServer) error {
	return status.Errorf(codes.Unimplemented, "method Terminate not implemented")
}

func RegisterTunnelingServer(s *grpc.Server, srv TunnelingServer) {
	s.RegisterService(&_Tunneling_serviceDesc, srv)
}

func _Tunneling_Create_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelingServer).Create(&tunnelingCreateServer{stream})
}

type Tunneling_CreateServer interface {
	Send(*CreateResponse) error
	Recv() (*CreateRequest, error)
	grpc.ServerStream
}

type tunnelingCreateServer struct {
	grpc.ServerStream
}

func (x *tunnelingCreateServer) Send(m *CreateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelingCreateServer) Recv() (*CreateRequest, error) {
	m := new(CreateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Tunneling_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelingServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tunneling.Tunneling/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelingServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tunneling_Pause_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelingServer).Pause(&tunnelingPauseServer{stream})
}

type Tunneling_PauseServer interface {
	Send(*PauseResponse) error
	Recv() (*PauseRequest, error)
	grpc.ServerStream
}

type tunnelingPauseServer struct {
	grpc.ServerStream
}

func (x *tunnelingPauseServer) Send(m *PauseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelingPauseServer) Recv() (*PauseRequest, error) {
	m := new(PauseRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Tunneling_Resume_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelingServer).Resume(&tunnelingResumeServer{stream})
}

type Tunneling_ResumeServer interface {
	Send(*ResumeResponse) error
	Recv() (*ResumeRequest, error)
	grpc.ServerStream
}

type tunnelingResumeServer struct {
	grpc.ServerStream
}

func (x *tunnelingResumeServer) Send(m *ResumeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelingResumeServer) Recv() (*ResumeRequest, error) {
	m := new(ResumeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Tunneling_Terminate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelingServer).Terminate(&tunnelingTerminateServer{stream})
}

type Tunneling_TerminateServer interface {
	Send(*TerminateResponse) error
	Recv() (*TerminateRequest, error)
	grpc.ServerStream
}

type tunnelingTerminateServer struct {
	grpc.ServerStream
}

func (x *tunnelingTerminateServer) Send(m *TerminateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelingTerminateServer) Recv() (*TerminateRequest, error) {
	m := new(TerminateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Tunneling_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tunneling.Tunneling",
	HandlerType: (*TunnelingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Tunneling_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Create",
			Handler:       _Tunneling_Create_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Pause",
			Handler:       _Tunneling_Pause_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Resume",
			Handler:       _Tunneling_Resume_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Terminate",
			Handler:       _Tunneling_Terminate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service/tunneling/tunneling.proto",
}
