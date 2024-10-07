package main

import "fmt"

func main(){
	var beratParsel, totalBiaya, tambahanBiaya, biayaRibuan int
	
	fmt.Print("Berat parsel(gram): ")
	fmt.Scan(&beratParsel)
	
	ribuan := beratParsel / 1000
	ratusan := beratParsel % 1000
	if ratusan < 500{
		tambahanBiaya = ratusan * 15
	}else{
		tambahanBiaya = ratusan * 5
	}
	biayaRibuan = ribuan * 10000

	fmt.Printf("Detail berat: %d kg + %d gr\n", ribuan, ratusan)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaRibuan, tambahanBiaya)

	totalBiaya = biayaRibuan + tambahanBiaya
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
	

}