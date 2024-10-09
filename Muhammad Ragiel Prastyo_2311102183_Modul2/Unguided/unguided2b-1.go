// Muhammad Ragiel Prastyo
// IF-11-02
// 2311102183
package main

import (
    "fmt"
    "strings"
)

func main() {
    var ujiKimia = [5][4]string{}
    var warna = [4]string{"merah", "kuning", "hijau", "ungu"}
    var beda bool = true
    for i := 0; i < 5; i++ {
        fmt.Print("Percobaan ", i+1, ": ")
        fmt.Scan(&ujiKimia[i][0], &ujiKimia[i][1], &ujiKimia[i][2], &ujiKimia[i][3])
    }

    for i := 0; i < 5; i++ {
        for j := 0; j < 4; j++ {
            beda = strings.ToLower(ujiKimia[i][j]) == warna[j]
            if !beda {
                i = 5
                break
            }
        }
    }
    fmt.Println("Berhasil : ", beda)
}