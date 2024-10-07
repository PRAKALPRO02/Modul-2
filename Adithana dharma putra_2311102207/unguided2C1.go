package main

import "fmt"

func main() {
	//deklarasi variabel
    var beratGram int
    fmt.Print("Masukkan berat parsel (dalam gram): ")
    fmt.Scan(&beratGram)

	//mengekstrak kg dan gram
    beratKg := beratGram / 1000
    sisaGram := beratGram % 1000
    biaya := beratKg * 10000
	var total int
	var biayagram int

	//hitung dan periksa kondisi
    if beratKg > 10 {
        sisaGram = 0
    }

    if sisaGram >= 500 {
        biayagram = sisaGram * 5
    } else {
        biayagram = sisaGram * 15
    }
	total = biayagram+biaya

	//output
	fmt.Printf("Detail berat %v kg + %v gr\n",beratKg,beratGram)
	fmt.Printf("Detail biaya Rp.%v  + Rp.%v\n",biaya,biayagram)
    fmt.Printf("Total biaya pengiriman: Rp %d\n",total)
}