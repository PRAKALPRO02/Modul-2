package main

import "fmt"

func main() {
	var warna [4]string
	var warna_urutan [4]string
	warna_urutan = [4]string{"merah","kuning","hijau","ungu"}
	var berhasil bool
	berhasil = true
	for i := 0; i < 5; i++ {
		fmt.Print("Percobaan ",i+1,": ")
		for i := 0; i < 4; i++ {
			fmt.Scan(&warna[i])
		}
		for i := 0; i < 4; i++{
			if warna[i] != warna_urutan[i] {
				berhasil = false
			}
		}
	}
	fmt.Println("Berhasil : ",berhasil)
}
