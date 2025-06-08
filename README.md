# Ingestion Service

This service handles incoming messages via REST API and publishes them to Apache Kafka for further processing by other services like Persistence and SSE.

## 📦 Features

- REST API endpoint `/api/message` to receive messages.
- Publishes messages to Kafka topic `message-topic`.
- JWT-protected endpoint using Gin middleware.
- Modular folder structure.
- Unit tested using `httptest` and mock Kafka writer.

---

## ⚙️ Requirements

- Go 1.20+
- Kafka (running on `localhost:9092`)
- Zookeeper (optional, depending on Kafka setup)
- `.env` file with configuration
- Make sure Kafka topic `message-topic` exists

---


## 🔐 Environment Variables

Create a `.env` file:

```env
JWT_SECRET=secret
PORT=9091
KAFKA_BROKERS=localhost:9092


