package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type LogEntry struct {
	Timestamp string                 `json:"timestamp"`
	Level     string                 `json:"level"`
	Service   string                 `json:"service"`
	Module    string                 `json:"module"`
	Operation string                 `json:"operation"`
	OrderID   string                 `json:"order_id"`
	UserID    string                 `json:"user_id"`
	Message   string                 `json:"message"`
	Metadata  map[string]interface{} `json:"metadata"`
	TraceID   string                 `json:"trace_id"`
}

var levels = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
var modules = []string{"OrderService", "PaymentGateway", "InventoryService", "OrderValidator", "NotificationService"}
var operations = []string{"CreateOrder", "ChargeCard", "ValidateOrder", "SendConfirmation", "ReserveStock"}

var messages = map[string]string{
	"DEBUG": "Checking internal state before proceeding.",
	"INFO":  "Operation completed successfully.",
	"WARN":  "Potential issue detected, proceeding with caution.",
	"ERROR": "Operation failed due to unexpected issue.",
	"FATAL": "Critical failure encountered, terminating process.",
}

func randomLogEntry(level string) LogEntry {
	orderID := fmt.Sprintf("ORD-%06d", gofakeit.Number(100000, 999999))
	userID := fmt.Sprintf("USR-%03d", gofakeit.Number(100, 999))
	module := gofakeit.RandomString(modules)
	operation := gofakeit.RandomString(operations)

	return LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     level,
		Service:   "ExampleService",
		Module:    module,
		Operation: operation,
		OrderID:   orderID,
		UserID:    userID,
		Message:   messages[level],
		Metadata: map[string]interface{}{
			"user":         userID,
			"order_total":  gofakeit.Price(100000, 500000),
			"item_count":   gofakeit.Number(1, 5),
			"payment_type": gofakeit.CreditCardType(),
		},
		TraceID: gofakeit.UUID(),
	}
}

func main() {
	gofakeit.Seed(time.Now().UnixNano())

	for {
		level := gofakeit.RandomString(levels)
		entry := randomLogEntry(level)
		jsonLog, _ := json.MarshalIndent(entry, "", "  ")
		fmt.Println(string(jsonLog))
		time.Sleep(3 * time.Second)
	}
}
