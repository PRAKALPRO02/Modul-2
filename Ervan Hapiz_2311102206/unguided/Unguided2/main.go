package main

import "fmt"

func main() {
	var bil1, bil2, bil3, bil4, bil5 rune
	var char1, char2, char3 rune

	fmt.Scanf("%d%d%d%d%d", &bil1, &bil2, &bil3, &bil4, &bil5)
	fmt.Scanf("\n")
	fmt.Scanf("%c%c%c", &char1, &char2, &char3)
	fmt.Print("\n")
	fmt.Printf("%c%c%c%c%c\n", bil1, bil2, bil3, bil4, bil5)
	fmt.Printf("%c%c%c\n", (char1 + 1), (char2 + 1), (char3 + 1))
}
