# Product Requirement Document (PRD)

## Aplikasi Pelacakan Fruit Transport (TPH ke PKS)

### 1. Ringkasan Produk

Aplikasi sederhana untuk mengelola dan melacak proses pengangkutan Buah Sawit dari Tempat Pengumpulan Hasil (TPH) menuju Pabrik Kelapa Sawit (PKS). Sistem ini menggunakan arsitektur terpisah: Backend berbasis **Go (Golang)** dengan database **PostgreSQL**, serta Frontend berbasis **Flutter** (Mobile untuk Supir, Web untuk Admin PKS).

### 2. Komponen Arsitektur

- **Backend:** Go (Golang) menggunakan framework `Fiber` atau `Gin` dan ORM `GORM`.
- **Database:** PostgreSQL.
- **Frontend Mobile:** Flutter (Android/iOS) untuk input data di TPH oleh Supir.
- **Frontend Web Dashboard:** Flutter Web untuk monitoring data di PKS oleh Admin.
- **Komunikasi:** REST API dengan format pertukaran data JSON.

### 3. Struktur Database (PostgreSQL / GORM)

Tabel utama yang dibutuhkan:

#### Tabel: `transports`

- `id` (Primary Key, UUID atau Auto-Increment)
- `nomor_truk` (VARCHAR, Not Null) - Contoh: "BK 1234 AB"
- `id_tph` (VARCHAR, Not Null) - Contoh: "TPH-A01"
- `berat_estimasi` (NUMERIC/FLOAT, Not Null) - Dalam satuan Ton atau Kg
- `status` (VARCHAR) - Default: "DALAM_PERJALANAN" (Pilihan: "DALAM_PERJALANAN", "DITERIMA")
- `created_at` (TIMESTAMP) - Waktu input di TPH
- `updated_at` (TIMESTAMP) - Waktu pembaruan status di PKS

### 4. Spesifikasi API Endpoint (Backend Go)

Backend harus menyediakan 3 endpoint utama:

1. `POST /api/v1/transports`
   - **Fungsi:** Menyimpan data pengangkutan baru dari TPH (Aplikasi Mobile).
   - **Payload Request (JSON):**
     ```json
     {
       "nomor_truk": "BK 1234 AB",
       "id_tph": "TPH-A01",
       "berat_estimasi": 2.5
     }
     ```
   - **Response (JSON):** `201 Created` beserta data yang tersimpan.

2. `GET /api/v1/transports`
   - **Fungsi:** Mengambil semua data pengangkutan untuk ditampilkan di Web Dashboard.
   - **Response (JSON):** `200 OK` berupa array of objek data pengangkutan (diurutkan dari yang terbaru).

3. `PATCH /api/v1/transports/:id/status`
   - **Fungsi:** Mengubah status ketika truk sampai di PKS menjadi "DITERIMA" (Web Dashboard).
   - **Payload Request (JSON):**
     ```json
     {
       "status": "DITERIMA"
     }
     ```

### 5. Alur Kerja Pengguna (User Flow)

1. **Di Lapangan (TPH):** Supir membuka aplikasi Flutter Mobile, mengisi nomor truk, ID TPH, berat estimasi, lalu menekan "Kirim". Data terkirim ke Backend Go melalui HTTP POST.
2. **Di Kantor (PKS):** Admin membuka Flutter Web Dashboard. Halaman otomatis menampilkan daftar tabel truk yang sedang dalam perjalanan (HTTP GET).
3. **Penerimaan Buah:** Ketika truk sampai di PKS, Admin mengklik tombol "Terima Buah" pada tabel, mengirimkan HTTP PATCH untuk memperbarui status data di database.
