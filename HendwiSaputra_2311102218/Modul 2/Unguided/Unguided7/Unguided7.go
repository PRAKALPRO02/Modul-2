package main

import (
	"fmt"
	"math"
)

func cariFaktor(n int) []int {
	var faktor []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			faktor = append(faktor, i)
		}
	}
	return faktor
}

func isPrima(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var angka int
	fmt.Print("Bilangan: ")
	fmt.Scan(&angka)

	faktor := cariFaktor(angka)
	fmt.Println("Faktor:", faktor)

	if isPrima(angka) {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
