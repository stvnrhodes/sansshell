/* Copyright (c) 2019 Snowflake Inc. All rights reserved.

   Licensed under the Apache License, Version 2.0 (the
   "License"); you may not use this file except in compliance
   with the License.  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing,
   software distributed under the License is distributed on an
   "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
   KIND, either express or implied.  See the License for the
   specific language governing permissions and limitations
   under the License.
*/

// Package telemetry contains code for emitting telemetry
// from Sansshell processes.
package telemetry

import (
	"context"
	"io"
	"log/slog"
	"strings"

	"github.com/Snowflake-Labs/sansshell/auth/opa/rpcauth"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	sansshellMetadata   = "sansshell-"
	sansshellTraceIDKey = sansshellMetadata + "trace-id"
)

// UnaryClientLogInterceptor returns a new grpc.UnaryClientInterceptor that logs
// outgoing requests using the supplied logger, as well as injecting it into the
// context of the invoker.
func UnaryClientLogInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx = passAlongMetadata(ctx)
		slog.InfoContext(ctx, "new client request", "method", method, "target", cc.Target())
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			slog.ErrorContext(ctx, "client request", "err", err)
		}
		return err
	}
}

// StreamClientLogInterceptor returns a new grpc.StreamClientInterceptor that logs
// client requests using the supplied logger, as well as injecting it into the context
// of the created stream.
func StreamClientLogInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		ctx = passAlongMetadata(ctx)
		slog.InfoContext(ctx, "new client stream", "method", method, "target", cc.Target())
		stream, err := streamer(ctx, desc, cc, method, opts...)
		if err != nil {
			slog.ErrorContext(ctx, "create stream", "err", err)
			return nil, err
		}
		return &loggedClientStream{
			ClientStream: stream,
		}, nil
	}
}

func hasSpan(ctx context.Context) bool {
	return trace.SpanContextFromContext(ctx).IsValid()
}

// // Add trace ID to logger if there's an active span
// func logOtelTraceID(ctx context.Context, l logr.Logger) logr.Logger {
// 	if hasSpan(ctx) {
// 		spanCtx := trace.SpanContextFromContext(ctx)
// 		l = l.WithValues(sansshellTraceIDKey, spanCtx.TraceID().String())
// 	}

// 	return l
// }

// func logMetadata(ctx context.Context, l logr.Logger) logr.Logger {
// 	// Add any sansshell specific metadata from incoming context to the logging we do.
// 	md, ok := metadata.FromIncomingContext(ctx)
// 	if ok {
// 		for k, v := range md {
// 			if strings.HasPrefix(k, sansshellMetadata) {
// 				for _, val := range v {
// 					l = l.WithValues(k, val)
// 				}
// 			}
// 		}
// 	}
// 	l = logOtelTraceID(ctx, l)
// 	return l
// }

func passAlongMetadata(ctx context.Context) context.Context {
	// See if we got any metadata that has our prefix and pass it along
	// downstream (i.e. proxy case).
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for k, v := range md {
			if strings.HasPrefix(k, sansshellMetadata) {
				for _, val := range v {
					ctx = metadata.AppendToOutgoingContext(ctx, k, val)
				}
			}
		}
	}
	return ctx
}

type loggedClientStream struct {
	grpc.ClientStream
}

// See: grpc.ClientStream.SendMsg()
func (l *loggedClientStream) SendMsg(m interface{}) error {
	slog.DebugContext(l.Context(), "SendMsg", "msg", m)
	err := l.ClientStream.SendMsg(m)
	if err != nil {
		slog.ErrorContext(l.Context(), "SendMsg", "err", err)
	}
	return err
}

// See: grpc.ClientStream.RecvMsg()
func (l *loggedClientStream) RecvMsg(m interface{}) error {
	slog.DebugContext(l.Context(), "RecvMsg", "msg", m)
	err := l.ClientStream.RecvMsg(m)
	if err != nil && err != io.EOF {
		slog.ErrorContext(l.Context(), "RecvMsg", "err", err)
	}
	return err
}

// See: grpc.ClientStream.CloseSend()
func (l *loggedClientStream) CloseSend() error {
	slog.Info("CloseSend")
	err := l.ClientStream.CloseSend()
	if err != nil {
		slog.Error("CloseSend", "err", err)
	}
	return err
}

// UnaryServerLogInterceptor returns a new gprc.UnaryServerInterceptor that logs
// incoming requests using the supplied logger, as well as injecting it into the
// context of downstream handlers. If incoming calls require client side provided justification
// (which is logged) then the justification parameter should be true and a required
// key of ReqJustKey must be in the context when the interceptor runs.
func UnaryServerLogInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		slog.InfoContext(ctx, "new request", "method", info.FullMethod, "peer", rpcauth.PeerInputFromContext(ctx))
		resp, err := handler(ctx, req)
		if err != nil {
			slog.ErrorContext(ctx, "handler", "err", err)
		}
		return resp, err
	}
}

// StreamServerLogInterceptor returns a new grpc.StreamServerInterceptor that logs
// incoming streams using the supplied logger, and makes it available via the stream
// context to stream handlers. If incoming calls require client side provided justification
// (which is logged) then the justification parameter should be true and a required
// key of ReqJustKey must be in the context when the interceptor runs.
func StreamServerLogInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		slog.Info("new stream", "method", info.FullMethod, "peer", rpcauth.PeerInputFromContext(ss.Context()))
		stream := &loggedStream{
			ServerStream: ss,
		}
		err := handler(srv, stream)
		if err != nil {
			slog.Error("handler", "err", err)
		}
		return err
	}
}

// loggedStream wraps a grpc.ServerStream with additional logging.
type loggedStream struct {
	grpc.ServerStream
}

func (l *loggedStream) SendMsg(m interface{}) error {
	slog.DebugContext(l.Context(), "RecvMsg", "msg", m)
	err := l.ServerStream.SendMsg(m)
	if err != nil {
		slog.ErrorContext(l.Context(), "SendMsg", "err", err)
	}
	return err
}

func (l *loggedStream) RecvMsg(m interface{}) error {
	slog.DebugContext(l.Context(), "RecvMsg", "msg", m)
	err := l.ServerStream.RecvMsg(m)
	if err != nil && err != io.EOF {
		slog.ErrorContext(l.Context(), "RecvMsg", "err", err)
	}
	return err
}
