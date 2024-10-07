package main

import (
	"fmt"
)

func cariFaktor(b int) []int {
	var faktor []int
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			faktor = append(faktor, i)
		}
	}
	return faktor
}

func cekPrima(b int) bool {
	faktor := cariFaktor(b)

	return len(faktor) == 2
}

func main() {
	var b int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&b)

	faktor := cariFaktor(b)
	fmt.Printf("Faktor: ")
	for _, f := range faktor {
		fmt.Printf("%d ", f)
	}
	fmt.Println()

	if cekPrima(b) {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
