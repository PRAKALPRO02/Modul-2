package main

import "fmt"

func main() {

	warna := [4]string{"merah", "kuning", "hijau", "ungu"}

	berhasil := true

	for i := 1; i <= 5; i++ {
		var c1, c2, c3, c4 string
		fmt.Printf("Percobaan %d: ", i)

		_, err := fmt.Scan(&c1, &c2, &c3, &c4)
		if err != nil {
			berhasil = false
			break
		}

		if c1 != warna[0] || c2 != warna[1] || c3 != warna[2] || c4 != warna[3] {
			berhasil = false
			break
		}
	}

	fmt.Printf("BERHASIL : %v\n", berhasil)
}