import clickhouse_connect


class Repository:
    client: clickhouse_connect.Client = None
    host: str = ""
    database: str = ""

    def __init__(self, host: str = 'localhost', database: str = 'FireDogTraces'):
        self.host = host
        self.database = database
        self.client = clickhouse_connect.get_client(host=self.host,
                                                    database=self.database,
                                                    query_limit=0)

    def get_paths_array(self, trace_id: str = "") -> any:
        if self.client is not None:
            rule = ""
            if trace_id != "":
                rule = f"WHERE trace_id='{trace_id}'"
            data = self.client.query_np(f'SELECT pathsArray FROM dogdig {rule}')[
                "pathsArray"]
            return data
        else:
            raise Exception("No DB connected")

    def set_prediction(self, trace_id: str, anomaly_detected: bool,
                      span_name: str = "", span_id: str = "", expected_span_name: str = "") -> None:
        try:
            self.client.insert(table="predictions",
                               data=[[trace_id, anomaly_detected, span_name, span_id, expected_span_name]])
        except Exception as e:
            print(e)


if __name__ == "__main__":
    x = Repository()
    xD = (True, "1", "2", "3")
    x.set_prediction("22", *xD)
