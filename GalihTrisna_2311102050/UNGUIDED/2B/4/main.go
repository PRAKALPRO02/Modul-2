package main

import "fmt"

// # before modification #

// import "math"
// func main() {
// 	var k int
//     fmt.Print("Masukkan nilai K: ")
//     fmt.Scanln(&k)
// 	fmt.Printf("Nilai f(K) = %.10f\n",math.Pow(float64(4*k+2), 2)/(float64(4*k+1)*float64(4*k+3)))
// }

// # after modification #

func main() {
	var k int
	fmt.Print("Nilai K = ")
	fmt.Scanln(&k)
	var result float64 = 1.0
	for i := 0; i <= k; i++ {
		result *= float64((4*i + 2) * (4*i + 2)) / float64((4*i + 1) * (4*i + 3))
	}
	fmt.Printf("Nilai akar 2 = %.10f\n", result)
}