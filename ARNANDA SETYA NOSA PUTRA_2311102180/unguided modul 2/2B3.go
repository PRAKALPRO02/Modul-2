package main

import "fmt"

func checkWeight(a, b float32) bool {
	totalWeight := a + b
	return totalWeight > 150 || a < 0 || b < 0
}

func isUnbalanced(a, b float32) bool {
	return a <= b-9.0 || b <= a-9.0
}

func main() {
	var weight1, weight2 float32
	var unbalanced bool

	for {
		fmt.Print("Masukkan berat di kedua kantong: ")
		fmt.Scan(&weight1, &weight2)

		if checkWeight(weight1, weight2) {
			break
		}

		unbalanced = isUnbalanced(weight1, weight2)
		fmt.Println("Sepeda motor pak Andi akan oleng:", unbalanced)
	}

	fmt.Println("Program selesai")
}
