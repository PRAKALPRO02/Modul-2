package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Bilangan: ")
	fmt.Scan(&bilangan)

	fmt.Print("Faktor: ")
	var faktor []int
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			faktor = append(faktor, i)
		}
	}
	for _, f := range faktor {
		fmt.Printf("%d ", f)
	}
	fmt.Println()

	if len(faktor) == 2 {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}

//Zahrina Antika Malahati_2311102109
