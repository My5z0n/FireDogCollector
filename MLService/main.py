import signal
from multiprocessing import Process, Queue, Event
from MLComponent.mlComponent import MLComponent
from API.APIService import APIService
from MsgReceiver.rabbitmqReceiver import RabbitmqReceiver
from time import sleep
from threading import Thread


def work_ml_component(span_queue: Queue, model_queue: Queue) ->None:
    stop_event = Event()
    signal.signal(signal.SIGINT, lambda a, b: stop_event.set())
    MLComponent(span_queue, model_queue, stop_event).process_messages()


def work_API(model_queue: Queue) -> None:
    api = APIService(model_queue)
    signal.signal(signal.SIGINT, lambda a,b: Thread(
        target=lambda: api.server.shutdown()).start())
    api.Run()


def work_rabbit(span_queue: Queue) -> None:
    rabbit = RabbitmqReceiver(span_queue)
    signal.signal(signal.SIGINT, lambda a,b:
                  Thread(target=lambda: rabbit.channel.stop_consuming()).start())
    rabbit.lisen()


class MainObj:

    def close_handler(self, a, b) -> None:
        print("Closing")
        self.ml_process.join()
        self.api_process.join()
        self.rabbit_process.join()
        print("Exiting")
        exit()

    def init(self) -> None:
        print("Init")
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
        print("Completed")
        self.start_model_queue.put(("LOAD_MODEL", "main"))
        



if __name__ == "__main__":
    main = MainObj()
    main.init()
