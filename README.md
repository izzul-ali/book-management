# Book Management API – Golang + Chi

A simple RESTful API for managing books, built using **Go**, **Chi router**, and **in-memory storage**.  

- Singleton pattern for centralized storage
- Dependency Injection for handlers
- DRY principle for reusable response and validation logic
- Thread-safe operations with `sync.RWMutex`

---

## Features

- Create, Read, Update, Delete (CRUD) operations for books
- In-memory storage using a map
- Search books by title & author
- Clean error and success response formatting
- Middleware: logger and method not allowed handling

---

## Requirements

- Go 1.20 or newer

---

## How to Run

1. **Clone the repository:**

```bash
git clone https://github.com/izzul-ali/book-management.git
cd book-management
```

2. **Run the server**

```bash
go run main.go
```

## Result

1. **Find All Book**
<img width="1007" alt="Screenshot 2025-05-07 at 14 03 34" src="https://github.com/user-attachments/assets/4d69020c-ae84-46cd-9bba-06bd1edeb9ba" />
<img width="1020" alt="Screenshot 2025-05-07 at 14 10 37" src="https://github.com/user-attachments/assets/4668fc14-b6da-4b9e-af53-7592f9b9e2c5" />


2. **Find One Book**
<img width="1011" alt="Screenshot 2025-05-07 at 14 03 55" src="https://github.com/user-attachments/assets/828f4b76-c43d-4dea-ba5f-a306a6326d62" />

3. **Create Book**
<img width="1007" alt="Screenshot 2025-05-07 at 14 04 15" src="https://github.com/user-attachments/assets/732b69d4-551c-45fc-b31f-a7d683718527" />

4. **Update Book**
<img width="1007" alt="Screenshot 2025-05-07 at 14 05 11" src="https://github.com/user-attachments/assets/7e655400-8def-40fb-aec1-2d0c80087a1a" />

5. **Delete Book**
<img width="1011" alt="Screenshot 2025-05-07 at 14 05 34" src="https://github.com/user-attachments/assets/3fe8f86d-ac6d-4791-89fe-03c823864217" />
