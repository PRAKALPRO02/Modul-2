package main

import(
	"fmt"
	"math"
)

func main(){
	var luasBola float64
	var volumeBola float64
	var jajari float64
	fmt.Print("Masukkan jari-jari: ")
	fmt.Scanln(&jajari)
	volumeBola = 4.0/3.0 * math.Pi * math.Pow(jajari, 3) 
	luasBola = 4.0 * math.Pi * math.Pow(jajari, 2)

	fmt.Println("Volume Bola = ", volumeBola)
	fmt.Println("Luas Bola = ", luasBola)
}