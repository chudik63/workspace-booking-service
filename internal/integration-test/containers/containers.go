package containers

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	postgresModule "github.com/testcontainers/testcontainers-go/modules/postgres"
)

func WorkspaceBookingPostgres(ctx context.Context) (*postgresModule.PostgresContainer, error) {
	options := []testcontainers.ContainerCustomizer{
		postgresModule.WithDatabase("workspace-booking"),
		testcontainers.WithExposedPorts("5430/tcp"),
		postgresModule.BasicWaitStrategies(),
	}

	ctn, err := postgresModule.Run(ctx, "postgres:latest", options...)
	if err != nil {
		return nil, fmt.Errorf("run postgres container: %w", err)
	}

	connStr, err := ctn.ConnectionString(ctx)
	if err != nil {
		return nil, fmt.Errorf("get connection string: %w", err)
	}

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("open postgres connection: %w", err)
	}

	err = goose.Up(db, "../../../migrations")
	if err != nil {
		return nil, fmt.Errorf("up migrations: %w", err)
	}

	return ctn, nil
}
