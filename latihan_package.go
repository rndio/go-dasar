package main

import (
	"belajar-go/rumus"
	"fmt"
)

func main() {
	hasil := rumus.Penjumlahan(10, 2)

	fmt.Println("Hasil:", hasil)

	fmt.Println("Penjumlahan:", rumus.Penjumlahan(5, 3))
	fmt.Println("Pengurangan:", rumus.Pengurangan(5, 3))
	fmt.Println("Perkalian:", rumus.Perkalian(5, 3))
	fmt.Println("Pembagian:", rumus.Pembagian(5, 3))
}
