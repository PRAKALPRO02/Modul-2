package main

import "fmt"

func reamur(cel float64) {
	r := (4.0 / 5.0) * cel
	fmt.Println("Derajat Reamur : ", r)
}

func fahrenheit(cel float64) {
	f := (9.0 / 5.0 * cel) + 32
	fmt.Println("Derajat Fahrenheit : ", f)
}

func kelvin(cel float64) {
	k := cel + 273.15
	fmt.Println("Derajat Kelvin : ", k)
}

func main() {
	var cel float64
	fmt.Print("Drajat Celcius : ")
	fmt.Scan(&cel)
	reamur(cel)
	fahrenheit(cel)
	kelvin(cel)
}
