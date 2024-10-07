package main

import (
	"fmt"
)

func main() {
	var E int

	fmt.Print("Banyak bunga : ")
	fmt.Scan(&E)

	bunga := make([]string, E)
	var pita string
	if E <= 0 {
		fmt.Print("Pita : ", pita)
	} else {

		for i := 0; i < E; i++ {
			fmt.Printf("Bunga %v : ", i+1)
			fmt.Scan(&bunga[i])
		}
		for i := 0; i < E; i++ {
			pita += bunga[i] + "-"
		}
		fmt.Printf("Pita : %v  ", pita)
	}
	//MODIFIKASI
	fmt.Print("\n\nAFTER MODIFIKASI\n\n")

	var pita1 string
	var Temp string
	var R = 0
	fmt.Printf("Bunga %v : ", R+1)
	fmt.Scan(&Temp)
	if Temp == "SELESAI" {
		fmt.Printf("Pita : %v  \n", pita1)
		fmt.Println("Bunga :", R)
	} else {
		for Temp != "SELESAI" {

			if Temp != "SELESAI" {
				pita1 += Temp + "-"
			}
			fmt.Printf("Bunga %v : ", R+2)
			fmt.Scan(&Temp)
			R++
		}
		fmt.Printf("Pita : %v  \n", pita1)
		fmt.Println("Bunga :", R)
	}

}
