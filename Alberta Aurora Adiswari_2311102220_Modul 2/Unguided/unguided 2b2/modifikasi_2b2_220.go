package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var bunga_220, pita_220 string
	for i := 0; ; i++ {
		n = i
		fmt.Print("Bunga ", i+1, ": ")
		fmt.Scan(&bunga_220)
		if strings.ToLower(bunga_220) == "selesai" {
			break
		}
		pita_220 += bunga_220 + " - "
	}
	fmt.Println("Pita: ", pita_220)
	fmt.Println("Bunga: ", n)
}
