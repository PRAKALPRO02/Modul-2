package main

import (
	"fmt"
)

func main() {
	var bilangan int

	// Meminta input dari pengguna
	fmt.Print("Bilangan: ")
	fmt.Scanln(&bilangan)

	// Memeriksa apakah bilangan lebih besar dari 1
	if bilangan <= 1 {
		fmt.Println("Silakan masukkan bilangan bulat b > 1.")
		return
	}

	// Mencetak faktor dan memeriksa apakah bilangan prima
	fmt.Print("Faktor: ")
	faktorPrima := true

	for i := 2; i <= bilangan/2; i++ { // Mengurangi iterasi hingga bilangan/2
		if bilangan%i == 0 {
			fmt.Print(i, " ")
			faktorPrima = false
		}
	}
	fmt.Print(bilangan) // Mencetak bilangan itu sendiri sebagai faktor
	fmt.Println()

	// Menentukan apakah bilangan tersebut adalah bilangan prima
	if faktorPrima {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
