package main

import "fmt"
import "math"

func main() {
//Meminta Input
var jariJari float64
fmt.Print("Masukkan jari-jari bola: ")
fmt.Scan(&jariJari)

//hitung
volume := (4.0 / 3.0) * math.Pi * math.Pow(jariJari, 3)
luasPermukaan := 4 * math.Pi * math.Pow(jariJari, 2)

//Output
fmt.Printf("Volume bola: %.4f\n", volume)
fmt.Printf("Luas permukaan bola: %.4f\n",luasPermukaan)
}