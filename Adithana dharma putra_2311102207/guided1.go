
package main 
import "fmt"

func main() {
//deklarasi variabel
var (
	satu, dua, tiga string
	temp string
)

//meminta input
fmt.Print("Masukan input string: ")
fmt.Scanln(&satu)

fmt.Print("Masukan input string: ") 
fmt. Scanln(&dua)

fmt.Print("Masukan input string: ")
fmt. Scanln(&tiga)

//melakukan oprasi pertukaran
fmt.Println("Output awal = " + satu + " " + dua + " " + tiga + " ")
temp = satu
satu = dua
dua = tiga
tiga = temp

//Output
fmt.Println("Output akhir = "+ satu + " " + dua + " " + tiga)
}