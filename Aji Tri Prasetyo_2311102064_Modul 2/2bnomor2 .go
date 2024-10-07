package main

import "fmt"

func main() {
	var jumlahBunga int
	var namaBunga, rangkaianBunga string
	fmt.Print("N : ")
	fmt.Scan(&jumlahBunga)
	for i := 0; i < jumlahBunga; i++ {3
		fmt.Print("Bunga ke-", i+1, ": ")
		fmt.Scan(&namaBunga)
		rangkaianBunga += namaBunga + " - "
	}

	fmt.Print("Rangkaian: ", rangkaianBunga)
}
