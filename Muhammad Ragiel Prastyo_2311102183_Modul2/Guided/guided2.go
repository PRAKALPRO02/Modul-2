// Muhammad Ragiel Prastyo
// IF-11-02
// 2311102183
package main

import "fmt"

func main() {
    var year int
    fmt.Print("Masukkan tahun: ")
    fmt.Scan(&year)

    if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
        fmt.Println("Kabisat:", true)
    } else {
        fmt.Println("Kabisat:", false)
    }
}