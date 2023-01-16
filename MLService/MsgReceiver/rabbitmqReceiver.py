import pika
import json
from time import sleep
from multiprocessing import Queue

class RabbitmqReceiver():
    con: str = 'rabbitmq'
    exchange_name: str = 'newSpanToProcessNotification'
    process_queue: Queue = None
    queue_name: str = None

    def __init__(self, procesQueue):
        self.process_queue = procesQueue

    def lisen(self) -> None:
        connection = pika.BlockingConnection(
            pika.ConnectionParameters(self.con))
        self.channel = connection.channel()
        self.channel.exchange_declare(
            exchange=self.exchange_name, exchange_type='fanout', durable=True)
        self.queue_name = self.channel.queue_declare(
            queue='', exclusive=True).method.queue

        self.channel.queue_bind(
            exchange=self.exchange_name, queue=self.queue_name)

        self.channel.basic_consume(queue=self.queue_name,
                                   auto_ack=True,
                                   on_message_callback=self.callback)
        self.channel.start_consuming()

        exit()

    def callback(self, ch, method, properties, body) -> None:
        self.process_queue.put(json.loads(body)["trace_id"])
