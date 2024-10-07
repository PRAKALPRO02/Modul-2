package main

import(
	"fmt"
)

func main(){
	var
	(
		fahrenheit float64
		reamur float64
		celcius float64
		kelvin float64
	)
	fmt.Print("Temperatur Celcius: ")
	fmt.Scanln(&celcius)
	
	fahrenheit = (celcius * 9.0/5.0) + 32
	reamur = 4.0/5.0 * celcius
	kelvin = celcius + 273.15

	fmt.Println("Fahrenheit =",fahrenheit)
	fmt.Println("Reamur =",reamur)
	fmt.Println("Kelvin =",kelvin)
}