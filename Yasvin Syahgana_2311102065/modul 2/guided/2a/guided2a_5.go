package main

import (
	"fmt"
)

func main() {
	// Definisikan array untuk menyimpan input integer dan karakter
	var bilangan [5]int
	var karakter [3]rune

	// Input untuk 5 bilangan bulat
	fmt.Println("Masukkan 5 bilangan bulat (nilai antara 32 - 127):")
	for i := 0; i < 5; i++ {
		fmt.Scan(&bilangan[i])
	}

	// Input untuk 3 karakter tanpa spasi
	fmt.Println("Masukkan 3 karakter (tanpa spasi):")
	for i := 0; i < 3; i++ {
		fmt.Scanf("%c", &karakter[i])
	}

	// Menampilkan karakter hasil konversi dari nilai bilangan bulat (ASCII)
	fmt.Println("\nKeluaran:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%c", bilangan[i])
	}
	fmt.Println()

	// Menampilkan karakter yang diinput oleh pengguna tanpa spasi
	for i := 0; i < 3; i++ {
		fmt.Printf("%c", karakter[i])
	}
	fmt.Println()
}
