// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.0
// - protoc             v3.14.0
// source: helloworld.proto

package proto

import (
	context "context"
	fmt "fmt"
)

import (
	protocol "dubbo.apache.org/dubbo-go/v3/protocol"
	dubbo3 "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	invocation "dubbo.apache.org/dubbo-go/v3/protocol/invocation"
)

import (
	common "github.com/dubbogo/triple/pkg/common"
	constant "github.com/dubbogo/triple/pkg/common/constant"
	grpc "github.com/dubbogo/triple/pkg/grpc"
	codes "github.com/dubbogo/triple/pkg/grpc/codes"
	status "github.com/dubbogo/triple/pkg/grpc/status"
	triple "github.com/dubbogo/triple/pkg/triple"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHelloStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloStreamClient, error)
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*User, common.ErrorWithAttachment)
	SayHelloWithError(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*User, common.ErrorWithAttachment)
}

type greeterClient struct {
	cc *triple.TripleConn
}

type GreeterClientImpl struct {
	SayHelloStream    func(ctx context.Context) (Greeter_SayHelloStreamClient, error)
	SayHello          func(ctx context.Context, in *HelloRequest) (*User, error)
	SayHelloWithError func(ctx context.Context, in *HelloRequest) (*User, error)
}

func (c *GreeterClientImpl) GetDubboStub(cc *triple.TripleConn) GreeterClient {
	return NewGreeterClient(cc)
}

func NewGreeterClient(cc *triple.TripleConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHelloStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloStreamClient, error) {
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	stream, err := c.cc.NewStream(ctx, "/"+interfaceKey+"/SayHelloStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloStreamClient{stream}
	return x, nil
}

type Greeter_SayHelloStreamClient interface {
	Send(*HelloRequest) error
	Recv() (*User, error)
	grpc.ClientStream
}

type greeterSayHelloStreamClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSayHelloStreamClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*User, common.ErrorWithAttachment) {
	out := new(User)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/SayHello", in, out)
}

func (c *greeterClient) SayHelloWithError(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*User, common.ErrorWithAttachment) {
	out := new(User)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/SayHelloWithError", in, out)
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// Sends a greeting
	SayHelloStream(Greeter_SayHelloStreamServer) error
	SayHello(context.Context, *HelloRequest) (*User, error)
	SayHelloWithError(context.Context, *HelloRequest) (*User, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedGreeterServer) SayHelloStream(Greeter_SayHelloStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloStream not implemented")
}
func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) SayHelloWithError(context.Context, *HelloRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloWithError not implemented")
}
func (s *UnimplementedGreeterServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedGreeterServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SayHelloStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	_, ok := srv.(dubbo3.Dubbo3GrpcService)
	invo := invocation.NewRPCInvocation("SayHelloStream", nil, nil)
	if !ok {
		fmt.Println(invo)
		return nil
	}
	return srv.(GreeterServer).SayHelloStream(&greeterSayHelloStreamServer{stream})
}

type Greeter_SayHelloStreamServer interface {
	Send(*User) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterSayHelloStreamServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloStreamServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSayHelloStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	invo := invocation.NewRPCInvocation("SayHello", args, nil)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloWithError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	invo := invocation.NewRPCInvocation("SayHelloWithError", args, nil)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/SayHelloWithError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHelloWithError(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "SayHelloWithError",
			Handler:    _Greeter_SayHelloWithError_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloStream",
			Handler:       _Greeter_SayHelloStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}

func (s *UnimplementedGreeterServer) XXX_ServiceDesc() *grpc.ServiceDesc {
	return &Greeter_ServiceDesc
}