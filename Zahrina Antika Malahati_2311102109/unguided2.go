package main

import "fmt"

func main() {
	var pita string
	var bunga string
	var count int

	fmt.Print("Masukkan jumlah bunga: ")
	fmt.Scanln(&count)

	for i := 1; i <= count || bunga != "SELESAI"; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scanln(&bunga)

		if bunga == "SELESAI" {
			break
		}

		pita += bunga + " - "
	}

	// Hapus tanda "-" terakhir jika ada
	if len(pita) > 0 {
		pita = pita[:len(pita)-3]
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", count-1)
}

//Zahrina Antika Malahati_2311102109
