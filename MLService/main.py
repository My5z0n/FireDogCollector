from ast import While
from time import sleep
from random import random
from multiprocessing import Process
from multiprocessing import Queue

from certifi import where
from matplotlib.pyplot import prism
from MLComponent.mlComponent import MLComponent

import rabbitmqReceiver
import APIService

from repository.FireDogTracesRepository import Repository as TracesRepository


def main():
  

    spanNotificationQueue = Queue()
    startModelQueue = Queue()

    x = rabbitmqReceiver.RabbitmqReceiver(spanNotificationQueue)
    consumer_process = Process(target=x.Lisen)
    consumer_process.start()

    consumer_process = Process(target=APIService.Init,args=(startModelQueue,))
    consumer_process.start()

    x = MLComponent(spanNotificationQueue,startModelQueue,TracesRepository())
    consumer_process = Process(target=x.Process,args=())
    consumer_process.start()


    while True:
        s = queue.get()
        print("Got: ",s)






if __name__ == "__main__":
    main()