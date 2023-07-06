package certfetcher

import (
	"context"
	"crypto/tls"
)

// To regenerate the proto headers if the .proto changes, just run go generate
// and this encodes the necessary magic:
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative certfetcher.proto

type fetcher struct{}

func (f *fetcher) Get(ctx context.Context, req *GetCertRequest) (*GetCertReply, error) {
	// We intentionally skip verification because this is meant to return info about the certificate.
	conn, err := tls.Dial("tcp", req.Target, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	var certChain [][]byte
	for _, cert := range conn.ConnectionState().PeerCertificates {
		certChain = append(certChain, cert.Raw)
	}
	return &GetCertReply{
		RawCertificateChain: certChain,
	}, nil

}
