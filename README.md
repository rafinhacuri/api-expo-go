# 🚀 API Expo Go

A simple and secure API built with **Gin** and **MongoDB**, featuring authentication, task management, and user management.

---

## ✅ Project Roadmap

### 🛠️ Core
- [x] Initialize Gin routes
- [x] Create database connection
- [ ] Finalize user model
- [ ] Finalize task model

### 🔒 Security
- [x] Password hashing
- [x] Password verification
- [ ] JWT-based login authentication
- [ ] Protect routes for authenticated users only

### ⚙️ Utilities
- [x] Validate passwords
- [x] Validate emails

### 👤 User Management
- [x] Insert user
- [x] Update user
- [x] Fetch user
- [x] Fetch all users
- [ ] Delete user

### 📋 Task Management
- [ ] Insert task
- [ ] Update task
- [ ] Fetch task
- [ ] Fetch all tasks
- [ ] Delete task

### 📚 Docs
- [ ] Create project documentation

---

## 📌 Tech Stack
- [Gin](https://gin-gonic.com/) - Web Framework
- [MongoDB](https://www.mongodb.com/) - Database
- [JWT](https://jwt.io/) - Authentication
- [Go](https://go.dev/) - Language

---

## 📖 How to Use
1. Clone this repository:
   ```bash
   git clone https://github.com/rafinhacuri/api-expo-go.git
   ```
2. Navigate into the project:
   ```bash
   cd api-expo-go
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run the server:
   ```bash
   go run ./
   ```

---