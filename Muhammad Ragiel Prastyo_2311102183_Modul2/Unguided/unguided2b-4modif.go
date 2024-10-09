// Muhammad Ragiel Prastyo
// IF-11-02
// 2311102183
package main

import (
	"fmt"
	"math"
)

func main() {
	var jumlah int

	fmt.Print("Masukkan jumlah nilai K yang ingin dihitung: ")
	fmt.Scan(&jumlah)

	for i := 0; i < jumlah; i++ {
		var k int
		fmt.Printf("Masukkan nilai K ke-%d : ", i+1)
		fmt.Scan(&k)

		// Menghitung √2 (dalam konteks ini, nilai K tidak digunakan)
		sqrt2 := math.Sqrt(2)

		fmt.Printf("Nilai akar 2 dari %d adalah %.10f\n", k, sqrt2)
	}
}