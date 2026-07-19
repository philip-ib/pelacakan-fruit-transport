# 🚛 Pelacakan Fruit Transport

Aplikasi untuk mengelola dan melacak proses pengangkutan Buah Sawit dari **TPH** (Tempat Pengumpulan Hasil) menuju **PKS** (Pabrik Kelapa Sawit).

---

## 🏗️ Arsitektur

```
┌─────────────────────┐     ┌─────────────────────┐
│   📱 Flutter Mobile  │     │  🌐 React Web        │
│   (Supir di TPH)     │     │  (Admin di PKS)      │
│                      │     │                      │
│  • Input nomor truk  │     │  • Lihat tabel data  │
│  • Input ID TPH       │     │  • Tombol Terima     │
│  • Input berat       │     │    Buah              │
│  • Kirim data        │     │                      │
└─────────┬───────────┘     └──────────┬───────────┘
          │ POST                        │ GET / PATCH
          │                             │
          └──────────┬──────────────────┘
                     │
              ┌──────┴──────┐
              │  🚀 Go API  │
              │  Fiber :8080│
              └──────┬──────┘
                     │
              ┌──────┴──────┐
              │  🐘 PostgreSQL│
              └─────────────┘
```

## 🛠️ Teknologi

| Komponen | Teknologi |
|---|---|
| **Backend** | Go 1.21+ • Fiber v2 • GORM |
| **Database** | PostgreSQL 18 |
| **Mobile App** | Flutter 3.44+ • Dart 3.12+ • Material 3 |
| **Web Dashboard** | React 18 • Vite • CSS |
| **Komunikasi** | REST API • JSON |

## 📁 Struktur Folder

```
pelacakan_fruit_transport/
├── README.md
├── PRD.md                          # Product Requirement Document
├── backend/                        # Go REST API
│   ├── main.go                     # Entry point server
│   ├── .env                        # Environment variables
│   ├── go.mod / go.sum
│   ├── config/
│   │   └── config.go               # Load konfigurasi
│   ├── database/
│   │   └── database.go             # Koneksi & migrasi DB
│   ├── model/
│   │   └── transport.go            # Model & DTO
│   ├── repository/
│   │   └── transport_repository.go # Data access layer
│   ├── handler/
│   │   └── transport_handler.go    # HTTP handler
│   └── router/
│       └── router.go               # Routing & CORS
├── frontend_mobile/                # Flutter Mobile (Supir)
│   ├── pubspec.yaml
│   └── lib/
│       ├── main.dart
│       ├── models/transport.dart
│       ├── services/api_service.dart
│       └── screens/supir_form_screen.dart
└── frontend_dashboard/             # React Web (Admin PKS)
    ├── package.json
    ├── vite.config.js
    ├── index.html
    └── src/
        ├── main.jsx
        ├── App.jsx
        ├── App.css
        ├── services/api.js
        └── components/
            ├── Dashboard.jsx
            ├── StatusChip.jsx
            └── ConfirmDialog.jsx
```

## 🚀 Cara Menjalankan

### Prasyarat

- **Go** 1.21+
- **Flutter** 3.44+ dengan Dart 3.12+
- **PostgreSQL** (terinstall & berjalan di port 5432)
- **JDK 17** (untuk build Flutter Android)

---

### 1. Backend (Go API)

```bash
cd backend

# Sesuaikan .env jika diperlukan
# DB_HOST=localhost
# DB_PORT=5432
# DB_USER=postgres
# DB_PASSWORD=admin123
# DB_NAME=pelacakan_fruit
# APP_PORT=8080

# Buat database di PostgreSQL (jika belum ada)
psql -U postgres -c "CREATE DATABASE pelacakan_fruit;"

# Jalankan server
go run main.go
```

Server berjalan di `http://localhost:8080`

---

### 2. Mobile App (Flutter — untuk HP Android)

```bash
cd frontend_mobile

# Install dependencies
flutter pub get

# Cek device tersedia
flutter devices

# Jalankan di HP Android
flutter run -d <device_id>
```

> ⚠️ **Untuk HP fisik:** ubah `baseUrl` di `lib/services/api_service.dart` ke IP laptop Anda, contoh: `http://192.168.1.17:8080`

---

### 3. Web Dashboard (React — untuk Admin PKS)

```bash
cd frontend_dashboard

# Install dependencies
npm install

# Jalankan dev server
npm run dev
```

Buka **Mozilla Firefox** dan akses `http://localhost:5173`

> Vite akan otomatis mem-proxy `/api` ke backend Go di `localhost:8080` dan `/ws` untuk WebSocket.

---

## 📡 API Endpoints

| Method | Endpoint | Deskripsi |
|---|---|---|
| `POST` | `/api/v1/transports` | Simpan data pengangkutan baru |
| `GET` | `/api/v1/transports` | Ambil semua data (terbaru dulu) |
| `PATCH` | `/api/v1/transports/:id/status` | Ubah status transport |

### Contoh Request

**POST** — Input dari Supir
```json
{
  "nomor_truk": "BK 1234 AB",
  "id_tph": "TPH-A01",
  "berat_estimasi": 2.5
}
```

**PATCH** — Terima Buah dari Admin
```json
{
  "status": "DITERIMA"
}
```

---

## 🔧 Environment Variables (Backend)

| Variable | Default | Deskripsi |
|---|---|---|
| `DB_HOST` | `localhost` | Host PostgreSQL |
| `DB_PORT` | `5432` | Port PostgreSQL |
| `DB_USER` | `postgres` | Username database |
| `DB_PASSWORD` | `admin123` | Password database |
| `DB_NAME` | `pelacakan_fruit` | Nama database |
| `DB_SSLMODE` | `disable` | SSL mode |
| `APP_PORT` | `8080` | Port server HTTP |
