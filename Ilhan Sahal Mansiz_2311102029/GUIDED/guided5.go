package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel untuk menampung input
	var a, b, c, d, e int
	var char1, char2, char3 byte

	// Membaca input untuk 5 data integer
	fmt.Scanf("%d %d %d %d %d", &a, &b, &c, &d, &e)

	// Membaca input untuk 3 data karakter, mengabaikan newline
	fmt.Scanf("\n%c%c%c", &char1, &char2, &char3)

	// Mencetak hasil konversi ASCII dari integer
	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)

	// Mencetak karakter yang di=berikan
	fmt.Printf("%c%c%c\n", char1+1, char2+1, char3+1)
}
