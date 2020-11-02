package main

import (
	"github.com/streadway/amqp"
	"fmt"
	"encoding/json"
	"GoAssignment/database"
	"GoAssignment/utils"
	"log"
)

func main(){

	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	utils.GenerateLogs(err, "Error in making connection to rabbit")
	defer conn.Close()

	channel, err := conn.Channel()
	utils.GenerateLogs(err, "Error in opening channel on rabbit")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"HOTEL_INFO",
		false,
		false,
		false,
		false,
		nil,
	)
	utils.GenerateLogs(err, "Error in declaring queue on rabbit")

	msg, err := channel.Consume(
		queue.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.GenerateLogs(err, "Error in consuming messages from rabbit")

	forever := make(chan bool)
	func() {
		for {
			select {
			case d := <-msg:
				dbAction(d)
			default:
				utils.GenerateLogs(err, "Error In Rabbit Connection")
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func dbAction(d amqp.Delivery){
	var err error

	body := fmt.Sprintf("%s", d.Body)

	var message map[string][]map[string]interface{}
	err = json.Unmarshal([]byte(body), &message)
	if err != nil {
		utils.GenerateLogs(err, "Error in unmarshalling message packet")
	}

	_, err = database.PerformDBAction(message)
	if err != nil {
		utils.GenerateLogs(err, "Error in performing DB actions")
	}
	
}