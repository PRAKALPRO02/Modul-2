package main

import "fmt"

func main() {
	var number int

	fmt.Print("Bilangan: ")
	fmt.Scan(&number)
	if number <= 1 {
		fmt.Println("Bilangan harus lebih dari 1!")
		return
	}
	fmt.Print("Faktor :")
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			fmt.Print(" ", i)
		}

	}
}
