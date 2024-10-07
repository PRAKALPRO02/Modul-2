package main

import "fmt"

func main() {
	var angka1, angka2, angka3, angka4, angka5 int
	var char1, char2, char3 byte

	fmt.Scanln(&angka1, &angka2, &angka3, &angka4, &angka5)
	fmt.Scanf("%c%c%c", &char1, &char2, &char3)

	fmt.Printf("%c%c%c%c%c\n", angka1, angka2, angka3, angka4, angka5)
	fmt.Printf("%c%c%c\n", char1+1, char2+1, char3+1)
}
