package main

import "fmt"

func main() {
	var b, f int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	if b > 1 {
		fmt.Print("Faktor: ")
		for i := 1; i <= b; i++ {
			if b%i == 0 {
				fmt.Printf("%d ", i)
				f++
			}
		}
		fmt.Println()

		if f == 2 {
			fmt.Println("Prima: true")
		} else {
			fmt.Println("Prima: false")
		}
	} else {
		fmt.Println("Input harus lebih dari 1")
	}
}
