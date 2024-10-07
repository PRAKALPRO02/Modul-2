package main

import "fmt"

func main() {
	var nums [5]int
	var chars [3]rune
	for i := 0; i < 5; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Scanln()
	for i := 0; i < 3; i++ {
		fmt.Scanf("%c", &chars[i])
	}
	
	for i := 0; i < 5; i++ {
		fmt.Printf("%c", nums[i])
	}
	fmt.Println()
	for i := 0; i < 3; i++ {
		fmt.Printf("%c", chars[i]+1)
	}
}

