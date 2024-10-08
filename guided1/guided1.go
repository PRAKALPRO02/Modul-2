package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukkan input string: ")
	fmt.ScanIn(&satu)
	fmt.Print("Masukkan input string: ")
	fmt.ScanIn(&dua)
	fmt.Print("Masukkan input string: ")
	fmt.ScanIn(&tiga)
	fmt.PrintIn("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.PrintIn("Output akhir = " + satu + " " + dua + " " + tiga)
}
