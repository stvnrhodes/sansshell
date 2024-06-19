// Copyright (c) 2024 Snowflake Inc. All rights reserved.
//
//Licensed under the Apache License, Version 2.0 (the
//"License"); you may not use this file except in compliance
//with the License.  You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing,
//software distributed under the License is distributed on an
//"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
//KIND, either express or implied.  See the License for the
//specific language governing permissions and limitations
//under the License.

// P

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.6
// source: psql.proto

package psql

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Psql_Exec_FullMethodName  = "/Psql.Psql/Exec"
	Psql_Query_FullMethodName = "/Psql.Psql/Query"
)

// PsqlClient is the client API for Psql service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The Psql service definition.
type PsqlClient interface {
	// Exec runs a SQL statement against a postgres database.
	Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error)
	// Query runs a SQL query against a postgres database and returns results.
	// This returns all results in a single response, so it's unsuitable for
	// results larger than a few megabytes.
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type psqlClient struct {
	cc grpc.ClientConnInterface
}

func NewPsqlClient(cc grpc.ClientConnInterface) PsqlClient {
	return &psqlClient{cc}
}

func (c *psqlClient) Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExecResponse)
	err := c.cc.Invoke(ctx, Psql_Exec_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *psqlClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, Psql_Query_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PsqlServer is the server API for Psql service.
// All implementations should embed UnimplementedPsqlServer
// for forward compatibility
//
// The Psql service definition.
type PsqlServer interface {
	// Exec runs a SQL statement against a postgres database.
	Exec(context.Context, *ExecRequest) (*ExecResponse, error)
	// Query runs a SQL query against a postgres database and returns results.
	// This returns all results in a single response, so it's unsuitable for
	// results larger than a few megabytes.
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
}

// UnimplementedPsqlServer should be embedded to have forward compatible implementations.
type UnimplementedPsqlServer struct {
}

func (UnimplementedPsqlServer) Exec(context.Context, *ExecRequest) (*ExecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}
func (UnimplementedPsqlServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}

// UnsafePsqlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PsqlServer will
// result in compilation errors.
type UnsafePsqlServer interface {
	mustEmbedUnimplementedPsqlServer()
}

func RegisterPsqlServer(s grpc.ServiceRegistrar, srv PsqlServer) {
	s.RegisterService(&Psql_ServiceDesc, srv)
}

func _Psql_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PsqlServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Psql_Exec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PsqlServer).Exec(ctx, req.(*ExecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Psql_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PsqlServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Psql_Query_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PsqlServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Psql_ServiceDesc is the grpc.ServiceDesc for Psql service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Psql_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Psql.Psql",
	HandlerType: (*PsqlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exec",
			Handler:    _Psql_Exec_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Psql_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "psql.proto",
}
