package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var bunga string
	var pita strings.Builder

	fmt.Print("N: ")
	fmt.Scanln(&N)

	for i := 1; i <= N; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scanln(&bunga)
		pita.WriteString(bunga + " - ")
	}

	fmt.Println("Pita:", strings.TrimRight(pita.String(), " - "))
}

// Dwi Hesti Ariani_2311102094
