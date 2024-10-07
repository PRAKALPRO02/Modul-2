package main

import "fmt"

func main() {
    var nilai float64
    var nilaihuruf213 string

    fmt.Print("Nilai akhir mata kuliah: ")
    fmt.Scanln(&nilai)

    switch {
    case nilai >= 80:
        nilaihuruf213 = "A"
    case nilai >= 72.5:
        nilaihuruf213 = "AB"
    case nilai >= 65:
        nilaihuruf213 = "B"
    case nilai >= 57.5:
        nilaihuruf213 = "BC"
    case nilai >= 50:
        nilaihuruf213 = "C"
    case nilai >= 40:
        nilaihuruf213 = "D"
    default:
        nilaihuruf213 = "E"
    }

    fmt.Printf("Nilai mata kuliah: %s\n", nilaihuruf213)
}