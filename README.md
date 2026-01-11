# Golang Book API

API sederhana untuk manajemen **Categories** dan **Books** menggunakan Golang.

## Authentication

API ini menggunakan **Basic Auth**.

Gunakan kredensial berikut untuk mengakses endpoint:

- **Username:** `admin`
- **Password:** `admin123`

---

## Dokumentasi Postman

Dokumentasi lengkap endpoint bisa diakses di sini:  
👉 https://documenter.getpostman.com/view/21073811/2sBXVfiWdm#ab967fc1-8d24-4e9f-b584-f3cdf63ce1f7

---

## Daftar Route API

### Categories

| Method | Endpoint                    | Deskripsi                      |
| ------ | --------------------------- | ------------------------------ |
| GET    | `/api/categories`           | Mengambil semua kategori       |
| POST   | `/api/categories`           | Menambahkan kategori baru      |
| GET    | `/api/categories/:id`       | Detail kategori berdasarkan ID |
| DELETE | `/api/categories/:id`       | Menghapus kategori             |
| GET    | `/api/categories/:id/books` | Mengambil buku per kategori    |

### Books

| Method | Endpoint         | Deskripsi                  |
| ------ | ---------------- | -------------------------- |
| GET    | `/api/books`     | Mengambil semua buku       |
| POST   | `/api/books`     | Menambahkan buku baru      |
| GET    | `/api/books/:id` | Detail buku berdasarkan ID |
| DELETE | `/api/books/:id` | Menghapus buku             |
