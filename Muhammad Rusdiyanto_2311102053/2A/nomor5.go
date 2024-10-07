package main

import "fmt"

func main() {
	var nums = [5]int{}
	var chars string

	fmt.Scanln(&nums[0], &nums[1], &nums[2], &nums[3], &nums[4])
	fmt.Scanln(&chars)

	for i := 0; i < 5; i++ {
		fmt.Printf("%c", nums[i])
	}
	fmt.Print("\n")
	for i := 0; i < 3; i++ {
		fmt.Printf("%c", byte(chars[i])+1)
	}
}
