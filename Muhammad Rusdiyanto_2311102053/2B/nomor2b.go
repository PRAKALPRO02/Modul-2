package main

import "fmt"

func main() {
	var flower, flowers string
	var flowerCount int = 0

	for true {
		fmt.Printf("Bunga %v : ", flowerCount+1)
		fmt.Scanln(&flower)
		if flower == "SELESAI" {
			break
		}
		flowers += flower + " - "
		flowerCount++
	}
	fmt.Printf("Pita : %v\nBunga : %v", flowers, flowerCount)
}
