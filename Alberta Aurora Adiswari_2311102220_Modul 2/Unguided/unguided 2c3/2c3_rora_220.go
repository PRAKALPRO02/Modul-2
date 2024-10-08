package main

import "fmt"

func main() {
	var bil_220 int
	var prima bool = true
	fmt.Print("Bilangan: ")
	fmt.Scan(&bil_220)
	fmt.Print("Faktor: ")

	for i := 1; i <= bil_220; i++ {
		if bil_220%i == 0 {
			fmt.Print(i, " ")
			if !(i == 1 || i == bil_220) {
				prima = false
			}
		}
	}

	fmt.Println()
	fmt.Print("Prima: ", prima)
}
