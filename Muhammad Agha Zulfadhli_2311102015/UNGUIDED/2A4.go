package main

import "fmt"

func main() {
	var celsius float32

	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	reaumur := celsius * 4 / 5
	kelvin := (celsius + 273.15)

	fmt.Printf("Derajat Reaumur: %.f\n", reaumur)
	fmt.Printf("Derajat Fahrenheit: %.f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.f\n", kelvin)
}
