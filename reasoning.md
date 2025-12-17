# Go Backend Development Task – Reasoning

## Overview

This project is a RESTful API built using Go and Fiber to manage users with name and date of birth. The API stores the date of birth in the database and calculates the user's age dynamically at runtime.

The goal was to learn Go quickly while following clean backend architecture principles.

---

## Tech Stack

- Go
- Fiber (HTTP framework)
- PostgreSQL (Dockerized)
- SQLC (type-safe database access)
- Uber Zap (structured logging)
- godotenv (environment configuration)

---

## Project Structure

I followed a layered architecture to keep the code clean and maintainable:

- handler: Handles HTTP requests and responses
- service: Contains business logic such as age calculation
- repository: Handles database operations using SQLC
- db/migrations: SQL migrations for database schema
- cmd/server: Application entry point

This separation makes the code easier to test, extend, and reason about.

---

## Database Design

The users table stores:

- id (primary key)
- name
- dob (date of birth)

The age is not stored in the database. It is calculated dynamically in Go using the time package to ensure correctness over time.

---

## Configuration Management

Environment variables are managed using a `.env` file and loaded using `godotenv`.
This avoids hardcoding secrets and makes the application portable across environments.

---

## Logging

Uber Zap is used for structured logging. It logs key application events such as server startup and request handling, which is helpful for debugging and production readiness.

---

## Key Design Decisions

- Used SQLC to generate type-safe database queries
- Used service layer to keep business logic separate from HTTP logic
- Calculated age dynamically instead of storing it
- Used Docker for database consistency across environments

---

## What I Learned

- How to structure a Go backend project
- How to use SQLC with PostgreSQL
- Dependency injection in Go
- Clean separation of concerns in backend systems

---

## Bonus work

- Added pagination for listing users endponts-> /users?limit=10&page=1
