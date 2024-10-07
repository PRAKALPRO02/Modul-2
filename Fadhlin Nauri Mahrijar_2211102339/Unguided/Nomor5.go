package main

import "fmt"

func main() {
	var numbers [5]int
	var chars [3]byte

	fmt.Println("Masukkan 5 buah data integer (32-127):")
	for i := 0; i < 5; i++ {
		fmt.Scan(&numbers[i])
	}

	fmt.Scanf("%c%c%c", &chars[0], &chars[1], &chars[2])

	for _, num := range numbers {
		fmt.Printf("%c", num)
	}
	for _, char := range chars {
		fmt.Printf("%c", char+1)
	}
	fmt.Println()
}
