package main

import "fmt"

func main() {
        var beratKantong1, beratKantong2 float64
        var totalBerat float64

        for {
                fmt.Print("Masukan berat belanjaan di kedua kantong: ")
                _, err := fmt.Scan(&beratKantong1, &beratKantong2)

                if err != nil {
                        fmt.Println("Input tidak valid. Harap masukkan bilangan real positif.")
                        continue
                }

                totalBerat += beratKantong1 + beratKantong2

                if beratKantong1 < 0 || beratKantong2 < 0 {
                        fmt.Println("Berat tidak boleh negatif.")
                        break
                }

                if totalBerat > 150 {
                        fmt.Println("Total berat melebihi 150 kg. Proses selesai.")
                        break
                }

                selisihBerat := beratKantong1 - beratKantong2
                if selisihBerat < 0 {
                        selisihBerat = -selisihBerat
                }

                fmt.Printf("Sepeda motor pak Andi akan oleng: %t\n", selisihBerat >= 9)
        }
}