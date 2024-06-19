package server

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/Snowflake-Labs/sansshell/services"
	pb "github.com/Snowflake-Labs/sansshell/services/psql"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type scannableItem struct {
	*pb.Row_Item
}

func (i scannableItem) Scan(src any) error {
	if src == nil {
		i.Value = nil
	}
	switch val := src.(type) {
	case string:
		i.Value = &pb.Row_Item_StringValue{StringValue: val}
	case int64:
		i.Value = &pb.Row_Item_IntValue{IntValue: val}
	case float64:
		i.Value = &pb.Row_Item_NumberValue{NumberValue: val}
	case []byte:
		i.Value = &pb.Row_Item_ByteValue{ByteValue: val}
	case bool:
		i.Value = &pb.Row_Item_BoolValue{BoolValue: val}
	case time.Time:
		i.Value = &pb.Row_Item_TimeValue{TimeValue: timestamppb.New(val)}
	case nil:
		i.Value = nil
	}
	return nil
}

type server struct {
	mu sync.Mutex
	db *sql.DB
}

// getDB returns a connection to a SQL database, creating it if needed
func (s *server) getDB(ctx context.Context) (*sql.DB, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		db, err := connect(ctx)
		if err != nil {
			return nil, err
		}
		s.db = db
	}
	return s.db, nil
}

// Exec runs a SQL Exec against a postgres database
func (s *server) Exec(ctx context.Context, req *pb.ExecRequest) (*pb.ExecResponse, error) {
	db, err := s.getDB(ctx)
	if err != nil {
		return nil, err
	}
	var args []any
	for _, p := range req.Parameters {
		args = append(args, p)
	}
	res, err := db.ExecContext(ctx, req.Query, args...)
	if err != nil {
		return nil, err
	}
	resp := &pb.ExecResponse{}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		// TODO log
	} else {
		resp.RowsAffected = rowsAffected
	}
	return resp, nil
}

// Query runs a SQL query against a postgres database and returns results
func (s *server) Query(ctx context.Context, req *pb.QueryRequest) (*pb.QueryResponse, error) {
	db, err := s.getDB(ctx)
	if err != nil {
		return nil, err
	}
	var args []any
	for _, p := range req.Parameters {
		args = append(args, p)
	}
	rows, err := db.QueryContext(ctx, req.Query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resp := &pb.QueryResponse{}

	// Get the column metadata
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	for _, c := range cols {
		resp.Columninfo = append(resp.Columninfo, &pb.ColumnInfo{Name: c})
	}

	// We return the entire response in a single message.
	for rows.Next() {
		var row pb.Row
		var toScan []any
		for range cols {
			item := &pb.Row_Item{}
			row.Item = append(row.Item, item)
			toScan = append(toScan, scannableItem{item})
		}
		if err := rows.Scan(toScan...); err != nil {
			return nil, err
		}
		resp.Rows = append(resp.Rows, &row)
	}
	return resp, rows.Err()
}

// Register is called to expose this handler to the gRPC server
func (s *server) Register(gs *grpc.Server) {
	pb.RegisterPsqlServer(gs, s)
}

func init() {
	services.RegisterSansShellService(&server{})
}
