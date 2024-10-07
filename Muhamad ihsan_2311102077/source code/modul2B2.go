package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukan berapa bunga: ")
	fmt.Scan(&n)
	total := 0

	var bunga = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Print("masukan bunga ", i+1, ": ")
		fmt.Scan(&bunga[i])
		total++
		if bunga[i] == "selesai" {
			break
		}
	}

	fmt.Print("pita: ")
	for i := 0; i < n; i++ {
		fmt.Print(bunga[i], " - ")
	}
	fmt.Println(" ")
	fmt.Print("Bunga: ", total)
}
