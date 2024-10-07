package main
import "fmt"

func main(){
	var tahun int
	var kabisat bool
	kabisat = false
	fmt.Print("Tahun: ")
	fmt.Scanln(&tahun)
	if tahun%4==0 && tahun%100!=0 {
		kabisat = true
	}
	println("Kabisat: ",kabisat)
}