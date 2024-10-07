package main

import "fmt"

func main() {
	var oleng bool
	var kanan, kiri float64
	for {
		oleng = false
		fmt.Print("\nMasukan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&kanan, &kiri)
		if kiri > 0 && kanan > 0{
			if kiri > kanan && kiri-kanan >= 9 {
				oleng = true
			}
			if kanan > kiri && kanan-kiri >= 9 {
				oleng = true
			}
		}
		if kiri+kanan > 150 {
			fmt.Print("Program Selesai.")
			break
		}else {
			fmt.Print("Sepeda motor pak Andi akan oleng: ",oleng)
		}
	}
}

