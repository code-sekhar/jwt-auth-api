# JWT Authentication API (Golang + Gin + PostgreSQL)

A secure REST API built with Golang, Gin Framework, PostgreSQL, GORM, and JWT Authentication.

## Features

* User Registration
* User Login
* Password Hashing using bcrypt
* JWT Token Authentication
* Protected Routes
* PostgreSQL Integration
* GORM ORM
* UUID Primary Keys
* Environment Variable Configuration

## Tech Stack

* Golang
* Gin Framework
* PostgreSQL
* GORM
* JWT (JSON Web Token)
* bcrypt
* UUID
* Godotenv

## Project Structure

```bash
jwt-auth-api/

├── cmd/
│   └── main.go
│
├── config/
│   └── database.go
│
├── controllers/
│   └── auth_controller.go
│
├── middleware/
│   └── auth_middleware.go
│
├── models/
│   └── user.go
│
├── routes/
│   └── routes.go
│
├── utils/
│   └── jwt.go
│
├── .env
├── go.mod
└── README.md
```

## Installation

### Clone Repository

```bash
git clone https://github.com/your-username/jwt-auth-api.git

cd jwt-auth-api
```

### Install Dependencies

```bash
go mod tidy
```

### Create Database

```sql
CREATE DATABASE jwt_auth;
```

### Environment Variables

Create a `.env` file in the root directory:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=jwt_auth

JWT_SECRET=my_secret_key
```

## Run Application

```bash
go run cmd/main.go
```

Application runs on:

```text
http://localhost:8080
```

---

## API Endpoints

### Register User

**POST**

```http
/register
```

Request Body

```json
{
  "name": "Sekhar",
  "email": "sekhar@gmail.com",
  "password": "123456"
}
```

Response

```json
{
  "message": "User Registered"
}
```

---

### Login User

**POST**

```http
/login
```

Request Body

```json
{
  "email": "sekhar@gmail.com",
  "password": "123456"
}
```

Response

```json
{
  "token": "jwt_token_here"
}
```

---

### Get Profile (Protected Route)

**GET**

```http
/profile
```

Headers

```http
Authorization: Bearer jwt_token_here
```

Response

```json
{
  "message": "Welcome Authenticated User"
}
```

---

## Authentication Flow

```text
Register
    ↓
Login
    ↓
Generate JWT Token
    ↓
Send Token in Authorization Header
    ↓
Access Protected Routes
```

## Security Features

* Password Hashing with bcrypt
* JWT-based Authentication
* Protected Route Middleware
* Unique Email Validation
* Environment Variable Secrets

## HTTP Status Codes

| Code | Description           |
| ---- | --------------------- |
| 200  | Success               |
| 201  | Created               |
| 400  | Bad Request           |
| 401  | Unauthorized          |
| 409  | Conflict              |
| 500  | Internal Server Error |

## Future Improvements

* Refresh Token
* Role Based Access Control (RBAC)
* Email Verification
* Password Reset
* Swagger Documentation
* Docker Support
* Redis Session Management
* Clean Architecture

## Author

Sekhar Chaudhary

Web Application Developer

Skills:
Laravel, Golang, Node.js, React, Vue, Angular, PostgreSQL
