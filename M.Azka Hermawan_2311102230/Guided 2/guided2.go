package main

import "fmt"

func main() {
	var tahun int

	fmt.Print("Tahun : ")
	fmt.Scan(&tahun)
	fmt.Printf("Kabisat : %t \n", (tahun%4 == 0 && (tahun%100 != 0 || tahun%400 == 0)))
}
