package main

import "fmt"

func main() {
	var fahrenheit, celcius float64
	fmt.Print("Temperatur Celcius : ")
	fmt.Scan(&celcius)

	fahrenheit = 32 + celcius*9.0/5.0
	fmt.Println("Derajat Fahrenheit ", fahrenheit)

}
