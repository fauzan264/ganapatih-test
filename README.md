# Ganapatih Test

## 📌 Overview

Ganapatih Test is a simple fullstack social media project consisting of **Backend (Golang + Fiber + GORM + PostgreSQL)** and **Frontend (Next.js + TailwindCSS)**.

This project was built as a take-home test submission, focusing on **JWT authentication**, **model relationships** (User, Feed, Follow), and a **REST API** for social media needs.

---

## 🚀 Tech Stack

### Backend

- **Golang 1.24.0** - Programming language
- **Fiber** - Fast and expressive web framework
- **GORM** - ORM for database access
- **PostgreSQL** - Relational database
- **JWT (JSON Web Token)** - For authentication and authorization
- **bcrypt** - For password hashing

### Frontend

- **Next.js** - React framework with server-side rendering
- **TailwindCSS** - Utility-first CSS framework
- **Axios** - HTTP client for API calls
- **React Hooks** - State management

---

## ⚙️ Setup Instructions

### Prerequisites

- **Go 1.24.0** or higher
- **Node.js 18+** and npm/yarn
- **PostgreSQL 14+**
- **Git**

### 1. Clone Repository

```bash
git clone https://github.com/fauzan264/ganapatih-test.git
cd ganapatih-test
```

### 2. Setup Backend

#### a. Navigate to backend directory

```bash
cd backend
```

#### b. Install Go dependencies

```bash
go mod download
```

#### c. Setup environment variables

```bash
cp .env.example .env
```

#### d. Edit `.env` file with your configuration

```env
# Application
APP_HOST=127.0.0.1
APP_PORT=8080
APP_ENV=development

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=ganapatih_test
DB_SSLMODE=disable

# JWT
JWT_SECRET=your-super-secret-key-change-this-in-production
JWT_EXPIRE=24h
```

#### e. Create database

```bash
# Connect to PostgreSQL
psql -U postgres

# Create database
CREATE DATABASE ganapatih_test;
\q
```

#### f. Run database migration

```bash
# Migration will run automatically when the app starts
# Or run manual migration script if available
go run main.go
```

#### g. Start backend server

```bash
go run main.go
```

Backend server will run at: **http://localhost:8080**

---

### 3. Setup Frontend

#### a. Navigate to frontend directory

```bash
cd frontend
```

#### b. Install dependencies

```bash
npm install
# or
yarn install
```

#### c. Setup environment variables

```bash
cp .env.example .env.local
```

#### d. Edit `.env.local` file

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api
```

#### e. Run development server

```bash
npm run dev
# or
yarn dev
```

Frontend will run at: **http://localhost:3000**

---

## 📖 API Documentation

Base URL: `http://localhost:8080/api`

### Authentication

#### Register New User

```http
POST /api/register
Content-Type: application/json

{
  "username": "budi",
  "password": "password"
}
```

**Response (201 Created):**

```json
{
  "id": 1,
  "username": "budi"
}
```

**Response (409 Conflict) - Username already taken:**

```json
{
  "message": "Error: username \"budi\" already taken."
}
```

---

#### Login

```http
POST /api/login
Content-Type: application/json

{
  "username": "budi",
  "password": "password"
}
```

**Response (200 OK):**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

---

#### Get Session (Protected)

```http
GET /api/session
Authorization: Bearer {token}
```

**Response (200 OK):**

```json
{
  "id": 1,
  "username": "budi"
}
```

---

### Feeds

#### Get Feed (Protected)

```http
GET /api/feed?page=1&limit=10
Authorization: Bearer {token}
```

**Query Parameters:**

- `page` (optional): Page number, default 1
- `limit` (optional): Items per page, default 10

**Response (200 OK):**

```json
[
  {
    "id": 15,
    "userid": 2,
    "content": "post from user 2",
    "createdat": "2025-09-12T11:00:00Z"
  },
  {
    "id": 12,
    "userid": 3,
    "content": "post from user 3",
    "createdat": "2025-09-12T10:30:00Z"
  }
]
```

---

#### Create New Post (Protected)

```http
POST /api/posts
Authorization: Bearer {token}
Content-Type: application/json

{
  "content": "hello world!"
}
```

**Response (201 Created):**

```json
{
  "message": "Successfully created a post ≤ 200 characters",
  "data": {
    "id": 10,
    "userid": 1,
    "content": "hello world!",
    "createdat": "2025-09-12T10:00:00Z"
  }
}
```

**Response (422 Unprocessable Entity) - Content > 200 characters:**

```json
{
  "message": "Content must be 200 characters or less"
}
```

---

### Follow System

#### Follow a User (Protected)

```http
POST /api/follow/:userid
Authorization: Bearer {token}
```

**URL Parameters:**

- `userid`: ID of user to follow

**Response (200 OK):**

```json
{
  "message": "you are now following user 2"
}
```

**Response (404 Not Found) - User doesn't exist:**

```json
{
  "message": "Follow a non-existent user"
}
```

---

#### Unfollow a User (Protected)

```http
DELETE /api/follow/:userid
Authorization: Bearer {token}
```

**URL Parameters:**

- `userid`: ID of user to unfollow

**Response (200 OK):**

```json
{
  "message": "you unfollowed user 2"
}
```

---

### Error Responses

**400 Bad Request:**

```json
{
  "message": "Invalid request data"
}
```

**401 Unauthorized:**

```json
{
  "message": "Unauthorized - Invalid or missing token"
}
```

**404 Not Found:**

```json
{
  "message": "Resource not found"
}
```

**409 Conflict:**

```json
{
  "message": "Error: username \"budi\" already taken."
}
```

**422 Unprocessable Entity:**

```json
{
  "message": "Content must be 200 characters or less"
}
```

**500 Internal Server Error:**

```json
{
  "message": "Internal server error"
}
```

---

## 🧪 Test Cases

API documentation and test cases are available in `backend/docs/openapi.yaml`.

### Authentication Tests

✅ **Register with an existing username** → Fail (409 Conflict)  
✅ **Register with valid data** → Success (201 Created)  
✅ **Login with wrong username** → Fail (401 Unauthorized)  
✅ **Login with wrong password** → Fail (401 Unauthorized)  
✅ **Login with valid credentials** → Success, return JWT token  
✅ **Access protected endpoint without token** → Fail (401 Unauthorized)  
✅ **Access protected endpoint with invalid token** → Fail (401 Unauthorized)  
✅ **Access protected endpoint with valid token** → Success (200 OK)

### Posts/Feeds Tests

✅ **Create post without login** → Fail (401 Unauthorized)  
✅ **Create post with empty content** → Fail (400 Bad Request)  
✅ **Create post with content > 200 characters** → Fail (422 Unprocessable Entity)  
✅ **Create post with valid login and ≤ 200 chars content** → Success (201 Created)  
✅ **Get feed with pagination** → Success, return posts array  
✅ **Get feed without login** → Fail (401 Unauthorized)

### Follow System Tests

✅ **User follows themselves** → Fail (400 Bad Request)  
✅ **User follows another valid user** → Success (201 Created)  
✅ **User follows the same user twice** → Fail (400 Bad Request)  
✅ **User unfollows a user not being followed** → Fail (404 Not Found)  
✅ **User unfollows successfully** → Success (200 OK)  
✅ **Follow a non-existent user** → Fail (404 Not Found)

---

## 📌 Design Notes

### Architecture

This project uses a simple **Clean Architecture** with clear separation of layers:

1. **Handlers** - Handle HTTP requests/responses
2. **Services** - Business logic and orchestration
3. **Repositories** - Data access layer using GORM
4. **Models** - Database models and struct definitions
5. **Middleware** - JWT authentication and request validation

### Security Features

- Passwords hashed using **bcrypt** (cost factor 10)
- JWT token for stateless authentication
- Protected routes with middleware
- SQL injection prevention via GORM
- CORS configuration for frontend integration
- Input validation on all endpoints

### Best Practices

- Consistent error handling across the application
- Standardized response format with `success`, `message`, and `data`
- Pagination for list endpoints
- Database indexing on foreign keys and unique fields
- Environment-based configuration for flexible deployments
