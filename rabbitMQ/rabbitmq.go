package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ interface {
	PublishMessage(queueName, exchangeName string, message []byte) error
	ConsumeMessages(queueName string, handler func([]byte)) error
	DeclareQueue(queueName string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) error
	Close() error
}


type rabbitmq struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// triggerFunctions map[string]func([]byte)
}

type TriggerFunction struct {
	QueueName    string
	ExchangeName string
	// TriggerFunc  func(message []byte)
}

// func NewRabbitMQ(url string, triggerFunctions ...TriggerFunction) (RabbitMQ, error) {
func NewRabbitMQ(url string) (RabbitMQ, error) {
	conn, ch, err := connectToRabbitMQ(url)
	if err != nil {
		return nil, err
	}

	// triggerFunctionsMap := map[string]func([]byte){}

	// for _, triggerFunction := range triggerFunctions {
	// 	key := triggerFunction.QueueName + triggerFunction.ExchangeName
	// 	triggerFunctionsMap[key] = triggerFunction.TriggerFunc
	// }

	return &rabbitmq{
		conn:    conn,
		channel: ch,
		// triggerFunctions: triggerFunctionsMap,
	}, nil
}

func connectToRabbitMQ(url string) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	return conn, ch, nil
}

func (r *rabbitmq) DeclareQueue(queueName string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) error {
	_, err := r.channel.QueueDeclare(
		queueName,  // name
		durable,    // durable
		autoDelete, // delete when unused
		exclusive,  // exclusive
		noWait,     // no-wait
		args,       // arguments
	)
	if err != nil {
		log.Printf("Failed to declare queue %s: %s", queueName, err)
		return err
	}
	log.Printf("Queue declared: %s", queueName)
	return nil
}

func (r *rabbitmq) PublishMessage(queueName, exchangeName string, message []byte) error {
	err := r.channel.Publish(
		exchangeName, // exchange
		queueName,    // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
	if err != nil {
		log.Printf("Failed to publish message: %s", err)
		return err
	}
	log.Printf("Published message to %s: %s", queueName, string(message))
	return nil
}

func (r *rabbitmq) ConsumeMessages(queueName string, handler func([]byte)) error {
	msgs, err := r.channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			log.Printf("Received message: %s", msg.Body)
			handler(msg.Body)
		}
	}()

	return nil
}

func (r *rabbitmq) Close() error {
	if err := r.channel.Close(); err != nil {
		return err
	}
	if err := r.conn.Close(); err != nil {
		return err
	}
	return nil
}
