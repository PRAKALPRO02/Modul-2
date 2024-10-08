package main

import "fmt"

func main() {
    var k int
    var result float64 = 1.0 

    fmt.Print("Nilai K = ")
    fmt.Scanln(&k)

    for i := 0; i < k; i++ {
        result = 0.5 * (result + 2/result) 
    }

    fmt.Printf("Nilai akar 2 = %.10f\n", result)
}
