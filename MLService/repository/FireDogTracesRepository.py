import clickhouse_connect


class Repository:
    client = None
    host = ""
    database = ""

    def __init__(self, host='localhost', database='FireDogTraces'):
        self.host = host
        self.database = database
        self.client = clickhouse_connect.get_client(host=self.host,
                                                    database=self.database,
                                                    query_limit=0)

    def getPathsArray(self, trace_id=""):
        if self.client is not None:
            rule = ""
            if trace_id != "":
                rule = f"WHERE trace_id='{trace_id}'"
            data = self.client.query_np(f'SELECT pathsArray FROM dogdig {rule}')["pathsArray"]
            return data
        else:
            raise Exception("No DB connected")

    def setPrediction(self, trace_id: str, anomaly_detected: bool,
                      span_name: str = "", span_id: str = "", expected_span_name: str = ""):
        try:
            self.client.insert(table="predictions",
                               data=[[trace_id, anomaly_detected, span_name, span_id, expected_span_name]])
        except Exception as e:
            print(e)


if __name__ == "__main__":
    x = Repository()
    xD = (True, "1", "2", "3")
    x.setPrediction("22", *xD)
