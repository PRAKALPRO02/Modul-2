package main

import "fmt"



func main() {

	var (

		satu, dua, tiga string

		temp string  

	)

	fmt.Print ("MAsukan input string: ")

	fmt.Scanln (&satu)

	fmt.Print ("MAsukan input string: ")

	fmt.Scanln (&dua)

	fmt.Print ("MAsukan input string: ")

	fmt.Scanln (&tiga)

	fmt.Println ("Outpu awal = " + satu + " " + dua + " " + tiga)

	temp = satu

	satu = dua

	dua = tiga 

	tiga = temp

	fmt.Println ("Output akhir = " + satu + " " + dua + " " + tiga)

}
