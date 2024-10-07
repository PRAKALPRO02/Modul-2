package main

import (
	"fmt"
	"math"
)

func main(){
	var k float64

	fmt.Print("Masukkan nilai k: ")
	_, err := fmt.Scanf("%f", &k)


	if err != nil{
		fmt.Println("Input tidak valid. Harap masukkan angka")
		return
	}

	hasil := hitungF(k)

	fmt.Printf("f(%.2f) = %.6f\n", k, hasil)
}

func hitungF(k float64) float64{
	pembilang := math.Pow(4*k+2, 2)
	penyebut := (4*k + 1) * (4*k + 3)
	return pembilang / penyebut
}