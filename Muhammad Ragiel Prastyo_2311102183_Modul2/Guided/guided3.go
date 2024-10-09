// Muhammad Ragiel Prastyo
// IF-11-02
// 2311102183
package main

import (
        "fmt"
        "math"
)

func main() {
        var jariJari int
        fmt.Print("Jejari = ")
        fmt.Scan(&jariJari)

        volume := (4.0 / 3.0) * math.Pi * float64(jariJari*jariJari*jariJari)
        luas := 4 * math.Pi * float64(jariJari*jariJari)

        fmt.Printf("Bola dengan jari-jari %d memiliki volume %.4f dan luas kulit %.4f\n", jariJari, volume, luas)
}