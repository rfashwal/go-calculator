// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/event/event.proto

/*
Package go_micro_calculator is a generated protocol buffer package.

It is generated from these files:
	proto/event/event.proto

It has these top-level messages:
	Event
	Response
	EventFilter
	EventResponse
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

type Event struct {
	EventId       string `protobuf:"bytes,1,opt,name=eventId" json:"eventId,omitempty"`
	EventType     string `protobuf:"bytes,2,opt,name=eventType" json:"eventType,omitempty"`
	AggregateId   string `protobuf:"bytes,3,opt,name=aggregateId" json:"aggregateId,omitempty"`
	AggregateType string `protobuf:"bytes,4,opt,name=aggregateType" json:"aggregateType,omitempty"`
	EventData     string `protobuf:"bytes,5,opt,name=eventData" json:"eventData,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Event) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *Event) GetAggregateId() string {
	if m != nil {
		return m.AggregateId
	}
	return ""
}

func (m *Event) GetAggregateType() string {
	if m != nil {
		return m.AggregateType
	}
	return ""
}

func (m *Event) GetEventData() string {
	if m != nil {
		return m.EventData
	}
	return ""
}

type Response struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type EventFilter struct {
	EventId     string `protobuf:"bytes,1,opt,name=eventId" json:"eventId,omitempty"`
	AggregateId string `protobuf:"bytes,2,opt,name=aggregateId" json:"aggregateId,omitempty"`
}

func (m *EventFilter) Reset()                    { *m = EventFilter{} }
func (m *EventFilter) String() string            { return proto.CompactTextString(m) }
func (*EventFilter) ProtoMessage()               {}
func (*EventFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EventFilter) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *EventFilter) GetAggregateId() string {
	if m != nil {
		return m.AggregateId
	}
	return ""
}

type EventResponse struct {
	Events []*Event `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
}

func (m *EventResponse) Reset()                    { *m = EventResponse{} }
func (m *EventResponse) String() string            { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()               {}
func (*EventResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EventResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "go.micro.calculator.Event")
	proto.RegisterType((*Response)(nil), "go.micro.calculator.Response")
	proto.RegisterType((*EventFilter)(nil), "go.micro.calculator.EventFilter")
	proto.RegisterType((*EventResponse)(nil), "go.micro.calculator.EventResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EventStore service

type EventStoreClient interface {
	GetEvents(ctx context.Context, in *EventFilter, opts ...grpc.CallOption) (*EventResponse, error)
	CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error)
}

type eventStoreClient struct {
	cc *grpc.ClientConn
}

func NewEventStoreClient(cc *grpc.ClientConn) EventStoreClient {
	return &eventStoreClient{cc}
}

func (c *eventStoreClient) GetEvents(ctx context.Context, in *EventFilter, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := grpc.Invoke(ctx, "/go.micro.calculator.EventStore/GetEvents", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventStoreClient) CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/go.micro.calculator.EventStore/CreateEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EventStore service

type EventStoreServer interface {
	GetEvents(context.Context, *EventFilter) (*EventResponse, error)
	CreateEvent(context.Context, *Event) (*Response, error)
}

func RegisterEventStoreServer(s *grpc.Server, srv EventStoreServer) {
	s.RegisterService(&_EventStore_serviceDesc, srv)
}

func _EventStore_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStoreServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.calculator.EventStore/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStoreServer).GetEvents(ctx, req.(*EventFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventStore_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStoreServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.calculator.EventStore/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStoreServer).CreateEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.calculator.EventStore",
	HandlerType: (*EventStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvents",
			Handler:    _EventStore_GetEvents_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _EventStore_CreateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/event/event.proto",
}

func init() { proto.RegisterFile("proto/event/event.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4d, 0x4f, 0xb3, 0x40,
	0x10, 0xc7, 0x1f, 0xda, 0x87, 0x5a, 0x86, 0xf4, 0xb2, 0x9a, 0x48, 0x88, 0x26, 0x64, 0xe3, 0xa1,
	0x27, 0x4c, 0xf0, 0xe6, 0xb5, 0xbe, 0x04, 0x8f, 0xd4, 0x2f, 0xb0, 0xd2, 0x09, 0x69, 0x82, 0x5d,
	0x32, 0xbb, 0x35, 0xf1, 0x03, 0x79, 0xf4, 0x3b, 0x1a, 0x66, 0x29, 0xad, 0x5a, 0xb9, 0x10, 0xe6,
	0x3f, 0x6f, 0xbf, 0xff, 0xee, 0xc2, 0x79, 0x43, 0xda, 0xea, 0x6b, 0x7c, 0xc3, 0x8d, 0x75, 0xdf,
	0x94, 0x15, 0x71, 0x5a, 0xe9, 0xf4, 0x75, 0x5d, 0x92, 0x4e, 0x4b, 0x55, 0x97, 0xdb, 0x5a, 0x59,
	0x4d, 0xf2, 0xc3, 0x03, 0xff, 0xbe, 0x2d, 0x12, 0x11, 0x9c, 0x70, 0x75, 0xbe, 0x8a, 0xbc, 0xc4,
	0x9b, 0x07, 0xc5, 0x2e, 0x14, 0x17, 0x10, 0xf0, 0xef, 0xf3, 0x7b, 0x83, 0xd1, 0x88, 0x73, 0x7b,
	0x41, 0x24, 0x10, 0xaa, 0xaa, 0x22, 0xac, 0x94, 0xc5, 0x7c, 0x15, 0x8d, 0x39, 0x7f, 0x28, 0x89,
	0x2b, 0x98, 0xf5, 0x21, 0xcf, 0xf8, 0xcf, 0x35, 0xdf, 0xc5, 0x7e, 0xcb, 0x9d, 0xb2, 0x2a, 0xf2,
	0x0f, 0xb6, 0xb4, 0x82, 0xbc, 0x85, 0x69, 0x81, 0xa6, 0xd1, 0x1b, 0x83, 0x2d, 0xa9, 0xd9, 0x96,
	0x25, 0x1a, 0xc3, 0xa4, 0xd3, 0x62, 0x17, 0x8a, 0x33, 0xf0, 0x91, 0x48, 0x53, 0x47, 0xe9, 0x02,
	0x99, 0x43, 0xc8, 0x16, 0x1f, 0xd6, 0xb5, 0x45, 0x1a, 0x30, 0xfa, 0xc3, 0xca, 0xe8, 0x97, 0x15,
	0xb9, 0x80, 0x19, 0x8f, 0xea, 0x59, 0x32, 0x98, 0x70, 0x77, 0x8b, 0x32, 0x9e, 0x87, 0x59, 0x9c,
	0x1e, 0x39, 0xe5, 0xd4, 0xf5, 0x74, 0x95, 0xd9, 0xa7, 0x07, 0xc0, 0xca, 0xd2, 0x6a, 0x42, 0xb1,
	0x84, 0xe0, 0x11, 0x2d, 0x0b, 0x46, 0x24, 0x7f, 0xf7, 0x3b, 0xfc, 0x58, 0x0e, 0x6c, 0xe8, 0xa8,
	0xe4, 0x3f, 0xf1, 0x04, 0xe1, 0x82, 0x50, 0x59, 0x74, 0x97, 0x3b, 0x80, 0x15, 0x5f, 0x1e, 0xcd,
	0xed, 0x67, 0xbd, 0x4c, 0xf8, 0xfd, 0xdc, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0xef, 0x21, 0xde,
	0xfb, 0x5a, 0x02, 0x00, 0x00,
}
