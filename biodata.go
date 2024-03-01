package main

import (
	"fmt"
	"os"
)

// Struct untuk menyimpan data teman
type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Data teman-teman dalam kelas
var dataTeman = map[int]Teman{
	1:  {"ADITTIYA ASRIL", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	2:  {"DERI FAUZI", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	3:  {"FAIZAL AMRI", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	4:  {"ALIF PUTRA DEWANTARA", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	5:  {"MUHAMMAD AGUNG DWI PRASETIYO", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	6:  {"Ilham Arifin", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	7:  {"M. HALVI RAHMAN", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	8:  {"KHAERUL LATIF", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	9:  {"MUHAMMAD SEMAN", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	10: {"IRFAN ZUHDI ABDILLAH", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	11: {"RIDHO EMIR FAJAR ALAM", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	12: {"David Wahyu Pratomo", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	13: {"DANIEL SIPANGKAR", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	14: {"MUHAMMAD AKLI", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	15: {"SULAS SRI", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	16: {"YUSRIANTO", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	17: {"Muhammad Rifky Lukman", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	18: {"Maulana Rafael Irianto", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	19: {"FIRDA ZUhrotul Ma`wa", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	20: {"AHMAD IHSANUDIN", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	21: {"LINTANG TRISNADI", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	22: {"M FIQRI ANANDA", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	23: {"MUHAMMAD ZHIKRI", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	24: {"GIOVANNI GOVERT", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	25: {"A. HASBI ZULFIKAR", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	26: {"Ilham Dimas Prayudha", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	27: {"LAZARUS", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	28: {"DIAN ARIES ALFATAH", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	29: {"Muhammad Fatih Fahroji", "Surabaya", "Mahasiswa", "Ingin Belajar"},
	30: {"HERU PURNAMA", "Surabaya", "Mahasiswa", "Ingin Belajar"},
	31: {"IKHSAN ADI PUTRA", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	32: {"RIDHO TRI PRASETYO", "Bandung", "Mahasiswa", "Ingin Belajar"},
	33: {"Herisa Pratama Nur Baeti", "Surabaya", "Mahasiswa", ""},
	34: {"ALSANDY MAULANA", "Surabaya", "Mahasiswa", "Ingin Belajar"},
	35: {"REFANDA SETYAGUNA SUTRISNO", "Surabaya", "Mahasiswa", "Ingin Belajar"},
	36: {"SADAM ALFIAN PRADANA", "Surabaya", "Mahasiswa", "Ingin Belajar"},
	37: {"GUSTI MUHAMMAD AULIA NUR SULTHAN", "Surabaya", "Mahasiswa", "Ingin Belajar"},
	38: {"I MADE RADITYA BAYU PANGESTU", "Surabaya", "Mahasiswa", "Ingin Belajar"},
	39: {"Faaza Bil Amri", "Surabaya", "Mahasiswa", "Ingin Belajar"},
	40: {"BAYU AJI PUTRAWIBOWO", "Surabaya", "Mahasiswa", "Ingin Belajar"},
	41: {"Rian Bachtiar Ashidiqy", "Jakarta", "Mahasiswa", "Ingin Belajar"},
	42: {"Apriansyah Rizqi Setiawan", "Bandung", "Mahasiswa", "Ingin Belajar"},
	43: {"JOSUA HALOMOAN SIAHAAN", "Surabaya", "Mahasiswa", "Ingin Belajar"},
	44: {"MUHAMMAD RAIHAN NADY KHAIRULLAH", "Surabaya", "Mahasiswa", "Ingin Belajar"},
	45: {"Raka ArmadiranggaO", "Surabaya", "Mahasiswa", "Ingin Belajar"},
	46: {"Wahyu Syamsul A'lam", "Surabaya", "Mahasiswa", "Ingin Belajar"},
	47: {"Reza Muhammad Akbar", "Surabaya", "Mahasiswa", "Ingin Belajar"},
	48: {"THOMAS HARYO WIBISONO", "Surabaya", "Mahasiswa", "Ingin Belajar"},
	49: {"dimas reynaldi dwi santoso", "Surabaya", "Mahasiswa", "Ingin Belajar"},
	50: {"ICHSAN RAMADHAN MOKODOMPIT", "Surabaya", "Mahasiswa", "Ingin Belajar"},
}

// Function untuk menampilkan data teman berdasarkan nomor absen
func TampilkanData(nomorAbsen int) {
	teman, found := dataTeman[nomorAbsen]
	if found {
		fmt.Println("Data Teman dengan Nomor Absen", nomorAbsen)
		fmt.Println("Nama:", teman.Nama)
		fmt.Println("Alamat:", teman.Alamat)
		fmt.Println("Pekerjaan:", teman.Pekerjaan)
		fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
	} else {
		fmt.Println("Data Teman dengan Nomor Absen", nomorAbsen, "tidak ditemukan.")
	}
}

func main() {
	// Memeriksa jumlah argumen
	if len(os.Args) != 2 {
		fmt.Println("Gunakan: go run biodata.go [nomor_absen]")
		return
	}

	// Mendapatkan nomor absen dari argumen
	nomorAbsen := os.Args[1]

	// Konversi nomor absen ke dalam tipe data integer
	var nomorAbsenInt int
	_, err := fmt.Sscanf(nomorAbsen, "%d", &nomorAbsenInt)
	if err != nil {
		fmt.Println("Nomor absen harus berupa bilangan bulat.")
		return
	}

	// Memanggil fungsi untuk menampilkan data teman
	TampilkanData(nomorAbsenInt)
}
