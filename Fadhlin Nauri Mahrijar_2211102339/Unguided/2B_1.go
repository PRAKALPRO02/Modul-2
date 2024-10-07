package main

import "fmt"

func main() {
	var a, b, c, d string
	i := 1
	berhasil := true
	for {
		fmt.Print("percobaan ", i, ":")
		fmt.Scanln(&a, &b, &c, &d)
		if a != "merah" && b != "kuning" && c != "hijau" && d != "unggu" {
			berhasil = false
		}
		i++
		if i == 6 {
			break
		}
	}
	fmt.Print(berhasil)
}
