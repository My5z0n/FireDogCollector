import clickhouse_connect
import numpy as np

class Repository:
    client = None
    host = ""
    database = ""


    def __init__(self,host='localhost',database='FireDogTraces'):
        self.host = host
        self.database = database
        self.client = clickhouse_connect.get_client(host=self.host,
                                        database = self.database,
                                        query_limit=0)
    
    def getPatshArray(self):
        if self.client is not None:
             data = self.client.query_np('SELECT pathsArray FROM dogdig')["pathsArray"]
             return data
        else:
            raise Exception("No DB connected")
    

if __name__ == "__main__":
    x= Repository()
    x.connect()
    x.getPatshArray