package main

import "fmt"

func main() {
	warnaTrue := [4]string{"merah", "kuning", "hijau", "ungu"}
	benar := true
	for i := 1; i <= 5; i++ {
		var warna1, warna2, warna3, warna4 string
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scan(&warna1, &warna2, &warna3, &warna4)
		if warna1 != warnaTrue[0] || warna2 != warnaTrue[1] || warna3 != warnaTrue[2] || warna4 != warnaTrue[3] {
			benar = false
		}
	}
	if benar {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}
