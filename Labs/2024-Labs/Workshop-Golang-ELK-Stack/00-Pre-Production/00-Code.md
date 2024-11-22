## **Complete Workshop Guide: Building Go Applications with ELK Stack**

This workshop walks you through the process of developing a Go application with logging integrated into the ELK Stack (Elasticsearch, Logstash, and Kibana), covering everything from the basics to advanced setups.

### **Folder Structure**

```
/go-elk-workshop
|-- /app
|   |-- main.go
|   |-- user_handler.go
|   |-- logging
|   |   |-- logger.go
|   |-- /models
|   |   |-- user.go
|-- /config
|   |-- docker-compose.yml
|   |-- Dockerfile
|   |-- nginx.conf
|   |-- /scripts
|   |   |-- setup_elk.sh
|   |   |-- deploy.sh
|-- /elk
|   |-- logstash
|   |   |-- logstash.conf
|   |-- /kibana
|   |-- /elasticsearch
|   |   |-- elasticsearch.yml
|-- /scripts
|   |-- init_elk.sh
```

---

### **1. Go Application Code**

#### **main.go**

This file initializes the Go server and sets up the routes.

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "app/user_handler"
)

func main() {
    // Setup routes
    http.HandleFunc("/users", user_handler.GetUsers)

    fmt.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

#### **user_handler.go**

This file handles API requests to fetch user data.

```go
package user_handler

import (
    "encoding/json"
    "net/http"
    "app/models"
    "app/logging"
)

// GetUsers handles requests to fetch user data.
func GetUsers(w http.ResponseWriter, r *http.Request) {
    // Create sample users
    users := []models.User{
        {ID: 1, Name: "John Doe", Email: "john@example.com"},
        {ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
    }

    // Log the user request event
    logging.Info("Fetching user data", "user_count", len(users))

    // Set response headers
    w.Header().Set("Content-Type", "application/json")

    // Send JSON response
    json.NewEncoder(w).Encode(users)
}
```

#### **logging/logger.go**

This file sets up structured logging with `logrus`.

```go
package logging

import (
    "github.com/sirupsen/logrus"
    "os"
)

// Logger instance
var log = logrus.New()

func init() {
    log.SetFormatter(&logrus.JSONFormatter{})
    log.SetOutput(os.Stdout)
    log.SetLevel(logrus.InfoLevel)
}

// Info logs an info level message
func Info(message string, fields ...interface{}) {
    log.WithFields(logrus.Fields{
        "fields": fields,
    }).Info(message)
}

// Error logs an error level message
func Error(message string, fields ...interface{}) {
    log.WithFields(logrus.Fields{
        "fields": fields,
    }).Error(message)
}
```

#### **models/user.go**

This file defines the user model used in the application.

```go
package models

// User represents a user in the system
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

---

### **2. ELK Stack Configuration**

#### **docker-compose.yml**

This file defines all the services in the Docker Compose setup, including Elasticsearch, Logstash, Kibana, and your Go application.

```yaml
version: "3"
services:
  elasticsearch:
    image: docker.elastic.co/elasticsearch/elasticsearch:7.10.0
    environment:
      - discovery.type=single-node
    ports:
      - "9200:9200"

  logstash:
    image: docker.elastic.co/logstash/logstash:7.10.0
    volumes:
      - ./elk/logstash/logstash.conf:/usr/share/logstash/pipeline/logstash.conf
    ports:
      - "5044:5044"

  kibana:
    image: docker.elastic.co/kibana/kibana:7.10.0
    ports:
      - "5601:5601"
    environment:
      - ELASTICSEARCH_URL=http://elasticsearch:9200

  go-api:
    build:
      context: ./app
    ports:
      - "8080:8080"
    depends_on:
      - elasticsearch
      - logstash
```

#### **Dockerfile**

This Dockerfile builds and runs the Go application.

```dockerfile
FROM golang:1.16-alpine

WORKDIR /app

# Copy Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main .

# Run the application
CMD ["./main"]
```

#### **nginx.conf**

This Nginx configuration acts as a reverse proxy to handle API and Kibana traffic.

```nginx
server {
    listen 80;

    server_name go-elk.local;

    location / {
        proxy_pass http://go-api:8080;
    }

    location /kibana {
        proxy_pass http://kibana:5601;
    }
}
```

#### **logstash.conf**

This configuration tells Logstash how to handle incoming logs and where to send them.

```bash
input {
  tcp {
    port => 5044
    codec => json
  }
}

filter {
  # Add filters to parse logs
}

output {
  elasticsearch {
    hosts => ["http://elasticsearch:9200"]
    index => "go-logs-%{+YYYY.MM.dd}"
  }
}
```

#### **elasticsearch.yml**

Configuration for the Elasticsearch service to ensure it runs as a single-node cluster.

```yaml
cluster.name: "docker-cluster"
network.host: 0.0.0.0
discovery.type: single-node
```

#### **setup_elk.sh**

This shell script starts up the ELK Stack using Docker Compose.

```bash
#!/bin/bash

# Install and configure ELK stack
echo "Setting up ELK Stack..."
docker-compose up -d --build
```

#### **deploy.sh**

This script deploys both the Go application and the ELK Stack services.

```bash
#!/bin/bash

# Deploy Go application and ELK Stack
echo "Deploying Go application and ELK Stack..."

# Build and deploy containers
docker-compose up -d --build

# Wait for services to initialize
sleep 10

# Verify if Elasticsearch and Kibana are running
curl http://localhost:9200
curl http://localhost:5601
```

#### **init_elk.sh**

This script prepares the environment for the ELK Stack.

```bash
#!/bin/bash
# Initialize the ELK Stack

docker-compose up -d
```

---

### **3. How to Run**

1. **Set up the ELK Stack**:

   - Run `setup_elk.sh` to start the ELK Stack services using Docker Compose.
   - Ensure that Elasticsearch, Logstash, and Kibana are running by checking `localhost:9200`, `localhost:5044`, and `localhost:5601`.

2. **Run the Go Application**:

   - Execute `docker-compose up -d` to start the Go application along with the ELK Stack.
   - The Go application will be accessible at `localhost:8080`.

3. **Test the API**:

   - Make a request to `localhost:8080/users` to test the Go application.
   - View the logs in Kibana at `localhost:5601` to ensure that logs are being sent correctly to Elasticsearch.

4. **Kibana Dashboard**:
   - Use Kibana to create a dashboard and visualize the logs sent from your Go application to Elasticsearch via Logstash.

---

### **4. Conclusion**

By following this guide, you will have set up a complete Go application with logging integrated into the ELK Stack. You will have learned how to deploy the application, structure logs properly, and visualize logs in Kibana. With this setup, you can scale the system further by adding features like advanced filtering, alerting, and more.
