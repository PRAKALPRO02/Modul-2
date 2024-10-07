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
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

	// Cek bilangan prima
	prima := len(faktor) == 2
	fmt.Printf("Prima: %t\n", prima)
}
