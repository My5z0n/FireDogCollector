CREATE TABLE FireDogTraces.dogdig
(
    trace_id String,
    paths String,
    pathsArray Array(Array(Tuple(span_name String, span_id String))),
    firedog_test1 String,
    firedog_test2 String,
    firedog_test3 String
)
ENGINE = MergeTree()
PRIMARY KEY (trace_id, paths);

SET allow_experimental_object_type=1;
CREATE TABLE FireDogTraces.traces
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
    anomaly_detected Bool,
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
AS SELECT trace_id from dogdig;
-----------------------------------------------------