package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/IBM/sarama"
)

type Message struct {
	UserId     int    `json:"user_id"`
	PostId     string `json:"post_id"`
	UserAction string `json:"user_action"`
}

func main() {
	brokers := []string{"localhost:9093", "localhost:9094", "localhost:9095"}
	producer, err := sarama.NewSyncProducer(brokers, nil)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
		os.Exit(1)
	}
	defer producer.Close()

	postId := [5]string{"POST00001", "POST00002", "POST00003", "POST00004", "POST00005"}
	userId := [5]int{100001, 100002, 100003, 100004, 100005}
	userAction := [5]string{"love", "like", "hate", "smile", "cry"}

	for {
		message := Message{
			UserId:     userId[rand.Intn(len(userId))],
			PostId:     postId[rand.Intn(len(postId))],
			UserAction: userAction[rand.Intn(len(userAction))],
		}

		jsonMessage, err := json.Marshal(message)
		if err != nil {
			log.Fatalln("Failed to marshal message:", err)
			os.Exit(1)
		}

		msg := &sarama.ProducerMessage{
			Topic: "post-likes",
			Key:   sarama.StringEncoder(message.PostId),
			Value: sarama.StringEncoder(jsonMessage),
		}

		_, _, err = producer.SendMessage(msg)
		if err != nil {
			log.Fatalln("Failed to send message:", err)
			os.Exit(1)
		}

		log.Println("Message sent:", string(jsonMessage))

		// wait 5s
		time.Sleep(5 * time.Second)
	}
}
