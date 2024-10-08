package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	pita := ""
	jumlahBunga := 0

	for {
		fmt.Printf("Bunga %d: ", jumlahBunga+1)
		scanner.Scan()
		namaBunga := scanner.Text()
		if strings.ToUpper(namaBunga) == "SELESAI" {
			break
		}
		if jumlahBunga == 0 {
			pita += namaBunga
		} else {
			pita += " - " + namaBunga
		}
		jumlahBunga++
	}
	if jumlahBunga == 0 {
		fmt.Println("Pita:")
	} else {
		fmt.Println("Pita:", pita)
		fmt.Println("Bunga:", jumlahBunga)
	}
}
