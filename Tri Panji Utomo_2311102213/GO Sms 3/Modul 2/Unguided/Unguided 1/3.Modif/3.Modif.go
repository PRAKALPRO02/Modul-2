package main

import "fmt"

var berat_kantong213, berat_kantong2 float32

func main() {
    for {
        fmt.Print("Masukan berat belanjaan di kedua kantong: ")
        fmt.Scan(&berat_kantong213, &berat_kantong2)

        if berat_kantong213 >= 9 || berat_kantong2 >= 9 {
            fmt.Println("Print selesai.")
            break
        }
    }
}
