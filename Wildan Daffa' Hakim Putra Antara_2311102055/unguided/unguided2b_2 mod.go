package main

import "fmt"

func main() {
	var pita string
	var bunga string
	var count int
	for {
		fmt.Print("Bunga ", count+1, " : ")
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}
		pita += bunga + " - "
		count++
	}
	fmt.Println("Pita : ", pita)
	fmt.Println("Bunga : ", count)

}
