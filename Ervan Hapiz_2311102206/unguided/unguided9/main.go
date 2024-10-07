package main

import "fmt"

func main() {
	var bil int

	fmt.Print("Bilangan : ")
	fmt.Scan(&bil)

	fmt.Print("Faktor : ")
	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			fmt.Printf("%d ", i)
		}
	}

	//menetukan bilangan prima
	var prima bool

	for i := 2; i*i <= bil; i++ {
		if bil%i == 0 {
			prima = false
		} else {
			prima = true
		}
	}

	fmt.Print("\nPrima : ", prima)
}
