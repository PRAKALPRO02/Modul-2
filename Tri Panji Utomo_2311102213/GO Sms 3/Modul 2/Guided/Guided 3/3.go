package main

import (
    "fmt"
    "math"
)

func main() {
    var jejari213, vol213, luas_kulit float64
    fmt.Print("jejari = ")
    fmt.Scan(&jejari213)
    vol213 = math.Pi * 4.0 / 3.0 * math.Pow(jejari213, 3)
    luas_kulit = math.Pi * 4 * math.Pow(jejari213, 2)
    fmt.Printf("Bola dengan jejari %v memiliki vol %.4f dan luas kulit %.4f", jejari213, vol213, luas_kulit)
}