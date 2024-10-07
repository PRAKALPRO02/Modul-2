package main
import "fmt"

func main (){
var celsius float64
var fahrenheit float64
var reamur float64
var kelvin  float64

fmt.Print ("Temperatur Celsius: ")
fmt.Scan(&celsius)

reamur = celsius*4/5
fmt.Println("Derajat reamur: ", reamur)

fahrenheit = (celsius * 9/5  ) + 32
fmt.Println("Derajat Fahrenheit: ", fahrenheit)

kelvin = celsius +273
fmt.Println("Derajat kelvin: ", kelvin)
}
