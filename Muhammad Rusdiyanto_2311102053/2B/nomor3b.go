package main

import "fmt"

func main() {
	for true {
		var kantongKanan, kantongKiri float64
		fmt.Print("Masukkan berat belanjaan kantong : ")
		fmt.Scanln(&kantongKanan, &kantongKiri)
		if kantongKanan+kantongKiri > 150 {
			fmt.Println("Proses selesai.")
			break
		}
		fmt.Printf("Sepeda motor Pak Andi akan oleng : %v\n", kantongKanan-kantongKiri >= 9 || kantongKanan-kantongKiri <= -9)
	}
}
