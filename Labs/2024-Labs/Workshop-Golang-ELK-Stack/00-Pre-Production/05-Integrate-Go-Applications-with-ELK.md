### **5. Integrate Go Applications with ELK**

Integrating your Go application with the **ELK Stack** (Elasticsearch, Logstash, Kibana) provides a powerful framework for log aggregation, analysis, and visualization. By sending logs from your Go application to Elasticsearch through Logstash, and visualizing them in Kibana, you can gain deep insights into your application's performance and behavior. In this section, we’ll walk through the process of integrating Go applications with the ELK Stack. We’ll focus on using structured logging libraries like `logrus` or `zap`, configuring Logstash to forward logs to Elasticsearch, and creating Kibana dashboards to visualize logs. Additionally, we’ll guide you through a mini-project to build a real-time log viewer in Kibana for monitoring user activity in your API.

---

#### **Skills You'll Learn**

1. **Sending Logs from Go to Logstash**: Learn how to use structured logging in Go applications to send logs to Logstash.
2. **Configuring Logstash to Forward Logs to Elasticsearch**: Configure Logstash to process logs from your Go application and forward them to Elasticsearch for indexing.
3. **Visualizing Logs in Kibana**: Create basic Kibana dashboards to visualize logs and monitor user activity in real-time.

---

### **Step 1: Sending Logs from a Go Application to Logstash**

The first step in integrating ELK with your Go application is ensuring that your Go app generates structured logs. Using logging libraries like `logrus` or `zap` will allow you to generate logs in structured formats (e.g., JSON), making it easier to process and forward them to Elasticsearch.

#### **1.1 Install and Set Up Logrus or Zap for Structured Logging**

Here’s how you can set up structured logging in your Go application using **Logrus** or **Zap**:

**Using Logrus**:

1. **Install Logrus**:

   First, install Logrus by running:

   ```bash
   go get github.com/sirupsen/logrus
   ```

2. **Basic Logrus Configuration**:

   ```go
   package main

   import (
       "github.com/sirupsen/logrus"
   )

   func main() {
       log := logrus.New()

       // Set log format to JSON (for Logstash compatibility)
       log.SetFormatter(&logrus.JSONFormatter{})

       // Example log entries with fields
       log.WithFields(logrus.Fields{
           "event": "user_login",
           "user":  "john_doe",
           "status": "success",
       }).Info("User logged in")
   }
   ```

**Using Zap**:

1. **Install Zap**:

   Install Zap using the following command:

   ```bash
   go get go.uber.org/zap
   ```

2. **Basic Zap Configuration**:

   ```go
   package main

   import (
       "go.uber.org/zap"
   )

   func main() {
       // Create a new zap logger
       logger, _ := zap.NewProduction()
       defer logger.Sync()

       // Example log entries with fields
       logger.Info("User logged in", zap.String("user", "john_doe"), zap.String("status", "success"))
   }
   ```

In both cases, the logs are being generated in JSON format, which is ideal for processing by Logstash and indexing into Elasticsearch.

---

### **Step 2: Configure Logstash to Forward Logs to Elasticsearch**

Now that your Go application is generating structured logs, we need to configure **Logstash** to collect these logs and forward them to **Elasticsearch**.

#### **2.1 Install and Configure Logstash**

1. **Install Logstash** on your Ubuntu server:

   ```bash
   sudo apt update
   sudo apt install logstash
   ```

2. **Configure Logstash**:

   Logstash uses configuration files to define the pipeline for processing logs. Create a configuration file that listens for logs on a specified port (or file) and forwards them to Elasticsearch.

   Create a Logstash configuration file at `/etc/logstash/conf.d/go-logs.conf` with the following content:

   ```bash
   input {
     tcp {
       port => 5044
       codec => json
     }
   }

   filter {
     # Optional: Add any custom log processing or parsing rules here
     # Example: Parse a custom log field
   }

   output {
     elasticsearch {
       hosts => ["http://localhost:9200"]
       index => "go-logs-%{+YYYY.MM.dd}"
     }
   }
   ```

   - **Input**: The configuration listens on port `5044` for incoming JSON logs from your Go application. You can adjust this port as needed.
   - **Filter**: Here, you can add any parsing or enrichment rules (optional). In this case, the logs are already structured, so no parsing is needed.
   - **Output**: Logs are forwarded to Elasticsearch, where they are indexed in the `go-logs` index. Each log entry will be indexed with a date-based suffix (e.g., `go-logs-2024.11.22`).

3. **Start Logstash**:

   ```bash
   sudo systemctl start logstash
   sudo systemctl enable logstash
   ```

Logstash will now listen for logs from your Go application and forward them to Elasticsearch.

---

### **Step 3: Visualizing Logs in Kibana**

Once Logstash is forwarding logs to Elasticsearch, you can use **Kibana** to visualize and monitor the logs. Kibana provides a user-friendly interface to search, analyze, and create dashboards based on the data stored in Elasticsearch.

#### **3.1 Install and Configure Kibana**

1. **Install Kibana**:

   ```bash
   sudo apt install kibana
   ```

2. **Start Kibana**:

   ```bash
   sudo systemctl start kibana
   sudo systemctl enable kibana
   ```

3. **Access Kibana**:

   You can now access Kibana via `http://localhost:5601` in your browser.

#### **3.2 Create an Index Pattern in Kibana**

Before you can visualize your logs, you need to create an **index pattern** in Kibana that corresponds to the index used by Logstash (`go-logs-*`).

1. In Kibana, go to the **Management** section and click on **Index Patterns**.
2. Click **Create Index Pattern**.
3. Enter `go-logs-*` as the index pattern and click **Next Step**.
4. Select the **@timestamp** field (if it exists in your logs) as the time filter field.
5. Click **Create Index Pattern**.

#### **3.3 Build Dashboards to Monitor User Activity**

1. **Explore Logs**:

   - Go to **Discover** in Kibana and select your `go-logs-*` index.
   - You should see your logs listed in a table format. You can use the search bar to filter logs by fields like `user`, `status`, or `event`.

2. **Create Visualizations**:

   - Go to the **Visualize** section in Kibana and create visualizations based on your log data.
   - For example, you can create:
     - A bar chart that shows the number of successful and failed user logins.
     - A pie chart to represent the distribution of events (e.g., `user_login`, `user_logout`, etc.).
     - A line graph that shows login activity over time.

3. **Create a Dashboard**:
   - Once you have a few visualizations, go to **Dashboard** and click **Create new dashboard**.
   - Add your visualizations to the dashboard to get a comprehensive view of user activity in real time.

---

### **Mini-Project: Build a Real-Time Log Viewer in Kibana**

To put your knowledge into practice, let's build a **real-time log viewer** in Kibana to monitor user activity in your Go application. This mini-project will involve the following steps:

1. **Generate Real-Time Logs**: Modify your Go application to log user activity (e.g., logins, page views, errors).
2. **Send Logs to Logstash**: Use structured logging (`logrus` or `zap`) to send logs to Logstash.
3. **Forward Logs to Elasticsearch**: Configure Logstash to index logs into Elasticsearch.
4. **Create a Kibana Dashboard**: Build a dashboard in Kibana to visualize user activity, such as login attempts, user status, and activity over time.

The real-time log viewer will allow you to monitor user interactions in your application, track errors, and identify trends. You can use this project as a starting point for building more advanced monitoring and alerting systems.

---

### **Conclusion**

In this section, we covered how to integrate Go applications with the **ELK Stack** for real-time log monitoring and analysis. We walked through sending structured logs from a Go application to Logstash using `logrus` or `zap`, configuring Logstash to forward logs to Elasticsearch, and creating Kibana dashboards for visualizing user activity. This integration enables you to leverage the powerful tools of the ELK Stack for centralized log management, monitoring, and troubleshooting, providing you with deeper insights into your Go application’s behavior and performance.
