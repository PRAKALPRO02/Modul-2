package main

import "fmt"

var berat_kantong1, berat_kantong2 float32

func main() {
    for {
        fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
        fmt.Scan(&berat_kantong1, &berat_kantong2)

        if berat_kantong1 < 0 || berat_kantong2 < 0 || berat_kantong1+berat_kantong2 > 150 {
            fmt.Println("Print selesai.")
            break
        }

        selisih := berat_kantong1 - berat_kantong2
        if selisih < 0 {
            selisih = -selisih
        }

        fmt.Printf("Sepeda motor Pak Andi akan oleng: %t\n", selisih >= 9)
    }
}