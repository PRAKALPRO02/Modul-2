package main

import "fmt"

func main() {
	for {
		var tahun int
		fmt.Print("Tahun (masukkan 0 untuk keluar): ")
		fmt.Scan(&tahun)

		if tahun == 0 {
			fmt.Println("Proses selesai.")
			break
		}

		kabisat := (tahun%4 == 0 && (tahun%100 != 0 || tahun%400 == 0))
		fmt.Printf("Kabisat: %t\n", kabisat)
	}
}
