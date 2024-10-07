package main

import (
	"fmt"
)

func main() {
	var j int
	fmt.Print("Masukan Jumlah : ")
	fmt.Scan(&j)

	if j == 0 {
		fmt.Print("Pita : ")
	} else {
		bunga := make([]string, j)
		for i := 0; i < j; i++ {
			fmt.Printf("Bunga %v : ", i+1)
			fmt.Scan(&bunga[i])
		}
		var gabung string
		for i := 0; i < j; i++ {
			gabung += bunga[i] + " - "
		}
		fmt.Print(gabung)
	}
}
