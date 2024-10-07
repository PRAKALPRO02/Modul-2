package main

import "fmt"

func main() {
	var R, F, K, suhu float64

	fmt.Print("Temperatur Celcius : ")
	fmt.Scan(&suhu)

	R = suhu * (4.00 / 5.00)
	F = (9.00/5.00)*suhu + 32
	K = suhu + 273

	fmt.Printf("Derajat Reamur : %.2f\n", R)
	fmt.Printf("TDerajat Fahrenheit : %.2f\n", F)
	fmt.Printf("Derajat Kelvin : %.2f\n", K)

}
