package main

import "fmt"

func main() {

	// Contoh 1
	var names [4]string
	names[0] = "John"
	names[1] = "Paul"
	names[2] = "George"
	names[3] = "Ringo"
	fmt.Println("Element Array 0:", names[0])

	// Contoh 2
	no_urut1 := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println("Element Array 1:", no_urut1[1])

	// Contoh 3
	var total int32
	score := [5]int32{90, 80, 70, 60, 50}

	for i := 0; i < len(score); i++ {
		total += score[i]
	}

	// for _, value := range score {
	// 	total += value
	// }

	fmt.Println("Total Nilai:", total)
}
