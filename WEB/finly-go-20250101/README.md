# **Finly - Invoicing Web Application**  

**Finly** is a modern invoicing application designed to help businesses easily generate and manage invoices online. It includes features like invoice creation, management, and data visualization through interactive charts.  

This project has been built using **Go (Golang)** with the **Fiber** framework, **MongoDB**, and **Chart.js** for data visualization. The frontend is styled with **Tailwind CSS** and **DaisyUI**.

---

## **Features**  
- Create and manage invoices online.  
- View detailed invoice statistics.  
- Interactive data visualization with charts.  
- User-friendly and responsive design powered by Tailwind CSS and DaisyUI.  

---

## **Requirements**  
- **Go** (v1.18 or higher)  
- **MongoDB** (v4.4 or higher)  

---

## **Getting Started**  

### **Installation**  

1. Clone this repository:  
   ```bash  
   git clone https://github.com/mzs-house/finly.git  
   cd finly  
   ```  

2. Install dependencies:  
   ```bash  
   go mod tidy  
   ```  

3. Copy the environment variables file:  
   ```bash  
   cp .env.sample .env  
   ```  

4. Ensure MongoDB is ready. You can use Docker for easy setup.

5. Start the application:  
   ```bash  
   go run main.go  
   ```  

The server will start at `http://localhost:8080`. Open the link in your browser to use the application.  

---

## **Usage**  

Once the application is running, you can:  
- **Create invoices**: Add new invoices with relevant details.  
- **Manage invoices**: View, update, or delete existing invoices.  
- **Analyze data**: Gain insights from charts and invoice statistics.  

---

## **Technologies Used**  
- **Go (Golang)**: Backend programming language.  
- **Fiber**: Web application framework for APIs and routing.  
- **MongoDB**: NoSQL database for data storage.  
- **Chart.js**: Interactive charting library for data visualization.  
- **Tailwind CSS** and **DaisyUI**: For modern, responsive UI design.  
- **Docker** and **Docker Compose**: For containerization and simplified deployment.

---

## **Contributing**  

Contributions are welcome! Follow these steps to contribute:  
1. Fork this repository.  
2. Create a new branch for your feature:  
   ```bash  
   git checkout -b feature/your-feature-name  
   ```  
3. Commit your changes:  
   ```bash  
   git commit -m "Add your feature"  
   ```  
4. Push to your fork:  
   ```bash  
   git push origin feature/your-feature-name  
   ```  
5. Open a pull request to this repository.  

---

## **License**  
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
