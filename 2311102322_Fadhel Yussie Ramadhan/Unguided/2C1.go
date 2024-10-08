package main

import "fmt"

func main() {
    var berat int
    var biayagr int
    var biayakg int
    var biayatotal int
    const biayaPerKg = 10000

    fmt.Print("Masukkan berat parsel (dalam gram): ")
    fmt.Scanln(&berat)

    kg := berat / 1000
    Gram := berat % 1000

    biayakg = kg * biayaPerKg

    if kg < 10 {
        if Gram > 500 {
        biayagr += Gram * 15
        } else {
        biayagr += Gram * 5
        }
    }

    biayatotal = biayakg + biayagr

    fmt.Println("Detail Berat :", kg,"kg +", Gram,"gr" )
    fmt.Println("Detail Biaya : Rp.",biayakg, "+", biayagr )
    fmt.Println("Total biaya pengiriman: Rp.",biayatotal)
}
