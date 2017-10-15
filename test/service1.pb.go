// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test/service1.proto

/*
Package proto_test is a generated protocol buffer package.

It is generated from these files:
	test/service1.proto

It has these top-level messages:
	TestRequest1
	TestResponse1
*/
package proto_test

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type TestRequest1 struct {
	Query         string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	PageNumber    int32  `protobuf:"varint,2,opt,name=page_number,json=pageNumber" json:"page_number,omitempty"`
	ResultPerPage int32  `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage" json:"result_per_page,omitempty"`
}

func (m *TestRequest1) Reset()                    { *m = TestRequest1{} }
func (m *TestRequest1) String() string            { return proto.CompactTextString(m) }
func (*TestRequest1) ProtoMessage()               {}
func (*TestRequest1) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TestRequest1) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *TestRequest1) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *TestRequest1) GetResultPerPage() int32 {
	if m != nil {
		return m.ResultPerPage
	}
	return 0
}

type TestResponse1 struct {
	Query         string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	PageNumber    int32  `protobuf:"varint,2,opt,name=page_number,json=pageNumber" json:"page_number,omitempty"`
	ResultPerPage int32  `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage" json:"result_per_page,omitempty"`
}

func (m *TestResponse1) Reset()                    { *m = TestResponse1{} }
func (m *TestResponse1) String() string            { return proto.CompactTextString(m) }
func (*TestResponse1) ProtoMessage()               {}
func (*TestResponse1) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TestResponse1) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *TestResponse1) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *TestResponse1) GetResultPerPage() int32 {
	if m != nil {
		return m.ResultPerPage
	}
	return 0
}

func init() {
	proto.RegisterType((*TestRequest1)(nil), "proto_test.TestRequest1")
	proto.RegisterType((*TestResponse1)(nil), "proto_test.TestResponse1")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestService1 service

type TestService1Client interface {
	SayHello(ctx context.Context, in *TestRequest1, opts ...grpc.CallOption) (*TestResponse1, error)
}

type testService1Client struct {
	cc *grpc.ClientConn
}

func NewTestService1Client(cc *grpc.ClientConn) TestService1Client {
	return &testService1Client{cc}
}

func (c *testService1Client) SayHello(ctx context.Context, in *TestRequest1, opts ...grpc.CallOption) (*TestResponse1, error) {
	out := new(TestResponse1)
	err := grpc.Invoke(ctx, "/proto_test.TestService1/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestService1 service

type TestService1Server interface {
	SayHello(context.Context, *TestRequest1) (*TestResponse1, error)
}

func RegisterTestService1Server(s *grpc.Server, srv TestService1Server) {
	s.RegisterService(&_TestService1_serviceDesc, srv)
}

func _TestService1_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestService1Server).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_test.TestService1/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestService1Server).SayHello(ctx, req.(*TestRequest1))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto_test.TestService1",
	HandlerType: (*TestService1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _TestService1_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test/service1.proto",
}

func init() { proto.RegisterFile("test/service1.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x49, 0x2d, 0x2e,
	0xd1, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x35, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x02, 0x53, 0xf1, 0x20, 0x29, 0xa5, 0x5c, 0x2e, 0x9e, 0x90, 0xd4, 0xe2, 0x92, 0xa0, 0xd4,
	0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x43, 0x21, 0x11, 0x2e, 0xd6, 0xc2, 0xd2, 0xd4, 0xa2, 0x4a, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48, 0x9e, 0x8b, 0xbb, 0x20, 0x31, 0x3d, 0x35,
	0x3e, 0xaf, 0x34, 0x37, 0x29, 0xb5, 0x48, 0x82, 0x49, 0x81, 0x51, 0x83, 0x35, 0x88, 0x0b, 0x24,
	0xe4, 0x07, 0x16, 0x11, 0x52, 0xe3, 0xe2, 0x2f, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x89, 0x2f, 0x48,
	0x2d, 0x8a, 0x07, 0x49, 0x48, 0x30, 0x83, 0x15, 0xf1, 0x42, 0x84, 0x03, 0x52, 0x8b, 0x02, 0x12,
	0xd3, 0x53, 0x95, 0xf2, 0xb8, 0x78, 0x21, 0xd6, 0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0xd2, 0xda,
	0x3e, 0xa3, 0x40, 0x88, 0xf7, 0x82, 0xa1, 0x01, 0x20, 0xe4, 0xc8, 0xc5, 0x11, 0x9c, 0x58, 0xe9,
	0x91, 0x9a, 0x93, 0x93, 0x2f, 0x24, 0xa1, 0x87, 0x08, 0x07, 0x3d, 0xe4, 0x40, 0x90, 0x92, 0xc4,
	0x94, 0x81, 0xba, 0x57, 0x89, 0x21, 0x89, 0x0d, 0x2c, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x2e, 0x31, 0x3e, 0xb5, 0x5b, 0x01, 0x00, 0x00,
}
