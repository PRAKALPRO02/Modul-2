package main

import "fmt"

func main() {
	var N int
	fmt.Print("N: ")
	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Println("Terjadi kesalahan saat membaca input.")
		return
	}

	if N <= 0 {
		fmt.Println("Bilangan N harus positif dan lebih dari nol.")
		return
	}

	pita := ""

	for i := 1; i <= N; i++ {
		var bunga string
		fmt.Printf("Bunga %d: ", i)
		_, err := fmt.Scan(&bunga)
		if err != nil {
			fmt.Println("Terjadi kesalahan saat membaca input bunga.")
			return
		}

		if pita == "" {
			pita = bunga
		} else {
			pita += "-" + bunga
		}
	}

	fmt.Println("Pita:", pita)
}
