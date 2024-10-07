package main
import "fmt"

func main(){
	var celcius, reamur, fahrenheit, kelvin float64
	fmt.Print("Temperatur Celsius: ")
	fmt.Scanln(&celcius)
	reamur = celcius/5*4
	fahrenheit = celcius*9/5+32
	kelvin = (fahrenheit+459.67)/9*5
	fmt.Println("Derajat Reamur: ", reamur)
	fmt.Println("Derajat Fahrenheit: ", fahrenheit)
	fmt.Println("Derajat Kelvin: ", kelvin)
}