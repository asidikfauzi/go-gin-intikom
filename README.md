# Petunjuk Instalasi
## 1. Instalasi Docker
   Pastikan Anda telah menginstal Docker di sistem Anda. Jika belum, Anda dapat mengunduhnya dari situs web resmi Docker: https://www.docker.com/.

## 2. Menjalankan Docker Compose
   Untuk menjalankan aplikasi, gunakan perintah Docker Compose berikut dari direktori proyek:

```bash
docker-compose up -d --build
```
Perintah ini akan membangun dan menjalankan kontainer Docker untuk aplikasi dan database PostgreSQL.

## 3. Membuat Basis Data
   Setelah kontainer berjalan, buat basis data PostgreSQL dengan nama `intikom`.

## 4. Mengisi Basis Data
   Jalankan perintah berikut untuk mengisi basis data dengan tabel-tabel dan data yang diperlukan:

### Untuk Pengguna Mac OS / Linux:

```bash
make migrate
make seed
```
### Untuk Pengguna Windows:

```bash
Copy code
go mod vendor -v
rm -f cmd/migrate/migrate
go build -o cmd/migrate/migrate cmd/migrate/migrate.go
./cmd/migrate/migrate
```

```bash
go mod vendor -v
rm -f cmd/seed/seed
go build -o cmd/seed/seed cmd/seed/seed.go
./cmd/seed/seed
```
## 5. Menjalankan Aplikasi
   Terakhir, jalankan aplikasi dengan perintah:

### Untuk Pengguna Mac OS / Linux:

```bash
make all
```
### Untuk Pengguna Windows:

Copy code
```bash
go mod vendor -v
rm -f cmd/app/app
go build -o cmd/app/app cmd/app/app.go
./cmd/app/app
```
## Jalankan dengan Pengembangan Docker Compose (Live Reload dengan Air Toml)
Jalankan 
```bash
docker compose -f docker-compose.yml -f dev.docker-compose.yml up
```
atau
```bash 
docker compose -f docker-compose.yml -f dev.docker-compose.yml up --build
```
untuk membangun kembali gambar buruh pelabuhan (kompilasi ulang layanan biner).

docker composer akan menjalankan go-gin-intikom dan live reload dengan Air di port :8080, dan pengaturan dari `dev.docker-compose.yml` digabungkan dengan `docker-compose.yml`.

## Penting: Penggunaan Air Toml
Jika Anda menggunakan Air Toml, pastikan untuk mengubah variabel lingkungan DB_HOST menjadi `DB_HOST="host.docker.internal"` dalam file .env. Jika tidak menggunakan Air Toml, biarkan variabel lingkungan DB_HOST tetap `DB_HOST="localhost"`.

Ini akan memastikan bahwa aplikasi dapat berkomunikasi dengan basis data yang berjalan di kontainer Docker dari dalam kontainer tersebut.

## Link Dokumentasi Postman
Untuk dokumentasi API, silakan lihat koleksi Postman di [Link Postman](https://crimson-flare-213787.postman.co/workspace/Team-Workspace~79f6082a-b6ff-401a-8e2d-348a27ef9881/collection/34821541-960f09ce-9bdc-467c-b611-06276471ad24?action=share&creator=34821541&active-environment=34821541-f862945e-f2f1-4218-b80b-2e9625d22140).

Dengan langkah-langkah ini, Anda harus dapat menjalankan dan menggunakan aplikasi Go Gin Intikom dengan lancar. Jangan ragu untuk bertanya jika Anda memiliki pertanyaan lebih lanjut!