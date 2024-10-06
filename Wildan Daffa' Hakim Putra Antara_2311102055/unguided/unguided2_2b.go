package main

import "fmt"

func main() {
	var pita string
	var bunga string
	var banyak int
	fmt.Print("N : ")
	fmt.Scan(&banyak)
	for i := 0; i < banyak; i++ {
		fmt.Print("Bunga ", i+1, " : ")
		fmt.Scan(&bunga)
		pita += bunga + " - "
	}
	fmt.Println("Pita : ", pita)

}
