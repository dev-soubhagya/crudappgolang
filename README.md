
# CRUD App in Golang

This project is a CRUD application built using Golang with the following features:
- Upload an Excel file and import data into MySQL.
- Cache data in Redis for fast retrieval.
- Perform CRUD operations on imported data.

## Folder Structure

```
crudappgolang/
├── cmd/
│   └── main.go              # Entry point for the application
├── config/
│   └── config.go            # Configuration setup (MySQL, Redis, etc.)
├── controllers/
│   └── user.go              # Handlers for API endpoints
├── environment/
│   └── cred.env             # Environment variables
├── models/
│   └── user.go              # Database model and table initialization
├── routes/
│   └── routes.go            # API route definitions
├── utils/
│   └── response.go          # Utility for consistent API responses
├── go.mod                   # Module file
├── go.sum                   # Dependencies file
```

## Features

1. **Excel File Upload**
   - Parse data from an uploaded Excel file.
   - Validate file structure for expected columns (`Name` and `Email`).
   - Insert data into MySQL asynchronously.

2. **Redis Caching**
   - Cache imported data for quick access.
   - Cache expires after 5 minutes.

3. **CRUD Operations**
   - Create, Read, Update, and Delete user records.
   - Ensure database and cache consistency.

4. **Graceful Error Handling**
   - Validate inputs and handle errors gracefully at every stage.

## Setup Instructions

### Prerequisites

- Go 1.23 or higher
- MySQL database
- Redis server
- Postman (or any other API testing tool)

### Step 1: Clone the repository
```bash
git clone <repository_url>
cd crudappgolang
```

### Step 2: Install dependencies
```bash
go mod tidy
```

### Step 3: Configure Environment Variables
Create a file named `cred.env` in the `environment/` directory with the following content:
```
DB_DSN = "user:password@tcp(localhost:3306)/mydb"
REDIS_HOST = "localhost:6379"

```

### Step 4: Run the Application
```bash
go run cmd/main.go
```

The server will start at `http://localhost:8080`.

### Step 5: Create MySQL Table
The application automatically creates the `users` table if it doesn't exist.

## API Endpoints

### 1. Upload Excel File
**POST** `/upload`

Upload an Excel file with columns `Name` and `Email`.

**Request**:
- File: Multipart form-data (`file` field)

**Response**:
- Success: `200 OK`
- Error: Appropriate error message.

---

### 2. View Data
**GET** `/data`

Fetch user data from Redis or MySQL.

**Response**:
- Success: List of users.
- Error: Appropriate error message.

---

### 3. Edit a Record
**PUT** `/edit/:id`

Update user details.

**Request Body**:
```json
{
  "name": "Updated Name",
  "email": "updatedemail@example.com"
}
```

**Response**:
- Success: `200 OK`
- Error: Appropriate error message.

---

## Technologies Used
- **Framework**: [Gin](https://github.com/gin-gonic/gin)
- **Database**: MySQL [Go Library](database/sql)
- **Cache**: Redis [Go Library](github.com/gomodule/redigo/redis)
- **Excel Parsing**: [Excelize](https://github.com/xuri/excelize)
