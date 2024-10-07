package main

import "fmt"

func inputBunga() []string {
	var bunga []string
	var namaBunga string

	for i := 1; ; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scanln(&namaBunga)
		if namaBunga == "SELESAI" {
			break
		}
		bunga = append(bunga, namaBunga)
	}
	return bunga
}

func tampilkanPitaDanJumlah(bunga []string) {
	pita := ""
	for i, b := range bunga {
		if i > 0 {
			pita += " - "
		}
		pita += b
	}

	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Jumlah bunga dalam pita: %d\n", len(bunga))
}

func main() {
	bunga := inputBunga()
	tampilkanPitaDanJumlah(bunga)
}
