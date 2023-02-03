package server

import (
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/My5z0n/FireDogCollector/Backend/Settings"
	"log"
)

func CreateDBConnection(c Settings.Config) driver.Conn {

	addrCon := c.DbUrl + ":" + c.DbPort
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{addrCon},
		Auth: clickhouse.Auth{
			Database: "FireDogTraces",
		},
	})

	if err != nil {
		log.Panicf("Error during CreateDBConnection: %v", err)
	}
	return conn
}
