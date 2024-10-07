package main

import "fmt"

func main() {
    var tahun213 int
    fmt.Print("tahun : ")
    fmt.Scan(&tahun213)
    if tahun213%4 == 0 && tahun213%100 != 0 {
        fmt.Print("Kabisat : True")
    } else {
        fmt.Print("Kabisat : False")
    }
}