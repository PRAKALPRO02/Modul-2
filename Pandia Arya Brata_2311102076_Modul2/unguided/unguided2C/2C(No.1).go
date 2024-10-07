package main

import "fmt"

func main() {
	var gr, kg, sisa, harga, harga2, total int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&gr)
	kg = gr / 1000
	sisa = gr % 1000
	fmt.Printf("Detail berat : %d kg + %d gr", kg, sisa)
	fmt.Println("")
	harga = 10000 * kg
	if sisa >= 500 {
		if kg > 10 {
			harga2 = 5 * sisa
			fmt.Printf("Detail biaya: Rp.%d kg + Rp.%d", harga, harga2)
			total = harga
			fmt.Println("")
			fmt.Println("Total biaya : Rp.", total)
		} else {
			harga2 = 5 * sisa
			fmt.Printf("Detail biaya: Rp.%d kg + Rp.%d", harga, harga2)
			total = harga + harga2
			fmt.Println("")
			fmt.Println("Total biaya : Rp.", total)
		}

	} else {
		if kg > 10 {
			harga2 = 5 * sisa
			fmt.Printf("Detail biaya: Rp.%d kg + Rp.%d", harga, harga2)
			total = harga
			fmt.Println("")
			fmt.Println("Total biaya : Rp.", total)
		} else {
			harga2 = 15 * sisa
			fmt.Printf("Detail biaya: Rp.%d kg + Rp.%d", harga, harga2)
			total = harga + harga2
			fmt.Println("")
			fmt.Println("Total biaya : Rp.", total)
		}
	}

}
