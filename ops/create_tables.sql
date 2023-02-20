CREATE DATABASE IF NOT EXISTS FireDogTraces;


CREATE TABLE IF NOT EXISTS FireDogTraces.traces
(
    trace_id String,
    paths String,
    paths_array Array(Array(Tuple(span_name String, span_id String))),
    start_time DATETIME,
    json_spans String
)
ENGINE = MergeTree()
PRIMARY KEY (trace_id, paths);

SET allow_experimental_object_type=1;
CREATE TABLE  IF NOT EXISTS FireDogTraces.spans
(
    trace_id String,
    span_id String,
    parent_span_id String,
    span_name String,
    start_time DATETIME,
    end_time DATETIME
)
ENGINE = MergeTree()
PRIMARY KEY (trace_id, span_id);

CREATE TABLE  IF NOT EXISTS FireDogTraces.predictions
(
    trace_id String,
    anomaly_detected Nullable(Bool),
    span_name String,
    span_id String,
    expected_span_name String
)
ENGINE = MergeTree()
PRIMARY KEY (trace_id);

----------------------------------------------------
CREATE MATERIALIZED VIEW  IF NOT EXISTS mqtraceprocess
ENGINE = RabbitMQ SETTINGS rabbitmq_address = 'amqp://guest:guest@rabbitmq:5672/',
                            rabbitmq_exchange_name = 'newSpanToProcessNotification',
                            rabbitmq_format = 'JSONEachRow',
                            rabbitmq_exchange_type = 'fanout'
AS SELECT trace_id from FireDogTraces.traces;