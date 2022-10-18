


import queue
import time
from turtle import Turtle


class MLComponent:
    modelReady = False
    newSpanNoticationQueue = None
    startLearningQueue = None
    repository = None
    
    def __init__(self,spanqueue,mlqueue,repository) -> None:
        self.repository = repository
        self.newSpanNoticationQueue=spanqueue
        self.startLearningQueue=mlqueue

    def Process(self):
        
        while True:
            ret = ""
            try:
                ret = self.startLearningQueue.get_nowait()
            except queue.Empty:
                pass
            try:
                ret = self.newSpanNoticationQueue.get_nowait()
                if self.modelReady is True:
                    self.calculateSpan(ret)
                else:
                    pass
            except queue.Empty:
                time.Sleep(0.5)



    def LearnModel(self):
        pass

    def calculateSpan(self,id):
        pass

