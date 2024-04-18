# KeynandZ Image Gallery

KeynandZ Image Gallery adalah aplikasi web sederhana yang memungkinkan pengguna untuk melihat dan mengunggah gambar ke galeri. Aplikasi ini menggunakan teknologi Go untuk server backend, Minio sebagai penyimpanan file, dan HTML/CSS/JavaScript untuk tampilan frontend.

## Fitur Utama

- Melihat galeri gambar yang telah diunggah.
- Mengunggah gambar baru ke galeri.

## Prasyarat

Sebelum menjalankan proyek ini, pastikan Anda telah memenuhi prasyarat berikut:

- Go (versi 1.16 atau lebih baru) telah terpasang di komputer Anda.
- Minio server telah terpasang dan dikonfigurasi.
- Go module telah diaktifkan.

## Cara Menggunakan

1. Clone repositori ini ke komputer Anda:

    ```
    git clone https://github.com/Keynandz/image-gallery.git
    ```

2. Masuk ke direktori proyek:

    ```
    cd keynandz-image-gallery
    ```

3. Buat file `.env` dan atur konfigurasi seperti yang dibutuhkan. Contoh isi `.env`:

    ```
    MINIO_ENDPOINT=http://localhost:9000
    MINIO_ACCESS_KEY=minio_access_key
    MINIO_SECRET_KEY=minio_secret_key
    MINIO_BUCKET=image-gallery
    ```

4. Install dependensi Go menggunakan Go module:

    ```
    go mod tidy
    ```

5. Jalankan aplikasi:

    ```
    go run main.go
    ```

6. Buka browser dan akses `http://localhost:9000` untuk melihat galeri gambar.

## INTERFACE
![IMG](https://i.ibb.co/sPP8hWz/image-2024-04-18-102232265.png)
