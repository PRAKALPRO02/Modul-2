package main

import "fmt"

func main() {
	var bil int
	fmt.Print("Bilangan: ")
	fmt.Scan(&bil)
	fmt.Print("Faktor : ")
	for i := 1; i < bil+1; i++ {
		if bil%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
