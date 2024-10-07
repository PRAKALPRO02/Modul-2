package main

import (
	"fmt"
)

func main() {
	var b int
	var count int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&b)

	fmt.Printf("Bilangan: %d\n", b)
	fmt.Print("Faktor: ")

	for f := 1; f <= b; f++ {
		if b%f == 0 {
			fmt.Print(f, " ")
			count++
		}
	}
	fmt.Println()

	if count == 2 {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
