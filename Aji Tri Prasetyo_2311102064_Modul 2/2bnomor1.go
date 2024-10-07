package main

import "fmt"

func main() {
	var warnaInput [4]string
	warnaDiharapkan := [4]string{"merah", "kuning", "hijau", "ungu"}
	var adaPerbedaan bool = false

	for percobaan := 1; percobaan <= 5; percobaan++ {
		fmt.Print("Percobaan ", percobaan, " : ")
		fmt.Scan(&warnaInput[0], &warnaInput[1], &warnaInput[2], &warnaInput[3])

		for j := 0; j < 4; j++ {
			if (warnaInput[j] != warnaDiharapkan[j]) && !adaPerbedaan {
				adaPerbedaan = true

			}
		}

	}
	fmt.Println("HASIL : ", !adaPerbedaan)
}
