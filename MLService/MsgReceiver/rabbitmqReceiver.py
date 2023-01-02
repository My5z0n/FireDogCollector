import pika
import json
from time import sleep


class RabbitmqReceiver():
    con = 'localhost'
    exchange_name = 'newSpanToProcessNotification'
    process_queue = None
    queue_name = None

    def __init__(self, procesQueue) -> None:
        self.process_queue = procesQueue

    def Lisen(self):
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

    def callback(self, ch, method, properties, body):
        self.process_queue.put(json.loads(body)["trace_id"])
