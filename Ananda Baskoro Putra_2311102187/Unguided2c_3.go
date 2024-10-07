package main

import "fmt"

func main() {
	var bil int
	var prima bool = true
	fmt.Print("Bilangan: ")
	fmt.Scan(&bil)
	fmt.Print("Faktor: ")

	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			fmt.Print(i, " ")
			if !(i == 1 || i == bil) {
				prima = false
			}
		}
	}

	fmt.Println()
	fmt.Print("Prima: ", prima)
}