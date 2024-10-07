package main

import "fmt"

func main() {
	var a, b, val float32
	var isTrue bool

	for {
		fmt.Print("masukan berat di kedua kantong: ")
		fmt.Scan(&a, &b)
		val = a + b
		if val > 150 || a < 0 || b < 0 {
			break
		}

		if a <= b-9.0 || b <= a-9.0 {
			isTrue = true
			fmt.Println("Sepeda motor pak andi akan oleng: ", isTrue)

		} else {
			isTrue = false
			fmt.Println("Sepeda motor pak andi akan oleng: ", isTrue)
		}
	}

	fmt.Print("program selesai")
}
