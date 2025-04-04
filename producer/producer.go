package producer

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main(){
	producer := &kafka.Writer{
		Addr: kafka.TCP("localhost:9092"),
		Topic: "email-topic",
		AllowAutoTopicCreation: true,
	}
	defer producer.Close()

	msg := kafka.Message{
		Key:  []byte("email-key"),
		Value: []byte(`{"email": "user@example.com", "subject": "Newsletter Issue #1"}`),
	}

	err := producer.WriteMessages(context.Background(), msg)

	if err != nil {
		panic(err)
	}

	fmt.Println("Email event sent to Kafka!")

}