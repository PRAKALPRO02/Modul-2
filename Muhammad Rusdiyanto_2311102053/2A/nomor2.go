package main

import "fmt"

func main() {
	var tahun int
	fmt.Scanln(&tahun)
	fmt.Printf("Kabisat: %v", tahun%4 == 0)
}
