package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var pita string
	jumlahBunga := 0

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Bunga %d: ", jumlahBunga+1)
		scanner.Scan()
		flower := scanner.Text()

		if strings.ToUpper(flower) == "SELESAI" {
			break
		}

		if pita == "" {
			pita += flower
		} else {
			pita += " - " + flower
		}

		jumlahBunga++
	}

	fmt.Println("Pita:", pita)

	fmt.Printf("Bunga: %d\n", jumlahBunga)
}
