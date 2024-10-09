// Muhammad Ragiel Prastyo
// IF-11-02
// 2311102183
package main

import (
	"fmt"
	"math"
)

func main() {
	var k float64

	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&k)

	fk := math.Pow(4*k+2, 2) / ((4*k + 1) * (4*k + 3))

	fmt.Printf("Nilai f(k) = %.10f\n",fk)
}