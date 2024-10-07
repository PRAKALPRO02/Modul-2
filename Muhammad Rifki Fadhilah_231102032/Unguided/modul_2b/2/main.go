package main

import "fmt"

func main() {
	// var n int

	// fmt.Printf("N: ")
	// fmt.Scan(&n)
	// bunga := make([]string, n)

	// for i := 0; i < n; i++ {
	// 	fmt.Print("Bunga ", i+1, ": ")
	// 	fmt.Scan(&bunga[i])
	// }

	// fmt.Print("Pita : ")
	// for i := 0; i < n; i++ {
	// 	fmt.Print(bunga[i], " - ")
	// }

	//Modifikasi
	var n int = 0
	var bunga[20] string

	for {
		fmt.Print("Bunga ", n+1, ": ")
		fmt.Scan(&bunga[n])
		if bunga[n] == "SELESAI" {
			break
		}
		n++
	}

	fmt.Print("Pita :")
	for i := 0; i < n; i++ {
		fmt.Print(bunga[i], " - ")
	}
	fmt.Print("\nBunga: ", n)
}
