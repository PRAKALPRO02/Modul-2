package main

import "fmt"

func main() {
	var b int
	var prisma bool
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)
	fmt.Print("Faktor: ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
	if b < 2 {
		prisma = false
	} else {
		prisma = true
	}
	for i := 2; i*i <= b; i++ {
		if b%i == 0 {
			prisma = false
		}
	}
	if prisma == true {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
