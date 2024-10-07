package main

import "fmt"

func main() {
	var warna1, warna2, warna3, warna4 string
	attempt := 1
	success := true

	for attempt <= 5 {
		fmt.Printf("Percobaan %d: ", attempt)
		fmt.Scanln(&warna1, &warna2, &warna3, &warna4)

		if warna1 != "merah" || warna2 != "kuning" || warna3 != "hijau" || warna4 != "ungu" {
			success = false
		}

		attempt++
		fmt.Println(success)
	}
}
