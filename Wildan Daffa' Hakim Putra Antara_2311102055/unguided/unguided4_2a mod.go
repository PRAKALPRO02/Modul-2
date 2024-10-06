package main

import "fmt"

func main() {
	var fahrenheit, celcius, reamur, kelvin float64
	fmt.Print("Temperatur Celcius : ")
	fmt.Scan(&celcius)

	fahrenheit = 32 + celcius*9.0/5.0
	reamur = celcius * 4.0 / 5.0
	kelvin = (fahrenheit + 459.67) * 5.0 / 9.0
	fmt.Println("Derajat Reamur", reamur)
	fmt.Println("Derajat Fahrenheit ", fahrenheit)
	fmt.Println("Derajat Kelvin ", int(kelvin))

}
