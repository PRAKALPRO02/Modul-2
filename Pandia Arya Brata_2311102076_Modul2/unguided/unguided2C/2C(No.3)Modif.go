package main

import "fmt"

func main() {
	var n int
	var prime bool
	fmt.Print("Bilangan: ")
	fmt.Scan(&n)
	fmt.Print("faktor : ")
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(" ", i)
		}
	}
	fmt.Print("\n")
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			prime = false
		} else {
			prime = true
		}
	}
	fmt.Print("prima : ", prime)

}
