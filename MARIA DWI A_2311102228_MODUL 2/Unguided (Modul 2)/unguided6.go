package main

import "fmt"

func main(){
	var nam float64 
	var nmk string

	// Meminta input nilai akhir mata kuliah
	for {
	fmt.Print("Nilai akhir mata kuliah: ")
	_, err := fmt.Scanln(&nam)
	if err != nil{
		fmt.Println("Input tidak valid. Mohon masukkan angka.")
		var discard string
		fmt.Scanln(&discard)
		continue
	}

	if nam == 0{
		fmt.Println("Program Selesai")
		break
	}

	// Menentukan nilai huruf berdasarkan nilai numerik
	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "AB"
	} else if nam > 65 {
		nmk = "B"
	} else if nam > 57.5 {
		nmk = "BC"
	} else if nam > 50 {
		nmk = "C"
	} else if nam > 40{
		nmk = "D"
	} else {
		nmk = "E"
	}
	

	// Menampilkan hasil nilai huruf
	fmt.Println("Nilai mata kuliah: ", nmk)
	fmt.Println()

	}
}
