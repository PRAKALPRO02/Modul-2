package main

import "fmt"

func main() {
	for {
		var number int

		fmt.Print("Bilangan (masukkan 0 untuk keluar): ")
		fmt.Scan(&number)

		// Memeriksa apakah pengguna ingin keluar
		if number == 0 {
			fmt.Println("Proses selesai.")
			break
		}

		// Memeriksa apakah bilangan valid
		if number <= 1 {
			fmt.Println("Bilangan harus lebih dari 1!")
			continue
		}

		// Menampilkan faktor-faktor bilangan
		fmt.Print("Faktor :")
		for i := 1; i <= number; i++ {
			if number%i == 0 {
				fmt.Print(" ", i)
			}
		}
		fmt.Println() // Baris baru untuk tampilan yang rapi
	}
}
