package main

import "fmt"

func main() {
	var tahun int
		
	fmt.Print ("Tahun: ")
	fmt.Scanln(&tahun)

	if tahun % 4 == 0 {
		if tahun % 100 == 0 {
			if tahun % 400 == 0{
				fmt.Print("Kabisat: True")
			}else {
				fmt.Print("Kabisat: False")
			}
		} else {
			fmt.Print("Kabisat: True")
		}
	} else {
		fmt.Print("Kabisat: False")
	}
}