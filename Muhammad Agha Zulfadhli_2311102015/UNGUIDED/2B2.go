package main

import "fmt"

func main() {
	var n int
	var bunga, pita string
	fmt.Print("N : ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Print("Bunga ", i+1, ": ")
		fmt.Scan(&bunga)
		pita += bunga + " - "
	}
	fmt.Print("Pita: ", pita)
}
