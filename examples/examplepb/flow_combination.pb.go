// Code generated by protoc-gen-go.
// source: examples/examplepb/flow_combination.proto
// DO NOT EDIT!

package examplepb

import proto "github.com/golang/protobuf/proto"

// discarding unused import google_api1 "google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type EmptyProto struct {
}

func (m *EmptyProto) Reset()         { *m = EmptyProto{} }
func (m *EmptyProto) String() string { return proto.CompactTextString(m) }
func (*EmptyProto) ProtoMessage()    {}

type NonEmptyProto struct {
	A string `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
	B string `protobuf:"bytes,2,opt,name=b" json:"b,omitempty"`
	C string `protobuf:"bytes,3,opt,name=c" json:"c,omitempty"`
}

func (m *NonEmptyProto) Reset()         { *m = NonEmptyProto{} }
func (m *NonEmptyProto) String() string { return proto.CompactTextString(m) }
func (*NonEmptyProto) ProtoMessage()    {}

type UnaryProto struct {
	Str string `protobuf:"bytes,1,opt,name=str" json:"str,omitempty"`
}

func (m *UnaryProto) Reset()         { *m = UnaryProto{} }
func (m *UnaryProto) String() string { return proto.CompactTextString(m) }
func (*UnaryProto) ProtoMessage()    {}

type NestedProto struct {
	A *UnaryProto `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
	B string      `protobuf:"bytes,2,opt,name=b" json:"b,omitempty"`
	C string      `protobuf:"bytes,3,opt,name=c" json:"c,omitempty"`
}

func (m *NestedProto) Reset()         { *m = NestedProto{} }
func (m *NestedProto) String() string { return proto.CompactTextString(m) }
func (*NestedProto) ProtoMessage()    {}

func (m *NestedProto) GetA() *UnaryProto {
	if m != nil {
		return m.A
	}
	return nil
}

type SingleNestedProto struct {
	A *UnaryProto `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
}

func (m *SingleNestedProto) Reset()         { *m = SingleNestedProto{} }
func (m *SingleNestedProto) String() string { return proto.CompactTextString(m) }
func (*SingleNestedProto) ProtoMessage()    {}

func (m *SingleNestedProto) GetA() *UnaryProto {
	if m != nil {
		return m.A
	}
	return nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for FlowCombination service

type FlowCombinationClient interface {
	RpcEmptyRpc(ctx context.Context, in *EmptyProto, opts ...grpc.CallOption) (*EmptyProto, error)
	RpcEmptyStream(ctx context.Context, in *EmptyProto, opts ...grpc.CallOption) (FlowCombination_RpcEmptyStreamClient, error)
	StreamEmptyRpc(ctx context.Context, opts ...grpc.CallOption) (FlowCombination_StreamEmptyRpcClient, error)
	StreamEmptyStream(ctx context.Context, opts ...grpc.CallOption) (FlowCombination_StreamEmptyStreamClient, error)
	RpcBodyRpc(ctx context.Context, in *NonEmptyProto, opts ...grpc.CallOption) (*EmptyProto, error)
	RpcPathSingleNestedRpc(ctx context.Context, in *SingleNestedProto, opts ...grpc.CallOption) (*EmptyProto, error)
	RpcPathNestedRpc(ctx context.Context, in *NestedProto, opts ...grpc.CallOption) (*EmptyProto, error)
	RpcBodyStream(ctx context.Context, in *NonEmptyProto, opts ...grpc.CallOption) (FlowCombination_RpcBodyStreamClient, error)
	RpcPathSingleNestedStream(ctx context.Context, in *SingleNestedProto, opts ...grpc.CallOption) (FlowCombination_RpcPathSingleNestedStreamClient, error)
	RpcPathNestedStream(ctx context.Context, in *NestedProto, opts ...grpc.CallOption) (FlowCombination_RpcPathNestedStreamClient, error)
}

type flowCombinationClient struct {
	cc *grpc.ClientConn
}

func NewFlowCombinationClient(cc *grpc.ClientConn) FlowCombinationClient {
	return &flowCombinationClient{cc}
}

func (c *flowCombinationClient) RpcEmptyRpc(ctx context.Context, in *EmptyProto, opts ...grpc.CallOption) (*EmptyProto, error) {
	out := new(EmptyProto)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcEmptyRpc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationClient) RpcEmptyStream(ctx context.Context, in *EmptyProto, opts ...grpc.CallOption) (FlowCombination_RpcEmptyStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FlowCombination_serviceDesc.Streams[0], c.cc, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcEmptyStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowCombinationRpcEmptyStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlowCombination_RpcEmptyStreamClient interface {
	Recv() (*EmptyProto, error)
	grpc.ClientStream
}

type flowCombinationRpcEmptyStreamClient struct {
	grpc.ClientStream
}

func (x *flowCombinationRpcEmptyStreamClient) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationClient) StreamEmptyRpc(ctx context.Context, opts ...grpc.CallOption) (FlowCombination_StreamEmptyRpcClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FlowCombination_serviceDesc.Streams[1], c.cc, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/StreamEmptyRpc", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowCombinationStreamEmptyRpcClient{stream}
	return x, nil
}

type FlowCombination_StreamEmptyRpcClient interface {
	Send(*EmptyProto) error
	CloseAndRecv() (*EmptyProto, error)
	grpc.ClientStream
}

type flowCombinationStreamEmptyRpcClient struct {
	grpc.ClientStream
}

func (x *flowCombinationStreamEmptyRpcClient) Send(m *EmptyProto) error {
	return x.ClientStream.SendMsg(m)
}

func (x *flowCombinationStreamEmptyRpcClient) CloseAndRecv() (*EmptyProto, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmptyProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationClient) StreamEmptyStream(ctx context.Context, opts ...grpc.CallOption) (FlowCombination_StreamEmptyStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FlowCombination_serviceDesc.Streams[2], c.cc, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/StreamEmptyStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowCombinationStreamEmptyStreamClient{stream}
	return x, nil
}

type FlowCombination_StreamEmptyStreamClient interface {
	Send(*EmptyProto) error
	Recv() (*EmptyProto, error)
	grpc.ClientStream
}

type flowCombinationStreamEmptyStreamClient struct {
	grpc.ClientStream
}

func (x *flowCombinationStreamEmptyStreamClient) Send(m *EmptyProto) error {
	return x.ClientStream.SendMsg(m)
}

func (x *flowCombinationStreamEmptyStreamClient) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationClient) RpcBodyRpc(ctx context.Context, in *NonEmptyProto, opts ...grpc.CallOption) (*EmptyProto, error) {
	out := new(EmptyProto)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcBodyRpc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationClient) RpcPathSingleNestedRpc(ctx context.Context, in *SingleNestedProto, opts ...grpc.CallOption) (*EmptyProto, error) {
	out := new(EmptyProto)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcPathSingleNestedRpc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationClient) RpcPathNestedRpc(ctx context.Context, in *NestedProto, opts ...grpc.CallOption) (*EmptyProto, error) {
	out := new(EmptyProto)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcPathNestedRpc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationClient) RpcBodyStream(ctx context.Context, in *NonEmptyProto, opts ...grpc.CallOption) (FlowCombination_RpcBodyStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FlowCombination_serviceDesc.Streams[3], c.cc, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcBodyStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowCombinationRpcBodyStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlowCombination_RpcBodyStreamClient interface {
	Recv() (*EmptyProto, error)
	grpc.ClientStream
}

type flowCombinationRpcBodyStreamClient struct {
	grpc.ClientStream
}

func (x *flowCombinationRpcBodyStreamClient) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationClient) RpcPathSingleNestedStream(ctx context.Context, in *SingleNestedProto, opts ...grpc.CallOption) (FlowCombination_RpcPathSingleNestedStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FlowCombination_serviceDesc.Streams[4], c.cc, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcPathSingleNestedStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowCombinationRpcPathSingleNestedStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlowCombination_RpcPathSingleNestedStreamClient interface {
	Recv() (*EmptyProto, error)
	grpc.ClientStream
}

type flowCombinationRpcPathSingleNestedStreamClient struct {
	grpc.ClientStream
}

func (x *flowCombinationRpcPathSingleNestedStreamClient) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationClient) RpcPathNestedStream(ctx context.Context, in *NestedProto, opts ...grpc.CallOption) (FlowCombination_RpcPathNestedStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FlowCombination_serviceDesc.Streams[5], c.cc, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcPathNestedStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowCombinationRpcPathNestedStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlowCombination_RpcPathNestedStreamClient interface {
	Recv() (*EmptyProto, error)
	grpc.ClientStream
}

type flowCombinationRpcPathNestedStreamClient struct {
	grpc.ClientStream
}

func (x *flowCombinationRpcPathNestedStreamClient) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for FlowCombination service

type FlowCombinationServer interface {
	RpcEmptyRpc(context.Context, *EmptyProto) (*EmptyProto, error)
	RpcEmptyStream(*EmptyProto, FlowCombination_RpcEmptyStreamServer) error
	StreamEmptyRpc(FlowCombination_StreamEmptyRpcServer) error
	StreamEmptyStream(FlowCombination_StreamEmptyStreamServer) error
	RpcBodyRpc(context.Context, *NonEmptyProto) (*EmptyProto, error)
	RpcPathSingleNestedRpc(context.Context, *SingleNestedProto) (*EmptyProto, error)
	RpcPathNestedRpc(context.Context, *NestedProto) (*EmptyProto, error)
	RpcBodyStream(*NonEmptyProto, FlowCombination_RpcBodyStreamServer) error
	RpcPathSingleNestedStream(*SingleNestedProto, FlowCombination_RpcPathSingleNestedStreamServer) error
	RpcPathNestedStream(*NestedProto, FlowCombination_RpcPathNestedStreamServer) error
}

func RegisterFlowCombinationServer(s *grpc.Server, srv FlowCombinationServer) {
	s.RegisterService(&_FlowCombination_serviceDesc, srv)
}

func _FlowCombination_RpcEmptyRpc_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(EmptyProto)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(FlowCombinationServer).RpcEmptyRpc(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FlowCombination_RpcEmptyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyProto)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowCombinationServer).RpcEmptyStream(m, &flowCombinationRpcEmptyStreamServer{stream})
}

type FlowCombination_RpcEmptyStreamServer interface {
	Send(*EmptyProto) error
	grpc.ServerStream
}

type flowCombinationRpcEmptyStreamServer struct {
	grpc.ServerStream
}

func (x *flowCombinationRpcEmptyStreamServer) Send(m *EmptyProto) error {
	return x.ServerStream.SendMsg(m)
}

func _FlowCombination_StreamEmptyRpc_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlowCombinationServer).StreamEmptyRpc(&flowCombinationStreamEmptyRpcServer{stream})
}

type FlowCombination_StreamEmptyRpcServer interface {
	SendAndClose(*EmptyProto) error
	Recv() (*EmptyProto, error)
	grpc.ServerStream
}

type flowCombinationStreamEmptyRpcServer struct {
	grpc.ServerStream
}

func (x *flowCombinationStreamEmptyRpcServer) SendAndClose(m *EmptyProto) error {
	return x.ServerStream.SendMsg(m)
}

func (x *flowCombinationStreamEmptyRpcServer) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FlowCombination_StreamEmptyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlowCombinationServer).StreamEmptyStream(&flowCombinationStreamEmptyStreamServer{stream})
}

type FlowCombination_StreamEmptyStreamServer interface {
	Send(*EmptyProto) error
	Recv() (*EmptyProto, error)
	grpc.ServerStream
}

type flowCombinationStreamEmptyStreamServer struct {
	grpc.ServerStream
}

func (x *flowCombinationStreamEmptyStreamServer) Send(m *EmptyProto) error {
	return x.ServerStream.SendMsg(m)
}

func (x *flowCombinationStreamEmptyStreamServer) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FlowCombination_RpcBodyRpc_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(NonEmptyProto)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(FlowCombinationServer).RpcBodyRpc(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FlowCombination_RpcPathSingleNestedRpc_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(SingleNestedProto)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(FlowCombinationServer).RpcPathSingleNestedRpc(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FlowCombination_RpcPathNestedRpc_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(NestedProto)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(FlowCombinationServer).RpcPathNestedRpc(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FlowCombination_RpcBodyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NonEmptyProto)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowCombinationServer).RpcBodyStream(m, &flowCombinationRpcBodyStreamServer{stream})
}

type FlowCombination_RpcBodyStreamServer interface {
	Send(*EmptyProto) error
	grpc.ServerStream
}

type flowCombinationRpcBodyStreamServer struct {
	grpc.ServerStream
}

func (x *flowCombinationRpcBodyStreamServer) Send(m *EmptyProto) error {
	return x.ServerStream.SendMsg(m)
}

func _FlowCombination_RpcPathSingleNestedStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SingleNestedProto)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowCombinationServer).RpcPathSingleNestedStream(m, &flowCombinationRpcPathSingleNestedStreamServer{stream})
}

type FlowCombination_RpcPathSingleNestedStreamServer interface {
	Send(*EmptyProto) error
	grpc.ServerStream
}

type flowCombinationRpcPathSingleNestedStreamServer struct {
	grpc.ServerStream
}

func (x *flowCombinationRpcPathSingleNestedStreamServer) Send(m *EmptyProto) error {
	return x.ServerStream.SendMsg(m)
}

func _FlowCombination_RpcPathNestedStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NestedProto)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowCombinationServer).RpcPathNestedStream(m, &flowCombinationRpcPathNestedStreamServer{stream})
}

type FlowCombination_RpcPathNestedStreamServer interface {
	Send(*EmptyProto) error
	grpc.ServerStream
}

type flowCombinationRpcPathNestedStreamServer struct {
	grpc.ServerStream
}

func (x *flowCombinationRpcPathNestedStreamServer) Send(m *EmptyProto) error {
	return x.ServerStream.SendMsg(m)
}

var _FlowCombination_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gengo.grpc.gateway.examples.examplepb.FlowCombination",
	HandlerType: (*FlowCombinationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RpcEmptyRpc",
			Handler:    _FlowCombination_RpcEmptyRpc_Handler,
		},
		{
			MethodName: "RpcBodyRpc",
			Handler:    _FlowCombination_RpcBodyRpc_Handler,
		},
		{
			MethodName: "RpcPathSingleNestedRpc",
			Handler:    _FlowCombination_RpcPathSingleNestedRpc_Handler,
		},
		{
			MethodName: "RpcPathNestedRpc",
			Handler:    _FlowCombination_RpcPathNestedRpc_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RpcEmptyStream",
			Handler:       _FlowCombination_RpcEmptyStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamEmptyRpc",
			Handler:       _FlowCombination_StreamEmptyRpc_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamEmptyStream",
			Handler:       _FlowCombination_StreamEmptyStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RpcBodyStream",
			Handler:       _FlowCombination_RpcBodyStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RpcPathSingleNestedStream",
			Handler:       _FlowCombination_RpcPathSingleNestedStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RpcPathNestedStream",
			Handler:       _FlowCombination_RpcPathNestedStream_Handler,
			ServerStreams: true,
		},
	},
}