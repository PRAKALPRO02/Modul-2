package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var bunga, pita string
	for i := 0; ; i++ {
		n = i
		fmt.Print("Bunga ", i+1, ": ")
		fmt.Scan(&bunga)
		if strings.ToLower(bunga) == "selesai" {
			break
		}
		pita += bunga + " - "
	}
	fmt.Println("Pita: ", pita)
	fmt.Println("Bunga: ", n)
}
