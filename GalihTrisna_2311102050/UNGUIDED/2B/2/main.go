package main

import "fmt"

func main() {

	// # before modification #

	// var n int 
	// fmt.Print("N : ")
	// fmt.Scan(&n)
	// var bunga [20]string
	// for i := 0; i < n; i++ {
	// 	fmt.Print("Bunga ",i+1,": ")
	// 	fmt.Scan(&bunga[i])
	// }
	// fmt.Println("Pita : ")
	// for i := 0; i < n; i++ {
	// 	fmt.Print(bunga[i]," - ")
	// }

	// # after modification #

	var n int
	var bunga [20]string
	n = 0
	for {
		fmt.Print("Bunga ", n+1,": ")
		fmt.Scan(&bunga[n])
		if bunga[n] == "SELESAI" {
			break
		}
		n++
	}
	fmt.Print("Pita : ")
	for i := 0; i < n; i++ {
		fmt.Print(bunga[i], " - ")
	}
	fmt.Println("\nBunga : ", n)
}
