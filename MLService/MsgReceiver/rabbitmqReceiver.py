import pika
import json

class RabbitmqReceiver():
    con = 'localhost'
    exchange_name = 'newSpanToProcessNotification'
    process_queue = None
    queue_name = None

    def __init__(self, procesQueue) -> None:
        self.process_queue = procesQueue

    def Lisen(self):
        connection = pika.BlockingConnection(pika.ConnectionParameters(self.con))
        channel = connection.channel()
        channel.exchange_declare(exchange=self.exchange_name, exchange_type='fanout',durable=True)
        self.queue_name = channel.queue_declare(queue='',exclusive=True).method.queue
        
        channel.queue_bind(exchange=self.exchange_name, queue=self.queue_name)

        channel.basic_consume(queue=self.queue_name,
                              auto_ack=True,
                              on_message_callback=self.callback)
        channel.start_consuming()

    def callback(self, ch, method, properties, body):
        self.process_queue.put(json.loads(body)["trace_id"])
