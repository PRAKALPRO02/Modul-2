// Muhammad Ragiel Prastyo
// IF-11-02
// 2311102183
package main

import "fmt"

func main() {
    var (
        satu, dua, tiga string
        temp string
    )
    fmt.Println("Masukan input string: ")
    fmt.Scanln(&satu)
    fmt.Println("Masukan input string: ")
    fmt.Scanln(&dua)
    fmt.Println("Masukan input string: ")
    fmt.Scanln(&tiga)
    fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
    temp = satu
    satu = dua
    dua = tiga
    tiga = temp
    fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}