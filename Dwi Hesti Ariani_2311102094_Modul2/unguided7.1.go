package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&bilangan)

	fmt.Println("Bilangan:", bilangan)
	fmt.Print("Faktor: ")
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	prima := true
	if bilangan <= 1 {
		prima = false
	} else {
		for i := 2; i*i <= bilangan; i++ {
			if bilangan%i == 0 {
				prima = false
				break
			}
		}
	}

	if prima {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}

// Dwi Hesti Ariani_2311102094
