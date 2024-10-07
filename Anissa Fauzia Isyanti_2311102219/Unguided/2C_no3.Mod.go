package main

import (
	"fmt"
)

func main() {
	var b int
	var banyakFaktor int

	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	if b <= 1 {
		fmt.Println("Bilangan harus lebih dari 1.")
		return
	}

	fmt.Print("Faktor: ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Printf("%d ", i)
			banyakFaktor++
		}
	}
	fmt.Print("\nPrima: ", banyakFaktor == 2, "\n")
}
