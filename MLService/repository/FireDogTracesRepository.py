import clickhouse_connect
from pandas import DataFrame
import numpy as np

class Repository:
    client = None
    def connect(self) -> DataFrame:
        self.client = clickhouse_connect.get_client(host='localhost',
                                        database = 'FireDogTraces',
                                        query_limit=0)
        data = self.client.query_np('SELECT pathsArray FROM dogdig')["pathsArray"]

        new_array  = np.array([v[0] for v in data])


        print(new_array)
        return new_array




if __name__ == "__main__":
    x= Repository()
    x.connect()