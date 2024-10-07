package main

import "fmt"

func main(){
	var b int
	var f [] int

	fmt.Print ("Bilangan : ")
	fmt.Scan(&b)

	if b > 1 {
		for i := 1; i <= b; i++ {
			if b % i == 0{
				f = append(f, i)
			}
		}
	}
	fmt.Println("Faktor : ", f)
	if len(f) == 2 {
		fmt.Println("Prima: true")
	} else {
		fmt.Println ("Prima: false")
	} 
}