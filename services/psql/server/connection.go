package server

import (
	"context"
	"crypto/x509"
	"database/sql"
	"fmt"
	"net"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"

	_ "embed"
)

var (
	// DatabaseUser is the username to use when connecting to the database. Binding this to a flag is often useful.
	DatabaseUser string

	// DatabaseHost is the hostname of the database to connect to. Binding this to a flag is often useful.
	DatabaseHost string

	// DatabaseName is the name of the database to use. Binding this to a flag is often useful.
	DatabaseName string

	// AuthMethod is the way to authenticate to a database. Binding this to a flag is often useful.
	AuthMethod string

	// AWSRegion specifies the region when using AWS authentication. Binding this to a flag is often useful.
	AWSRegion string
)

const (
	// AuthMethodAWS authenticates to an RDS instance using AWS IAM Authentication.
	// AWSRegion must be provided.
	AuthMethodAWS = "aws"
	// AuthMethodAzure authenticates to an Azure Database for PostgreSQL instance using Azure AD authentication.
	AuthMethodAzure = "azure"
	// AuthMethodGCP authenticates to a CloudSQL instance using GCP IAM Authentication and the Cloud SQL Go Connector.
	// Database host should look like "project-name:region:instance-name".
	AuthMethodGCP = "gcp"
	// AuthMethodNoAuth attempts to directly connect to a SQL database without using a password.
	AuthMethodNoAuth = "noauth"
)

//go:embed certs/rds-combined-ca-bundle.pem
var rdsCombinedCABundle []byte

func connect(ctx context.Context) (*sql.DB, error) {
	switch AuthMethod {
	case AuthMethodAWS:
		return connectAWS(ctx)
	case AuthMethodAzure:
		return connectAzure(ctx)
	case AuthMethodGCP:
		return connectGCP(ctx)
	case AuthMethodNoAuth:
		return connectNoAuth()
	case "":
		return nil, status.Errorf(codes.FailedPrecondition, "postgres access is not configured on this server")
	default:
		return nil, status.Errorf(codes.FailedPrecondition, "unknown postgres authentication method %v", AuthMethod)
	}
}

// From https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.Connecting.Go.html#UsingWithRDS.IAMDBAuth.Connecting.GoV2
func connectAWS(ctx context.Context) (*sql.DB, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load aws config: %v", err)
	}

	authenticationToken, err := auth.BuildAuthToken(
		ctx, net.JoinHostPort(DatabaseHost, "5432"), AWSRegion, DatabaseUser, cfg.Credentials)
	if err != nil {
		return nil, fmt.Errorf("failed to create authentication token: %v", err)
	}

	dsn := fmt.Sprintf("host=%s  user=%s password=%s dbname=%s",
		DatabaseHost, DatabaseUser, authenticationToken, DatabaseName,
	)
	connConfig, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	rdsCerts := x509.NewCertPool()
	rdsCerts.AppendCertsFromPEM(rdsCombinedCABundle)
	connConfig.TLSConfig.RootCAs = rdsCerts
	connStr := stdlib.RegisterConnConfig(connConfig)
	return sql.Open("pgx", connStr)
}

// From https://learn.microsoft.com/en-us/azure/service-connector/tutorial-passwordless?tabs=user%2Cgo%2Csql-me-id-dotnet%2Cappservice&pivots=postgresql#connect-to-a-database-with-microsoft-entra-authentication
func connectAzure(ctx context.Context) (*sql.DB, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, fmt.Errorf("could not load azure config: %v", err)
	}
	authenticationToken, err := cred.GetToken(ctx, policy.TokenRequestOptions{
		Scopes: []string{"https://ossrdbms-aad.database.windows.net/.default"},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create authentication token: %v", err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=verify-full",
		DatabaseHost, DatabaseUser, authenticationToken, DatabaseName,
	)
	return sql.Open("pgx", dsn)
}

// From https://cloud.google.com/sql/docs/postgres/iam-logins#go
func connectGCP(ctx context.Context) (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s", DatabaseUser, DatabaseName)
	connConfig, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	dialer, err := cloudsqlconn.NewDialer(ctx, cloudsqlconn.WithIAMAuthN())
	if err != nil {
		return nil, err
	}
	connConfig.DialFunc = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return dialer.Dial(ctx, DatabaseHost)
	}

	connStr := stdlib.RegisterConnConfig(connConfig)
	return sql.Open("pgx", connStr)
}

// Useful for local postgres connections
func connectNoAuth() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s", DatabaseHost, DatabaseUser, DatabaseName)
	return sql.Open("pgx", dsn)
}
