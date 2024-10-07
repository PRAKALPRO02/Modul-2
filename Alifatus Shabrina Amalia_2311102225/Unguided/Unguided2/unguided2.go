package main

import (
	"fmt"
)

func main() {
	var integers [5]int
	var chars [3]rune

	fmt.Print("Masukkan 5 buah data integer (antara 32 hingga 127): ")
	for i := 0; i < 5; i++ {
		fmt.Scan(&integers[i])
	}

	fmt.Print("Masukkan 3 karakter (tanpa spasi): ")
	for i := 0; i < 3; i++ {
		fmt.Scanf("%c", &chars[i])

		fmt.Scanf("%c")
	}

	fmt.Print("Hasil karakter dari data integer: ")
	for _, val := range integers {
		fmt.Printf("%c", val)
	}
	fmt.Println()

	fmt.Print("Karakter setelah karakter input: ")
	for _, char := range chars {
		fmt.Printf("%c", char+1)
	}
	fmt.Println()
}
