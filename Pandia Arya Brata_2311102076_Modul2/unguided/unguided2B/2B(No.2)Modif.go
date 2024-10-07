package main

import (
	"fmt"
)

func main() {
	var input string
	var gabung string
	var j = 0
	fmt.Printf("Bunga %v= : ", j+1)
	fmt.Scan(&input)
	if input == "selesai" {
		fmt.Println("Pita : ")
		fmt.Println("Bunga : 0")

	} else {
		for input != "selesai" {
			if input != "selesai" {
				gabung += input + " - "
			}
			fmt.Printf("Bunga %v= : ", j+2)
			fmt.Scan(&input)
			j++
		}

		fmt.Println(gabung)
		fmt.Print("bunga : ", j)
	}

}
