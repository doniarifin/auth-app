# Go Auth App with Swagger API Doc (using PostgreSQL)

Simple authentication API built with **Go**, **Gin**, and **PostgreSQL**.

---

## Features

- User Registration
- User Login with JWT
- Password Hashing (bcrypt)
- PostgreSQL integration
- JWT-based authentication
- Swagger API Documentation

---

## Quick Start

### 1. Clone Repo

```
git clone https://github.com/doniarifin/go-auth-app.git
cd go-auth-app
```

### 2. Rename/Copy .env_copy to .env

Copy

```
cp .env_copy .env
```

### 3. Install Dependencies & Run the App

Ensure you have **PostgreSQL** installed.

Then run the app:

```
go mod tidy
go run main.go
```

### 4. Swagger API Documentation

Ensure you have **Swagger** installed.

Install

```
go install github.com/swaggo/swag/cmd/swag@latest
```

Run `swag init` in the project root containing `main.go` to generate the docs folder and `docs/docs.go` file from your comments.

```
swag init
```

**[Click here](https://github.com/swaggo/swag)** to view the Swaggo documentation

Then open in browser:

```
http://localhost:8080/swagger/index.html
```

_note: set port according to your local port_

---

## Endpoints

| Method | Endpoint                 | Description       |
| ------ | ------------------------ | ----------------- |
| POST   | `/register`              | Register new user |
| POST   | `/login`                 | Login with JWT    |
| GET    | `/api/v1/GetCurrentUser` | Get current user  |
| PUT    | `/api/v1/UpdateUser`     | Update User       |

## Project Structure

```
auth-app/
├── config/          # Load env & config
├── controllers/     # Auth handlers
├── database/        # DB connection & migration
├── middleware/      # JWT Middleware
├── models/          # User model
├── routes/          # Route setup
├── utils/           # Hashing & JWT
├── .env             # Environment variables
├── main.go          # Entry point
```

## Stack

- Go
- Gin
- GORM
- PostgreSQL
- JWT
- bcrypt
- Swagger

---
