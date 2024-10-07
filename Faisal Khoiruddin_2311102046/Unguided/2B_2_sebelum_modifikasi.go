package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("N: ")
	scanner.Scan()
	N := 0
	fmt.Sscanf(scanner.Text(), "%d", &N)

	pita := ""

	for i := 1; i <= N; i++ {
		fmt.Printf("Bunga %d: ", i)
		scanner.Scan()
		namaBunga := scanner.Text()

		pita += namaBunga

		if i <= N {
			pita += " - "
		}
	}
	fmt.Println("Pita: ", pita)
}
