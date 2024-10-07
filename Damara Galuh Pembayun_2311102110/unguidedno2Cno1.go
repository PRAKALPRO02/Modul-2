package main
//Damara Galuh Pembayun _2311102110

import "fmt"

func main() {
        var beratParsel int

        fmt.Print("Berat parsel (gram): ")
        fmt.Scanln(&beratParsel)

        // Konversi ke kilogram dan gram
        kg := beratParsel / 1000
        gram := beratParsel % 1000

        // Hitung biaya dasar
        biayaDasar := kg * 10000

        // Hitung biaya tambahan
        var biayaTambahan int
        if gram >= 500 && kg <= 10 {
                biayaTambahan = gram * 15
        } else if gram >= 500 {
                biayaTambahan = gram * 5
        }

        // Hitung total biaya
        totalBiaya := biayaDasar + biayaTambahan

        // Tampilkan hasil
        fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gram)
        fmt.Printf("Detail biaya: Rp %d + Rp %d\n", biayaDasar, biayaTambahan)
        fmt.Printf("Total biaya: Rp %d\n", totalBiaya)
}