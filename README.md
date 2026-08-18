# 🚀 Full-Stack To-Do Planner (Backend API)

A robust, high-performance RESTful API for a To-Do List & Planner application, built with **Go (Golang)** and **Gin Framework**. 

🌍 **Live API Endpoint:** [https://todo-backend-z6fv.onrender.com/api](https://todo-backend-z6fv.onrender.com/api)

## ✨ Features

- **High Performance:** Developed using Go and Gin for maximum efficiency and speed.
- **Secure Authentication:** Implements secure JWT-based authentication for user registration, login, and protected routes.
- **Data Persistence:** Integrated with a **PostgreSQL** database to securely store user and task data.
- **RESTful Endpoints:** Full support for CRUD operations (Create, Read, Update, Delete) to manage categorized tasks.
- **CORS Configured:** Securely configured to accept requests from the frontend client.

## 🛠 Tech Stack

- **Language:** Go (Golang)
- **Web Framework:** Gin (github.com/gin-gonic/gin)
- **Database:** PostgreSQL (using `database/sql` & `lib/pq`)
- **Authentication:** JWT (JSON Web Tokens)
- **Password Hashing:** bcrypt

## 🚀 Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/rabiaozden/todo-backend.git
   ```
2. Set up your `.env` file with PostgreSQL credentials:
   ```env
   PORT=5001
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_user
   DB_PASSWORD=your_password
   DB_NAME=tododb
   JWT_SECRET=your_super_secret_key
   ```
3. Run the application:
   ```bash
   go run cmd/api/main.go
   ```
4. The API will be available at `http://localhost:5001`.

## 🔗 Frontend Repository
The Next.js & Chakra UI frontend for this project can be found here: [Todo Frontend](https://github.com/rabiaozden2/todo-frontend)
