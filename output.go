package main

import "fmt"

func main() {
	str := "Hello, World!"
	fmt.Printf("String: %s\n", str)

	pi := 3.14159
	fmt.Printf("Pi (float): %f\n", pi)

	isTrue := true
	fmt.Printf("Is True? %t\n", isTrue)

	char := 'A'
	fmt.Printf("Character: %c\n", char)

	num := 10
	fmt.Printf("Binary: %b\n", num)
	fmt.Printf("Octal: %o\n", num)
	fmt.Printf("Lowercase Hex: %x\n", num)
	fmt.Printf("Uppercase Hex: %X\n", num)
}
