from multiprocessing import Process
from multiprocessing import Queue

from MLComponent.mlComponent import MLComponent

import MsgReceiver.rabbitmqReceiver as rabbitmqReceiver
import API.APIService as APIService


def workMLComponent(q1, q2):
    MLComponent(q1, q2).Process()


def work_API(q1):
    APIService.Init(q1)


def work_Rabbit(q1):
    rabbitmqReceiver.RabbitmqReceiver(q1).Lisen()


def main():
    span_notification_queue = Queue()
    start_model_queue = Queue()

    consumer_process = Process(target=work_Rabbit, args=(span_notification_queue,))
    consumer_process.start()

    consumer_process = Process(target=work_API, args=(start_model_queue,))
    consumer_process.start()

    consumer_process = Process(target=workMLComponent,
                               args=(span_notification_queue, start_model_queue))
    consumer_process.start()

    print("LOL")


if __name__ == "__main__":
    main()
