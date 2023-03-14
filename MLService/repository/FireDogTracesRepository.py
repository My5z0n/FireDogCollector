import clickhouse_connect
import os

class Repository:
    client = None
    host: str = ""
    database: str = ""

    def __init__(self):
        db_host = os.getenv("DB_SERVER_URL")
        if db_host is None:
            self.host = "localhost"
        else:
            self.host = db_host        
        self.database = 'FireDogTraces'
        self.client = clickhouse_connect.get_client(host=self.host,
                                                    database=self.database,
                                                    query_limit=0)

    def get_paths_array(self, trace_id: str = "") -> any:
        if self.client is not None:
            rule = ""
            if trace_id != "":
                rule = f"WHERE trace_id='{trace_id}'"
            data = self.client.query_np(f'SELECT paths_array FROM traces {rule}')[
                "paths_array"]
            return data
        else:
            raise Exception("No DB connected")

    def set_prediction(self, trace_id: str, anomaly_detected: bool,
                      span_name: str = "", span_id: str = "", expected_span_name: str = "", position: int = 0) -> None:
        try:
            self.client.insert(table="predictions",
                               data=[[trace_id, anomaly_detected, span_name, span_id, expected_span_name,position]])
        except Exception as e:
            print(e)


if __name__ == "__main__":
    pass
   #Debug
