# 🛡️ Go Auth App (using SQL Server)

Simple authentication API built with **Go**, **Gin**, and **SQL Server**.

---

## 📦 Features

- ✅ User Registration
- ✅ User Login with JWT
- ✅ Protected Route (`/me`)
- ✅ Password Hashing (bcrypt)
- ✅ SQL Server integration
- ✅ JWT-based authentication
- ✅ Clean Project Structure

---

## 🚀 Quick Start

### 1. Clone Repo

```
git clone https://github.com/doniarifin/go-auth-app.git
cd go-auth-app
```

### 2. Rename .env

```
mv env_copy .env
```

### 3. Install Dependencies & Run the App

```
go mod tidy
go run main.go
```

---

## 🧪 Endpoints

| Method | Endpoint    | Description       |
| ------ | ----------- | ----------------- |
| POST   | `/register` | Register new user |
| POST   | `/login`    | Login with JWT    |
| GET    | `/api/me`   | Get current user  |


## 🗂️ Project Structure

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

## 🛠️ Stack

- ✅ Go
- ✅ Gin
- ✅ SQL Server
- ✅ JWT
- ✅ bcrypt

---