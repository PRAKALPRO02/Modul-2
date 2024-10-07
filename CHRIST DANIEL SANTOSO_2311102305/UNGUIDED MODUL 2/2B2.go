package main

import "fmt"

func main() {
	var jumlahBunga int
	fmt.Print("Masukkan jumlah bunga: ")
	fmt.Scan(&jumlahBunga)

	bungaList := make([]string, 0, jumlahBunga)
	total := 0

	for i := 0; i < jumlahBunga; i++ {
		var bunga string
		fmt.Print("Masukkan bunga ", i+1, ": ")
		fmt.Scan(&bunga)

		if bunga == "selesai" {
			break
		}
		bungaList = append(bungaList, bunga)
		total++
	}

	fmt.Print("Pita: ")
	for _, bunga := range bungaList {
		fmt.Print(bunga, " - ")
	}
	fmt.Println()
	fmt.Print("Bunga: ", total)
}
