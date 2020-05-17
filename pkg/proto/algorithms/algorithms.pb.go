// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/proto/algorithms/algorithms.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type FindMinimumNumberRequest struct {
	Numbers              []float32 `protobuf:"fixed32,1,rep,packed,name=numbers,proto3" json:"numbers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FindMinimumNumberRequest) Reset()         { *m = FindMinimumNumberRequest{} }
func (m *FindMinimumNumberRequest) String() string { return proto.CompactTextString(m) }
func (*FindMinimumNumberRequest) ProtoMessage()    {}
func (*FindMinimumNumberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b9be65ee1798038, []int{0}
}

func (m *FindMinimumNumberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMinimumNumberRequest.Unmarshal(m, b)
}
func (m *FindMinimumNumberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMinimumNumberRequest.Marshal(b, m, deterministic)
}
func (m *FindMinimumNumberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMinimumNumberRequest.Merge(m, src)
}
func (m *FindMinimumNumberRequest) XXX_Size() int {
	return xxx_messageInfo_FindMinimumNumberRequest.Size(m)
}
func (m *FindMinimumNumberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMinimumNumberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMinimumNumberRequest proto.InternalMessageInfo

func (m *FindMinimumNumberRequest) GetNumbers() []float32 {
	if m != nil {
		return m.Numbers
	}
	return nil
}

type FindMinimumNumberResponse struct {
	MinimumNumber        float32  `protobuf:"fixed32,1,opt,name=minimum_number,json=minimumNumber,proto3" json:"minimum_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMinimumNumberResponse) Reset()         { *m = FindMinimumNumberResponse{} }
func (m *FindMinimumNumberResponse) String() string { return proto.CompactTextString(m) }
func (*FindMinimumNumberResponse) ProtoMessage()    {}
func (*FindMinimumNumberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b9be65ee1798038, []int{1}
}

func (m *FindMinimumNumberResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMinimumNumberResponse.Unmarshal(m, b)
}
func (m *FindMinimumNumberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMinimumNumberResponse.Marshal(b, m, deterministic)
}
func (m *FindMinimumNumberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMinimumNumberResponse.Merge(m, src)
}
func (m *FindMinimumNumberResponse) XXX_Size() int {
	return xxx_messageInfo_FindMinimumNumberResponse.Size(m)
}
func (m *FindMinimumNumberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMinimumNumberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMinimumNumberResponse proto.InternalMessageInfo

func (m *FindMinimumNumberResponse) GetMinimumNumber() float32 {
	if m != nil {
		return m.MinimumNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*FindMinimumNumberRequest)(nil), "proto.FindMinimumNumberRequest")
	proto.RegisterType((*FindMinimumNumberResponse)(nil), "proto.FindMinimumNumberResponse")
}

func init() {
	proto.RegisterFile("pkg/proto/algorithms/algorithms.proto", fileDescriptor_1b9be65ee1798038)
}

var fileDescriptor_1b9be65ee1798038 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0xc8, 0x4e, 0xd7,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0xcc, 0x49, 0xcf, 0x2f, 0xca, 0x2c, 0xc9, 0xc8, 0x2d,
	0x46, 0x62, 0xea, 0x81, 0xe5, 0x84, 0x58, 0xc1, 0x94, 0x94, 0x7c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e,
	0x2a, 0x44, 0x43, 0x52, 0x69, 0x9a, 0x7e, 0x49, 0x66, 0x6e, 0x6a, 0x71, 0x49, 0x62, 0x6e, 0x01,
	0x44, 0x9d, 0x94, 0x0c, 0x54, 0x41, 0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49, 0x62,
	0x49, 0x66, 0x7e, 0x1e, 0xd4, 0x14, 0x25, 0x13, 0x2e, 0x09, 0xb7, 0xcc, 0xbc, 0x14, 0xdf, 0xcc,
	0xbc, 0xcc, 0xdc, 0xd2, 0x5c, 0xbf, 0xd2, 0xdc, 0xa4, 0xd4, 0xa2, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4,
	0xe2, 0x12, 0x21, 0x09, 0x2e, 0xf6, 0x3c, 0xb0, 0x40, 0xb1, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0x53,
	0x10, 0x8c, 0xab, 0xe4, 0xc4, 0x25, 0x89, 0x45, 0x57, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90,
	0x2a, 0x17, 0x5f, 0x2e, 0x44, 0x22, 0x1e, 0xa2, 0x5e, 0x82, 0x51, 0x81, 0x51, 0x83, 0x29, 0x88,
	0x37, 0x17, 0x59, 0xb9, 0xd1, 0x24, 0x46, 0x2e, 0x2e, 0x47, 0xb8, 0xa7, 0x84, 0x9a, 0x19, 0xb9,
	0x04, 0x31, 0xcc, 0x14, 0x92, 0x87, 0x38, 0x53, 0x0f, 0x97, 0x1b, 0xa5, 0x14, 0x70, 0x2b, 0x80,
	0x38, 0x47, 0x49, 0xbb, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0xaa, 0x4a, 0x0a, 0xfa, 0x65, 0x86, 0xc8,
	0x01, 0x8a, 0xa1, 0xc3, 0x8a, 0x51, 0x2b, 0x89, 0x0d, 0x6c, 0x9a, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x3b, 0x12, 0xa9, 0x8c, 0x84, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlgorithmsClient is the client API for Algorithms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlgorithmsClient interface {
	FindMinimumNumber(ctx context.Context, in *FindMinimumNumberRequest, opts ...grpc.CallOption) (*FindMinimumNumberResponse, error)
}

type algorithmsClient struct {
	cc *grpc.ClientConn
}

func NewAlgorithmsClient(cc *grpc.ClientConn) AlgorithmsClient {
	return &algorithmsClient{cc}
}

func (c *algorithmsClient) FindMinimumNumber(ctx context.Context, in *FindMinimumNumberRequest, opts ...grpc.CallOption) (*FindMinimumNumberResponse, error) {
	out := new(FindMinimumNumberResponse)
	err := c.cc.Invoke(ctx, "/proto.Algorithms/FindMinimumNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlgorithmsServer is the server API for Algorithms service.
type AlgorithmsServer interface {
	FindMinimumNumber(context.Context, *FindMinimumNumberRequest) (*FindMinimumNumberResponse, error)
}

// UnimplementedAlgorithmsServer can be embedded to have forward compatible implementations.
type UnimplementedAlgorithmsServer struct {
}

func (*UnimplementedAlgorithmsServer) FindMinimumNumber(ctx context.Context, req *FindMinimumNumberRequest) (*FindMinimumNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindMinimumNumber not implemented")
}

func RegisterAlgorithmsServer(s *grpc.Server, srv AlgorithmsServer) {
	s.RegisterService(&_Algorithms_serviceDesc, srv)
}

func _Algorithms_FindMinimumNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindMinimumNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgorithmsServer).FindMinimumNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Algorithms/FindMinimumNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgorithmsServer).FindMinimumNumber(ctx, req.(*FindMinimumNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Algorithms_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Algorithms",
	HandlerType: (*AlgorithmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindMinimumNumber",
			Handler:    _Algorithms_FindMinimumNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/algorithms/algorithms.proto",
}
