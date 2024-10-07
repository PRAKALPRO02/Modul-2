package main
import "fmt"
import "math"

func main(){
	var jari, luas, volume float64
	fmt.Print("Jejari = ")
	fmt.Scanln(&jari)
	volume = 4*math.Pi*math.Pow(jari,3)/3
	luas = 4*math.Pi*math.Pow(jari,2)
	fmt.Println("Bola dengan jejari ",jari," memiliki volume ",volume," dan luas kulit ",luas)
}