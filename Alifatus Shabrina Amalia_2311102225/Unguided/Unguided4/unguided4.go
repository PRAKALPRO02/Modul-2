package main

import (
	"fmt"
	"strings"
)

func main() {
	var pita string
	var count int

	var N int
	fmt.Print("N: ")
	fmt.Scanln(&N)

	if N > 0 {
		for i := 1; i <= N; i++ {
			var bunga string
			fmt.Printf("Bunga %d: ", i)
			fmt.Scanln(&bunga)
			pita += bunga + " - "
		}
		fmt.Println("Pita:", strings.TrimSpace(pita))
	} else {

		fmt.Println("Pita: ")
	}

	pita = ""
	count = 0

	for {
		var bunga string
		fmt.Print("Bunga: ")
		fmt.Scanln(&bunga)

		if strings.ToUpper(bunga) == "SELESAI" {
			break
		}

		pita += bunga + " - "
		count++
	}

	fmt.Println("Pita:", strings.TrimSpace(pita))
	fmt.Printf("Bunga: %d\n", count)
}
