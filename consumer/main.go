package consumer

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/HETIC-MT-P2021/Go_Mail_Groupe05/consumer/mailing"
	"github.com/joho/godotenv"
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
	defer channel.Close()

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

func main() {
	println("BOI")
	env, _ := godotenv.Read(".env")

	ConnectToRabbit(env["RABBIT_HOST"],
		env["RABBIT_PORT"],
		env["RABBIT_USER"],
		env["RABBIT_PASSWORD"])

	smtpPort, _ := strconv.Atoi(env["SMTP_PORT"])

	mailing.InitSMTPCon(env["SMTP_USER"], env["SMTP_PASSWORD"], env["SMTP_HOST"], smtpPort)

	Receive()
}
