package main

import (
	"fmt"
	"strings"
)

func main() {
	var bunga, pita string
	var count int

	fmt.Println("Masukkan nama bunga (input 'SELESAI' untuk berhentiin program):")

	for {
		fmt.Printf("Bunga %d: ", count+1)
		fmt.Scan(&bunga)

		
		if strings.ToLower(bunga) == "selesai" {
			break
		}

	
		pita += bunga + " - "
		count++
	}

	if count > 0 {
		fmt.Println("Pita:", pita)
		fmt.Printf("Banyaknya bunga dalam pita: %d\n", count)
	} else {
		fmt.Println("Pita :")
	}
}
