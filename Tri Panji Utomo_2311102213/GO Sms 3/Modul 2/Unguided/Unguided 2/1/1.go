package main

import "fmt"

func main() {
        var beratGr213 int

        fmt.Print("Berat parsel (gr): ")
        fmt.Scanln(&beratGr213)

        kg := beratGr213 / 1000
        gr := beratGr213 % 1000

        biayaDasar := kg * 10000

        var biayaTambahan int
        if gr >= 500 {
                biayaTambahan = gr * 5
        } else {
                biayaTambahan = gr * 15
        }

        if kg > 10 {
                biayaTambahan = 0
        }

        totalBiaya := biayaDasar + biayaTambahan

        fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gr)
        fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaDasar, biayaTambahan)
        fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}