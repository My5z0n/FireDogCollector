from multiprocessing import Queue
import time

import queue

from MLComponent.mlModel import MLModel
from repository.FireDogTracesRepository import Repository


class MLComponent:
    modelReady = False
    newSpanNotificationQueue = None
    startLearningQueue = None
    repository = None
    model = None

    def __init__(self, span_queue: Queue, ml_queue: Queue) -> None:
        self.repository = Repository()
        self.newSpanNotificationQueue = span_queue
        self.startLearningQueue = ml_queue
        self.model = MLModel()

    def Process(self):

        while True:
            ret = ""
            try:
                ret = self.startLearningQueue.get_nowait()
                print("Learning model!")
                #self.LearnModel()
            except queue.Empty:
                pass
            try:
                ret = self.newSpanNotificationQueue.get_nowait()
                print("Calculate Path")
                #if self.modelReady is True:
                #    self.calculateSpan(ret)
                #else:
                #    pass
            except queue.Empty:
                time.sleep(0.5)
                print("time end xdd")

    def LearnModel(self):
        data = self.repository.getPathsArray()

        data = [v[0]["span_name"] for v in data]
        self.model.learn(data)

    def calculateSpan(self, trace_id):
        data = self.repository.getPathsArray(trace_id)[0][0]
        result = self.model.predict(data)
        self.repository.setPrediction(trace_id, *result)
