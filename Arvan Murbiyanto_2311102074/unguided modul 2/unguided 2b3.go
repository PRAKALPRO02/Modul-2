package main

import "fmt"

func cekOleng(beratKanan, beratKiri float32) bool {
	selisih := beratKanan - beratKiri
	return selisih >= 9 || selisih <= -9
}

func inputBerat() (float32, float32) {
	var beratKanan, beratKiri float32
	fmt.Print("Masukkan berat di kedua kantong: ")
	fmt.Scan(&beratKanan, &beratKiri)
	return beratKanan, beratKiri
}

func main() {
	for {

		beratKanan, beratKiri := inputBerat()
		
		totalBerat := beratKanan + beratKiri
	
		if totalBerat > 150 || beratKanan < 0 || beratKiri < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		isOleng := cekOleng(beratKanan, beratKiri)

	
		fmt.Printf("Sepeda motor Pak Andi akan oleng: %v\n", isOleng)
	}
}
