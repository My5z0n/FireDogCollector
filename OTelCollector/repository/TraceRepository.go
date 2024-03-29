package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/My5z0n/FireDogCollector/OtelCollector/models"
	"github.com/My5z0n/FireDogCollector/OtelCollector/utils"
	spanProcess "github.com/My5z0n/FireDogCollector/OtelCollector/utils/spanPathsProcessing"
	"regexp"
	"time"
)

type TraceRepository struct {
	connection clickhouse.Conn
}

func NewTraceRepository(config utils.Config) (TraceRepository, error) {
	t := TraceRepository{}

	return t, t.openConn(config)
}

func (r *TraceRepository) openConn(c utils.Config) error {

	conAdr := c.DbUrl + ":" + c.DbPort
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{conAdr},
		Auth: clickhouse.Auth{
			Database: "FireDogTraces",
		},
	})
	if err != nil {
		fmt.Print(err)
		return err
	}
	r.connection = conn

	return nil
}

func (r *TraceRepository) buildInsertSpanQuery(model models.ClickHouseSpan) (string, []any) {
	query := `INSERT INTO spans (trace_id,span_id,parent_span_id,span_name,start_time,end_time`

	keys := model.ExternalColNames
	args := []any{model.Trace_id, model.Span_id, model.Parent_span_id, model.Span_name, model.Start_time, model.End_time}

	for _, v := range keys {
		args = append(args, model.Attributes[v])
		query = fmt.Sprintf("%s,%s", query, v)
	}
	query = fmt.Sprintf("%s)", query)

	return query, args

}

func (r *TraceRepository) prepareSpanBatch(query string, model models.ClickHouseSpan) (driver.Batch, error) {
	var batch driver.Batch
	var err error
	for {
		batch, err = r.connection.PrepareBatch(context.Background(), query)
		if err != nil {
			ret := regexp.MustCompile(`No such column (.*?) in table`).FindAllStringSubmatch(err.Error(), -1)

			if len(ret) == 0 {
				return nil, err
			}

			missingColName := ret[0][1]
			if v, ok := model.Attributes[missingColName]; !ok {
				return nil, errors.New("Unable to create Span Batch")
			} else {
				var t string
				err, t = utils.GetTypeName(v)

				if err != nil {
					return nil, err
				}

				err = r.connection.Exec(context.Background(), fmt.Sprintf("ALTER TABLE spans ADD COLUMN IF NOT EXISTS %s Nullable(%s);", missingColName, t))
				if err != nil {
					return nil, err
				}
			}
		} else {
			return batch, nil
		}
	}

}

func (r *TraceRepository) SaveSpan(model models.ClickHouseSpan) error {
	query, argList := r.buildInsertSpanQuery(model)

	batch, err := r.prepareSpanBatch(query, model)
	if err != nil {
		return err
	}

	err = batch.Append(argList...)
	if err != nil {
		return err
	}

	return batch.Send()
}

func (r *TraceRepository) SaveTrace(paths [][]map[string]string, traceId string, startTime time.Time, jsonSpans string) error {

	flatSpan := spanProcess.FlattenSpansList(paths)[0]
	//TODO: Handle many paths

	batch, err := r.connection.PrepareBatch(context.Background(), "INSERT INTO traces ")
	if err != nil {
		return err
	}
	err = batch.AppendStruct(&struct {
		Trace     string                `ch:"trace_id"`
		Paths     string                `ch:"paths"`
		Array     [][]map[string]string `ch:"paths_array"`
		StartTime time.Time             `ch:"start_time"`
		JsonSpans string                `ch:"json_spans"`
	}{
		Trace:     traceId,
		Paths:     flatSpan,
		Array:     paths,
		StartTime: startTime,
		JsonSpans: jsonSpans,
	})

	if err != nil {
		return err
	}

	return batch.Send()
}
