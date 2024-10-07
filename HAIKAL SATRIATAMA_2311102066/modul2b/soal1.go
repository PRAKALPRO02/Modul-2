package main

import "fmt"

func main() {
	var bunga1, bunga2, bunga3, bunga4 string
	var beda bool = false

	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, " : ")
		fmt.Scan(&bunga1, &bunga2, &bunga3, &bunga4)
		if (bunga1 != "merah" || bunga2 != "kuning" || bunga3 != "hijau" || bunga4 != "ungu") && !beda {
			beda = true
		}
	}
	fmt.Println("BERHASIL : ", !beda)
}
