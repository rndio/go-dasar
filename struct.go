package main

import "fmt"

type Siswa struct {
	nama   string
	nisn   int
	berat  float32
	tinggi float32
}

func main() {
	rendio := new(Siswa)
	rendio.nama = "Rendio"
	rendio.nisn = 123456789
	rendio.berat = 55.5
	rendio.tinggi = 173.5

	fmt.Println("Nama:", rendio.nama)
	fmt.Println("NISN:", rendio.nisn)
	fmt.Println("Berat:", rendio.berat)
	fmt.Println("Tinggi:", rendio.tinggi)
}
