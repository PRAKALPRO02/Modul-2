package main

import (
	"fmt"
)

func main() {
	var b int

	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	fmt.Printf("Faktor dari %d adalah: ", b)
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
}
