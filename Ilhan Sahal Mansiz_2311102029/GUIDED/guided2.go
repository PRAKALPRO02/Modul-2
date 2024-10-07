package main

import "fmt"

func isLeapYear(year int) bool {
    if year%400 == 0 {
        return true
    }
    if year%100 == 0 {
        return false
    }
    if year%4 == 0 {
        return true
    }
    return false
}

func main() {
    var year int
    fmt.Print("Masukkan tahun: ")
    fmt.Scan(&year)
    fmt.Println("Tahun :", year)
    if isLeapYear(year) {
        fmt.Println("Kabisat : true")
    } else {
        fmt.Println("Kabisat : false")
    }
}