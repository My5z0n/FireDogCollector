package repository

import (
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/My5z0n/FireDogCollector/models"
)

type TraceRepository struct {
	port       string
	database   string
	connection clickhouse.Conn
}

func NewTraceRepository(port string, database string) (TraceRepository, error) {
	t := TraceRepository{
		port:     port,
		database: database,
	}

	return t, t.openConn()
}

func (r *TraceRepository) openConn() error {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"localhost:9001"},
		Auth: clickhouse.Auth{
			Database: "FiredogTraces",
		},
	})
	if err != nil {
		return err
	}
	r.connection = conn

	return nil
}

func (r *TraceRepository) SaveSpan(model models.SaveSpan) error{
	r.connection.
}
