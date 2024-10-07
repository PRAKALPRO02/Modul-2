package main

import "fmt"

func main() {
	var pita string
	var banyak int

	fmt.Print("Masukkan jumlah (N) : ")
	fmt.Scan(&banyak)

	// Loop untuk meminta input bunga sebanyak jumlah yang ditentukan
	for i := 0; i < banyak; i++ {
		var bunga string
		fmt.Printf("Bunga %d : ", i+1)
		fmt.Scan(&bunga)

		// Menambahkan nama bunga ke pita dengan pemisah " - "
		if i == banyak-1 {
			pita += bunga // Menghindari tanda " - " di akhir
		} else {
			pita += bunga + " - "
		}
	}

	// Menampilkan hasil
	fmt.Println("Pita : ", pita)
}
