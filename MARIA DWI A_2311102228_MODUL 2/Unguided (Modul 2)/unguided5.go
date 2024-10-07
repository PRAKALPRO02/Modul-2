package main

import (
	"fmt"
	//"os"
)

func hitungBiayaPos(beratGram int) (int, int, int, int){ // int 4 kali akan menunjukkan apa yg direturn 
	beratKg := beratGram / 1000 
	sisaGram := beratGram % 1000


	biayaPerKg := beratKg * 10000

	var biayaSisa int
	if beratKg >= 10{
		biayaSisa = 0
	} else if sisaGram >= 500{
		biayaSisa = sisaGram * 5
	} else {
		biayaSisa = sisaGram * 15
	}

	totalBiaya := biayaPerKg + biayaSisa

	return beratKg, sisaGram, biayaPerKg, totalBiaya
} 

func main(){
	for {
		var berat int
		fmt.Print("Masukkan berat parsel dalam gram (0 untuk keluar) : ")
		_, err := fmt.Scanln(&berat) // _ untuk mengabaikan nilai yang dikembalikan oleh fungsi 
		if err != nil{ // memeriksa apakah error atau tidak 
			fmt.Println("Input tidak valid. Mohon masukkan angka.")
			continue
		}

		if berat == 0{
			fmt.Println("Terima kasih telah menggunakan program ini.")
			break
		}
		if berat < 0{
			fmt.Println("Berat tidak boleh negatif. Silahkan coba lagi")
			continue
		}

		beratKg, sisaGram, biayaPerKg, totalBiaya := hitungBiayaPos(berat)

		fmt.Println("\nHasil Perhitungan: ")
		fmt.Printf("Berat parsel\t: %d\n", berat)
		fmt.Printf("Detail berat\t: %d kg + %d gr\n", beratKg, sisaGram)
		fmt.Printf("Detail Biaya\t: Rp. %d +  Rp. %d\n", biayaPerKg, totalBiaya-biayaPerKg)
		fmt.Printf("Total biaya\t: Rp. %d\n\n", totalBiaya)
	}
}

