package consumer

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)


func main(){
	var ctx = context.Background()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "email-topic",
		GroupID: "email-consumers",
	})

	defer reader.Close()

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	defer redisClient.Close()

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Fatal("Failed to read message:", err)
		}

		emailData := string(msg.Value)
		fmt.Println("Received Email Event:", emailData)

		err = redisClient.RPush(ctx, "email_queue", emailData).Err()
		if err != nil {
			log.Fatal("Failed to push to Redis queue:", err)
		}

		fmt.Println("Email task added to Redis queue!")
	}
}
