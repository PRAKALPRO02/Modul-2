package main

import "fmt"

func main() {
	var i int
	var bunga, pita string
	for {
		fmt.Print("Bunga ", i+1, ": ")
		fmt.Scan(&bunga)
		if bunga == "SELESAI" {
			break
		}
		pita = pita + bunga + " - "
		i++
	}
	fmt.Print("Pita: ", pita)
	fmt.Print("Bunga: ", i)
}
