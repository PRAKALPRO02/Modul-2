package main

import "fmt"

func main() {
	var x int

	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)

	fmt.Print("Faktor-faktor dari ", x, ": ")
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
