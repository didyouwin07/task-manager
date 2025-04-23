# Task Manager

An in-memory Task Management microservice that supports creating, reading, updating, and deleting tasks, along with pagination and filtering support.

---

## Problem Breakdown

Build a **Task Management System** with the following requirements:
- CRUD operations on tasks.
- Pagination on `GET /tasks`.
- Filtering tasks by status (`Completed`, `Pending`, etc.).
- Microservice-oriented design with separation of concerns.
- Scalable architecture that can evolve to support user management, authentication, etc.

---

## Design Decisions

### Microservice Principles Applied:
- **Single Responsibility Principle**: Handlers, services, and repositories each manage a separate layer of the application.
- **Loose Coupling**: Components are loosely coupled and can be independently tested or replaced.
- **Scalability**: In-memory storage now, but repository pattern allows easy switching to persistent storage later (e.g., Postgres, MongoDB).
- **Extensibility**: Easily extendable to multiple services like User Management, Notifications, etc.

### Architecture:
```
main.go
internal/
├── handler/      # HTTP handlers (e.g., CreateTask, GetTasks)
├── service/      # Business logic
├── repository/   # In-memory data storage
├── model/        # Data model definition
└── router/       # HTTP routing (mux)
```

---

## How to Run the Service

### Prerequisites
- Go 1.24 or higher installed.

### Steps to run

```bash
# Clone the repo
git clone https://github.com/didyouwin07/task-manager.git
cd task-manager

# Run the service
go run ./cmd/server/
```

The server will start on `http://localhost:8080`.

---

## API Documentation

### Create Task
**POST** `/tasks`

#### Request:
```json
{
  "title": "Buy groceries",
  "description": "Milk, eggs, bread",
  "status": "Pending"
}
```

#### Response:
**Status:** 201 Created
```json
{
  "id": "generated-uuid",
  "title": "Buy groceries",
  "description": "Milk, eggs, bread",
  "status": "Pending"
}
```

### Get All Tasks with Pagination & Filtering
**GET** `/tasks?page=1&limit=10&status=Completed`

Returns a list of tasks filtered by status and paginated.

#### Response:
```json
[
  {
    "id": "uuid-1",
    "title": "Buy groceries",
    "description": "Milk, eggs, bread",
    "status": "Completed"
  },
  ...
]
```

### Get Task by ID
**GET** `/tasks/{id}`

#### Response:
```json
{
  "id": "uuid-1",
  "title": "Buy groceries",
  "description": "Milk, eggs, bread",
  "status": "Completed"
}
```

### Update Task
**PUT** `/tasks/{id}`

#### Request:
```json
{
  "title": "Updated title",
  "description": "Updated description",
  "status": "Completed"
}
```

### Delete Task
**DELETE** `/tasks/{id}`

Returns **204 No Content** on success.

---

## How It Demonstrates Microservices Concepts

- **Isolation of Concerns**: Business logic, HTTP handlers, and data storage are modular and can be swapped independently.
- **Scalability**: The design can easily be extended to distributed systems by adding message queues or switching to gRPC for communication.
- **Future Inter-service Communication**: If we add a User service, it can interact with this Task service using REST or gRPC. For high throughput systems, RabbitMQ can also be integrated.

---

## Tech Stack

- **Language**: Go
- **Routing**: Gorilla Mux
- **UUID Generation**: `google/uuid`
- **Storage**: In-memory

---

## Author

Vishwas Bordia (didyouwin07)