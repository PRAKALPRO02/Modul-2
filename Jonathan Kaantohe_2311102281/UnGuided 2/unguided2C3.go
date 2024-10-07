package main

import (
	"fmt"
)

func main() {
	var b int

	fmt.Print("Bilangan: ")
	fmt.Scanln(&b)

	if b <= 1 {
		fmt.Println("Silakan masukkan bilangan bulat b > 1.")
		return
	}

	fmt.Print("Faktor: ")
	isPrima := true
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			if i > 1 && i < b {
				isPrima = false
			}
		}
	}
	fmt.Println()

	fmt.Printf("Prima: %t\n", isPrima)
}