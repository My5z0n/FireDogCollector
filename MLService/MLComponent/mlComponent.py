from asyncio import QueueEmpty
from multiprocessing import Queue
import time

from repository.FireDogTracesRepository import Repository



class MLComponent:
    modelReady = False
    newSpanNoticationQueue = None
    startLearningQueue = None
    repository = None
    
    def __init__(self, spanqueue: Queue, mlqueue: Queue, repository : Repository) -> None:
        self.repository = repository
        self.newSpanNoticationQueue=spanqueue
        self.startLearningQueue=mlqueue

    def Process(self):
        
        while True:
            ret = ""
            try:
                ret = self.startLearningQueue.get_nowait()
                self.LearnModel()
            except QueueEmpty:
                pass
            try:
                ret = self.newSpanNoticationQueue.get_nowait()
                if self.modelReady is True:
                    self.calculateSpan(ret)
                else:
                    pass
            except QueueEmpty:
                time.Sleep(0.5)



    def LearnModel(self):
        data = self.repository.getPatshArray()

        data = [v[0] for v in data]


    def calculateSpan(self,id):
        pass