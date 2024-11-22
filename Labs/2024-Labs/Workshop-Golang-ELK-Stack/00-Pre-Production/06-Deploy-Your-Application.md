### **6. Deploy Your Application**

Deploying your Go application and ELK Stack setup to production is a crucial step in ensuring that your application is scalable, secure, and easy to manage. In this section, we will explore how to containerize your Go application using **Docker**, use **Docker Compose** to deploy both your Go application and the **ELK Stack**, and set up a **reverse proxy** with **Nginx** to secure access to Kibana and your APIs. The outcome of this section will be a **production-ready deployment setup** that can be used for scalable and secure application hosting.

---

### **Steps for Deployment**

#### **1. Containerize Your Go Application Using Docker**

Containerizing your Go application ensures that it can be deployed consistently across different environments. Docker allows you to bundle your Go application with all its dependencies into a lightweight container, making it easy to deploy and scale.

##### **1.1 Create a Dockerfile for Your Go Application**

First, create a `Dockerfile` in the root of your Go project. This file defines the steps Docker will use to build and run your Go application.

Here is an example of a `Dockerfile` for a Go application:

```dockerfile
# Step 1: Use an official Go image as a builder
FROM golang:1.20-alpine AS builder

# Step 2: Set the Current Working Directory inside the container
WORKDIR /app

# Step 3: Copy go mod and sum files
COPY go.mod go.sum ./

# Step 4: Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Step 5: Copy the source code into the container
COPY . .

# Step 6: Build the Go app
RUN go build -o main .

# Step 7: Start a new stage from scratch to reduce image size
FROM alpine:latest

# Step 8: Install necessary dependencies (e.g., libc)
RUN apk --no-cache add ca-certificates

# Step 9: Set the Current Working Directory inside the container
WORKDIR /root/

# Step 10: Copy the pre-built binary from the builder stage
COPY --from=builder /app/main .

# Step 11: Expose the application port
EXPOSE 8080

# Step 12: Run the binary when the container starts
CMD ["./main"]
```

##### **1.2 Build and Run the Docker Image**

To build and run your Go application in a Docker container:

1. **Build the Docker image**:

   ```bash
   docker build -t my-go-app .
   ```

2. **Run the Docker container**:

   ```bash
   docker run -d -p 8080:8080 --name my-go-app my-go-app
   ```

This will run your Go application inside a Docker container, accessible on port `8080`.

---

#### **2. Deploy the ELK Stack Using Docker Compose**

Now that your Go application is containerized, you can deploy the **ELK Stack** (Elasticsearch, Logstash, and Kibana) alongside your Go application using **Docker Compose**. Docker Compose allows you to define and manage multi-container Docker applications with ease.

##### **2.1 Create a `docker-compose.yml` File**

Create a `docker-compose.yml` file that defines the services for your Go application and the ELK Stack. Here is an example `docker-compose.yml` for deploying both:

```yaml
version: "3.7"

services:
  # Go application service
  go-app:
    build: .
    container_name: go-app
    ports:
      - "8080:8080"
    networks:
      - elk-net

  # Elasticsearch service
  elasticsearch:
    image: docker.elastic.co/elasticsearch/elasticsearch:8.6.2
    container_name: elasticsearch
    environment:
      - discovery.type=single-node
      - ES_JAVA_OPTS=-Xmx512m -Xms512m
    ports:
      - "9200:9200"
    networks:
      - elk-net

  # Logstash service
  logstash:
    image: docker.elastic.co/logstash/logstash:8.6.2
    container_name: logstash
    volumes:
      - ./logstash.conf:/usr/share/logstash/pipeline/logstash.conf
    networks:
      - elk-net

  # Kibana service
  kibana:
    image: docker.elastic.co/kibana/kibana:8.6.2
    container_name: kibana
    ports:
      - "5601:5601"
    networks:
      - elk-net

networks:
  elk-net:
    driver: bridge
```

##### **2.2 Explanation of the `docker-compose.yml` File**

- **Go Application (`go-app`)**: This service is built from your Go application’s Dockerfile and exposed on port `8080`.
- **Elasticsearch**: Elasticsearch is configured to run as a single-node instance (ideal for development and small-scale environments). It’s exposed on port `9200`.
- **Logstash**: This service is responsible for processing logs from your Go application and forwarding them to Elasticsearch. The configuration file `logstash.conf` needs to be created in the same directory.
- **Kibana**: Kibana is exposed on port `5601` and will provide the dashboard to visualize logs in real-time.

##### **2.3 Start the Docker Compose Services**

Once the `docker-compose.yml` file is ready, you can start all services:

```bash
docker-compose up -d
```

This will build and run your Go application, Elasticsearch, Logstash, and Kibana as Docker containers in the background.

---

#### **3. Set Up a Reverse Proxy with Nginx**

A **reverse proxy** is essential for securely managing access to your Go API and Kibana. Nginx will act as a proxy between external users and your services, providing secure access, load balancing, and SSL/TLS termination.

##### **3.1 Install Nginx**

If you don’t already have Nginx installed on your server, you can install it as follows:

```bash
sudo apt update
sudo apt install nginx
```

##### **3.2 Configure Nginx as a Reverse Proxy**

Create an Nginx configuration file (`/etc/nginx/sites-available/go-app`) to route requests to your Go application and Kibana. Below is an example configuration:

```nginx
server {
    listen 80;

    server_name example.com;

    # Reverse Proxy for Go Application
    location /api/ {
        proxy_pass http://localhost:8080/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }

    # Reverse Proxy for Kibana
    location /kibana/ {
        proxy_pass http://localhost:5601/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}

```

This configuration sets up Nginx to proxy requests to your Go API at `/api/` and Kibana at `/kibana/`.

##### **3.3 Enable the Nginx Configuration**

Create a symbolic link to enable the site:

```bash
sudo ln -s /etc/nginx/sites-available/go-app /etc/nginx/sites-enabled/
```

Then restart Nginx to apply the changes:

```bash
sudo systemctl restart nginx
```

You can now access your Go API at `http://example.com/api/` and Kibana at `http://example.com/kibana/`.

---

### **Outcome: Production-Ready Deployment Setup**

By completing these steps, you will have successfully containerized your Go application, deployed it alongside the ELK Stack using Docker Compose, and secured access to your services with Nginx. This setup is production-ready and can be easily scaled by adding more instances of your Go application or the ELK Stack components, as needed.

You now have a robust, scalable, and secure deployment for both your Go application and the ELK Stack, providing powerful monitoring, visualization, and log aggregation capabilities.
