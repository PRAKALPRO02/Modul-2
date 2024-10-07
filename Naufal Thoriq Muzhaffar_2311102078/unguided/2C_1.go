package main
import (
	"fmt"
)

func main() {
	var berat, biaya, biayaTambahan, totalBiaya int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanf("%d", &berat)
	kg := berat / 1000
	gram := berat % 1000
	biaya = kg * 10000
	if gram > 0 {
		if gram >= 500 {
			biayaTambahan = gram * 5
		} else if gram < 500 {
			biayaTambahan = gram * 15
		}
		if kg >= 10 {
			biayaTambahan = 0
		}
	}
	totalBiaya = biaya + biayaTambahan
	fmt.Printf("Berat parsel: %d kg + %d gram\n", kg, gram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biaya, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
