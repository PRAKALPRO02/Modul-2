package main
import "fmt"


func main() {
var (jari, luas , volume float64) 

fmt.Print ("jejari : ")
fmt.Scanln(&jari)

volume = (4.0/3.0) * (3.1415926535) * jari * jari * jari 
luas = 4 * (3.1415926535) * jari * jari

fmt.Print("Bola dengan jejari " ,jari, " memiliki volume " ,volume, " dan luas kulit ", luas)
}
