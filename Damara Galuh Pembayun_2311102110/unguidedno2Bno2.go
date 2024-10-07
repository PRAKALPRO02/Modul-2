package main
//Damara Galuh Pembayun _2311102110

import "fmt"

func main() {
        var n int
        var bunga string
        var pita string

        fmt.Print("N: ")
        fmt.Scanln(&n)

        for i := 1; i <= n; i++ {
                fmt.Printf("Bunga %d: ", i)
                fmt.Scanln(&bunga)
                if bunga == "SELESAI" {
                        break
                }
                pita += bunga + " - "
        }

        // Hapus spasi dan tanda hubung berlebih di akhir
        pita = pita[:len(pita)-3]

        fmt.Println("Pita:", pita)
}