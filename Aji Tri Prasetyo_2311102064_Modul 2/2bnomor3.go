package main

import "fmt"

func main() {
	var beratKantongA, beratKantongB float64

	for {
		fmt.Print("Masukkan berat belanjaan kedua kantong: ")
		fmt.Scan(&beratKantongA, &beratKantongB)

		if beratKantongA >= 9 || beratKantongB >= 9 {
			fmt.Println("Proses selesai")
			break
		}
	}
}
