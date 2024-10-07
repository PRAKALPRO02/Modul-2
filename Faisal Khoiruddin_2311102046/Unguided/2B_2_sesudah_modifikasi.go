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

	count := 1

	for {
		fmt.Println("Bunga ", count, ": ")
		scanner.Scan()
		namaBunga := scanner.Text()

		if strings.ToUpper(namaBunga) == "SELESAI" {
			break
		}

		pita += namaBunga

		count++
		if count > 1 {
			pita += " - "
		}
	}
	fmt.Println("Pita: ", pita)
	fmt.Println("Bunga: ", count-1)
}
