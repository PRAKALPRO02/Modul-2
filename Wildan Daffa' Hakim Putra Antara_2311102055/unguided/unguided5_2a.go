package main

import "fmt"

func main() {
	var char1, char2, char3, char4, char5, char6, char7, char8 byte

	fmt.Scanf("%d%d%d%d%d", &char1, &char2, &char3, &char4, &char5)
	fmt.Scanf("%c%c%c", &char6, &char7, &char8)
	fmt.Printf("%c%c%c%c%c \n", char1, char2, char3, char4, char5)
	fmt.Printf("%c%c%c \n", (char6 + 1), (char7 + 1), (char8 + 1))

}
