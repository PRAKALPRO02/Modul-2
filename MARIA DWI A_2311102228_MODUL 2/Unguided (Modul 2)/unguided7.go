package main

import (
	"fmt"
	"strings"
)


func cariFaktor(bilangan int) []int{
	var faktor [] int 
	for i := 1; i <= bilangan; i++{
		if bilangan%i == 0{
			faktor = append(faktor, i)
		}
	}
	return faktor
}

func isPrima(bilangan int) bool{
	if bilangan <= 1{
		return false
	}
	for i := 2; i*i <= bilangan; i++{
		if bilangan%i == 0{
			return false
		}
	}
	return true
}

func main(){
	var bilangan int
	fmt.Print("Masukkan bilangan bulat (b > 1) : ")
	_, err := fmt.Scanf("%d", &bilangan)


	if err != nil || bilangan <= 1{
		fmt.Println("Input tidak valid. Masukkan bilangan bulat lebih besar dari 1.")
		return
	}

	faktor := cariFaktor(bilangan)
	prima := isPrima(bilangan)


	fmt.Printf("\nBilangan: %d\n", bilangan)
	fmt.Printf("Faktor: %s\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(faktor)), " "), "[]"))
	fmt.Printf("Prima : %t\n", prima)
	fmt.Println("\n\n")
}