package main

import (
	"fmt"
	"math"
)






func main(){
	var r int
	fmt.Printf("\nMasukkan jari-jari bola : ")
	fmt.Scanln(&r)

	volume := (4.0 / 3.0) * math.Pi * math.Pow(float64(r), 3)
	luasPermukaan := 4 * math.Pi * math.Pow(float64(r), 2)

	// menampilkan hasil
	fmt.Printf("\nVolume bola : %.2f\n", volume)
	fmt.Printf("Luas permukaan bola: %.2f\n", luasPermukaan)
	fmt.Println()

}