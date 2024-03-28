package main

import (
	"fmt"
	"os"
)

func main() {
	var x int = 10
	var y int

	if x < 20 {
		y = 20
		fmt.Println("Nilai x =", x)
		fmt.Println("Nilai y =", y)
	}

	// Penggunaan if dengan statement
	if x = 50; x > y {
		fmt.Println("x > y")
	}

	osName, err := os.Hostname()
	if err == nil {
		fmt.Println("\nOs Name :", osName)
	}
}
