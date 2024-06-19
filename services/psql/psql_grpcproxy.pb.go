// Auto generated code by protoc-gen-go-grpcproxy
// DO NOT EDIT

// Adds OneMany versions of RPC methods for use by proxy clients

package psql

import (
	context "context"
	proxy "github.com/Snowflake-Labs/sansshell/proxy/proxy"
	grpc "google.golang.org/grpc"
)

import (
	"fmt"
)

// PsqlClientProxy is the superset of PsqlClient which additionally includes the OneMany proxy methods
type PsqlClientProxy interface {
	PsqlClient
	ExecOneMany(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (<-chan *ExecManyResponse, error)
	QueryOneMany(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (<-chan *QueryManyResponse, error)
}

// Embed the original client inside of this so we get the other generated methods automatically.
type psqlClientProxy struct {
	*psqlClient
}

// NewPsqlClientProxy creates a PsqlClientProxy for use in proxied connections.
// NOTE: This takes a proxy.Conn instead of a generic ClientConnInterface as the methods here are only valid in proxy.Conn contexts.
func NewPsqlClientProxy(cc *proxy.Conn) PsqlClientProxy {
	return &psqlClientProxy{NewPsqlClient(cc).(*psqlClient)}
}

// ExecManyResponse encapsulates a proxy data packet.
// It includes the target, index, response and possible error returned.
type ExecManyResponse struct {
	Target string
	// As targets can be duplicated this is the index into the slice passed to proxy.Conn.
	Index int
	Resp  *ExecResponse
	Error error
}

// ExecOneMany provides the same API as Exec but sends the same request to N destinations at once.
// N can be a single destination.
//
// NOTE: The returned channel must be read until it closes in order to avoid leaking goroutines.
func (c *psqlClientProxy) ExecOneMany(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (<-chan *ExecManyResponse, error) {
	conn := c.cc.(*proxy.Conn)
	ret := make(chan *ExecManyResponse)
	// If this is a single case we can just use Invoke and marshal it onto the channel once and be done.
	if len(conn.Targets) == 1 {
		go func() {
			out := &ExecManyResponse{
				Target: conn.Targets[0],
				Index:  0,
				Resp:   &ExecResponse{},
			}
			err := conn.Invoke(ctx, "/Psql.Psql/Exec", in, out.Resp, opts...)
			if err != nil {
				out.Error = err
			}
			// Send and close.
			ret <- out
			close(ret)
		}()
		return ret, nil
	}
	manyRet, err := conn.InvokeOneMany(ctx, "/Psql.Psql/Exec", in, opts...)
	if err != nil {
		return nil, err
	}
	// A goroutine to retrive untyped responses and convert them to typed ones.
	go func() {
		for {
			typedResp := &ExecManyResponse{
				Resp: &ExecResponse{},
			}

			resp, ok := <-manyRet
			if !ok {
				// All done so we can shut down.
				close(ret)
				return
			}
			typedResp.Target = resp.Target
			typedResp.Index = resp.Index
			typedResp.Error = resp.Error
			if resp.Error == nil {
				if err := resp.Resp.UnmarshalTo(typedResp.Resp); err != nil {
					typedResp.Error = fmt.Errorf("can't decode any response - %v. Original Error - %v", err, resp.Error)
				}
			}
			ret <- typedResp
		}
	}()

	return ret, nil
}

// QueryManyResponse encapsulates a proxy data packet.
// It includes the target, index, response and possible error returned.
type QueryManyResponse struct {
	Target string
	// As targets can be duplicated this is the index into the slice passed to proxy.Conn.
	Index int
	Resp  *QueryResponse
	Error error
}

// QueryOneMany provides the same API as Query but sends the same request to N destinations at once.
// N can be a single destination.
//
// NOTE: The returned channel must be read until it closes in order to avoid leaking goroutines.
func (c *psqlClientProxy) QueryOneMany(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (<-chan *QueryManyResponse, error) {
	conn := c.cc.(*proxy.Conn)
	ret := make(chan *QueryManyResponse)
	// If this is a single case we can just use Invoke and marshal it onto the channel once and be done.
	if len(conn.Targets) == 1 {
		go func() {
			out := &QueryManyResponse{
				Target: conn.Targets[0],
				Index:  0,
				Resp:   &QueryResponse{},
			}
			err := conn.Invoke(ctx, "/Psql.Psql/Query", in, out.Resp, opts...)
			if err != nil {
				out.Error = err
			}
			// Send and close.
			ret <- out
			close(ret)
		}()
		return ret, nil
	}
	manyRet, err := conn.InvokeOneMany(ctx, "/Psql.Psql/Query", in, opts...)
	if err != nil {
		return nil, err
	}
	// A goroutine to retrive untyped responses and convert them to typed ones.
	go func() {
		for {
			typedResp := &QueryManyResponse{
				Resp: &QueryResponse{},
			}

			resp, ok := <-manyRet
			if !ok {
				// All done so we can shut down.
				close(ret)
				return
			}
			typedResp.Target = resp.Target
			typedResp.Index = resp.Index
			typedResp.Error = resp.Error
			if resp.Error == nil {
				if err := resp.Resp.UnmarshalTo(typedResp.Resp); err != nil {
					typedResp.Error = fmt.Errorf("can't decode any response - %v. Original Error - %v", err, resp.Error)
				}
			}
			ret <- typedResp
		}
	}()

	return ret, nil
}
