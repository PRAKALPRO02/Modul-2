package main

import "fmt"

func main(){
	var celcius, reamur, kelvin, fahrenheit float64

	fmt.Print("Temperatur Celcius: ")
	fmt.Scanln(&celcius)

	fahrenheit = celcius * 9 / 5 + 32
	reamur = celcius / 5 * 4
	kelvin = (fahrenheit + 459.67) / 9 * 5

	fmt.Println("Derajat Reamur: ", reamur)
	fmt.Println("Derajat Fahrenheit: ", fahrenheit)
	fmt.Println("Derajat Kelvin: ", kelvin)
}