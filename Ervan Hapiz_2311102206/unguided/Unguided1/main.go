package main

import "fmt"

func Konversi(n float64) {
	var fahrenhait, kelvin, reamur float64

	fahrenhait = n*9.0/5.0 + 32
	kelvin = n + 273.15
	reamur = n * 4.0 / 5.0

	fmt.Println("Derajat Reamur     :", reamur)
	fmt.Println("Derajat fahrenhait :", fahrenhait)
	fmt.Println("Derajat kelvin     :", kelvin)
}

func main() {
	var celcius float64

	fmt.Print("Tempratur Celcius  : ")
	fmt.Scan(&celcius)

	Konversi(celcius)
}
