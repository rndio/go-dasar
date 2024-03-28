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

func (s *Siswa) ToString() {
	fmt.Println("Nama:", s.nama)
	fmt.Println("NISN:", s.nisn)
	fmt.Println("Berat:", s.berat)
	fmt.Println("Tinggi:", s.tinggi)
	fmt.Println("Nama Ayah:", s.ortu.nama_ayah)
	fmt.Println("Nama Ibu:", s.ortu.nama_ibu)
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

	astrid := new(Siswa)
	astrid.nama = "Astrid Sunandar"

	asep.ToString()
	fmt.Println()
	astrid.ToString()
}
