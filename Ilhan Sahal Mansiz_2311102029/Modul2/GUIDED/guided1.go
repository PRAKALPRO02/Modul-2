package main
import "fmt"

func main() {
    var (
        satu, dua, tiga string
        temp string
    )
    fmt.Print("Masukan input string: ")
    fmt.Scanln(&satu)
    fmt.Print("Masukan input string: ")
    fmt.Scanln(&dua)
    fmt.Print("Masukan input string: ")
    fmt.Scanln(&tiga)
    fmt.Printf("Output awal = %s %s %s\n", satu, dua, tiga)
    temp = satu
    satu = dua
    dua = tiga
    tiga = temp
    fmt.Printf("Output akhir = %s %s %s\n", satu, dua, tiga)
}