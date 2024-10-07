package main
import (
    "fmt"
    "math"
)
func main(){
    var jari float64
    fmt.Print("Jejari = ")
    fmt.Scanln(&jari)
    volume := (4.0 / 3.0) * math.Pi * math.Pow(jari, 3)
    luas := 4 * math.Pi * math.Pow(jari, 2)
    fmt.Printf("Bola dengan jejari %v memiliki volume %.4f dan luas kulit %.4f", jari, volume, luas)
}
