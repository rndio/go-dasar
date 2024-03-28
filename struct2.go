package main

import "fmt"

type Siswa struct {
	nama   string
	nisn   int
	berat  float32
	tinggi float32
	ortu   OrangTua
}

type OrangTua struct {
	nama_ayah string
	nama_ibu  string
	usia_ayah int
	usia_ibu  int
}

func main() {
	asep := new(Siswa)
	asep.nama = "Asep Sunandar"
	asep.nisn = 123456789
	asep.berat = 55.5
	asep.tinggi = 173.5

	asep.ortu.nama_ayah = "Budi"
	asep.ortu.nama_ibu = "Siti"
	asep.ortu.usia_ayah = 45
	asep.ortu.usia_ibu = 40

	fmt.Println("Nama:", asep.nama)
	fmt.Println("NISN:", asep.nisn)
	fmt.Println("Berat:", asep.berat)
	fmt.Println("Tinggi:", asep.tinggi)
	fmt.Println("Nama Ayah:", asep.ortu.nama_ayah)
	fmt.Println("Nama Ibu:", asep.ortu.nama_ibu)
}
