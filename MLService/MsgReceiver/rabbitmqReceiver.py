import pika


class RabbitmqReceiver():
    con = 'localhost'
    queueName = 'insertion_queue'
    procesQueue = None

    def __init__(self,procesQueue) -> None:
        self.procesQueue = procesQueue

    def Lisen(self):
        connection = pika.BlockingConnection(pika.ConnectionParameters(self.con))
        channel = connection.channel()
        channel.queue_declare(queue=self.queueName)
         
        channel.basic_consume(queue=self.queueName,
                            auto_ack=True,
                            on_message_callback=self.callback)
        channel.start_consuming()

    
    def callback(self,ch, method, properties, body):
        self.procesQueue.put(body.decode('UTF-8'))

