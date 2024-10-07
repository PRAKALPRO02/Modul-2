package main

import "fmt"

func main() {
	var kantongA, kantongB, total float32
	var oleng bool

	for {
		fmt.Print("Masukkan berat di kedua kantong: ")
		fmt.Scan(&kantongA, &kantongB)
		total = kantongA + kantongB

		if total > 150 || kantongA < 0 || kantongB < 0 {
			break
		}

		if kantongA <= kantongB-9.0 || kantongB <= kantongA-9.0 {
			oleng = true
		} else {
			oleng = false
		}
		fmt.Println("Sepeda motor pak Andi akan oleng: ", oleng)
	}

	fmt.Println("Program selesai")
}
