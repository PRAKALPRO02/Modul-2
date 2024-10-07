package main

import "fmt"

func main() {
	var angka int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)
	cetakFaktor(angka)
}
func cetakFaktor(n int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
