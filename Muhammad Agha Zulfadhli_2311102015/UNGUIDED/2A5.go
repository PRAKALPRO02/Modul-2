package main

import "fmt"

func main() {
	// 1. input 5 buah byte, output 5 buah karakter DONE
	// 2. input 3 buah rune, output karakter setelahnya
	var c1, c2, c3, c4, c5 byte
	var b1, b2, b3 rune
	fmt.Scanf("%d%d%d%d%d\n", &c1, &c2, &c3, &c4, &c5)
	fmt.Scanf("%c%c%c", &b1, &b2, &b3)
	fmt.Printf("%c%c%c%c%c\n", c1, c2, c3, c4, c5)
	fmt.Printf("%c%c%c", b1+1, b2+1, b3+1)
}
