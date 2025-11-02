# Go Auth App with Swagger API Doc (using PostgreSQL)

Simple authentication API built with **Go**, **Gin**, and **PostgreSQL**.

[Live Demo](https://go-auth-app.koyeb.app/swagger/index.html)

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
git clone https://github.com/doniarifin/auth-app.git
cd auth-app
```

### 2. Rename/Copy .env_copy to .env

Setup the port and DB in `.env` file

```
cp .env_copy .env
```

### 3. Install Dependencies & Run the App

Ensure you have **PostgreSQL** installed.

- Install Depedencies:

```
go mod tidy
```

- Entry point is inside of `cmd` directory, change dir to `cmd` dir, then run the app:

```
cd cmd
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
swag init -g cmd/main.go
```

**[Click here](https://github.com/swaggo/swag)** to view the Swaggo documentation

Then open in browser:

```
http://localhost:8080/swagger/index.html
```

_note: set port according to your local port_

---

## Endpoints

| Method | Endpoint                 | Description                    |
| ------ | ------------------------ | ------------------------------ |
| POST   | `/register`              | Register new user              |
| POST   | `/login`                 | Login with JWT                 |
| GET    | `/api/v1/GetCurrentUser` | Get current user               |
| GET    | `/api/v1/GetAllUsers`    | Get all users                  |
| PUT    | `/api/v1/Update/{id}`    | Update User                    |
| DELETE | `/api/v1/Delete/{id}`    | Delete User with Authorization |

## Project Structure

```
auth-app/
├── cmd/              # entry point
├── config/           # config
├── docs/             # swagger docs
├── internal
│   ├── database/     # database
│   ├── dto/          # data tranfer object
│   ├── handler/      # handler
│   ├── middleware/   # middleware and authorization
│   ├── model/        # model
│   ├── pkg
│   │   ├── jwt/      # jwt
│   │   └── logger/
│   ├── repository/   # repository
│   ├── routes/       # setup routes
│   ├── service/      # service / bussines logic
│   └── utils/        # utilities
└── .env              # environment variables
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
