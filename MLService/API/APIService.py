from http.server import BaseHTTPRequestHandler, HTTPServer

hostName = "localhost"
serverPort = 8090


class MyServer(BaseHTTPRequestHandler):
    procesQueue = None


    def do_GET(self):

        if self.path == '/a':
            print("A")
        elif self.path == '/b':
            print("B")
        self.send_response(200)
        self.send_header("Content-type", "text/html")
        self.end_headers()
        self.procesQueue.put("START_TRAIN")



def Init(q): 
    x = MyServer
    x.procesQueue = q
    webServer = HTTPServer((hostName, serverPort), MyServer)
    
    webServer.serve_forever()

