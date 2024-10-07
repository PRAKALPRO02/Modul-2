package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var bunga string
	var pita []string

	// Meminta input jumlah bunga N
	fmt.Print("N: ")
	fmt.Scan(&n)

	// Proses input bunga sebanyak N kali
	for i := 1; i <= n; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)
		// Menyimpan nama bunga ke dalam slice pita
		pita = append(pita, bunga)
	}

	// Menampilkan pita setelah input selesai
	fmt.Println("Pita:", strings.Join(pita, " - "))

	// Modifikasi: Berhenti jika input adalah "SELESAI"
	pita = nil // Reset pita untuk sesi berikutnya
	for {
		fmt.Printf("Bunga %d: ", len(pita)+n+1)
		fmt.Scan(&bunga)

		// Jika user mengetik "SELESAI", proses berhenti
		if strings.ToUpper(bunga) == "SELESAI" {
			break
		}

		// Menyimpan nama bunga ke dalam slice pita
		pita = append(pita, bunga)
	}

	// Menampilkan pita terakhir dan jumlah bunga
	fmt.Println("Pita:", strings.Join(pita, " - "))
	fmt.Println("Total Bunga:", len(pita)+n)
}
