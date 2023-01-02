from http.server import BaseHTTPRequestHandler, HTTPServer
from urllib.parse import urlparse
from time import sleep
from multiprocessing import Queue
hostName = "localhost"
serverPort = 9181


class MainServer(BaseHTTPRequestHandler):
    process_queue = None

    def do_GET(self) -> None:
        parseObj = urlparse(self.path)
        path = parseObj.path

        # Parse Query Params
        parsedParamas = [x.split("=") for x in parseObj.query.split("&")]
        queryParams = {}
        if len(parsedParamas[0]) > 1:
            queryParams = {x[0]: x[1] for x in parsedParamas}

        modelNameParam = "model"
        if "modelName" in queryParams:
            modelNameParam = queryParams["modelName"]

        # Route
        if path == '/START_TRAIN':
            self.process_queue.put(("START_TRAIN", modelNameParam))

        if path == '/LOAD_MODEL':
            self.process_queue.put(("LOAD_MODEL", modelNameParam))

        self.send_response(200)
        self.send_header("Content-type", "text/html")
        self.end_headers()


class APIService:

    def __init__(self, output_queue: Queue):
        server_handler = MainServer
        server_handler.process_queue = output_queue
        self.server = HTTPServer((hostName, serverPort), server_handler)

    def Run(self) -> None:
        self.server.serve_forever()
        exit()