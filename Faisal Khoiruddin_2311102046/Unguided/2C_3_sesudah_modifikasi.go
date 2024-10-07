package main

import (
	"fmt"
	"strings"
)

func main() {
	var factorsnum int
	fmt.Print("Bilangan: ")
	fmt.Scanln(&factorsnum)

	var factors []string
	for i := 1; i <= factorsnum; i++ {
		if factorsnum%i == 0 {
			factors = append(factors, fmt.Sprintf("%d", i))
		}
	}

	fmt.Println("Faktor:", strings.Join(factors, " "))

	isPrime := true
	if factorsnum <= 1 {
		isPrime = false
	} else {
		for i := 2; i*i <= factorsnum; i++ {
			if factorsnum%i == 0 {
				isPrime = false
				break
			}
		}
	}

	fmt.Println("Prima:", isPrime)
}
