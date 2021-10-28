# Simple Golang Product API Server
Layanan API untuk menambah, merubah informasi, mengambil data dan menghapus produk.
Implementasi kode terinspirasi oleh pola arsitektur hexagonal yang memisahkan proses manjadi beberapa domain.

### Layanan API Server
-    Autentikasi user
-    CRUD produk dan foto produk

### Inisialisasi API Server
Database yang digunakan `PostgreSQL`, ORM menggunakan pustaka `gorm` dan web framework menggunakan ```echo```. Sebelum memulai pastikan
file ```.env``` sudah disesuaikan dengan konfigurasi yang digunakan. Stuktur database sudah otomatis ter-migrasi ke database yang sudah diset di file ```.env``` . Untuk setiap kali request penambahan,
perubahan informasi, pengambilan data dan penghapusan produk menggunakan autentikasi ``JWT`` dimana ```Token``` didapat dari proses login 
dan user untuk login didapat dari proses registrasi user.

### Menjalankan API Server
-    Menggunakan ***Golang***, cukup jalankan perintah ini di terminal tempat sources di simpan.
     ```console
     go run .
     ```
-    Menggunakan ***Docker***, pastikan anda sudah menginstall docker terlebih dahulu dan jalankan perintah dibawah ini. 
     Docker akan mendownload image yang sudah dibuat dengan resource tersebut dan akan membuat container yang langsung siap untuk digunakan.
     - Membuat dan menjalankan container dibelakang layar
     
        ```console
        docker-compose up -d
        ```
        
        API server secara default berjalan di: ```http://localhost:8000/v1```
        
     - Menon-aktifkan daftar container
     
        ```console
        docker-compose stop
        ```
        
     - Men-aktifkan daftar container
     
        ```console
        docker-compose start
        ```
        
     - Menghapus daftar container
     
        ```console
        docker-compose down
        ```
-    ***Build image docker***, Anda juga dapat melakukan build image docker API server dengan menjalankan perintah diterminal pada source.

        ```console
        docker build -t <<nama_image_anda>> .
        ```

### Melakukan pengujian API server
Untuk dokumentasi singkat penggunaan API server yang berkaitan dengan nama end point, parameter dan respon data dapat dilihat pada ```Postman-2021-10-28.json``` .

### Kontak
Widya Ade Bagus - https://www.linkedin.com/in/widya-ade-bagus-3a660716b/
