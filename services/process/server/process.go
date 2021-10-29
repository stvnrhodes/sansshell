package server

import (
	"bytes"
	"context"
	"log"
	"os/exec"

	"github.com/Snowflake-Labs/sansshell/services"
	pb "github.com/Snowflake-Labs/sansshell/services/process"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// server is used to implement the gRPC server
type server struct {
}

func (s *server) List(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	log.Printf("Received request for List: %+v", req)

	cmdName := *psBin
	options := psOptions()

	// We gather all the processes up and then filter by pid if needed at the end.
	cmd := exec.CommandContext(ctx, cmdName, options...)
	var stderrBuf, stdoutBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	cmd.Stdin = nil

	if err := cmd.Start(); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := cmd.Wait(); err != nil {
		return nil, status.Errorf(codes.Internal, "command exited with error: %v\n%s", err, stderrBuf.String())
	}

	errBuf := stderrBuf.Bytes()
	if len(errBuf) != 0 {
		return nil, status.Errorf(codes.Internal, "unexpected error output:\n%s", stderrBuf.String())
	}

	entries, err := parser(&stdoutBuf)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected parsing error: %v", err)
	}

	reply := &pb.ListReply{}
	if len(req.Pids) != 0 {
		for _, pid := range req.Pids {
			if _, ok := entries[pid]; !ok {
				return nil, status.Errorf(codes.InvalidArgument, "pid %d does not exist", pid)
			}

			reply.ProcessEntries = append(reply.ProcessEntries, entries[pid])
		}
		return reply, nil
	}

	// If not filtering fill everything in and return. We don't guarentee any ordering.
	for _, e := range entries {
		reply.ProcessEntries = append(reply.ProcessEntries, e)
	}
	return reply, nil
}

func (s *server) GetStacks(ctx context.Context, req *pb.GetStacksRequest) (*pb.GetStacksReply, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *server) GetJavaStacks(ctx context.Context, req *pb.GetJavaStacksRequest) (*pb.GetJavaStacksReply, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *server) GetCore(req *pb.GetCoreRequest, stream pb.Process_GetCoreServer) error {
	return status.Error(codes.Unimplemented, "not implemented")
}

func (s *server) GetJavaHeapDump(req *pb.GetJavaHeapDumpRequest, stream pb.Process_GetJavaHeapDumpServer) error {
	return status.Error(codes.Unimplemented, "not implemented")
}

// Register is called to expose this handler to the gRPC server
func (s *server) Register(gs *grpc.Server) {
	pb.RegisterProcessServer(gs, s)
}

func init() {
	services.RegisterSansShellService(&server{})
}