package main

import "fmt"

func main() {
	var bilfaktor int
	fmt.Print("Bilangan: ")
	fmt.Scan(&bilfaktor)

	fmt.Print("Faktor: ")
	i := 1
	for i <= bilfaktor {
		if bilfaktor%i == 0 {
			fmt.Print(i, " ")
		}
		i++
	}
	fmt.Println()
}
