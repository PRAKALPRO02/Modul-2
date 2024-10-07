package main
import "fmt"

func main(){
    year := 0
    fmt.Print("Tahun: ")
    fmt.Scan(&year)
    fmt.Println("Kabisat: ", year%4 == 0 && year%1000 != 0)
}
