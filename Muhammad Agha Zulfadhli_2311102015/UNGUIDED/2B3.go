package main

import "fmt"

func main() {
	var kiri, kanan float32
	for kiri < 9 && kanan < 9 {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kiri, &kanan)
	}
	fmt.Println("Proses selesai.")
}
