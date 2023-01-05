CREATE TABLE FireDogTraces.traces
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
CREATE TABLE FireDogTraces.spans
(
    trace_id String,
    span_id String,
    parent_span_id String,
    span_name String,
    start_time DATETIME,
    end_time DATETIME,
    attributes JSON
)
ENGINE = MergeTree()
PRIMARY KEY (trace_id, span_id);

CREATE TABLE FireDogTraces.predictions
(
    trace_id String,
    anomaly_detected Nullable(Bool),
    span_name String,
    span_id String,
    expected_span_name String
)
ENGINE = MergeTree()
PRIMARY KEY (trace_id,anomaly_detected);

----------------------------------------------------
CREATE MATERIALIZED VIEW mqtraceprocess
ENGINE = RabbitMQ SETTINGS rabbitmq_address = 'amqp://guest:guest@some-rabbit:5672/',
                            rabbitmq_exchange_name = 'newSpanToProcessNotification',
                            rabbitmq_format = 'JSONEachRow',
                            rabbitmq_exchange_type = 'fanout'
AS SELECT trace_id from traces;