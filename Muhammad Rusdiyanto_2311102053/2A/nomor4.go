package main

import "fmt"

func main() {
	var celcius float64
	fmt.Print("Temperatur Celcius : ")
	fmt.Scanln(&celcius)
	fmt.Printf("Derajat Reamur : %v\n", (celcius * 4 / 5))
	fmt.Printf("Derajat Fahrenheir : %v\n", (celcius*9/5)+32)
	fmt.Printf("Derajat Kelvin : %v\n", celcius+273.15)
}
