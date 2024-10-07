package main

import "fmt"

func main() {
	var berat, sisaBerat, biayaBerat, biayaSisaBerat, total int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&berat)
	sisaBerat = berat % 1000
	berat -= sisaBerat
	biayaBerat = berat * 10
	fmt.Printf("Detail berat : %v kg + %v gr\n", berat/1000, sisaBerat)
	if sisaBerat >= 500 {
		biayaSisaBerat = sisaBerat * 5
	} else {
		biayaSisaBerat = sisaBerat * 15
	}
	fmt.Printf("Detail biaya : Rp. %v + Rp. %v\n", biayaBerat, biayaSisaBerat)
	if berat > 10000 {
		total = biayaBerat
	} else {
		total = biayaBerat + biayaSisaBerat
	}
	fmt.Printf("Total biaya : Rp. %v", total)
}