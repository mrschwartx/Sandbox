# **10. Build a Logging Framework for Go**

In modern software development, effective logging is crucial for monitoring application health, tracking issues, and debugging. A custom logging framework tailored to the specific needs of your organization can provide the flexibility and control necessary for managing logs in a production environment. In this section, we will guide you through building a reusable logging framework for your Go applications. This framework will support multiple output destinations (e.g., files, Logstash), contextual logging, dynamic log levels, and advanced features like log rotation.

---

## **Key Objectives**

1. **Support for Multiple Outputs**: The logging framework will allow you to log to multiple destinations such as local files, Logstash (for centralized logging), or external systems (e.g., a monitoring service).
2. **Contextual Logging**: By supporting contextual information such as user IDs, request IDs, and other relevant details, your logs will provide deeper insights into specific application flow, making it easier to trace issues and analyze system behavior.

3. **Dynamic Log Levels**: You’ll be able to adjust log verbosity on the fly based on the environment (e.g., `DEBUG` for development, `ERROR` for production), providing fine-grained control over logging during debugging and live operation.

4. **Log Rotation and Retention**: We will also implement log rotation to ensure that log files don’t grow indefinitely, which can lead to storage issues.

---

### **1. Creating the Basic Logging Package**

Start by creating a basic logging framework that supports essential logging functionality. The framework will allow you to output logs to different destinations (e.g., file, console) and adjust log levels dynamically.

#### **1.1 Creating a Logger Struct**

The first step is to define the basic structure for the logger.

1. **Define the `Logger` struct**:

   ```go
   package logger

   import (
       "fmt"
       "log"
       "os"
   )

   // Define log levels
   const (
       DEBUG = iota
       INFO
       WARN
       ERROR
   )

   var currentLevel = INFO

   // Logger struct to hold log configuration
   type Logger struct {
       level   int
       out     *log.Logger
       context map[string]interface{}
   }

   // NewLogger creates a new logger instance with a specified output destination
   func NewLogger(level int, output *os.File) *Logger {
       return &Logger{
           level:   level,
           out:     log.New(output, "", log.Ldate|log.Ltime|log.Lshortfile),
           context: make(map[string]interface{}),
       }
   }

   // SetLevel dynamically adjusts the log level
   func SetLevel(level int) {
       currentLevel = level
   }

   // Log method to log messages at the appropriate level
   func (l *Logger) Log(level int, message string) {
       if level >= currentLevel {
           switch level {
           case DEBUG:
               l.out.SetPrefix("[DEBUG] ")
           case INFO:
               l.out.SetPrefix("[INFO] ")
           case WARN:
               l.out.SetPrefix("[WARN] ")
           case ERROR:
               l.out.SetPrefix("[ERROR] ")
           }
           logMessage := fmt.Sprintf("%s%s", message, formatContext(l.context))
           l.out.Println(logMessage)
       }
   }

   func formatContext(context map[string]interface{}) string {
       contextStr := ""
       for key, value := range context {
           contextStr += fmt.Sprintf("%s=%v ", key, value)
       }
       return contextStr
   }
   ```

#### **Explanation of Key Features**:

1. **Logger Struct**: The `Logger` struct holds the log level, output destination, and context data.
2. **Log Levels**: We define four log levels (`DEBUG`, `INFO`, `WARN`, `ERROR`) to control the verbosity of logs.
3. **SetLevel Function**: This function allows dynamic adjustment of the logging level (e.g., switch between `DEBUG` in development and `ERROR` in production).
4. **Logging Method**: The `Log` method is used to print logs with appropriate prefixes based on the log level. It also appends contextual information, such as user IDs or request IDs, if available.

---

### **2. Supporting Multiple Output Destinations**

One of the key features of a good logging framework is the ability to log to multiple destinations. This can be useful for tracking logs locally, sending them to an external system like Logstash, or even pushing them to a cloud-based monitoring service.

#### **2.1 Logging to a File**

You can modify the logger to support file-based logging.

```go
file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
    log.Fatal(err)
}
defer file.Close()

logger := logger.NewLogger(logger.INFO, file)
logger.Info("Application started")
```

This configuration will write log entries to `app.log`. You can also configure log rotation, ensuring that logs don’t fill up the storage.

#### **2.2 Sending Logs to Logstash**

Centralized logging systems like **Logstash** help aggregate logs for easier monitoring. You can send logs to Logstash using a TCP or UDP connection.

```go
import (
    "net"
    "fmt"
    "log"
)

func sendToLogstash(logMessage string) {
    conn, err := net.Dial("tcp", "logstash.example.com:5044")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    fmt.Fprintf(conn, logMessage)
}

logger := logger.NewLogger(logger.INFO, os.Stdout)
logger.Info("Log message sent to Logstash")
sendToLogstash("Log message sent to Logstash")
```

This setup ensures that logs are sent to Logstash, which can then forward them to Elasticsearch for storage and analysis.

---

### **3. Contextual Logging for Tracing Requests**

Contextual logging is essential for tracing requests across distributed systems. You can add context to your logs, such as user IDs, session IDs, or any other identifiers that will help you track the flow of requests.

#### **3.1 Adding Context**

To add context, you can store additional data in the logger instance. For example, logging a user’s actions can include their user ID and session ID.

```go
func (l *Logger) AddContext(key string, value interface{}) {
    l.context[key] = value
}

logger := logger.NewLogger(logger.INFO, os.Stdout)
logger.AddContext("user_id", "12345")
logger.AddContext("request_id", "abcde")
logger.Info("User login successful")
```

This will log the message with the additional context, making it easy to trace specific user activity.

---

### **4. Implementing Dynamic Log Levels for Different Environments**

In production, you typically only want to log errors or warnings. In development, you may need more detailed logs to debug the application. Implementing dynamic log levels allows you to adjust the verbosity of logs depending on the environment.

#### **4.1 Environment-based Log Level Configuration**

You can configure your logging framework to use different log levels depending on the environment:

```go
// Set log level based on environment
if os.Getenv("ENV") == "production" {
    logger.SetLevel(logger.ERROR)
} else {
    logger.SetLevel(logger.DEBUG)
}
```

This ensures that in a development environment, you get detailed logs (e.g., `DEBUG` level), while in production, only critical issues (e.g., `ERROR`) are logged.

---

### **5. Implementing Log Rotation**

For applications that generate a significant amount of log data, log rotation is essential to prevent log files from growing indefinitely. You can use a third-party Go package like `lumberjack` to implement log rotation:

```go
import (
    "gopkg.in/natefinch/lumberjack.v2"
)

func rotateLogs() {
    logger := logger.NewLogger(logger.INFO, &lumberjack.Logger{
        Filename:   "app.log",
        MaxSize:    10, // megabytes
        MaxBackups: 3,
        MaxAge:     28, // days
    })
    logger.Info("Log file rotation implemented")
}
```

This ensures that logs are rotated once they reach a certain size (in this case, 10MB), and old log files are automatically deleted after 28 days.

---

### **Mini-Project: Build a Complete Logging Framework for Your Go Application**

#### **Goal**:

Design a reusable, robust logging framework that will integrate with your Go applications and meet your organizational needs.

#### **Steps**:

1. **Create a Basic Logger**:  
   Develop the core logger with support for multiple log levels and outputs (e.g., file, console, Logstash).

2. **Implement Contextual Logging**:  
   Add support for tracing user actions and request flow with contextual information.

3. **Implement Dynamic Log Levels**:  
   Allow environment-based dynamic adjustment of log verbosity (e.g., DEBUG in development, ERROR in production).

4. **Log Rotation**:  
   Implement log rotation to manage disk space usage and avoid log file bloating.

5. **Integration with External Systems**:  
   Send logs to Logstash or other centralized logging systems to aggregate and monitor application logs.

6. **Deploy and Test**:  
   Test the logging framework in various environments (development, staging, production) to ensure it functions as expected and meets performance requirements.

---

### **Conclusion**

A custom logging framework tailored for your Go applications provides flexibility and control over how logs are generated, stored, and processed. By supporting multiple outputs, adding contextual information, and allowing dynamic log levels, you can improve observability and simplify troubleshooting across your systems. Implementing features like log rotation ensures that

logs are managed efficiently in production environments, while centralized logging solutions like Logstash allow for real-time monitoring and alerting. By building this framework, you create a powerful tool that enhances your Go application’s scalability, maintainability, and resilience.
