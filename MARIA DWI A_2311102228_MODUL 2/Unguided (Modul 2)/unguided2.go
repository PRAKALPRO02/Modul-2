package main

import (
	"fmt"
	"strings"
)

func main(){

	var bunga string
	var pita string
	var jumlah int

	for {
		fmt.Printf("Bunga %d: ", jumlah+1)
		fmt.Scan(&bunga)


		if strings.ToUpper(bunga) == "SELESAI"{
			break
		}

		if jumlah > 0{
			pita += " - "
		}
		pita += bunga
		jumlah++
	}

	fmt.Printf("\nPita: %s\n", pita)
	fmt.Printf("Bunga: %d\n", jumlah)
}