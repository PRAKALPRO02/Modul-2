package main

import "fmt"

func main() {
	var nilaiN int
	var bunga, pita string
	fmt.Print("N : ")
	fmt.Scan(&nilaiN)
	for i := 0; i < nilaiN; i++ {
		fmt.Print("Bunga ", i+1, ": ")
		fmt.Scan(&bunga)
		pita = pita + bunga + " - "
	}
	fmt.Print("Pita: ", pita)
}
