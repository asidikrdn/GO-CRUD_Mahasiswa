package main

import (
	"fmt"
	"html/template"
	"net/http"
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
	fmt.Fprintln(w, "tambah mahasiswa")
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
