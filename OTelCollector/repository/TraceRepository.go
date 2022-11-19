package repository

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/My5z0n/FireDogCollector/OtelCollector/models"
	"time"
)

type TraceRepository struct {
	port       string
	database   string
	connection clickhouse.Conn
	mapper     map[string]string
}

func NewTraceRepository(port string, database string) (TraceRepository, error) {
	t := TraceRepository{
		port:     port,
		database: database,
	}
	t.mapper = map[string]string{
		"firedog.test1": "firedog_test1",
		"firedog.test2": "firedog_test2",
		"firedog.test3": "firedog_test3",
	}

	return t, t.openConn()
}

func (r *TraceRepository) openConn() error {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"localhost:9001"},
		Auth: clickhouse.Auth{
			Database: "FireDogTraces",
		},
	})
	if err != nil {
		return err
	}
	r.connection = conn

	return nil
}

func (r *TraceRepository) SaveSpan(model models.ClickHouseSpan) error {
	//comand := fmt.Sprintf("INSERT INTO trace ()")
	//r.connection.AsyncInsert()
	batch, err := r.connection.PrepareBatch(context.Background(), "INSERT INTO spans")
	if err != nil {
		return err
	}
	err = batch.AppendStruct(&model)
	if err != nil {
		return err
	}

	return batch.Send()
}

func (r *TraceRepository) SaveTrace(paths [][]map[string]string, traceId string, startTime time.Time) error {

	//tmpSpan := flattenSpansList(paths)
	flatSpan := flattenSpansList(paths)[0]
	//TODO: Handle many paths
	fmt.Println(flatSpan)

	batch, err := r.connection.PrepareBatch(context.Background(), "INSERT INTO traces ")
	if err != nil {
		return err
	}
	err = batch.AppendStruct(&struct {
		trace_id    string
		paths       string
		paths_array [][]map[string]string
		start_time  time.Time
	}{
		trace_id:    traceId,
		paths:       flatSpan,
		paths_array: paths,
		start_time:  startTime,
	})

	if err != nil {
		return err
	}

	return batch.Send()
}

func flattenSpansList(paths [][]map[string]string) []string {

	flattenPaths := make([]string, 0)

	for _, v := range paths {
		str := v[0]["span_name"]
		for i := 1; i < len(v); i++ {
			str = fmt.Sprintf("%s#%s", str, v[i]["span_name"])
		}
		flattenPaths = append(flattenPaths, str)
	}

	return flattenPaths
}
