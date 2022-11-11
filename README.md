# GO-CRUD_Mahasiswa

Aplikasi Web CRUD Mahasiswa menggunakan Golang dengan Web-Templatenya.

## Cara Menggunakan Aplikasi

- Pastikan sudah menginstall Go Compiler di komputer anda
- Clone repository ini
- Masuk ke folder `GO-CRUD_Mahasiswa`
- Buka terminal dan jalankan perintah `go run main.go`

## Fitur Yang Tersedia

- `index` : Berisi daftar mahasiswa yang tersimpan
- `tambah mahasiswa` : Digunakan untuk menambahkan data mahasiswa baru, pengguna harus mengisi seluruh form yang tersedia mulai dari `Nama`, `Nim`, `Prodi`, dan `Foto Profil`. Pada kolom `Nim` wajib diisi dengan 12 digit angka, dan pada kolom `Foto Profil` hanya bisa menerima berkas berformat `jpg` atau `png`.
- `edit mahasiswa` : Digunakan untuk mengubah data mahasiswa yang ada, data yang dapat diubah adalah `Nama`, `Prodi`, dan `Foto Profil`.
- `hapus mahasiswa` : Digunakan untuk menghapus data mahasiswa.
