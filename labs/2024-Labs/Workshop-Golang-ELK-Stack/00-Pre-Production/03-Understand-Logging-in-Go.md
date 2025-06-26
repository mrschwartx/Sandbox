### **3. Understand Logging in Go**

Logging is an essential part of any application, as it provides visibility into the system's behavior, helps identify issues, and offers insights into application performance. In Go, logging is a straightforward yet powerful feature, and understanding how to log events and errors effectively is crucial for building robust applications. This section will guide you through logging in Go, including using the standard `log` package, exploring third-party libraries for structured logging, and preparing logs in JSON format for integration with the ELK stack.

---

#### **Goal: Learn How to Log Events and Errors in Go Applications**

The primary goal of logging in Go is to keep track of the application's execution flow, record any issues or errors that occur, and provide useful information to help with debugging and monitoring. Logs should be both human-readable and machine-readable, especially when you integrate your application with centralized logging systems like the ELK stack (Elasticsearch, Logstash, Kibana).

In this section, we will:

1. Cover Go’s built-in logging facilities.
2. Introduce third-party libraries like `logrus` and `zap` for advanced logging features.
3. Demonstrate how to structure logs in JSON format, which is ideal for integration with ELK.

---

#### **Step 1: Basic Logging with the Standard `log` Package**

Go provides a simple logging package called `log` that allows you to output logs with different severity levels such as `info`, `warning`, and `error`. By default, the `log` package outputs logs to the standard output (stdout) with timestamps, which is useful for basic logging needs.

Here’s how to use the `log` package for basic logging:

```go
package main

import (
	"log"
)

func main() {
	// Simple logging with log package
	log.Println("This is an info message.")
	log.Fatal("This is a fatal error.") // Terminates the program after logging
}
```

**Key Functions**:

- `log.Println`: Outputs an informational message with a newline.
- `log.Fatal`: Logs a message and then calls `os.Exit(1)`, terminating the application.
- `log.Panic`: Logs a message and then panics, halting the execution.

While the standard `log` package is sufficient for small applications, it lacks advanced features such as log levels, structured logging, and flexibility in log output formats (e.g., JSON or key-value pairs). For larger projects or production applications, third-party logging libraries are recommended.

---

#### **Step 2: Advanced Logging with `logrus`**

[Logrus](https://github.com/sirupsen/logrus) is one of the most popular third-party logging libraries in Go. It is more feature-rich than the standard `log` package, offering support for log levels (e.g., `Info`, `Warn`, `Error`), structured logging, and various output formats like JSON.

**Installation**:
To use Logrus in your Go project, you need to install it via `go get`:

```bash
go get github.com/sirupsen/logrus
```

**Example Usage**:

```go
package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	// Create a new logrus logger
	log := logrus.New()

	// Set log level (default is InfoLevel)
	log.SetLevel(logrus.DebugLevel)

	// Log a message at different levels
	log.Info("This is an info message.")
	log.Warn("This is a warning message.")
	log.Error("This is an error message.")
	log.Debug("This is a debug message.")

	// Structured logging with fields
	log.WithFields(logrus.Fields{
		"user": "john_doe",
		"status": "active",
	}).Info("User logged in.")
}
```

**Features of Logrus**:

- **Log Levels**: You can set different log levels like `Debug`, `Info`, `Warn`, `Error`, `Fatal`, and `Panic`.
- **Structured Logging**: Log entries can include custom fields, making it easier to capture and track contextual information.
- **Log Formatting**: You can output logs in various formats (text, JSON, or even log entries formatted as key-value pairs).

For instance, logging in JSON format with Logrus:

```go
log.SetFormatter(&logrus.JSONFormatter{})
log.Info("User logged in")
```

This would produce log output like:

```json
{ "level": "info", "msg": "User logged in", "time": "2024-11-22T13:25:03Z" }
```

This format is well-suited for ingestion into log management systems like ELK, which expect logs in JSON format.

---

#### **Step 3: Structured Logging with `zap`**

[Zap](https://github.com/uber-go/zap) is another popular logging library in Go, developed by Uber. Zap is designed for high-performance logging, offering structured logging with a minimalistic API. It focuses on speed and memory efficiency, making it a great choice for high-volume logging scenarios.

**Installation**:
To use Zap in your Go project, install it via `go get`:

```bash
go get go.uber.org/zap
```

**Example Usage**:

```go
package main

import (
	"go.uber.org/zap"
)

func main() {
	// Create a new zap logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // Flushes buffer, if any

	// Simple logging
	logger.Info("This is an info message.")
	logger.Warn("This is a warning message.")
	logger.Error("This is an error message.")

	// Structured logging
	logger.Info("User logged in", zap.String("user", "john_doe"), zap.String("status", "active"))
}
```

**Features of Zap**:

- **High Performance**: Zap is optimized for low-latency, high-throughput logging and is much faster than other libraries, such as Logrus.
- **Structured Logging**: Just like Logrus, Zap allows you to add key-value pairs to logs. This makes it easier to extract useful information when analyzing logs.
- **Log Levels**: Zap supports `Info`, `Warn`, `Error`, `Fatal`, and `Debug` log levels.
- **Log Output Formats**: You can log in different formats (JSON, text, etc.), and the default production configuration outputs logs in JSON format, ideal for integration with systems like ELK.

Here’s an example of structured logging in JSON format with Zap:

```go
logger, _ := zap.NewProduction()
defer logger.Sync() // Ensure logs are written before the program exits

logger.Info("Processed request", zap.Int("status_code", 200), zap.String("endpoint", "/tasks"))
```

This would output:

```json
{
  "level": "info",
  "msg": "Processed request",
  "status_code": 200,
  "endpoint": "/tasks",
  "time": "2024-11-22T13:25:03Z"
}
```

This structured approach is perfect for aggregating logs in systems like Elasticsearch, where you can easily query specific fields (like `status_code` or `endpoint`) to filter and analyze logs.

---

#### **Step 4: Writing Logs in JSON Format for ELK Integration**

To integrate your Go application with the ELK stack (Elasticsearch, Logstash, and Kibana), you must ensure that your logs are in a machine-readable format like JSON. Both Logrus and Zap support logging in JSON format, which is ideal for this purpose.

Here’s a sample of how you might configure your application to output logs in JSON format:

1. **Using Logrus**:

```go
log.SetFormatter(&logrus.JSONFormatter{})
log.Info("Task completed", "task_id", 123)
```

2. **Using Zap**:

```go
logger, _ := zap.NewProduction()
defer logger.Sync()
logger.Info("Task completed", zap.Int("task_id", 123))
```

Once your logs are structured in JSON format, you can easily ship them to Logstash for processing and index them into Elasticsearch. From there, Kibana can be used to visualize the logs, enabling real-time monitoring and analysis.

---

### **Conclusion**

Understanding logging in Go is an essential skill for building maintainable and observable applications. While the standard `log` package is useful for basic logging, third-party libraries like Logrus and Zap provide enhanced features such as structured logging, log levels, and JSON output—features that are invaluable for debugging and integration with centralized logging systems like ELK.

As your application scales and requires more sophisticated logging capabilities, adopting libraries like Logrus or Zap will give you greater flexibility and performance, helping you monitor and analyze your application more effectively. By preparing logs in JSON format, you are ready to integrate with tools like Logstash, Elasticsearch, and Kibana to ensure your application is observable, efficient, and scalable.
