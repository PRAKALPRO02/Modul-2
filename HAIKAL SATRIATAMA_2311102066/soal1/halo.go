package main

import "fmt"

func main() {

	var nilai int
	fmt.Scan(&nilai)
	if nilai >= 90 {
		fmt.Print("A")

	} else if nilai >= 85 && nilai <= 90 {
		fmt.Print("B")
	} else if nilai >= 70 && nilai <= 85 {
		fmt.Print("C")
	} else {
		fmt.Print("D")

	}

}
