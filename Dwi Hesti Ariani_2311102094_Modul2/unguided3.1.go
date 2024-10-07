package main

import "fmt"

func main() {
	var kantongKiri, kantongKanan float64

	for {
		fmt.Print("Masukkan berat kantong kiri: ")
		fmt.Scanln(&kantongKiri)

		fmt.Print("Masukkan berat kantong kanan: ")
		fmt.Scanln(&kantongKanan)

		if kantongKiri < 0 || kantongKanan < 0 {
			fmt.Println("Berat kantong tidak boleh negatif!")
			break
		}

		totalBerat := kantongKiri + kantongKanan
		if totalBerat > 150 {
			fmt.Println("Total berat melebihi 150 kg!")
			break
		}

		selisih := kantongKiri - kantongKanan
		if selisih < 0 {
			selisih = -selisih
		}

		fmt.Println("Sepeda motor pak Andi akan oleng :", selisih >= 9)
	}
}

// Dwi Hesti Ariani_2311102094
