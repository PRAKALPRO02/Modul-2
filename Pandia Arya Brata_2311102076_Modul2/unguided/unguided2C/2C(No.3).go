package main

import "fmt"

func main() {
	var n int

	fmt.Print("Bilangan: ")
	fmt.Scan(&n)
	fmt.Print("faktor : ")
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(" ", i)
		}
	}

}
