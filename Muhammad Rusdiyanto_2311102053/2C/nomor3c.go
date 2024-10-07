package main

import "fmt"

func main() {
	var num int
	var isPrime bool = true
	fmt.Print("Bilangan : ")
	fmt.Scanln(&num)
	fmt.Print("Faktor : ")
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Printf("%v ", i)
			if i != 1 && i != num {
				isPrime = false
			}
		}
	}
	if num < 2 {
		isPrime = false
	}
	fmt.Printf("\nPrima : %v", isPrime)
}
