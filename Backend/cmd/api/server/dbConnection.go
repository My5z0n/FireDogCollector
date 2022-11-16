package server

import (
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"log"
)

type dbConnectionSettings struct {
	Name string
}

func CreateDBConnection() *driver.Conn {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"localhost:9001"},
		Auth: clickhouse.Auth{
			Database: "FireDogTraces",
		},
	})

	if err != nil {
		log.Panicf("Error during CreateDBConnection: %v", err)
	}
	return &conn
}
