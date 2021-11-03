// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proxy

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProxyClient is the client API for Proxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProxyClient interface {
	// Proxy represents a bidirectional stream of requests from a sanshell
	// client to one or more target sanshell instances reachable from the
	// proxy server.
	Proxy(ctx context.Context, opts ...grpc.CallOption) (Proxy_ProxyClient, error)
}

type proxyClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyClient(cc grpc.ClientConnInterface) ProxyClient {
	return &proxyClient{cc}
}

func (c *proxyClient) Proxy(ctx context.Context, opts ...grpc.CallOption) (Proxy_ProxyClient, error) {
	stream, err := c.cc.NewStream(ctx, &Proxy_ServiceDesc.Streams[0], "/Proxy.Proxy/Proxy", opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyProxyClient{stream}
	return x, nil
}

type Proxy_ProxyClient interface {
	Send(*ProxyRequest) error
	Recv() (*ProxyReply, error)
	grpc.ClientStream
}

type proxyProxyClient struct {
	grpc.ClientStream
}

func (x *proxyProxyClient) Send(m *ProxyRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *proxyProxyClient) Recv() (*ProxyReply, error) {
	m := new(ProxyReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProxyServer is the server API for Proxy service.
// All implementations should embed UnimplementedProxyServer
// for forward compatibility
type ProxyServer interface {
	// Proxy represents a bidirectional stream of requests from a sanshell
	// client to one or more target sanshell instances reachable from the
	// proxy server.
	Proxy(Proxy_ProxyServer) error
}

// UnimplementedProxyServer should be embedded to have forward compatible implementations.
type UnimplementedProxyServer struct {
}

func (UnimplementedProxyServer) Proxy(Proxy_ProxyServer) error {
	return status.Errorf(codes.Unimplemented, "method Proxy not implemented")
}

// UnsafeProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProxyServer will
// result in compilation errors.
type UnsafeProxyServer interface {
	mustEmbedUnimplementedProxyServer()
}

func RegisterProxyServer(s grpc.ServiceRegistrar, srv ProxyServer) {
	s.RegisterService(&Proxy_ServiceDesc, srv)
}

func _Proxy_Proxy_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProxyServer).Proxy(&proxyProxyServer{stream})
}

type Proxy_ProxyServer interface {
	Send(*ProxyReply) error
	Recv() (*ProxyRequest, error)
	grpc.ServerStream
}

type proxyProxyServer struct {
	grpc.ServerStream
}

func (x *proxyProxyServer) Send(m *ProxyReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *proxyProxyServer) Recv() (*ProxyRequest, error) {
	m := new(ProxyRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Proxy_ServiceDesc is the grpc.ServiceDesc for Proxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Proxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Proxy.Proxy",
	HandlerType: (*ProxyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Proxy",
			Handler:       _Proxy_Proxy_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proxy.proto",
}