package main

import "fmt"

var (
    N, jumlahBunga int
    namaBunga, pita string
)

func main() {
    fmt.Print("Masukkan jumlah bunga: ")
    fmt.Scanln(&N)

    if N <= 0 {
        fmt.Println("Tidak dapat membuat pita bunga.")
    } else {
        for i := 1; i <= N; i++ {
            fmt.Print("Bunga ", i, ": ")
            fmt.Scanln(&namaBunga)

            if namaBunga == "Selesai" {
                break
            }

            pita += " - " + namaBunga
            jumlahBunga++
        }

        fmt.Println("Pita:", pita)
        fmt.Println("Bunga:", jumlahBunga)
    }
}