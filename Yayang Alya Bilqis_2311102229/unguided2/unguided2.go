package main

import "fmt"

func main() {
	var n int
	var bunga, pita string

	fmt.Print("N: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)
		pita += bunga + " - "
	}

	// Hapus spasi dan tanda "-" berlebih di akhir pita
	pita = pita[:len(pita)-3]

	fmt.Println("Pita:", pita)
}
