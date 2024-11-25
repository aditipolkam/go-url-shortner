# **Go HTTP Server with MongoDB**

This project is a production-ready Go HTTP server that connects to MongoDB, featuring a modular folder structure and a scalable architecture. The server includes two example routes: one for fetching records (GET) and another for adding records (POST).

## **Features**

- RESTful HTTP server built with Go.
- MongoDB integration for persistent data storage.
- Clean, modular folder structure for scalability and maintainability.
- Graceful shutdown handling.
- Configurable via environment variables.

## **Setup and Installation**

### **1. Prerequisites**

- Go (>= 1.20)
- MongoDB (running instance)

### **2. Clone the Repository**

```bash
git clone https://github.com/aditipolkam/go-url-shortner.git
cd go-url-shortner
```

### **3. Configure Environment Variables**

Create a `.env` file in the root directory with the following variables:

```env
DB_URL=mongodb://localhost:27017
SERVER_ADDRESS=:8080
```

### **4. Run the Server**

Install dependencies and start the server:

```bash
go mod tidy
go run cmd/server/main.go
```

The server will run on `http://localhost:8080`.

---

## **Routes**

### **1. GET `/:shortcode`**

For redirecting shortened url to the original url

#### **Request**

```bash
curl -X GET http://localhost:8080/{{shortcode}}
```

#### **Response**

Redirected to the original URL.

---

### **2. POST `/`**

Add a new record to the database.

#### **Request**

```bash
curl -X POST http://localhost:8080/ \
-H "Content-Type: application/json" \
-d '{"url": "https://blog.aditipolkam.me/efficient-database-indexes"}'
```

#### **Response**

```json
{
  "short_url": "http://localhost:8080/xvlixG"
}
```

---

## **Graceful Shutdown**

The server supports graceful shutdown on receiving `SIGINT` (Ctrl+C) to ensure cleanup tasks, such as database disconnection, are performed.

---

## **Contributing**

1. Fork the repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes and push:
   ```bash
   git push origin feature-name
   ```
4. Open a pull request.

---

## **License**

This project is licensed under the [MIT License](LICENSE).

---

## **Acknowledgments**

- [Go Documentation](https://golang.org/doc/)
- [MongoDB Go Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)

---
