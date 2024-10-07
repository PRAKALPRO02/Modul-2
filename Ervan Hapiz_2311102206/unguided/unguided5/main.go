package main

import "fmt"

func main() {
	var berat_kiri, berat_kanan float64
	var kondisi = false
	var kondisi1 = false

	for kondisi == false {
		fmt.Print("Masukan berat belanjaan di kedua kantong : ")
		fmt.Scan(&berat_kiri, &berat_kanan)
		if berat_kiri == 9 || berat_kanan == 9 {
			kondisi = true
		}
	}
	fmt.Print("proses selesai")
	//MODISIKASI
	fmt.Print("\n\nAFTER MODIFIKASI\n\n")
	var oleng bool
	for kondisi1 == false {
		fmt.Print("Masukan berat belanjaan di kedua kantong : ")
		fmt.Scan(&berat_kiri, &berat_kanan)
		if berat_kiri < 0 || berat_kanan < 0 || berat_kiri+berat_kanan > 150 {
			kondisi1 = true
		} else {
			if berat_kiri-berat_kanan >= 9 || berat_kanan-berat_kiri >= 9 {
				oleng = true
				fmt.Println("Sepeda motor pak Andi akan oleng :", oleng)
			} else {
				oleng = false
				fmt.Println("Sepeda motor pak Andi akan oleng :", oleng)
			}
		}

	}
	fmt.Print("proses selesai")

}
