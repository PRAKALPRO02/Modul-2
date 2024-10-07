package main

import "fmt"

func main() {
	var v1, v2 float64
	var ulang = true
	for ulang != false {
		fmt.Print("Masukan belanjaan di kedua kantong : ")
		fmt.Scan(&v1, &v2)
		if v1 == 9 || v2 == 9 {
			ulang = false
		}
	}
	fmt.Print("Proses selesai")
}
