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
- [x] Delete user

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

## ✅ Quick Start Checklist

### 🔧 Step 1 — Install Prerequisites
- [ ] Install **Docker**  
- [ ] Install **Docker Compose**  
👉 Official guide: [Get Docker](https://docs.docker.com/get-started/get-docker/)

---

### 📦 Step 2 — Get the `docker-compose.yaml`
Choose one of the options below to download the configuration file:

<details>
<summary>🔽 Using curl</summary>

```bash
curl -L -o docker-compose.yaml https://raw.githubusercontent.com/rafinhacuri/api-expo-go/main/docker-compose.yaml
```
</details>

<details>
<summary>🔽 Using wget</summary>

```bash
wget -O docker-compose.yaml https://raw.githubusercontent.com/rafinhacuri/api-expo-go/main/docker-compose.yaml
```
</details>

Alternatively, copy it directly from the [example file](https://github.com/rafinhacuri/api-expo-go/blob/main/docker-compose.yaml).

---

### 📝 Step 3 — Configure Environment
Download and prepare your `.env` file:

<details>
<summary>🔽 Using curl</summary>

```bash
curl -L -o .env https://raw.githubusercontent.com/rafinhacuri/api-expo-go/main/.env.example
```
</details>

<details>
<summary>🔽 Using wget</summary>

```bash
wget -O .env https://raw.githubusercontent.com/rafinhacuri/api-expo-go/main/.env.example
```
</details>

Then **edit the `.env`** with your database, ports, and other settings.

---

### 🚀 Step 4 — Launch Services
Run the following commands:

```bash
docker compose pull
docker compose up -d --force-recreate
```

---

### 🔍 Step 5 — Verify Installation
Check running containers:

```bash
docker compose ps
```

If all services show as `Up`, you’re ready! 🎉

---

## 📜 License

> Licensed under the [MIT License](https://github.com/rafinhacuri/api-expo-go/blob/main/LICENSE)  
> © 2025 [Rafael Curi Leonardo](https://github.com/rafinhacuri)  

![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)