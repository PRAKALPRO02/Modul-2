package main

import (
	"fmt"
)

func main() {
	// Variabel untuk menyimpan pita (kumpulan nama bunga) dan jumlah bunga
	var pita string
	var jumlahBunga int

	// Loop untuk terus meminta input sampai pengguna mengetik 'SELESAI'
	for {
		var namaBunga string
		fmt.Print("Masukkan nama bunga : ")
		fmt.Scan(&namaBunga)

		// Jika input adalah 'SELESAI', keluar dari loop
		if namaBunga == "SELESAI" {
			break
		}

		// Jika ini bukan input pertama, tambahkan tanda '-' sebelum menambahkan nama bunga
		if pita != "" {
			pita += " - "
		}

		// Menambahkan nama bunga ke dalam pita
		pita += namaBunga
		jumlahBunga++
	}

	// Menampilkan pita dan jumlah bunga setelah semua input selesai
	fmt.Println("\nIsi pita setelah proses input:")
	fmt.Println(pita)
	fmt.Printf("Jumlah bunga dalam pita: %d\n", jumlahBunga)
}
