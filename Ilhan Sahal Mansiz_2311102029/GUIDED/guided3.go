package main

import (
    "fmt"
    "math"
)

func main() {
    var r float64
    const pi = 3.1415926535

    fmt.Print("Jejari = ")
    fmt.Scan(&r)

    volume := (4.0 / 3.0) * pi * math.Pow(r, 3)

    luas := 4 * pi * math.Pow(r, 2)

    fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}