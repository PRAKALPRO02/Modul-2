package main

import "fmt"

func main() {
	var tahun int
	fmt.Printf("Tahun: ")
	fmt.Scan(&tahun)

	cekTahunKabisat := (tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0

	fmt.Println("Kabisat:", cekTahunKabisat)
}
