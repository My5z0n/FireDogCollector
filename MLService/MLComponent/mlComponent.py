import time
import queue
from multiprocessing import Queue
from MLComponent.mlModel import MLModel
from repository.FireDogTracesRepository import Repository
from multiprocessing import Event


class MLComponent:
    model_ready: bool = False
    new_span_notification_queue: Queue = None
    start_learning_queue: Queue = None
    repository: Repository = None
    model: MLModel = None
    stop_event: Event  = None

    def __init__(self, span_queue: Queue, ml_queue: Queue, stop_event: Event):
        self.repository = Repository()
        self.new_span_notification_queue = span_queue
        self.start_learning_queue = ml_queue
        self.model = MLModel()
        self.stop_event = stop_event

    def process_messages(self) -> None:
        ret = ""

        while not self.stop_event.is_set():
            try:
                ret = self.start_learning_queue.get_nowait()
                if ret[0] == 'START_TRAIN':
                    self.learn_model(ret[1])
                elif ret[0] == 'LOAD_MODEL':
                    self.LoadModel(ret[1])
            except queue.Empty:
                pass
            try:
                ret = self.new_span_notification_queue.get_nowait()
                if self.model_ready is True:
                    self.calculateSpan(ret)
                else:
                    pass
            except queue.Empty:
                time.sleep(0.5)
        exit()

    def learn_model(self, modelName) -> None:
        data = self.repository.get_paths_array()

        data = [v[0]["span_name"] for v in data]

        err = self.model.learn(data, modelName)
        if err != None:
            self.model_ready = True
        else:
            print("Error during model learning: "+err)

    def LoadModel(self, modelName):
        err = self.model.loadModel(modelName)
        if err != None:
            self.model_ready = True
        else:
            print("Error during model load: "+err)

    def calculateSpan(self, trace_id):
        data = self.repository.get_paths_array(trace_id)[0][0]
        result = self.model.predict(data)
        self.repository.setPrediction(trace_id, *result)
