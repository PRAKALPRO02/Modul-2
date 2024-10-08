package main

import "fmt"

func main() {
	var kiri_220, kanan_220 float32
	for kiri_220 < 9 && kanan_220 < 9 {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kiri_220, &kanan_220)
	}
	fmt.Println("Proses selesai.")
}
