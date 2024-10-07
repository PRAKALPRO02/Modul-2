package main

import "fmt"

func main() {
	var N int
	fmt.Print("N: ")
	fmt.Scanln(&N)
	var pita string
	var bunga string
	var count int

	for {
		fmt.Printf("Bunga %d: ", count+1)
		fmt.Scanln(&bunga)

		if bunga == "Selesai" {
			break
		}

		pita += bunga + " - "
		count++
	}

	// Hapus tanda "-" yang berlebihan di akhir
	pita = pita[:len(pita)-3]

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", count)
}
