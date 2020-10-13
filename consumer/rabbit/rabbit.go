package rabbit

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// RabbitMQ connection global instance
var RabbitMQ *amqp.Connection
var MailQueue *amqp.Queue
var MailChannel *amqp.Channel

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// ConnectToRabbit connects to RabbitMQ instance
func ConnectToRabbit(host string, port string, user string, password string) {
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)
	instanceTmp, err := amqp.Dial(connectionString)

	numberOfTest := 0

	for err != nil && numberOfTest < 5 {
		fmt.Println(err)
		fmt.Println("Connection to the rabbitMQ did not succeed, new try")

		time.Sleep(5 * time.Second)
		instanceTmp, err = amqp.Dial(connectionString)

		numberOfTest++
	}

	failOnError(err, "Failed to connect to RabbitMQ")
	log.Println("Connected to RabbitMQ server successfullyed!")

	channel, err := instanceTmp.Channel()

	failOnError(err, "Failed to open a channel")

	q, err := channel.QueueDeclare(
		"mails", // name
		false,   // durable
		true,    // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	RabbitMQ = instanceTmp
	MailQueue = &q
	MailChannel = channel
}

// Receive gets consume message from a queue
func Receive() {
	msgs, err := MailChannel.Consume(
		MailQueue.Name, // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
