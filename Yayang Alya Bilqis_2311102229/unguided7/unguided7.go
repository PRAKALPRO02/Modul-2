package main

import "fmt"

func findFactors(num int) []int {
	var factors []int
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var number int

	for {
		fmt.Print("Masukkan bilangan (0 untuk berhenti): ")
		fmt.Scan(&number)

		if number == 0 {
			break
		}

		factors := findFactors(number)
		fmt.Println("Faktor:", factors)

		if isPrime(number) {
			fmt.Println("Prima: true")
		} else {
			fmt.Println("Prima: false")
		}
	}
}
