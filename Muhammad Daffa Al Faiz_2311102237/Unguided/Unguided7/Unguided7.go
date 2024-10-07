package main

import "fmt"

func main() {
	var bil int

	fmt.Print("Bilangan: ")
	fmt.Scan(&bil)

	fmt.Print("Faktor: ")
	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			fmt.Printf("%d ", i)
		}
	}

	var isPrime bool
	if bil <= 1 {
		isPrime = false
	} else {
		isPrime = true
		for i := 2; i*i <= bil; i++ {
			if bil%i == 0 {
				isPrime = false
				break
			}
		}
	}

	fmt.Print("\nPrima: ", isPrime)
}
