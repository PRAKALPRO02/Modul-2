package main

import "fmt"

func main() {
	for {
		var number int
		var countFaktor int

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

		// Mencari faktor
		fmt.Print("Faktor :")
		for i := 1; i <= number; i++ {
			if number%i == 0 {
				fmt.Print(" ", i)
				countFaktor++
			}
		}

		// Menentukan apakah bilangan prima
		fmt.Print("\nPrima : ", countFaktor == 2, " \n\n")
	}
}
