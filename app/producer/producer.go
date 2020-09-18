package producer

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
	log.Println("Connected to RabbitMQ server successfully!")

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

// PublishMailData sends mail data to message broker
func PublishMailData(subject string, content string, from string, to []string) {

	for i := 0; i < len(to); i++ {
		body := fmt.Sprintf("{customerEmail: %s,content: %s,from: %s,to: %s}", subject, content, from, to[i])

		err := MailChannel.Publish(
			"",             // exchange
			MailQueue.Name, // routing key
			false,          // mandatory
			false,          // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(body),
			})

		log.Printf(" [x] Sent %s", body)
		failOnError(err, "Failed to publish a message")
	}
}
