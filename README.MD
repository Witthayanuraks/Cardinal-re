#  ;LKS - Cyber Security Backend (CTF System) [  UNFINISHED ]
# KEEP IT THIS IF WANT CONTRIBUTION 

```
git clone https://github.com/Witthayanuraks/Cardinal-re
```

Backend ini adalah sistem untuk **Capture The Flag (CTF)** berbasis **Golang (Gin Framework)** dan **PostgreSQL**.  
Mendukung **login dengan JWT**, **leaderboard**, **manajemen tim & tantangan**, serta **monitoring sistem**.

---

## 🚀 **Fitur**
- ✅ **Login & Register** dengan hashing password menggunakan **bcrypt**  
- ✅ **Autentikasi JWT** untuk akses API  
- ✅ **Leaderboard** berdasarkan skor tim  
- ✅ **Manajemen Tim** (Menambahkan tim dengan token unik)  
- ✅ **Manajemen Soal/Tantangan** (Soal dengan flag & poin)  
- ✅ **Monitoring Sistem** (CPU, Memori, & Koneksi jaringan)

---

## 📌 **Struktur Direktori**
```
project/
│── internal/
│   ├── handlers/            # Berisi handler API
│   │   ├── auth.go          # Login & register
│   │   ├── leaderboard.go   # Leaderboard
│   │   ├── team.go          # Manajemen tim
│   ├── models/              # Model database
│   │   ├── database.go      # Koneksi database
│   │   ├── user.go          # Model user
│   │   ├── team.go          # Model tim
│   │   ├── challenge.go     # Model tantangan
│   ├── middleware/          # Middleware JWT
│   │   ├── auth_middleware.go 
│── main.go                  # Entry point server
│── .env                     # Environment variables
│── go.mod                    # Dependencies Golang
│── go.sum
│── README.md                 # Dokumentasi
```

---

## ⚙ **Instalasi & Konfigurasi**
### **1️⃣ Clone Repository**
```sh
git clone https://github.com/your-repo/ctf-backend.git
cd ctf-backend
```

### **2️⃣ Setup Database PostgreSQL**
Buat database dan tabel dengan perintah berikut:
```sql
CREATE DATABASE ctf_db;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role VARCHAR(20) CHECK (role IN ('admin', 'user')) NOT NULL
);

CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    token VARCHAR(255) NOT NULL,
    score INT DEFAULT 0
);

CREATE TABLE challenges (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    flag TEXT NOT NULL,
    points INT NOT NULL
);
```

### **3️⃣ Konfigurasi `.env`**
Buat file **.env** dan isi dengan konfigurasi database:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=ctf_db
SECRET_KEY=your_secret_key
```

### **4️⃣ Install Dependency**
Jalankan perintah berikut untuk menginstal semua package yang diperlukan:
```sh
go mod tidy
```

### **5️⃣ Jalankan Server**
```sh
go run main.go
```
API akan berjalan di **http://localhost:8080**

---

## 📌 **Endpoint API**
| Method | Endpoint           | Deskripsi              | Autentikasi |
|--------|-------------------|----------------------|-------------|
| POST   | `/api/login`      | Login user          | ❌ |
| POST   | `/api/register`   | Register user       | ❌ |
| GET    | `/api/leaderboard` | Ambil leaderboard    | ✅ |
| POST   | `/api/add-team`   | Tambah tim baru      | ✅ |
| POST   | `/api/add-challenge` | Tambah soal tantangan | ✅ |

### **Contoh JSON Login**
**Request:**  
```json
{
  "username": "admin",
  "password": "password123"
}
```
**Response (Success):**  
```json
{
  "token": "eyJhbGciOiJIUzI1..."
}
```

---

## 📦 **Library & Dependencies**
Berikut adalah daftar library yang digunakan dalam project ini:
| Library | Kegunaan |
|---------|----------|
| [Gin](https://github.com/gin-gonic/gin) | Web framework untuk Golang |
| [PostgreSQL Driver](https://github.com/lib/pq) | Driver PostgreSQL |
| [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) | Hashing password |
| [JWT-Go](https://github.com/dgrijalva/jwt-go) | Autentikasi dengan JWT |

---

## 💡 **Fitur Selanjutnya**
- 🔜 Reset Password  
- 🔜 Mode Live Monitoring  

📢 Jika ada bug atau ide fitur baru, silakan **open issue** di repo ini.  
🚀 **Enjoy Coding!** 🔥
