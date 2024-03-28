package main

import (
	"fmt"
	"reflect"
)

func main() {
	println("\n\nVariabel eksplisit\n\n")
	// Variabel eksplisit
	var name string = "Rendio"
	fmt.Println("Nama saya adalah:", name)

	var age int = 20
	fmt.Println("Umur saya:", age)

	var height float32 = 174.2
	fmt.Println("Tinggi saya:", height, "cm")

	var isCool bool = true
	fmt.Println("Cool:", isCool)

	println("\n\nVariabel implisit\n\n")

	// Variabel implisit
	namee := "Rendio"
	fmt.Println("Nama saya adalah:", reflect.TypeOf(namee), namee)

	agee := 20
	fmt.Println("Umur saya:", reflect.TypeOf(agee), agee)

	heightt := 174.2
	fmt.Println("Tinggi saya:", reflect.TypeOf(heightt), heightt, "cm")

	isCooll := true
	fmt.Println("Cool:", reflect.TypeOf(isCooll), isCooll)

}
