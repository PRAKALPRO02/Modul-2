package main

import "fmt"

func main() {
	var beratKiri, beratKanan float64
	for (beratKiri+beratKanan) <= 150 && beratKiri >= 0 && beratKanan >= 0 {
		fmt.Printf("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&beratKiri, &beratKanan)

		if (beratKiri+beratKanan) > 150 || beratKiri < 0 || beratKanan < 0 {
			break
		}

		sepedaMotorOleng := beratKiri-beratKanan >= 9 || beratKanan-beratKiri >= 9
		fmt.Printf("Sepeda motor pak Andi akan oleng: %v\n", sepedaMotorOleng)
	}
	fmt.Println("Proses selesai.")
}
