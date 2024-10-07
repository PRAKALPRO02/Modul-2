package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var pita string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("N : ")
	fmt.Scan(&n)

	scanner.Scan()

	for i := 1; i <= n; i++ {
		fmt.Print("Bunga ", i, " : ") 
		scanner.Scan()
		bunga := scanner.Text()

		if i == 1 {
			pita = bunga
		} else {
			pita += " - " + bunga
		}
	}

	fmt.Println("Pita :", pita) 
}