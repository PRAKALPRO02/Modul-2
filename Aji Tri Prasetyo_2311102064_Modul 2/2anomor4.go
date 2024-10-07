package main

import "fmt"

func main() {
	var celcius, reamur, fahren, kelvin float64
	fmt.Print("Tempratur Celcius : ")
	fmt.Scan(&celcius)

	fahren = celcius*9.0/5.0 + 32
	reamur = celcius * 4.0 / 5.0
	kelvin = 5.0 / 9.0 * (fahren + 459.67)

	fmt.Println("Derajat Reamur: ", reamur)
	fmt.Println("Derajat Fahrenheit: ", fahren)
	fmt.Println("Derajat Kelvin: ", int(kelvin))
}
