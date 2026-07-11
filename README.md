# 🚀 OrderFlow

**OrderFlow** is an event-driven backend system built using **Go**, **Gin**, **Kafka**, **Redis**, and **MySQL**. The project demonstrates how modern backend applications process orders asynchronously while maintaining scalability, reliability, and clean architecture.

It is designed as a portfolio project to showcase backend engineering concepts such as authentication, asynchronous messaging, caching, idempotency, API documentation, and containerization.

---

# ✨ Features

- JWT Authentication
- User Registration & Login
- Order Creation & Retrieval
- Event-driven Order Processing with Kafka
- Redis Caching
- Redis Rate Limiting
- Idempotent Order APIs
- MySQL Persistent Storage
- Layered Architecture (Handler → Service → Repository)
- Swagger API Documentation
- Docker Support

---

# 🏗️ Tech Stack

| Technology | Purpose |
|------------|---------|
| Go | Backend Language |
| Gin | HTTP Web Framework |
| MySQL | Relational Database |
| Redis | Caching & Rate Limiting |
| Kafka | Asynchronous Event Processing |
| JWT | Authentication |
| Swagger | API Documentation |
| Docker | Containerization |

---

# 📂 Project Structure

```text
orderflow
├── cache/
├── cmd/
│   └── api/
├── config/
├── database/
├── docs/
├── dto/
├── events/
├── handlers/
├── kafka/
├── middleware/
├── models/
├── repositories/
├── routes/
├── services/
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

---

# ⚙️ Architecture

```text
                Client
                   │
                   ▼
             Gin REST API
                   │
         JWT Authentication
                   │
                   ▼
      Handler → Service → Repository
                   │
      ┌────────────┼────────────┐
      ▼            ▼            ▼
   MySQL        Kafka        Redis
(Database)   (Async Queue) (Cache)
                   │
                   ▼
            Kafka Consumer
                   │
                   ▼
        Order Processing Service
```

---

# 🔄 Order Processing Flow

1. User registers and logs in.
2. A JWT token is generated.
3. User creates an order.
4. Order is stored in MySQL.
5. Order event is published to Kafka.
6. Kafka consumer processes the order asynchronously.
7. Redis caches frequently accessed order data.
8. Order status is updated after processing.

---

# 🔐 Authentication

OrderFlow uses JWT-based authentication.

Protected endpoints require:

```text
Authorization: Bearer <JWT_TOKEN>
```

---

# 📖 API Documentation

Swagger UI:

```text
http://localhost:8080/swagger/index.html
```

---

# 📡 API Endpoints

## Authentication

| Method | Endpoint | Description |
|---------|----------|-------------|
| POST | /register | Register User |
| POST | /login | Login User |

## User

| Method | Endpoint | Description |
|---------|----------|-------------|
| GET | /profile | User Profile |

## Orders

| Method | Endpoint | Description |
|---------|----------|-------------|
| POST | /orders | Create Order |
| GET | /orders | Get User Orders |

## Metrics

| Method | Endpoint | Description |
|---------|----------|-------------|
| GET | /metrics | Order Metrics |

---

# 🚀 Running Locally

## Clone Repository

```bash
git clone https://github.com/<your-username>/orderflow.git

cd orderflow
```

## Install Dependencies

```bash
go mod tidy
```

## Configure Environment

Create a `.env` file using `.env.example`.

## Start Required Services

- MySQL
- Redis
- Kafka
- Zookeeper

(using Docker or Docker Compose)

## Run

```bash
go run cmd/api/main.go
```

---

# 🐳 Docker

Build

```bash
docker build -t orderflow .
```

Run

```bash
docker run -p 8080:8080 orderflow
```

For the complete application stack (Go + MySQL + Redis + Kafka), use Docker Compose.

---

# 📈 Backend Concepts Demonstrated

- REST API Development
- Layered Architecture
- Repository Pattern
- JWT Authentication
- Event-driven Architecture
- Asynchronous Processing
- Kafka Producers & Consumers
- Redis Caching
- Rate Limiting
- Idempotency
- API Documentation
- Docker Containerization

---

# 🔮 Future Improvements

- Docker Compose orchestration
- Prometheus Metrics
- Health Check Endpoint
- Structured Logging
- CI/CD Pipeline with GitHub Actions
- Unit & Integration Testing
- gRPC Support
- Kubernetes Deployment

---
