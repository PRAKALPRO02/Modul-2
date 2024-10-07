package main

import "fmt"

func main() {
	var kantong1, kantong2 float64
	var totalBerat float64

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantong1, &kantong2)

		totalBerat = kantong1 + kantong2

		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Berat tidak boleh negatif.")
			break
		} else if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		} else if totalBerat >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}
	}
}
