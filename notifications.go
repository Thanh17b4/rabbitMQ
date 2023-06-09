package main

//
//import (
//	"encoding/json"
//	"github.com/streadway/amqp"
//	"log"
//)
//
//func rabbitMQ(data interface{}) error {
//	// Kết nối tới rabbitMQ
//	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
//	if err != nil {
//		log.Println("Failed to connect to rabbitMQ:", err.Error())
//		return err
//	}
//	defer conn.Close()
//
//	// create chanel
//	ch, err := conn.Channel()
//	if err != nil {
//		log.Println("Failed to open a channel:", err)
//		return err
//	}
//	defer ch.Close()
//
//	// create queue
//	q, err := ch.QueueDeclare(
//		"create_user_queue",
//		false,
//		false,
//		false,
//		false,
//		nil,
//	)
//	if err != nil {
//		log.Println("Failed to declare a queue:", err)
//		return err
//	}
//
//	// create massage
//	message, err := json.Marshal(data)
//	if err != nil {
//		log.Println("Failed to marshal user to JSON:", err)
//		return err
//	}
//
//	err = ch.Publish(
//		"topic", // Exchange
//		q.Name,  // Routing key
//		false,   // Mandatory
//		false,   // Immediate
//		amqp.Publishing{
//			ContentType: "application/json",
//			Body:        message,
//		},
//	)
//	if err != nil {
//		log.Println("Failed to publish a message:", err)
//		return err
//	}
//
//	return nil
//}
