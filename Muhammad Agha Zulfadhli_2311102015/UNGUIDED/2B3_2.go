package main

import "fmt"

func main() {
	var kiri, kanan float32
	for kiri+kanan < 150 {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kiri, &kanan)
		fmt.Println("Sepeda motor pak Andi akan oleng: ", (kiri > kanan+9 || kanan > kiri+9))
	}
	fmt.Println("Proses selesai.")
}
