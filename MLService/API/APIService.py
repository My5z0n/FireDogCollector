from http.server import BaseHTTPRequestHandler, HTTPServer

hostName = "localhost"
serverPort = 9181


class MyServer(BaseHTTPRequestHandler):
    process_queue = None

    def do_GET(self):

        if self.path == '/START_TRAIN':
            self.process_queue.put("START_TRAIN")

        self.send_response(200)
        self.send_header("Content-type", "text/html")
        self.end_headers()
        
def Init(q):
    x = MyServer
    x.process_queue = q
    web_server = HTTPServer((hostName, serverPort), MyServer)

    web_server.serve_forever()
