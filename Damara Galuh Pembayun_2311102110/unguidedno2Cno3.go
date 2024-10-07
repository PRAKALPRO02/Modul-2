package main
//Damara Galuh Pembayun _2311102110

import "fmt"

func main() {
        var bilangan int
        var jumlahFaktor int = 0

        fmt.Print("Bilangan: ")
        fmt.Scanln(&bilangan)

        fmt.Print("Faktor: ")
        for i := 1; i <= bilangan; i++ {
                if bilangan%i == 0 {
                        fmt.Printf("%d ", i)
                        jumlahFaktor++
                }
        }
        fmt.Println()

        if jumlahFaktor == 2 {
                fmt.Println("Prima: true")
        } else {
                fmt.Println("Prima: false")
        }
}