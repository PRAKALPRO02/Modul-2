package main

import "fmt"

func main() {
	var f float64 = 1
	var k int

	fmt.Print("Nilai K : ")
	fmt.Scan(&k)
	for i := 0; i <= k; i++ {
		f *= float64((4.0*i+2.0)*(4.0*i+2.0)) / float64((4.0*i+1.0)*(4.0*i+3.0))
	}
	fmt.Printf("Nilai akar 2 : %.10f", f)
}
