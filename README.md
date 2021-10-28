# Simple Golang Product API Server
Layanan API untuk menambah, merubah informasi, mengambil data dan menghapus produk.
Implementasi kode terinspirasi oleh pola arsitektur hexagonal yang memisahkan proses manjadi beberapa domain.

### Layanan API Server
-    Autentikasi user
-    CRUD produk dan foto produk

### Inisialisasi API Server
Database yang digunakan `PostgreSQL`, ORM menggunakan pustaka `gorm` dan web framework menggunakan ```echo```. Sebelum memulai pastikan
file ```.env``` sudah disesuaikan dengan konfigurasi yang digunakan. Untuk setiap kali request penambahan,
perubahan informasi, pengambilan data dan penghapusan produk menggunakan autentikasi ``JWT`` dimana ```Token``` didapat dari proses login 
dan user untuk login didapat dari proses registrasi user.

### Menjalankan API Server
-    Menggunakan ***Golang***, cukup jalankan perintah ini di terminal tempat sources di simpan.
     ```console
     go run .
     ```
-    Menggunakan ***Docker***, pastikan anda sudah menginstall docker terlebih dahulu dan jalankan perintah ini di tempat source di simpan:<br>
     - Membuat dan menjalankan container dibelakang layar
     
        ```console
        docker-compose up -d
        ```
     
     - Menon-aktifkan daftar container
     
        ```console
        docker-compose stop
        ```
        
     - Men-aktifkan daftar container
     
        ```console
        docker-compose start
        ```
        
     - Menghapus hapus daftar container
     
        ```console
        docker-compose down
        ```

### Melakukan pengujian API server
Untuk dokumentasi singkat penggunaan API server yang berkaitan dengan nama end point, parameter dan respon data dapat dilihat pada ```Postman-2021-10-28.json``` .
