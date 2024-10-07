package main

import (
	"fmt"
)

func main() {
	var celsius float64

	// Input from user
	fmt.Print("Masukkan suhu dalam derajat Celsius: ")
	fmt.Scan(&celsius)

	// Konversi suhu
	reamur := celsius * 4 / 5
	fahrenheit := (celsius * 9 / 5) + 32
	kelvin := celsius + 273.15

	// Output hasil konversi
	fmt.Println("Temperatur Celsius:", celsius)
	fmt.Println("Derajat Reamur:", reamur)
	fmt.Println("Derajat Fahrenheit:", fahrenheit)
	fmt.Println("Derajat Kelvin:", kelvin)
}
