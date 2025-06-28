package main

import (
	"log"
	"math/rand"
	"os"
	"time"
)

var messages = []string{
	"User logged in",
	"User logged out",
	"File uploaded successfully",
	"Failed to connect to database",
	"Timeout while connecting to service",
	"Data processed successfully",
	"Invalid user input",
	"Job queued",
	"Email sent to user",
	"Unexpected error occurred",
}

func main() {
	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create logs directory: %v", err)
	}

	file, err := os.OpenFile("./logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	for {
		msg := messages[rand.Intn(len(messages))]
		log.Println(msg)
		time.Sleep(2 * time.Second)
	}
}
