package main

import "fmt"

func main() {
	var beratKiri, beratKanan float64

	for beratKiri < 9 && beratKanan < 9 {
		fmt.Println("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&beratKiri, &beratKanan)

		if beratKiri >= 9 {
			fmt.Println("Proses selesai")
		}
	}
}
