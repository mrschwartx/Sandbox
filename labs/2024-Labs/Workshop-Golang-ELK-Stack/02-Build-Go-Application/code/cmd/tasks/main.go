package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/tasks"
)

func main() {
	repo := tasks.NewTaskRepository()
	repo.InitData(context.Background())

	handler := tasks.NewTaskHandler(repo)

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
		log.Println("Starting server on :8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error occurred: %v", err)
		}
	}()

	// Wait for termination signal (like Ctrl+C)
	<-stopChan
	log.Println("Shutting down server gracefully...")

	// Create a context with a timeout for the shutdown process (e.g., 5 seconds)
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully shut down the server
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server stopped")
}
