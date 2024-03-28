package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, World!\n")
	scn := bufio.NewScanner(os.Stdin)
	scn.Scan()
	data := scn.Text()
	fmt.Printf("Selamat datang %s!\n", strings.TrimSpace(data))
}
