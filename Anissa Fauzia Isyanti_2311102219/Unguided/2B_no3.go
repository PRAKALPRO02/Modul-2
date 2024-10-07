package main

import (
	"fmt"
)

func main() {
	var kantongKiri, kantongKanan float64

	for {
		fmt.Printf("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantongKiri, &kantongKanan)

		if kantongKiri >= 9 || kantongKanan >= 9 {
			fmt.Printf("Program selesai")
			break
		}
	}
}
