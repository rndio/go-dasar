package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Perulangan ke-", i)
	}

	var j int = 0
	for j < 5 {
		fmt.Println("Perulangan ke-", j)
		j++
	}

	for pos, char := range "ABC" {
		fmt.Printf("Character %c di index %d\n", char, pos)
	}

	for {
		if j++; j <= 10 {
			fmt.Println("Perulangan ke-", j)
		} else {
			break
		}

	}
}
