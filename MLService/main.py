import signal
from multiprocessing import Process, Queue, Event
from MLComponent.mlComponent import MLComponent
from API.APIService import APIService
from MsgReceiver.rabbitmqReceiver import RabbitmqReceiver
from time import sleep
from threading import Thread
import debugpy

def work_ml_component(span_queue: Queue, model_queue: Queue) ->None:
    stop_event = Event()
    MLComponent(span_queue, model_queue, stop_event).process_messages()


def work_API(model_queue: Queue) -> None:
    api = APIService(model_queue)
    api.Run()


def work_rabbit(span_queue: Queue) -> None:
    rabbit = RabbitmqReceiver(span_queue)
    rabbit.lisen()


class MainObj:

    def close_handler(self, a, b,c=0) -> None:
        print("Closing")
        self.ml_process.terminate()
        self.api_process.terminate()
        self.rabbit_process.terminate()
        print("Exiting")
        exit(c)

    def init(self) -> None:
        self.span_notification_queue = Queue()
        self.start_model_queue = Queue()

        self.rabbit_process = Process(
            target=work_rabbit, args=(self.span_notification_queue, ))
        self.rabbit_process.start()

        self.api_process = Process(target=work_API,
                                   args=(self.start_model_queue, ))
        self.api_process.start()

        self.ml_process = Process(target=work_ml_component,
                                  args=(self.span_notification_queue, self.start_model_queue, ))
        self.ml_process.start()

        signal.signal(signal.SIGINT, self.close_handler)
        
        #Debug
        print("Starting main model")
        self.start_model_queue.put(("LOAD_MODEL", "main_model"))
        
        while True:
            if not self.rabbit_process.is_alive():
                print("Rabbit died")
                self.close_handler(None, None,-1)
            if not self.api_process.is_alive():
                print("API died")
                self.close_handler(None, None,-1)
            if not self.ml_process.is_alive():
                print("ML died")
                self.close_handler(None, None,-1)
            sleep(5)
        



if __name__ == "__main__":
    debugpy.listen(("0.0.0.0", 5678))
    #print("Waiting for client to attach...")
   # debugpy.wait_for_client()
    main = MainObj()
    main.init()
