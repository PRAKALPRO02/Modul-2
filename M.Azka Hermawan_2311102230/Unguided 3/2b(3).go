package main

import "fmt"

func main() {
	var kantong1, kantong2 float64
	var totalBerat float64

	for {

		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		_, err := fmt.Scan(&kantong1, &kantong2)

		if err != nil {
			fmt.Println("Input tidak valid, silakan coba lagi.")

			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		var selisih float64
		if kantong1 > kantong2 {
			selisih = kantong1 - kantong2
		} else {
			selisih = kantong2 - kantong1
		}

		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}

		totalBerat += kantong1 + kantong2

		if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}
