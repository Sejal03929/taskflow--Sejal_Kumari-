# TaskFlow Backend API

---

## 1. Overview

This is a backend API for a simple Task Management system.

It allows users to:

* Register and login (JWT authentication)
* Create and manage projects
* Create and manage tasks within projects

### Tech Stack

* Go (Gin framework)
* PostgreSQL
* Docker & Docker Compose

---

## 2. Architecture Decisions

### Structure

The project follows a modular structure:

* `cmd/` → Entry point (main.go)
* `internal/db/` → Database connection
* `internal/handlers/` → API handlers (auth, project, task)
* `internal/middleware/` → Authentication middleware
* `migrations/` → SQL schema

### Why this structure?

* Keeps code organized and scalable
* Separates concerns (DB, handlers, middleware)

### Tradeoffs

* Did not use ORM (used raw SQL for simplicity)
* Basic error handling (kept simple for assignment)

### What I left out

* Pagination
* Advanced validation
* Role-based access control

---

## 3. Running Locally

### Prerequisites

* Docker installed

### Steps

```bash
git clone https://github.com/Sejal0329/taskflow-backend
cd taskflow-backend
docker-compose up --build
```

Server runs at:
http://localhost:5000

---

## 4. Running Migrations

Migrations run automatically on container startup using:

```
migrations/001_init.up.sql
```

No manual steps required.

---

## 5. Test Credentials

Use this to login directly:

Email:    [test@example.com](mailto:test@example.com)
Password: password123

(If not present, register using API)


The API collection is available in the repo:

/bruno/taskflow-api.json

Import it into Bruno to test all endpoints easily.

## 6. API Reference

### Auth

#### Register

POST /auth/register

```json
{
  "name": "Sejal",
  "email": "test@example.com",
  "password": "password123"
}
```

#### Login

POST /auth/login

Response:

```json
{
  "token": "JWT_TOKEN"
}
```

---

### Projects

#### Create Project

POST /projects

Headers:
Authorization: Bearer <token>

```json
{
  "name": "Project 1",
  "description": "My first project"
}
```

---

#### Get Projects

GET /projects

Headers:
Authorization: Bearer <token>

---

### Tasks

#### Create Task

POST /projects/:id/tasks

Headers:
Authorization: Bearer <token>

```json
{
  "title": "Task 1",
  "description": "Do something",
  "status": "todo",
  "priority": "high"
}
```

---

#### Get Tasks

GET /projects/:id/tasks

Headers:
Authorization: Bearer <token>

---

## 7. What I'd Do With More Time

If I had more time, I would:

* Add input validation (using validator library)
* Implement pagination for projects and tasks
* Add update/delete APIs
* Improve error handling and logging
* Add unit tests
* Implement role-based access control
* Use an ORM like GORM for cleaner queries

---
