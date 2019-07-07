package mq

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

// MQ structure
type MQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// NewMQ to Initialize MQ
func NewMQ() MQ {
	return MQ{}
}

// Connect to MQ
func (mq *MQ) Connect(connString string) {
	conn, err := amqp.Dial(connString)
	failOnError(err, "Failed to connect to RabbitMQ")
	mq.conn = conn
}

// CreateChannel With MQ
func (mq *MQ) CreateChannel() {
	ch, err := mq.conn.Channel()
	failOnError(err, "Failed to open a channel")
	mq.ch = ch
}

// CreateExchange to create new exchange
func (mq *MQ) CreateExchange(name string, kind string, durable bool) {

	err := mq.ch.ExchangeDeclare(name, kind, durable, false, false, false, nil)
	failOnError(err, "Failed to declare exchange")

}

// Publish message
func (mq *MQ) Publish(exchange string, queue string, data interface{}) {
	body, merr := json.Marshal(data)
	failOnError(merr, "Unable to Marshal Data")

	err := mq.ch.Publish(
		exchange,
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failOnError(err, "Unable to Publish Message")
}

// CreateQueue in mq
func (mq *MQ) CreateQueue(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp.Table) {
	_, err := mq.ch.QueueDeclare(
		name,
		durable,
		autoDelete,
		exclusive,
		noWait,
		args,
	)
	failOnError(err, "Unable to Declare Queue")
}

// WriteToQueue function
func (mq *MQ) WriteToQueue(queue string, data interface{}) {
	body, merr := json.Marshal(data)
	failOnError(merr, "Unable to Marshal Data")

	err := mq.ch.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	failOnError(err, "Unable to write to queue")
}

// BindQueueWithExchange function
func (mq *MQ) BindQueueWithExchange(queue string, exchange string, key string, args amqp.Table) {
	err := mq.ch.QueueBind(
		queue,
		key,
		exchange,
		false,
		args,
	)
	failOnError(err, "Unable to bind queue to exchange")
}

// EstablishWorker method
func (mq *MQ) EstablishWorker(queue string) <-chan amqp.Delivery {

	msgs, err := mq.ch.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	failOnError(err, "Failed to register a consumer")

	return msgs
}
