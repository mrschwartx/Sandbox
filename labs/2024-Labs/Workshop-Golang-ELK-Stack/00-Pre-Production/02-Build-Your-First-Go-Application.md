### **2. Build Your First Go Application**

Building your first Go application is an exciting and essential step toward mastering the language. In this section, we will guide you through creating a simple **RESTful API** using Go. This API will simulate a **library** or **task management system**, where users can manage and interact with books or tasks. This project will help you get hands-on experience with core Go concepts such as API development, JSON encoding/decoding, and testing.

---

#### **Project Idea: A Simple RESTful API for a Library or Task Management System**

A RESTful API (Representational State Transfer) is an architectural style for designing networked applications. RESTful APIs typically use HTTP methods such as `GET`, `POST`, `PUT`, and `DELETE` to allow clients to interact with the server. In this case, we will create a RESTful API for managing tasks or books.

For simplicity, we will create endpoints to:

- **GET** all tasks or books.
- **POST** a new task or book.
- **DELETE** a task or book by ID.
- **PUT** to update a task or book's information.

---

#### **Skills You'll Gain**

- **Using the `net/http` Package**: The `net/http` package is part of Go's standard library and is the primary tool for building web servers and making HTTP requests. You'll learn to use it to define HTTP routes and handlers.
- **JSON Encoding/Decoding**: Most APIs today exchange data in the JSON format. Go provides the `encoding/json` package, which allows you to easily encode and decode JSON data to and from Go data types.

- **Testing with the `testing` Package**: Writing tests is crucial in Go development. Go's `testing` package offers built-in support for unit testing, making it easy to write tests for your API endpoints and business logic.

- **Tools**:
  - **Postman** or **cURL**: These tools are used for testing and interacting with APIs. Postman provides a user-friendly interface for making HTTP requests, while cURL allows you to interact with APIs via the command line.

---

#### **Step 1: Define the Project Structure**

Before we dive into coding, let's define the structure of the project. In a simple Go project, you can organize your application like this:

```
/go-task-app
  /main.go
  /models
    task.go
  /handlers
    task_handlers.go
  /tests
    task_handlers_test.go
```

- **main.go**: Contains the main entry point of the application, sets up the HTTP server, and routes.
- **models/task.go**: Defines the structure and behavior of the task (or book) model.
- **handlers/task_handlers.go**: Contains the logic to handle incoming HTTP requests.
- **tests/task_handlers_test.go**: Contains tests for your API endpoints.

---

#### **Step 2: Create the Task Model**

In Go, we use structs to define models. The `Task` struct will represent a task in our application.

```go
// models/task.go
package models

// Task defines the structure for a task in the task management system.
type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}
```

This struct will contain three fields:

- `ID`: A unique identifier for the task.
- `Title`: The name or title of the task.
- `Status`: The current status of the task (e.g., "in-progress", "completed").

The `json` tags indicate how Go should serialize and deserialize the struct fields when encoding or decoding JSON data.

---

#### **Step 3: Implement the Handlers**

Now, let’s create the handlers that will define how the server will respond to HTTP requests.

```go
// handlers/task_handlers.go
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"your_project/models"
)

// A global variable to store tasks in memory.
var tasks []models.Task

// GetTasks handles GET requests to fetch all tasks.
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// CreateTask handles POST requests to create a new task.
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = len(tasks) + 1
	tasks = append(tasks, task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// DeleteTask handles DELETE requests to remove a task by ID.
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, _ := strconv.Atoi(params["id"])

	for index, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusNoContent) // No content is returned upon successful deletion
}
```

- **GetTasks**: This handler will respond to `GET` requests and return all tasks in JSON format.
- **CreateTask**: This handler handles `POST` requests, decoding the JSON body into a `Task` struct, and adding it to the tasks slice.
- **DeleteTask**: This handler deletes a task based on the `ID` passed in the URL.

---

#### **Step 4: Set Up the HTTP Server**

In the `main.go` file, we will set up the HTTP server and configure the routes.

```go
// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"your_project/handlers"
)

func main() {
	// Initialize router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	// Start the server
	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
```

We use the **Gorilla Mux router** (`mux`) to handle routing. This allows us to map HTTP requests to the appropriate handler functions.

---

#### **Step 5: Testing Your API**

Now that the application is set up, you can start testing the endpoints using **Postman** or **cURL**.

- **Testing with Postman**:

  1. **GET** request to `http://localhost:8080/tasks` to retrieve all tasks.
  2. **POST** request to `http://localhost:8080/tasks` with a JSON body to create a new task. Example body:
     ```json
     {
       "title": "Learn Go",
       "status": "in-progress"
     }
     ```
  3. **DELETE** request to `http://localhost:8080/tasks/1` to delete the task with ID `1`.

- **Testing with cURL**:

  ```bash
  # GET all tasks
  curl http://localhost:8080/tasks

  # POST a new task
  curl -X POST -H "Content-Type: application/json" -d '{"title": "Learn Go", "status": "in-progress"}' http://localhost:8080/tasks

  # DELETE a task
  curl -X DELETE http://localhost:8080/tasks/1
  ```

---

#### **Step 6: Write Tests for the Handlers**

In Go, the `testing` package allows you to write unit tests for your application. Here's a simple test case to check the `GetTasks` handler.

```go
// tests/task_handlers_test.go
package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"your_project/handlers"
	"your_project/models"
)

func TestGetTasks(t *testing.T) {
	// Setup
	handlers.Tasks = append(handlers.Tasks, models.Task{ID: 1, Title: "Learn Go", Status: "in-progress"})

	// Create a request
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(handlers.GetTasks)
	handler.ServeHTTP(rr, req)

	// Check if the status code is correct
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check if the response body contains the task
	expected := `[{"id":1,"title":"Learn Go","status":"in-progress"}]`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
```

This test case verifies that the `GetTasks` handler returns the correct status and body when a request is made to `/tasks`.

---

### **Conclusion**

By building your first Go application—a simple RESTful API—you have gained hands-on experience with the essential Go concepts of API development, JSON encoding/decoding, and testing. You’ve also learned how to handle HTTP requests and organize code effectively in Go. Testing your API endpoints ensures that your application behaves as expected and lays a solid foundation for more complex Go applications in the future.
