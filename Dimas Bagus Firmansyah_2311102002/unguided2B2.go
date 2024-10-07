package main

import (
	"fmt"
)

func main() {
	var N int
	var bunga, pita string

	fmt.Print("N: ")
	fmt.Scan(&N)

	if N > 0 {
		for i := 1; i <= N; i++ {
			fmt.Printf("Bunga %d: ", i)
			fmt.Scan(&bunga)
			pita += bunga + " - "
		}
		fmt.Println("Pita:", pita)
	} else {
		fmt.Println("Pita :")
	}
}
