package main

import "fmt"

func main() {
        var bil213 int

        fmt.Print("bilangan: ")
        fmt.Scanln(&bil213)

        fmt.Print("Faktor: ")
        for i := 1; i <= bil213; i++ {
                if bil213%i == 0 {
                        fmt.Print(i, " ")
                }
        }
        fmt.Println()

        var jumlahFaktor int = 0
        for i := 1; i <= bil213; i++ {
                if bil213%i == 0 {
                        jumlahFaktor++
                }
        }

        if jumlahFaktor == 2 {
                fmt.Println("Prima: true")
        } else {
                fmt.Println("Prima: false")
        }
}