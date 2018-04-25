// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/calculator/calculator.proto

/*
Package go_micro_calculator is a generated protocol buffer package.

It is generated from these files:
	proto/calculator/calculator.proto

It has these top-level messages:
	CalculatorResponse
	OperationRequest
	GetAverageRequest
*/
package go_micro_calculator

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

type CalculatorResponse struct {
	CalculatedValue float32 `protobuf:"fixed32,1,opt,name=calculatedValue" json:"calculatedValue,omitempty"`
}

func (m *CalculatorResponse) Reset()                    { *m = CalculatorResponse{} }
func (m *CalculatorResponse) String() string            { return proto.CompactTextString(m) }
func (*CalculatorResponse) ProtoMessage()               {}
func (*CalculatorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CalculatorResponse) GetCalculatedValue() float32 {
	if m != nil {
		return m.CalculatedValue
	}
	return 0
}

type OperationRequest struct {
	Num1 float32 `protobuf:"fixed32,1,opt,name=num1" json:"num1,omitempty"`
	Num2 float32 `protobuf:"fixed32,2,opt,name=num2" json:"num2,omitempty"`
}

func (m *OperationRequest) Reset()                    { *m = OperationRequest{} }
func (m *OperationRequest) String() string            { return proto.CompactTextString(m) }
func (*OperationRequest) ProtoMessage()               {}
func (*OperationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OperationRequest) GetNum1() float32 {
	if m != nil {
		return m.Num1
	}
	return 0
}

func (m *OperationRequest) GetNum2() float32 {
	if m != nil {
		return m.Num2
	}
	return 0
}

type GetAverageRequest struct {
	Numbers []float32 `protobuf:"fixed32,1,rep,packed,name=numbers" json:"numbers,omitempty"`
}

func (m *GetAverageRequest) Reset()                    { *m = GetAverageRequest{} }
func (m *GetAverageRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAverageRequest) ProtoMessage()               {}
func (*GetAverageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetAverageRequest) GetNumbers() []float32 {
	if m != nil {
		return m.Numbers
	}
	return nil
}

func init() {
	proto.RegisterType((*CalculatorResponse)(nil), "go.micro.calculator.CalculatorResponse")
	proto.RegisterType((*OperationRequest)(nil), "go.micro.calculator.OperationRequest")
	proto.RegisterType((*GetAverageRequest)(nil), "go.micro.calculator.GetAverageRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CalculatorService service

type CalculatorServiceClient interface {
	AddNumbers(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
	SubtractNumbers(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
	GetAverage(ctx context.Context, in *GetAverageRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) AddNumbers(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := grpc.Invoke(ctx, "/go.micro.calculator.CalculatorService/AddNumbers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) SubtractNumbers(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := grpc.Invoke(ctx, "/go.micro.calculator.CalculatorService/SubtractNumbers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) GetAverage(ctx context.Context, in *GetAverageRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := grpc.Invoke(ctx, "/go.micro.calculator.CalculatorService/GetAverage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CalculatorService service

type CalculatorServiceServer interface {
	AddNumbers(context.Context, *OperationRequest) (*CalculatorResponse, error)
	SubtractNumbers(context.Context, *OperationRequest) (*CalculatorResponse, error)
	GetAverage(context.Context, *GetAverageRequest) (*CalculatorResponse, error)
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_AddNumbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).AddNumbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.calculator.CalculatorService/AddNumbers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).AddNumbers(ctx, req.(*OperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_SubtractNumbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).SubtractNumbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.calculator.CalculatorService/SubtractNumbers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).SubtractNumbers(ctx, req.(*OperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_GetAverage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAverageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).GetAverage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.calculator.CalculatorService/GetAverage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).GetAverage(ctx, req.(*GetAverageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNumbers",
			Handler:    _CalculatorService_AddNumbers_Handler,
		},
		{
			MethodName: "SubtractNumbers",
			Handler:    _CalculatorService_SubtractNumbers_Handler,
		},
		{
			MethodName: "GetAverage",
			Handler:    _CalculatorService_GetAverage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/calculator/calculator.proto",
}

func init() { proto.RegisterFile("proto/calculator/calculator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4e, 0xcc, 0x49, 0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0x42, 0x62, 0xea, 0x81,
	0xe5, 0x84, 0x84, 0xd3, 0xf3, 0xf5, 0x72, 0x33, 0x93, 0x8b, 0xf2, 0xf5, 0x10, 0x52, 0x4a, 0x76,
	0x5c, 0x42, 0xce, 0x70, 0x5e, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x06, 0x17,
	0x3f, 0x4c, 0x4d, 0x6a, 0x4a, 0x58, 0x62, 0x4e, 0x69, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x53,
	0x10, 0xba, 0xb0, 0x92, 0x15, 0x97, 0x80, 0x7f, 0x41, 0x6a, 0x51, 0x62, 0x49, 0x66, 0x7e, 0x5e,
	0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x69, 0xae, 0x21, 0x54,
	0x0b, 0x98, 0x0d, 0x15, 0x33, 0x92, 0x60, 0x82, 0x8b, 0x19, 0x29, 0xe9, 0x72, 0x09, 0xba, 0xa7,
	0x96, 0x38, 0x96, 0xa5, 0x16, 0x25, 0xa6, 0xa7, 0xc2, 0x34, 0x4b, 0x70, 0xb1, 0xe7, 0x95, 0xe6,
	0x26, 0xa5, 0x16, 0x15, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x30, 0x05, 0xc1, 0xb8, 0x46, 0x9b, 0x99,
	0xb8, 0x04, 0x11, 0x6e, 0x0d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x8a, 0xe1, 0xe2, 0x72,
	0x4c, 0x49, 0xf1, 0x83, 0xa8, 0x11, 0x52, 0xd5, 0xc3, 0xe2, 0x49, 0x3d, 0x74, 0x17, 0x4a, 0xa9,
	0x63, 0x55, 0x86, 0x25, 0x20, 0x12, 0xb9, 0xf8, 0x83, 0x4b, 0x93, 0x4a, 0x8a, 0x12, 0x93, 0x4b,
	0x68, 0x65, 0x45, 0x2c, 0x17, 0x17, 0x22, 0x14, 0x84, 0xd4, 0xb0, 0x6a, 0xc3, 0x08, 0x26, 0xa2,
	0x8d, 0x4f, 0x62, 0x03, 0x47, 0xbe, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xef, 0x2b, 0xcb, 0x98,
	0x21, 0x02, 0x00, 0x00,
}