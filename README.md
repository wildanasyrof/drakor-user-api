
# User Management API  

A **RESTful API** built with **Go (Golang)** using **Clean Architecture** principles.  
It provides user authentication, favorites management, and history tracking with secure **JWT-based authorization**.  

---

## 🚀 Features
- 🔐 User authentication (register, login, JWT)  
- 👤 User management  
- ⭐ Favorites system  
- 📜 History tracking  
- 🛡️ Middleware for logging & authentication  
- 📦 Environment-based configuration  
- 📂 Postman collection included for testing  

---

## 🛠 Tech Stack
- **Language:** Go (Golang)  
- **Architecture:** Clean Architecture (Domain, Repository, Service, Handler)  
- **Auth:** JWT (JSON Web Token)  
- **Database:** (configurable in `.env`)  
- **Tools:** Postman for API testing  

---

## 📂 Project Structure
```
.
├── cmd/                # Application entry point
├── internal/
│   ├── db/             # Database connection
│   ├── repository/     # Data access layer
│   ├── service/        # Business logic
│   ├── http/           # Handlers, routers, middleware
│   ├── domain/         # Entities & DTOs
│   ├── config/         # Configuration loader
│   └── di/             # Dependency injection
├── pkg/                # Utilities (jwt, logger, response, etc.)
├── .env.example        # Environment variables example
├── go.mod / go.sum     # Go modules
└── User.postman_collection.json # API testing collection
```

---

## ⚡ Getting Started

### 1️⃣ Clone the repo
```bash
git clone https://github.com/your-username/your-repo.git
cd your-repo
```

### 2️⃣ Install dependencies
```bash
go mod tidy
```

### 3️⃣ Setup environment
Copy `.env.example` to `.env` and update values:
```bash
cp .env.example .env
```

### 4️⃣ Run the server
```bash
go run cmd/main.go
```

---

## 📬 API Testing
Use the included **Postman collection**:  
`User.postman_collection.json`

---