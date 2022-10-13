package repository

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/My5z0n/FireDogCollector/models"
	"reflect"
	"regexp"
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

	columnNames := []string{"trace_id", "paths"}
	//params := fmt.Sprintf("%s,%s", traceId, flatSpan)

	list := []reflect.StructField{
		{
			Name: "C1",
			Type: reflect.TypeOf(""),
			Tag:  `ch:"trace_id"`,
		},
		{
			Name: "C2",
			Type: reflect.TypeOf(""),
			Tag:  `ch:"paths"`,
		},
	}

	var re = regexp.MustCompile(`\.`)
	counter := 3

	tmparr := []any{traceId, flatSpan}
	for k, v := range attributes {
		columnNames = append(columnNames, re.ReplaceAllString(k, "_"))
		tmparr = append(tmparr, v)
		list = append(list, reflect.StructField{
			Name: fmt.Sprintf("C%v", counter),
			Type: reflect.TypeOf(""),
			Tag:  reflect.StructTag(fmt.Sprintf(`ch:"%s"`, re.ReplaceAllString(k, "_"))),
		})
		counter += 1
	}

	obj := reflect.StructOf(list)
	if obj.Kind() != reflect.Struct {
		fmt.Println("Error")
	} else {
		fmt.Println("Good")
	}
	//fmt.Println(reflect.ValueOf(obj).Kind())

	//err := r.connection.Exec(context.Background(), "INSERT INTO dogdig (?) VALUES (?)", columnNames, tmparr)
	batch, err := r.connection.PrepareBatch(context.Background(), "INSERT INTO dogdig")
	err = batch.AppendStruct(obj)
	if err != nil {
		return err
	}

	return batch.Send()
}

func flattenSpansList(paths [][]string) []string {

	flattenPaths := make([]string, 0)

	for range paths {
		str := ""
		for _, k := range paths {
			str = fmt.Sprintf("%s'%s", str, k)
		}
		str = fmt.Sprintf("%s'!END|", str)
		flattenPaths = append(flattenPaths, str)
	}

	return flattenPaths
}
