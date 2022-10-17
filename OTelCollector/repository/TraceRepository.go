package repository

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/My5z0n/FireDogCollector/models"
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
	batch, err := r.connection.PrepareBatch(context.Background(), "INSERT INTO traces")
	if err != nil {
		return err
	}
	err = batch.AppendStruct(&model)
	if err != nil {
		return err
	}

	return batch.Send()
}

func (r *TraceRepository) SaveDogDig(paths [][]string, traceId string, attributes map[string]any) error {

	flatSpan := flattenSpansList(paths)[0]
	//TODO: Handle many paths
	fmt.Println(flatSpan)

	columnNames := []string{}

	tmparr := []any{traceId, flatSpan}

	for k, v := range attributes {
		columnNames = append(columnNames, r.mapper[k])
		tmparr = append(tmparr, v)

	}

	params := "trace_id, paths"
	for _, v := range columnNames {
		params = fmt.Sprintf("%s,%s", params, v)
	}

	err := r.connection.Exec(context.Background(), fmt.Sprintf("INSERT INTO dogdig (%s) VALUES (?)", params), tmparr)

	if err != nil {
		return err
	}

	return nil
}

func flattenSpansList(paths [][]string) []string {

	flattenPaths := make([]string, 0)

	for _, v := range paths {
		str := v[len(v)-1]
		for i := len(v) - 2; i >= 0; i-- {
			str = fmt.Sprintf("%s#%s", str, v[i])
		}
		str = fmt.Sprintf("%s'!END", str)
		flattenPaths = append(flattenPaths, str)
	}

	return flattenPaths
}
