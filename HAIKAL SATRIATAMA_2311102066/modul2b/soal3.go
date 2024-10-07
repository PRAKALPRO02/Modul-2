package main

import "fmt"

func main() {
	var kantonga, kantongb float64
	for {
		fmt.Print("Masukkan berat belanjaan kedua kantong: ")
		fmt.Scan(&kantonga, &kantongb)
		if kantonga >= 9 || kantongb >= 9 {
			fmt.Print("Proses selesai")
			break
		}
	}
}
