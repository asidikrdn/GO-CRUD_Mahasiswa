package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

type Mahasiswa struct {
	Nim   string
	Nama  string
	Prodi string
	Pict  string
}

type Data struct {
	DataMahasiswa []Mahasiswa
}

var dataValue = Data{
	DataMahasiswa: []Mahasiswa{
		{"201106041165", "Ahmad", "Teknik Informatika", "201106041165.jpg"},
		{"201106041166", "Sidik", "Sistem Informasi", "201106041166.jpg"},
		{"201106041167", "Rudini", "Teknik Elektro", "201106041167.jpg"},
	},
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", index)
	http.HandleFunc("/tambah", tambahMahasiswa)
	http.HandleFunc("/edit", editMahasiswa)
	http.HandleFunc("/hapus", hapusMahasiswa)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	server := new(http.Server)
	server.Addr = ":4135"

	fmt.Printf("Server running on http://localhost:%s\n", server.Addr)
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	tView := template.Must(template.ParseFiles("views/index.html"))
	tView.Execute(w, dataValue)
}

func tambahMahasiswa(w http.ResponseWriter, r *http.Request) {
	// Tampilan form tambah mahasiswa
	if r.Method == "GET" {
		tView := template.Must(template.ParseFiles("views/tambah_mahasiswa.html"))
		tView.Execute(w, nil)
	}

	// Saat form tambah mahasiswa di-submit
	if r.Method == "POST" {
		// Handling data dari form
		if err := r.ParseMultipartForm(1024); err != nil {
			panic(err.Error())
		}

		// mengambil data dari form
		nim := r.FormValue("nim")
		nama := r.FormValue("nama")
		prodi := r.FormValue("prodi")

		// fmt.Println(nim)

		// Jika nim tidak berisi 12 digit angka, maka tampilkan error
		regexNIM, _ := regexp.Compile(`^\d{12}$`)
		if !regexNIM.MatchString(nim) {
			http.Error(w, "NIM harus berisi 12 digit angka", http.StatusBadRequest)
			return
		}

		// fmt.Println(nama)
		// fmt.Println(prodi)

		// mengambil file dari form
		uploadedFile, handler, err := r.FormFile("pict")
		if err != nil {
			http.Error(w, "Please upload a JPEG or PNG image", http.StatusBadRequest)
			return
		}
		defer uploadedFile.Close()

		// Apabila format file bukan .jpg, maka tampilkan error
		if filepath.Ext(handler.Filename) != ".jpg" && filepath.Ext(handler.Filename) != ".png" {
			http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
			return
		}

		// mengambil direktori aktif
		dir, err := os.Getwd()
		if err != nil {
			panic(err.Error())
		}

		// memberi nama pada file
		filename := fmt.Sprintf("%s%s", nim, filepath.Ext(handler.Filename))
		// fmt.Println(filename)

		// menentukan lokasi file
		fileLocation := filepath.Join(dir, "assets/img", filename)
		// fmt.Println(fileLocation)

		// membuat file baru yang menjadi tempat untuk menampung hasil salinan file upload
		targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err.Error())
		}
		defer targetFile.Close()

		// Menyalin file hasil upload, ke file baru yang menjadi target
		if _, err := io.Copy(targetFile, uploadedFile); err != nil {
			panic(err.Error())
		}

		// Menambahkan data mahasiswa baru ke storage
		newMahasiswa := Mahasiswa{nim, nama, prodi, filename}
		dataValue.DataMahasiswa = append(dataValue.DataMahasiswa, newMahasiswa)

		// Redirect ke halaman index
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func editMahasiswa(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "edit mahasiswa")
}

func hapusMahasiswa(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hapus mahasiswa")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "login")
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "logout")
}
