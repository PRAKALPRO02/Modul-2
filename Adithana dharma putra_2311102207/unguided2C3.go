package main

import "fmt"
import "math"

// Fungsi untuk mencari faktor
func cariFaktor(n int) []int {
    var faktor []int
    for i := 1; i <= n; i++ {
        if n%i == 0 {
            faktor = append(faktor,i)
        }
    }
    return faktor
}

// Fungsi untuk mengecek
func apakahPrima(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func main() {
	//input
    var angka int
    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&angka)

	//panggil fungsi
    faktor := cariFaktor(angka)
    fmt.Println("Faktor:",faktor)

	//cek prima
    if apakahPrima(angka) {
        fmt.Println("Prima: true")
    } else {
        fmt.Println("Prima: false")}
}