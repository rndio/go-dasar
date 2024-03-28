package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func proses(menu string, jlh int) {
	for i := 0; i < jlh; i++ {
		a := time.Duration(rand.Intn(5) * 1e9)
		time.Sleep(a)

		fmt.Printf("%v : %v, waktu yang dibutuhkan : %v \n", menu, i, a)
	}
}

func main() {
	proses("Pesan Jus Jambu", 10)
	proses("Pesan Nasi Kuning", 10)
	proses("Pesan Jus Apel", 5)
	proses("Pesan Bubur Manado", 5)

	fmt.Println("Jumlah Goroutine yang berjalan : ", runtime.NumGoroutine())

	var input string
	fmt.Scanln(&input)
	fmt.Println("selesai")
}
