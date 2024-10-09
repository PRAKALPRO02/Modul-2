// Muhammad Ragiel Prastyo
// IF-11-02
// 2311102183
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    var pita string
    var TotalBunga int

    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Masukkan nama bunga dan ketik 'SELESAI' untuk mengakhiri. :\n")

    for {
        TotalBunga++
        fmt.Printf("Bunga %d: ", TotalBunga)

        scanner.Scan()
        input := scanner.Text()

        if strings.ToUpper(input) == "SELESAI" {
            TotalBunga-- 
            break
        }

        if pita == "" {
            pita = input
        } else {
            pita += " - " + input
        }
    }

    fmt.Println("Pita:", pita)
    fmt.Println("Total bunga:", TotalBunga )
}