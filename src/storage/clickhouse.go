package storage

import (
	"database/sql"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func NewConnection() *sql.DB {
	return clickhouse.OpenDB(&clickhouse.Options{
		Addr:     []string{"54.236.30.177:8123"},
		Protocol: clickhouse.HTTP,
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: ">tN#G$g*Mp;g",
		},
	})
}
