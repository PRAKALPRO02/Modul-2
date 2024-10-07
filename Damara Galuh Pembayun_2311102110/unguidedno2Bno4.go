package main
import "fmt"
//Damara Galuh Pembayun _2311102110
func hitungFungsiFK(K int) float64 { 
	return float64((4*K+2)*(4*K+2)) /
float64((4*K+1)*(4*K+3)) 
} 
func hitungAkar2Hampiran(K int) float64 { 
	akar2 := 1.0
	for x := 0; x <= K; x++ { 
		akar2 *= float64((4*x+2)*(4*x+2)) /
float64((4*x+1)*(4*x+3)) 
  }
  return akar2
} 
func main() { 
	var choice int
	var K int
    fmt.Println("Pilih opsi:") 
    fmt.Println("1. Hitung f(K)") 
    fmt.Println("2. Hitung hampiran akar 2") 
    fmt.Print("Masukkan choice (1/2): ") 
    fmt.Scan(&choice) 
    fmt.Print("Masukkan nilai K: ") 
    fmt.Scan(&K)

    if choice == 1 { 
       fK:= hitungFungsiFK(K) 
       fmt.Printf("Nilai f(K) = %.10f\n", fK) 
    } else if choice == 2 { 
      akar2 := hitungAkar2Hampiran(K) 
      fmt.Printf("Nilai hampiran akar 2 = %.10f\n", akar2) 
} else { 
    fmt.Println("choice tidak valid!") 
} 
}