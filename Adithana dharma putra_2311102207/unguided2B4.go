package main

import "fmt"
import "math"

func main() {
    //DEKLARASII VARIABEL
	var kr float64
    fmt.Print("Masukkan nilai K: ")
    fmt.Scan(&kr)

	//MASUKKAN NILAI KE FUNGSI
    f := (math.Pow(4*kr+2, 2)) / ((4*kr + 1) * (4*kr + 3))
    fmt.Printf("Nilai f(K) = %.10f\n",f)

//Modifikasi
fmt.Print("\n\nSETELAH MODIFIKASI\n\n")

//DEKLARASI VARIABEL
		var K int
		fmt.Print("Masukkan nilai K: ")
		fmt.Scan(&K)
	
		//MENGHITUNG NILAI SESUAI RUMUS
		sqrt2 := 1.0
		for k := 0; k <= K; k++ {
			term := math.Pow(4*float64(k)+2, 2) / ((4*float64(k) + 1) * (4*float64(k) + 3))
			sqrt2 *= term
		}
		
	//OUTPUT
		fmt.Printf("Nilai akar 2 = %.10f\n",sqrt2)

}
