package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	defer redisClient.Close()

	for {
		emailData, err := redisClient.LPop(ctx, "email_queue").Result()
		if err == redis.Nil {
			time.Sleep(2 * time.Second)
			continue
		} else if err != nil {
			log.Fatal("Failed to pop from Redis queue:", err)
		}

		// Simulate sending an email
		fmt.Println("Sending email:", emailData)
		time.Sleep(1 * time.Second)
		fmt.Println("Email sent successfully!")
	}
}
