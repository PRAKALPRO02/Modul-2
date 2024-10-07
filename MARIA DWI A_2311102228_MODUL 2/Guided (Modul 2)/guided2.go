package main
import (
	"fmt"
	
)

func tahunKabisat(tahun int) bool{
	if tahun % 400 == 0{
		return true
	} else if tahun%100 == 0{
		return false
	} else if tahun%4 == 0{
		return true
	}
	return false
}

func main(){
	var year int
	for {

		fmt.Printf("\nMasukkan tahun (0 untuk keluar) : ")
		fmt.Scanln(&year)
		

		if (year) == 0 {
			fmt.Println("\nProgram Selesai!\n")
			break
		}

	if tahunKabisat(year){
		fmt.Println("Kabisat :  true")
	} else {
		fmt.Println("Kabisat:  false")
	}
	
	}
	fmt.Println("\n")
}