// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gracefulexit.proto

package internalpb

import (
	context "context"
	fmt "fmt"
	math "math"
	time "time"

	proto "github.com/gogo/protobuf/proto"

	drpc "storj.io/drpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type GetNonExitingSatellitesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNonExitingSatellitesRequest) Reset()         { *m = GetNonExitingSatellitesRequest{} }
func (m *GetNonExitingSatellitesRequest) String() string { return proto.CompactTextString(m) }
func (*GetNonExitingSatellitesRequest) ProtoMessage()    {}
func (*GetNonExitingSatellitesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{0}
}
func (m *GetNonExitingSatellitesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNonExitingSatellitesRequest.Unmarshal(m, b)
}
func (m *GetNonExitingSatellitesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNonExitingSatellitesRequest.Marshal(b, m, deterministic)
}
func (m *GetNonExitingSatellitesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNonExitingSatellitesRequest.Merge(m, src)
}
func (m *GetNonExitingSatellitesRequest) XXX_Size() int {
	return xxx_messageInfo_GetNonExitingSatellitesRequest.Size(m)
}
func (m *GetNonExitingSatellitesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNonExitingSatellitesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNonExitingSatellitesRequest proto.InternalMessageInfo

type GetNonExitingSatellitesResponse struct {
	Satellites           []*NonExitingSatellite `protobuf:"bytes,1,rep,name=satellites,proto3" json:"satellites,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetNonExitingSatellitesResponse) Reset()         { *m = GetNonExitingSatellitesResponse{} }
func (m *GetNonExitingSatellitesResponse) String() string { return proto.CompactTextString(m) }
func (*GetNonExitingSatellitesResponse) ProtoMessage()    {}
func (*GetNonExitingSatellitesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{1}
}
func (m *GetNonExitingSatellitesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNonExitingSatellitesResponse.Unmarshal(m, b)
}
func (m *GetNonExitingSatellitesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNonExitingSatellitesResponse.Marshal(b, m, deterministic)
}
func (m *GetNonExitingSatellitesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNonExitingSatellitesResponse.Merge(m, src)
}
func (m *GetNonExitingSatellitesResponse) XXX_Size() int {
	return xxx_messageInfo_GetNonExitingSatellitesResponse.Size(m)
}
func (m *GetNonExitingSatellitesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNonExitingSatellitesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNonExitingSatellitesResponse proto.InternalMessageInfo

func (m *GetNonExitingSatellitesResponse) GetSatellites() []*NonExitingSatellite {
	if m != nil {
		return m.Satellites
	}
	return nil
}

// NonExitingSatellite contains information that's needed for a storagenode to start graceful exit.
type NonExitingSatellite struct {
	NodeId               NodeID   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	DomainName           string   `protobuf:"bytes,2,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty"`
	SpaceUsed            float64  `protobuf:"fixed64,3,opt,name=space_used,json=spaceUsed,proto3" json:"space_used,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NonExitingSatellite) Reset()         { *m = NonExitingSatellite{} }
func (m *NonExitingSatellite) String() string { return proto.CompactTextString(m) }
func (*NonExitingSatellite) ProtoMessage()    {}
func (*NonExitingSatellite) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{2}
}
func (m *NonExitingSatellite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonExitingSatellite.Unmarshal(m, b)
}
func (m *NonExitingSatellite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonExitingSatellite.Marshal(b, m, deterministic)
}
func (m *NonExitingSatellite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonExitingSatellite.Merge(m, src)
}
func (m *NonExitingSatellite) XXX_Size() int {
	return xxx_messageInfo_NonExitingSatellite.Size(m)
}
func (m *NonExitingSatellite) XXX_DiscardUnknown() {
	xxx_messageInfo_NonExitingSatellite.DiscardUnknown(m)
}

var xxx_messageInfo_NonExitingSatellite proto.InternalMessageInfo

func (m *NonExitingSatellite) GetDomainName() string {
	if m != nil {
		return m.DomainName
	}
	return ""
}

func (m *NonExitingSatellite) GetSpaceUsed() float64 {
	if m != nil {
		return m.SpaceUsed
	}
	return 0
}

type InitiateGracefulExitRequest struct {
	NodeId               NodeID   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitiateGracefulExitRequest) Reset()         { *m = InitiateGracefulExitRequest{} }
func (m *InitiateGracefulExitRequest) String() string { return proto.CompactTextString(m) }
func (*InitiateGracefulExitRequest) ProtoMessage()    {}
func (*InitiateGracefulExitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{3}
}
func (m *InitiateGracefulExitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitiateGracefulExitRequest.Unmarshal(m, b)
}
func (m *InitiateGracefulExitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitiateGracefulExitRequest.Marshal(b, m, deterministic)
}
func (m *InitiateGracefulExitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitiateGracefulExitRequest.Merge(m, src)
}
func (m *InitiateGracefulExitRequest) XXX_Size() int {
	return xxx_messageInfo_InitiateGracefulExitRequest.Size(m)
}
func (m *InitiateGracefulExitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitiateGracefulExitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitiateGracefulExitRequest proto.InternalMessageInfo

type GetExitProgressRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetExitProgressRequest) Reset()         { *m = GetExitProgressRequest{} }
func (m *GetExitProgressRequest) String() string { return proto.CompactTextString(m) }
func (*GetExitProgressRequest) ProtoMessage()    {}
func (*GetExitProgressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{4}
}
func (m *GetExitProgressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetExitProgressRequest.Unmarshal(m, b)
}
func (m *GetExitProgressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetExitProgressRequest.Marshal(b, m, deterministic)
}
func (m *GetExitProgressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetExitProgressRequest.Merge(m, src)
}
func (m *GetExitProgressRequest) XXX_Size() int {
	return xxx_messageInfo_GetExitProgressRequest.Size(m)
}
func (m *GetExitProgressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetExitProgressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetExitProgressRequest proto.InternalMessageInfo

type GetExitProgressResponse struct {
	Progress             []*ExitProgress `protobuf:"bytes,1,rep,name=progress,proto3" json:"progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetExitProgressResponse) Reset()         { *m = GetExitProgressResponse{} }
func (m *GetExitProgressResponse) String() string { return proto.CompactTextString(m) }
func (*GetExitProgressResponse) ProtoMessage()    {}
func (*GetExitProgressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{5}
}
func (m *GetExitProgressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetExitProgressResponse.Unmarshal(m, b)
}
func (m *GetExitProgressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetExitProgressResponse.Marshal(b, m, deterministic)
}
func (m *GetExitProgressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetExitProgressResponse.Merge(m, src)
}
func (m *GetExitProgressResponse) XXX_Size() int {
	return xxx_messageInfo_GetExitProgressResponse.Size(m)
}
func (m *GetExitProgressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetExitProgressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetExitProgressResponse proto.InternalMessageInfo

func (m *GetExitProgressResponse) GetProgress() []*ExitProgress {
	if m != nil {
		return m.Progress
	}
	return nil
}

type ExitProgress struct {
	DomainName           string   `protobuf:"bytes,1,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty"`
	NodeId               NodeID   `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	PercentComplete      float32  `protobuf:"fixed32,3,opt,name=percent_complete,json=percentComplete,proto3" json:"percent_complete,omitempty"`
	Successful           bool     `protobuf:"varint,4,opt,name=successful,proto3" json:"successful,omitempty"`
	CompletionReceipt    []byte   `protobuf:"bytes,5,opt,name=completion_receipt,json=completionReceipt,proto3" json:"completion_receipt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExitProgress) Reset()         { *m = ExitProgress{} }
func (m *ExitProgress) String() string { return proto.CompactTextString(m) }
func (*ExitProgress) ProtoMessage()    {}
func (*ExitProgress) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{6}
}
func (m *ExitProgress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExitProgress.Unmarshal(m, b)
}
func (m *ExitProgress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExitProgress.Marshal(b, m, deterministic)
}
func (m *ExitProgress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitProgress.Merge(m, src)
}
func (m *ExitProgress) XXX_Size() int {
	return xxx_messageInfo_ExitProgress.Size(m)
}
func (m *ExitProgress) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitProgress.DiscardUnknown(m)
}

var xxx_messageInfo_ExitProgress proto.InternalMessageInfo

func (m *ExitProgress) GetDomainName() string {
	if m != nil {
		return m.DomainName
	}
	return ""
}

func (m *ExitProgress) GetPercentComplete() float32 {
	if m != nil {
		return m.PercentComplete
	}
	return 0
}

func (m *ExitProgress) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

func (m *ExitProgress) GetCompletionReceipt() []byte {
	if m != nil {
		return m.CompletionReceipt
	}
	return nil
}

type GracefulExitFeasibilityRequest struct {
	NodeId               NodeID   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GracefulExitFeasibilityRequest) Reset()         { *m = GracefulExitFeasibilityRequest{} }
func (m *GracefulExitFeasibilityRequest) String() string { return proto.CompactTextString(m) }
func (*GracefulExitFeasibilityRequest) ProtoMessage()    {}
func (*GracefulExitFeasibilityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{7}
}
func (m *GracefulExitFeasibilityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GracefulExitFeasibilityRequest.Unmarshal(m, b)
}
func (m *GracefulExitFeasibilityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GracefulExitFeasibilityRequest.Marshal(b, m, deterministic)
}
func (m *GracefulExitFeasibilityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GracefulExitFeasibilityRequest.Merge(m, src)
}
func (m *GracefulExitFeasibilityRequest) XXX_Size() int {
	return xxx_messageInfo_GracefulExitFeasibilityRequest.Size(m)
}
func (m *GracefulExitFeasibilityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GracefulExitFeasibilityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GracefulExitFeasibilityRequest proto.InternalMessageInfo

type GracefulExitFeasibilityResponse struct {
	JoinedAt             time.Time `protobuf:"bytes,1,opt,name=joined_at,json=joinedAt,proto3,stdtime" json:"joined_at"`
	MonthsRequired       int32     `protobuf:"varint,2,opt,name=months_required,json=monthsRequired,proto3" json:"months_required,omitempty"`
	IsAllowed            bool      `protobuf:"varint,3,opt,name=is_allowed,json=isAllowed,proto3" json:"is_allowed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GracefulExitFeasibilityResponse) Reset()         { *m = GracefulExitFeasibilityResponse{} }
func (m *GracefulExitFeasibilityResponse) String() string { return proto.CompactTextString(m) }
func (*GracefulExitFeasibilityResponse) ProtoMessage()    {}
func (*GracefulExitFeasibilityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{8}
}
func (m *GracefulExitFeasibilityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GracefulExitFeasibilityResponse.Unmarshal(m, b)
}
func (m *GracefulExitFeasibilityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GracefulExitFeasibilityResponse.Marshal(b, m, deterministic)
}
func (m *GracefulExitFeasibilityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GracefulExitFeasibilityResponse.Merge(m, src)
}
func (m *GracefulExitFeasibilityResponse) XXX_Size() int {
	return xxx_messageInfo_GracefulExitFeasibilityResponse.Size(m)
}
func (m *GracefulExitFeasibilityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GracefulExitFeasibilityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GracefulExitFeasibilityResponse proto.InternalMessageInfo

func (m *GracefulExitFeasibilityResponse) GetJoinedAt() time.Time {
	if m != nil {
		return m.JoinedAt
	}
	return time.Time{}
}

func (m *GracefulExitFeasibilityResponse) GetMonthsRequired() int32 {
	if m != nil {
		return m.MonthsRequired
	}
	return 0
}

func (m *GracefulExitFeasibilityResponse) GetIsAllowed() bool {
	if m != nil {
		return m.IsAllowed
	}
	return false
}

func init() {
	proto.RegisterType((*GetNonExitingSatellitesRequest)(nil), "storagenode.gracefulexit.GetNonExitingSatellitesRequest")
	proto.RegisterType((*GetNonExitingSatellitesResponse)(nil), "storagenode.gracefulexit.GetNonExitingSatellitesResponse")
	proto.RegisterType((*NonExitingSatellite)(nil), "storagenode.gracefulexit.NonExitingSatellite")
	proto.RegisterType((*InitiateGracefulExitRequest)(nil), "storagenode.gracefulexit.InitiateGracefulExitRequest")
	proto.RegisterType((*GetExitProgressRequest)(nil), "storagenode.gracefulexit.GetExitProgressRequest")
	proto.RegisterType((*GetExitProgressResponse)(nil), "storagenode.gracefulexit.GetExitProgressResponse")
	proto.RegisterType((*ExitProgress)(nil), "storagenode.gracefulexit.ExitProgress")
	proto.RegisterType((*GracefulExitFeasibilityRequest)(nil), "storagenode.gracefulexit.GracefulExitFeasibilityRequest")
	proto.RegisterType((*GracefulExitFeasibilityResponse)(nil), "storagenode.gracefulexit.GracefulExitFeasibilityResponse")
}

func init() { proto.RegisterFile("gracefulexit.proto", fileDescriptor_8f0acbf2ce5fa631) }

var fileDescriptor_8f0acbf2ce5fa631 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xfd, 0xb9, 0xff, 0x7e, 0xc9, 0xb4, 0x6a, 0xcb, 0x82, 0xc0, 0x0a, 0xa2, 0x8e, 0x2c, 0x41,
	0xc3, 0xa1, 0x0e, 0x14, 0x21, 0xc1, 0xb1, 0x01, 0x5a, 0xe5, 0x40, 0x85, 0x16, 0xb8, 0x20, 0x21,
	0x6b, 0x63, 0x4f, 0xcd, 0x56, 0xf6, 0xae, 0xeb, 0x5d, 0x43, 0xb9, 0xf0, 0x11, 0x10, 0xdf, 0x81,
	0x2b, 0x1f, 0x84, 0x33, 0x47, 0x0e, 0xe5, 0xab, 0x20, 0xdb, 0xdb, 0x60, 0xda, 0xd8, 0x6a, 0xb9,
	0x25, 0x6f, 0x66, 0x9e, 0x67, 0xde, 0xbe, 0x07, 0x24, 0xca, 0x58, 0x80, 0x07, 0x79, 0x8c, 0xc7,
	0x5c, 0x7b, 0x69, 0x26, 0xb5, 0x24, 0xb6, 0xd2, 0x32, 0x63, 0x11, 0x0a, 0x19, 0xa2, 0x57, 0xaf,
	0xf7, 0x20, 0x92, 0x91, 0xac, 0xba, 0x7a, 0x4e, 0x24, 0x65, 0x14, 0xe3, 0xb0, 0xfc, 0x37, 0xc9,
	0x0f, 0x86, 0x9a, 0x27, 0xa8, 0x34, 0x4b, 0xd2, 0xaa, 0xc1, 0xed, 0xc3, 0xc6, 0x1e, 0xea, 0x7d,
	0x29, 0x9e, 0x1d, 0x73, 0xcd, 0x45, 0xf4, 0x92, 0x69, 0x8c, 0x63, 0xae, 0x51, 0x51, 0x3c, 0xca,
	0x51, 0x69, 0x37, 0x05, 0xa7, 0xb1, 0x43, 0xa5, 0x52, 0x28, 0x24, 0xcf, 0x01, 0xd4, 0x14, 0xb5,
	0xad, 0xfe, 0xfc, 0x60, 0x79, 0x7b, 0xcb, 0x6b, 0x5a, 0xd0, 0x9b, 0xc1, 0x45, 0x6b, 0x04, 0xee,
	0x27, 0xb8, 0x3a, 0xa3, 0x85, 0x6c, 0xc2, 0xff, 0x05, 0x97, 0xcf, 0x43, 0xdb, 0xea, 0x5b, 0x83,
	0x95, 0xd1, 0xea, 0xf7, 0x13, 0xe7, 0xbf, 0x9f, 0x27, 0xce, 0xd2, 0xbe, 0x0c, 0x71, 0xfc, 0x94,
	0x2e, 0x15, 0xe5, 0x71, 0x48, 0x1c, 0x58, 0x0e, 0x65, 0xc2, 0xb8, 0xf0, 0x05, 0x4b, 0xd0, 0x9e,
	0xeb, 0x5b, 0x83, 0x2e, 0x85, 0x0a, 0xda, 0x67, 0x09, 0x92, 0x5b, 0x00, 0x2a, 0x65, 0x01, 0xfa,
	0xb9, 0xc2, 0xd0, 0x9e, 0xef, 0x5b, 0x03, 0x8b, 0x76, 0x4b, 0xe4, 0xb5, 0xc2, 0xd0, 0xdd, 0x85,
	0x9b, 0x63, 0xc1, 0x35, 0x67, 0x1a, 0xf7, 0xcc, 0xde, 0xc5, 0x32, 0x46, 0x90, 0x0b, 0xef, 0xe1,
	0xda, 0x70, 0x7d, 0x0f, 0x75, 0x31, 0xfa, 0x22, 0x93, 0x51, 0x86, 0x6a, 0xaa, 0xe9, 0x5b, 0xb8,
	0x71, 0xae, 0x62, 0xb4, 0x1c, 0x41, 0x27, 0x35, 0x98, 0x51, 0xf2, 0x4e, 0xb3, 0x92, 0x7f, 0x31,
	0x4c, 0xe7, 0xdc, 0x1f, 0x16, 0xac, 0xd4, 0x4b, 0x67, 0x15, 0xb1, 0xce, 0x29, 0x52, 0xbb, 0x69,
	0xae, 0x55, 0xdb, 0xbb, 0xb0, 0x9e, 0x62, 0x16, 0xa0, 0xd0, 0x7e, 0x20, 0x93, 0x34, 0x46, 0x8d,
	0xa5, 0x80, 0x73, 0x74, 0xcd, 0xe0, 0x4f, 0x0c, 0x4c, 0x36, 0x00, 0x54, 0x1e, 0x04, 0xa8, 0xd4,
	0x41, 0x1e, 0xdb, 0x0b, 0x7d, 0x6b, 0xd0, 0xa1, 0x35, 0x84, 0x6c, 0x01, 0x31, 0x14, 0x5c, 0x0a,
	0x3f, 0xc3, 0x00, 0x79, 0xaa, 0xed, 0xc5, 0xe2, 0xf3, 0xf4, 0xca, 0x9f, 0x0a, 0xad, 0x0a, 0xee,
	0x18, 0x36, 0xea, 0xaf, 0xb1, 0x8b, 0x4c, 0xf1, 0x09, 0x8f, 0xb9, 0xfe, 0x78, 0xe9, 0x87, 0xf9,
	0x66, 0x81, 0xd3, 0xc8, 0x65, 0xde, 0x61, 0x07, 0xba, 0x87, 0x92, 0x0b, 0x0c, 0x7d, 0xa6, 0x4b,
	0xba, 0xe5, 0xed, 0x9e, 0x57, 0xa5, 0xc9, 0x3b, 0x4d, 0x93, 0xf7, 0xea, 0x34, 0x4d, 0xa3, 0x4e,
	0xf1, 0xa9, 0x2f, 0xbf, 0x1c, 0x8b, 0x76, 0xaa, 0xb1, 0x9d, 0x62, 0x9f, 0xb5, 0x44, 0x0a, 0xfd,
	0x4e, 0xf9, 0x19, 0x1e, 0xe5, 0x3c, 0xc3, 0x4a, 0xdc, 0x45, 0xba, 0x5a, 0xc1, 0xd4, 0xa0, 0x85,
	0x1f, 0xb9, 0xf2, 0x59, 0x1c, 0xcb, 0x0f, 0xc6, 0x8f, 0x1d, 0xda, 0xe5, 0x6a, 0xa7, 0x02, 0xb6,
	0xbf, 0x2e, 0xc0, 0x7a, 0x71, 0x41, 0x7d, 0x65, 0xf2, 0xd9, 0x2a, 0x3d, 0x34, 0x2b, 0x97, 0xe4,
	0x51, 0xb3, 0x63, 0xda, 0xc3, 0xde, 0x7b, 0xfc, 0x0f, 0x93, 0x46, 0xb0, 0x1c, 0xae, 0xcd, 0x4a,
	0x0d, 0x79, 0xd8, 0x4c, 0xd9, 0x92, 0xb2, 0xde, 0x05, 0x5d, 0x4f, 0xde, 0xc3, 0xda, 0x99, 0x28,
	0x91, 0x7b, 0xad, 0x47, 0xcc, 0xc8, 0x63, 0xef, 0xfe, 0x25, 0x26, 0xcc, 0xb9, 0xa5, 0xfe, 0xb3,
	0x3d, 0xd4, 0xaa, 0x7f, 0xab, 0x85, 0x5b, 0xf5, 0x6f, 0x37, 0xec, 0x68, 0xf3, 0xcd, 0xed, 0x62,
	0xf6, 0xd0, 0xe3, 0x72, 0x58, 0xfe, 0x18, 0xd6, 0xa8, 0x86, 0x5c, 0x68, 0xcc, 0x04, 0x8b, 0xd3,
	0xc9, 0x64, 0xa9, 0xb4, 0xef, 0x83, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x75, 0x55, 0x23,
	0x56, 0x06, 0x00, 0x00,
}

// --- DRPC BEGIN ---

type DRPCNodeGracefulExitClient interface {
	DRPCConn() drpc.Conn

	// GetSatellitesList returns a list of satellites that the storagenode has not exited.
	GetNonExitingSatellites(ctx context.Context, in *GetNonExitingSatellitesRequest) (*GetNonExitingSatellitesResponse, error)
	// InitiateGracefulExit updates one or more satellites in the storagenode's database to be gracefully exiting.
	InitiateGracefulExit(ctx context.Context, in *InitiateGracefulExitRequest) (*ExitProgress, error)
	// GetExitProgress returns graceful exit status on each satellite for a given storagenode.
	GetExitProgress(ctx context.Context, in *GetExitProgressRequest) (*GetExitProgressResponse, error)
	// GracefulExitFeasibility returns node's join date and satellites config's amount of months required for graceful exit to be allowed.
	GracefulExitFeasibility(ctx context.Context, in *GracefulExitFeasibilityRequest) (*GracefulExitFeasibilityResponse, error)
}

type drpcNodeGracefulExitClient struct {
	cc drpc.Conn
}

func NewDRPCNodeGracefulExitClient(cc drpc.Conn) DRPCNodeGracefulExitClient {
	return &drpcNodeGracefulExitClient{cc}
}

func (c *drpcNodeGracefulExitClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcNodeGracefulExitClient) GetNonExitingSatellites(ctx context.Context, in *GetNonExitingSatellitesRequest) (*GetNonExitingSatellitesResponse, error) {
	out := new(GetNonExitingSatellitesResponse)
	err := c.cc.Invoke(ctx, "/storagenode.gracefulexit.NodeGracefulExit/GetNonExitingSatellites", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcNodeGracefulExitClient) InitiateGracefulExit(ctx context.Context, in *InitiateGracefulExitRequest) (*ExitProgress, error) {
	out := new(ExitProgress)
	err := c.cc.Invoke(ctx, "/storagenode.gracefulexit.NodeGracefulExit/InitiateGracefulExit", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcNodeGracefulExitClient) GetExitProgress(ctx context.Context, in *GetExitProgressRequest) (*GetExitProgressResponse, error) {
	out := new(GetExitProgressResponse)
	err := c.cc.Invoke(ctx, "/storagenode.gracefulexit.NodeGracefulExit/GetExitProgress", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcNodeGracefulExitClient) GracefulExitFeasibility(ctx context.Context, in *GracefulExitFeasibilityRequest) (*GracefulExitFeasibilityResponse, error) {
	out := new(GracefulExitFeasibilityResponse)
	err := c.cc.Invoke(ctx, "/storagenode.gracefulexit.NodeGracefulExit/GracefulExitFeasibility", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCNodeGracefulExitServer interface {
	// GetSatellitesList returns a list of satellites that the storagenode has not exited.
	GetNonExitingSatellites(context.Context, *GetNonExitingSatellitesRequest) (*GetNonExitingSatellitesResponse, error)
	// InitiateGracefulExit updates one or more satellites in the storagenode's database to be gracefully exiting.
	InitiateGracefulExit(context.Context, *InitiateGracefulExitRequest) (*ExitProgress, error)
	// GetExitProgress returns graceful exit status on each satellite for a given storagenode.
	GetExitProgress(context.Context, *GetExitProgressRequest) (*GetExitProgressResponse, error)
	// GracefulExitFeasibility returns node's join date and satellites config's amount of months required for graceful exit to be allowed.
	GracefulExitFeasibility(context.Context, *GracefulExitFeasibilityRequest) (*GracefulExitFeasibilityResponse, error)
}

type DRPCNodeGracefulExitDescription struct{}

func (DRPCNodeGracefulExitDescription) NumMethods() int { return 4 }

func (DRPCNodeGracefulExitDescription) Method(n int) (string, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/storagenode.gracefulexit.NodeGracefulExit/GetNonExitingSatellites",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCNodeGracefulExitServer).
					GetNonExitingSatellites(
						ctx,
						in1.(*GetNonExitingSatellitesRequest),
					)
			}, DRPCNodeGracefulExitServer.GetNonExitingSatellites, true
	case 1:
		return "/storagenode.gracefulexit.NodeGracefulExit/InitiateGracefulExit",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCNodeGracefulExitServer).
					InitiateGracefulExit(
						ctx,
						in1.(*InitiateGracefulExitRequest),
					)
			}, DRPCNodeGracefulExitServer.InitiateGracefulExit, true
	case 2:
		return "/storagenode.gracefulexit.NodeGracefulExit/GetExitProgress",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCNodeGracefulExitServer).
					GetExitProgress(
						ctx,
						in1.(*GetExitProgressRequest),
					)
			}, DRPCNodeGracefulExitServer.GetExitProgress, true
	case 3:
		return "/storagenode.gracefulexit.NodeGracefulExit/GracefulExitFeasibility",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCNodeGracefulExitServer).
					GracefulExitFeasibility(
						ctx,
						in1.(*GracefulExitFeasibilityRequest),
					)
			}, DRPCNodeGracefulExitServer.GracefulExitFeasibility, true
	default:
		return "", nil, nil, false
	}
}

func DRPCRegisterNodeGracefulExit(mux drpc.Mux, impl DRPCNodeGracefulExitServer) error {
	return mux.Register(impl, DRPCNodeGracefulExitDescription{})
}

type DRPCNodeGracefulExit_GetNonExitingSatellitesStream interface {
	drpc.Stream
	SendAndClose(*GetNonExitingSatellitesResponse) error
}

type drpcNodeGracefulExitGetNonExitingSatellitesStream struct {
	drpc.Stream
}

func (x *drpcNodeGracefulExitGetNonExitingSatellitesStream) SendAndClose(m *GetNonExitingSatellitesResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCNodeGracefulExit_InitiateGracefulExitStream interface {
	drpc.Stream
	SendAndClose(*ExitProgress) error
}

type drpcNodeGracefulExitInitiateGracefulExitStream struct {
	drpc.Stream
}

func (x *drpcNodeGracefulExitInitiateGracefulExitStream) SendAndClose(m *ExitProgress) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCNodeGracefulExit_GetExitProgressStream interface {
	drpc.Stream
	SendAndClose(*GetExitProgressResponse) error
}

type drpcNodeGracefulExitGetExitProgressStream struct {
	drpc.Stream
}

func (x *drpcNodeGracefulExitGetExitProgressStream) SendAndClose(m *GetExitProgressResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCNodeGracefulExit_GracefulExitFeasibilityStream interface {
	drpc.Stream
	SendAndClose(*GracefulExitFeasibilityResponse) error
}

type drpcNodeGracefulExitGracefulExitFeasibilityStream struct {
	drpc.Stream
}

func (x *drpcNodeGracefulExitGracefulExitFeasibilityStream) SendAndClose(m *GracefulExitFeasibilityResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

// --- DRPC END ---
