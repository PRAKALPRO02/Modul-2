package main

import "fmt"

func main() {
	var c1, c2, c3, c4, c5, c6, c7, c8 byte

	fmt.Scanf("%d%d%d%d%d", &c1, &c2, &c3, &c4, &c5)
	fmt.Scanf("\n")
	fmt.Scanf("%c%c%c", &c6, &c7, &c8)
	fmt.Printf("%c%c%c%c%c \n", c1, c2, c3, c4, c5)
	fmt.Printf("%c%c%c \n", (c6 + 1), (c7 + 1), (c8 + 1))

}
