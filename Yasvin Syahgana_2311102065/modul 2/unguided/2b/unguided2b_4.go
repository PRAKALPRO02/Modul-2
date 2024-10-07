package main

import (
        "fmt"
        "math"
)

const (
        // Konstanta untuk faktor dalam rumus
        faktor = 4
)

// hitungAkarDua menghitung hampiran akar kuadrat dari 2 menggunakan rumus yang diberikan
func hitungAkarDua(k int) float64 {
        if k < 0 {
                panic("Nilai k harus bilangan bulat positif")
        }
        var akar2 float64 = 1
        for i := 0; i <= k; i++ {
                akar2 *= (faktor*i + 2) * (faktor*i + 2) / (float64(faktor*i + 1) * float64(faktor*i + 3))
        }
        return akar2
}

func main() {
        var k int

        fmt.Print("Masukkan nilai K (bilangan bulat positif): ")
        _, err := fmt.Scan(&k)
        if err != nil || k < 0 {
                fmt.Println("Input tidak valid. Harap masukkan bilangan bulat positif.")
                return
        }

        akar2Hampiran := hitungAkarDua(k)
        akar2Sebenarnya := math.Sqrt(2)

        fmt.Printf("Nilai akar 2 (hampiran dengan k=%d) = %.10f\n", k, akar2Hampiran)
        fmt.Printf("Nilai akar 2 (sebenarnya)         = %.10f\n", akar2Sebenarnya)
}