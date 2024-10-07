package main

import "fmt"

func main() {
	year := 0
	fmt.Scan(&year)
	fmt.Println((year%4 == 0 && year%100 != 0) || (year%400 == 0))
}