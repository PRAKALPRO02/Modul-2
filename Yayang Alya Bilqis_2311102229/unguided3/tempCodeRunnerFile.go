package main

import "fmt"

func main() {
	var beratKantong1, beratKantong2 float64

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&beratKantong1, &beratKantong2)

		if beratKantong1 >= 9 || beratKantong2 >= 9 || beratKantong1+beratKantong2 > 150 || beratKantong1 < 0 || beratKantong2 < 0 {
			break
		}

		akanOleng := false
		selisihBerat := beratKantong1 - beratKantong2
		if selisihBerat < 0 {
			selisihBerat = -selisihBerat
		}
		if selisihBerat >= 9 {
			akanOleng = true
		}

		fmt.Printf("Sepeda motor pak Andi akan oleng: %t\n", akanOleng)
	}

	fmt.Println("Proses selesai.")
}
