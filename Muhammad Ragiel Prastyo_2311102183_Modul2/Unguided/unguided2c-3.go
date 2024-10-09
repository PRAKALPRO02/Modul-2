// Muhammad Ragiel Prastyo
// IF-11-02
// 2311102183

package main

import "fmt"

func main() {
    var bil int

    fmt.Print("Bilangan: ")
    fmt.Scan(&bil)

    if bil <= 1 {
        fmt.Println("Error: Bilangan harus lebih besar dari 1!")
        return
    }

    fmt.Print("Faktor: ")
    for i := 1; i <= bil; i++ {
        if bil%i == 0 {
            fmt.Print(i, " ")
        }
    }
    fmt.Println()
}