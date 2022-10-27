package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//membuat type Mahasiswa dengan struktur

type mahasiswa struct {
	Nim      string
	Nama     string
	Prodi    string
	Semester int
}

//membuat type response dengan struktur

type response struct {
	Status bool
	Pesan  string
	Data   []mahasiswa
}

// connection my sql
func koneksi() (*sql.DB, error) {
	db, salahe := sql.Open("mysql", "root:your pass name@tcp(localhost:3306)/your db name")
	if salahe != nil {
		return nil, salahe
	}
	return db, nil
}

// fungsi tampil data
func tampil(pesane string) response {
	db, salahe := koneksi()
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal Koneksi : " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer db.Close()
	dataMhs, salahe := db.Query("select * from mahasiswa")
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal Query" + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer dataMhs.Close()
	var hasil []mahasiswa
	for dataMhs.Next() {
		var mhs = mahasiswa{}
		var salahe = dataMhs.Scan(&mhs.Nim, &mhs.Nama, &mhs.Prodi, &mhs.Semester)
		if salahe != nil {
			return response{
				Status: false,
				Pesan:  "Gagal Baca : " + salahe.Error(),
				Data:   []mahasiswa{},
			}

		}
		hasil = append(hasil, mhs)
	}

	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Kesalahan : " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	return response{
		Status: true,
		Pesan:  pesane,
		Data:   hasil,
	}

}

// fungsi tampil berdasarkan nim
func getMhs(nim string) response {
	db, salahe := koneksi()
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal Koneksi : " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer db.Close()
	dataMhs, salahe := db.Query("select * from mahasiswa where nim=?", nim)
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal Query" + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer dataMhs.Close()
	var hasil []mahasiswa
	for dataMhs.Next() {
		var mhs = mahasiswa{}
		var salahe = dataMhs.Scan(&mhs.Nim, &mhs.Nama, &mhs.Prodi, &mhs.Semester)
		if salahe != nil {
			return response{
				Status: false,
				Pesan:  "Gagal Baca : " + salahe.Error(),
				Data:   []mahasiswa{},
			}

		}
		hasil = append(hasil, mhs)
	}

	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Kesalahan : " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	return response{
		Status: true,
		Pesan:  "Berhasil tampil",
		Data:   hasil,
	}

}

// fungsi tambah data
func tambah(nim string, nama string, prodi string, semester string) response {
	db, salahe := koneksi()
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal koneksi : " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer db.Close()
	_, salahe = db.Exec("insert into mahasiswa values (?,?,?,?)", nim, nama, prodi, semester)
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal Query Insert : " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	return response{
		Status: true,
		Pesan:  "Berhasil Tambah",
		Data:   []mahasiswa{},
	}
}

// fungsi ubah data
func ubah(nim string, nama string, prodi string, semester string) response {
	db, salahe := koneksi()
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal koneksi : " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer db.Close()
	_, salahe = db.Exec("update mahasiswa set nama=?, prodi=?, semester=? where nim=?", nama, prodi, semester, nim)
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal Query Update : " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	return response{
		Status: true,
		Pesan:  "Berhasil Update",
		Data:   []mahasiswa{},
	}
}

// fungsi hapus
func hapus(nim string) response {
	db, salahe := koneksi()
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal koneksi : " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer db.Close()
	_, salahe = db.Exec("delete from mahasiswa where nim=?", nim)
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal Query Delete : " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	return response{
		Status: true,
		Pesan:  "Berhasil Hapus",
		Data:   []mahasiswa{},
	}
}

//controller

func kontroller(w http.ResponseWriter, r *http.Request) {
	var tampilHtml, salaheTampil = template.ParseFiles("template/tampil.html")
	if salaheTampil != nil {
		fmt.Println(salaheTampil.Error())
		return
	}

	var tambahHtml, salaheTambah = template.ParseFiles("template/tambah.html")
	if salaheTambah != nil {
		fmt.Println(salaheTambah.Error())
		return
	}

	var ubahHtml, salaheUbah = template.ParseFiles("template/ubah.html")
	if salaheUbah != nil {
		fmt.Println(salaheUbah.Error())
		return
	}

	var hapusHtml, salaheHapus = template.ParseFiles("template/hapus.html")
	if salaheHapus != nil {
		fmt.Println(salaheHapus.Error())
		return
	}

	switch r.Method {
	case "GET":
		aksi := r.URL.Query()["aksi"]
		if len(aksi) == 0 {
			tampilHtml.Execute(w, tampil("Berhasil Tampil"))
		} else if aksi[0] == "tambah" {
			tambahHtml.Execute(w, nil)
		} else if aksi[0] == "ubah" {
			nim := r.URL.Query()["nim"]
			ubahHtml.Execute(w, getMhs(nim[0]))
		} else if aksi[0] == "hapus" {
			nim := r.URL.Query()["nim"]
			hapusHtml.Execute(w, getMhs(nim[0]))
		} else {
			tampilHtml.Execute(w, tampil("Berhasil Tampil"))
		}
	case "POST":
		var salahe = r.ParseForm()
		if salahe != nil {
			fmt.Fprintln(w, "Kesalahan : ", salahe)
			return
		}
		var nim = r.FormValue("nim")
		var nama = r.FormValue("nama")
		var prodi = r.FormValue("prodi")
		var semester = r.FormValue("semester")
		var aksi = r.URL.Path
		if aksi == "/tambah" {
			var hasil = tambah(nim, nama, prodi, semester)
			tampilHtml.Execute(w, tampil(hasil.Pesan))
		} else if aksi == "/ubah" {
			var hasil = ubah(nim, nama, prodi, semester)
			tampilHtml.Execute(w, tampil(hasil.Pesan))
		} else if aksi == "/hapus" {
			var hasil = hapus(nim)
			tampilHtml.Execute(w, tampil(hasil.Pesan))
		} else {
			tampilHtml.Execute(w, tampil("Berhasil Tampil"))
		}

	default:
		fmt.Fprint(w, "Maaf. method yang di dukung hanya GET dan POST")
	}
}

func main() {
	http.HandleFunc("/", kontroller)
	fmt.Println("Server berjalan di port 8080")
	http.ListenAndServe(":8080", nil)
}
