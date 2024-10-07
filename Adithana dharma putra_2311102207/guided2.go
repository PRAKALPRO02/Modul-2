package main
import "fmt"

func main (){
	//deklarasi variabel
	var tahun int 

	//meminta input
	fmt.Print("tahun: ")
	fmt.Scan(&tahun)

	//hitung input
	Kabisat := tahun % 400 == 0 || tahun % 4 == 0 && tahun % 100 != 0
	fmt.Printf("kabisat: %v", Kabisat)
}