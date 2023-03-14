import clickhouse_connect


class ClickhouseRepository():
    host = "localhost"
    database = "FireDogTraces"
    unwanted_columns = ['trace_id', 'span_id',
                        'parent_span_id', 'span_name', 'start_time', 'end_time']

    def __init__(self):
        self.connection = None
        self.client = clickhouse_connect.get_client(host=self.host,
                                                    database=self.database,
                                                    query_limit=0)

    def getColumnNames(self, table):
        query = f"DESCRIBE TABLE {table}"
        result = self.client.query_df(query)
        column_names = result['name'].values.tolist()
        column_names.insert(0, "Test")
        column_names = [
            col for col in column_names if col not in self.unwanted_columns]
        return column_names


if __name__ == "__main__":
    x = ClickhouseRepository()
    x.getColumnNames("spans")
