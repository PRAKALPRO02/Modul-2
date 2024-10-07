package main

import "fmt"

func main() {
	var v1, v2 float64
	var jatuh = false
	var ulang = true
	for ulang != false {

		fmt.Print("Masukan belanjaan di kedua kantong : ")
		fmt.Scan(&v1, &v2)
		if v1-v2 >= 9 || v2-v1 >= 9 {
			jatuh = true
		} else {
			jatuh = false
		}
		if v1+v2 >= 150 || v1 < 0 || v2 < 0 {
			ulang = false
		} else {
			fmt.Println("Speda motor pak Andi akan oleng : ", jatuh)
		}

	}
	fmt.Print("Proses selesai")
}
