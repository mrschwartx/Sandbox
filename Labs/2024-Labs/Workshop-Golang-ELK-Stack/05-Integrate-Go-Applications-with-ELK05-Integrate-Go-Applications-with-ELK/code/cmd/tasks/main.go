package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/tasks"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	appLevel := "development"

	var logger *zap.Logger
	var cleanup func()
	var err error

	if appLevel == "production" {
		logger, cleanup, err = setupMultiFileLogger() // Use the appropriate logger for production
	} else {
		logger, cleanup, err = setupSingleFileLogger() // Same setup for development (or change if needed)
	}

	if err != nil {
		log.Fatalf("failed to set up logger: %v", err)
	}
	defer cleanup()

	repo := tasks.NewTaskRepository(logger)
	repo.InitData(context.Background())

	handler := tasks.NewTaskHandler(repo, logger)

	// Access: GET http://localhost:8080/tasks
	http.HandleFunc("/tasks", handler.GetAllTasksHandler) // GET all tasks

	// Access: POST http://localhost:8080/tasks/create
	// Expected body: {"title": "New Task", "status": "pending"}
	http.HandleFunc("/tasks/create", handler.CreateTaskHandler) // POST create task

	// Access: GET http://localhost:8080/tasks/get?id=1
	// Replace `1` with the ID of the task you want to retrieve.
	http.HandleFunc("/tasks/get", handler.GetTaskByIDHandler) // GET task by ID

	// Access: PUT http://localhost:8080/tasks/update?id=1
	// Expected body: {"title": "Updated Task", "status": "completed"}
	// Replace `1` with the ID of the task you want to update.
	http.HandleFunc("/tasks/update", handler.UpdateTaskHandler) // PUT update task

	// Access: DELETE http://localhost:8080/tasks/delete?id=1
	// Replace `1` with the ID of the task you want to delete.
	http.HandleFunc("/tasks/delete", handler.DeleteTaskHandler) // DELETE task by ID

	// Create a new HTTP server
	server := &http.Server{
		Addr:    ":8080", // Set the port to listen on
		Handler: nil,     // Default handler (we've already set up routes)
	}

	// Set up a channel to listen for OS signals (like SIGINT, SIGTERM)
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a goroutine to avoid blocking the main thread
	go func() {
		logger.Info("Starting server on :8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Server error", zap.Error(err))
		}
	}()

	// Wait for termination signal (like Ctrl+C)
	<-stopChan
	logger.Info("Shutting down server gracefully...")

	// Create a context with a timeout for the shutdown process (e.g., 5 seconds)
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully shut down the server
	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Fatal("Server shutdown failed", zap.Error(err))
	}

	logger.Info("Server stopped")
}

func setupMultiFileLogger() (*zap.Logger, func(), error) {
	logDir := "./logs"
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	logFile := fmt.Sprintf("%s/app-%s.log", logDir, timestamp)

	// Ensure the log directory exists
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, nil, fmt.Errorf("failed to create log directory: %v", err)
	}

	// Configure lumberjack for log rotation
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFile, // Log file with timestamp
		MaxSize:    10,      // Max size in MB
		MaxBackups: 7,       // Max number of backup files
		MaxAge:     1,       // Max age in days
		Compress:   false,   // Disable compression
	})
	// Console logging configuration
	consoleWriter := zapcore.Lock(os.Stdout)

	// Encoder configuration (shared for both console and file)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"                   // Add timestamp to logs
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Use ISO8601 time format
	encoderConfig.LevelKey = "level"                      // Log level key
	encoderConfig.MessageKey = "message"                  // Log message key

	// Create separate encoders for file and console
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)       // JSON for files
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig) // Pretty print for console

	// Define log levels for both
	fileCore := zapcore.NewCore(fileEncoder, fileWriter, zapcore.DebugLevel)
	consoleCore := zapcore.NewCore(consoleEncoder, consoleWriter, zapcore.DebugLevel)

	// Combine both cores
	combinedCore := zapcore.NewTee(fileCore, consoleCore)

	// Create a logger
	logger := zap.New(combinedCore)

	// Return the logger and a cleanup function
	cleanup := func() {
		logger.Sync() // Flush buffered logs
	}

	return logger, cleanup, nil
}

func setupSingleFileLogger() (*zap.Logger, func(), error) {
	logFile := "./logs/app.log" // Single log file path

	// Ensure the log directory exists
	if err := os.MkdirAll("./logs", 0755); err != nil {
		return nil, nil, fmt.Errorf("failed to create log directory: %v", err)
	}

	// Configure lumberjack for file logging
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFile, // Always log to the same file
		MaxSize:    10,      // Max size in MB before rotation
		MaxBackups: 7,       // Max number of backup files to keep
		MaxAge:     7,       // Max age of backup files in days
		Compress:   false,   // Disable compression
	})

	// Encoder configuration
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"                   // Add timestamp to logs
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Use ISO8601 time format
	encoderConfig.LevelKey = "level"                      // Log level key
	encoderConfig.MessageKey = "message"                  // Log message key

	// Create file encoder and core
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	fileCore := zapcore.NewCore(fileEncoder, fileWriter, zapcore.DebugLevel)

	// Create a logger
	logger := zap.New(fileCore)

	// Return the logger and a cleanup function
	cleanup := func() {
		logger.Sync() // Flush buffered logs
	}

	return logger, cleanup, nil
}
