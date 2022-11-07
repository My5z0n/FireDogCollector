from multiprocessing import Process
from multiprocessing import Queue


from MLComponent.mlComponent import MLComponent

import MsgReceiver.rabbitmqReceiver as rabbitmqReceiver
import API.APIService as APIService

from repository.FireDogTracesRepository import Repository as TracesRepository


def main():
  

    spanNotificationQueue = Queue()
    startModelQueue = Queue()

    tmp = rabbitmqReceiver.RabbitmqReceiver(spanNotificationQueue)
    consumer_process = Process(target=tmp.Lisen)
    consumer_process.start()

    consumer_process = Process(target=APIService.Init,args=(startModelQueue,))
    consumer_process.start()

    tmp = MLComponent(spanNotificationQueue,startModelQueue,TracesRepository())
    consumer_process = Process(target=tmp.Process,args=())
    consumer_process.start()


    while True:
        s = queue.get()
        print("Got: ",s)






if __name__ == "__main__":
    main()