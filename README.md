# Vinna

**Vinna** is a modular backend platform written in Go designed to help users build and maintain good habits, organize tasks, and plan their life effectively. The project embraces clean architecture, strong typing, and scalable design using modern Go features such as **generics**, **interfaces**, **prepared statements**, and **modular components**.

---

## Why the name *Vinna*?

The name **Vinna** is derived from the Icelandic verb **“vinna”**, which means **“to work”** or **“to win”**. This dual meaning reflects the purpose of the platform:

- **To Work**: Build consistency and structure into your life with planned habits and tasks.
- **To Win**: Achieve your goals through sustainable daily routines and intentional living.

---

## 📦 Modules Overview

The project follows a modular structure with strong separation of concerns:

- `internal/user`: User creation, validation, and persistence
- `internal/frequency`: Time-based scheduling rules for tasks and habits
- `internal/habit`: (coming soon) Habit tracking logic
- `internal/task`: (coming soon) Task planning and lifecycle
- `pkg/db`: Reusable generic SQL repository with prepared statements and type-safe operations

---

## ✅ Features

- Generic SQL repository (CRUD) using Go 1.18+ generics
- Fully testable domain-driven architecture
- Strict input validation
- Modular packages for habits, tasks, users, and frequency rules
- Clean interface and struct definitions
- Lightweight — no external ORMs like GORM

---

## 🛠 Tech Stack

- Language: Go 1.22+
- Database: PostgreSQL
- Dependencies:
  - `github.com/google/uuid`
  - Standard `database/sql` for persistence
  - Regex and time utilities for input validation

---
