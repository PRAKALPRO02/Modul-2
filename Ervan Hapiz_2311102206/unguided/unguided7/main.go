package main

import "fmt"

func main() {
	var berat, kg, gr int
	var biaya_kg, biaya_gr, total int

	fmt.Print("Berat parsel (gram) : ")
	fmt.Scan(&berat)

	kg = berat / 1000
	gr = berat % 1000

	fmt.Printf("Detail berat : %v kg + %v gr\n", kg, gr)

	biaya_kg = kg * 10000
	if gr < 500 {
		biaya_gr = gr * 15
	} else {
		biaya_gr = gr * 5
	}

	if kg > 10 {
		total = biaya_kg
	} else {
		total = biaya_kg + biaya_gr
	}

	fmt.Printf("Detai biaya : RP.%v + RP.%v\n", biaya_kg, biaya_gr)
	fmt.Printf("Total biaya : RP.%v", total)

}
