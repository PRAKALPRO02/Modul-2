package main

import "fmt"

func main() {
	var kantong1, kantong2 float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantong1, &kantong2)

		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		selisih := kantong1 - kantong2
		if selisih < 0 {
			selisih = -selisih
		}

		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}

		totalBerat := kantong1 + kantong2
		if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}
