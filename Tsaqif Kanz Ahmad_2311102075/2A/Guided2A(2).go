package main 
import "fmt" 
func main() { 
    var tahun int 

    fmt.Print("Masukkan tahun: ") 
    fmt.Scanln(&tahun) 

    kabisat := (tahun%400 == 0) || (tahun%4 == 0 && 
tahun%100 != 0) 
  
    fmt.Printf("%d adalah tahun kabisat: %t\n", tahun, 
kabisat) 
} 
