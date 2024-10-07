package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Urutan warna yang benar
	correctOrder := []string{"merah", "kuning", "hijau", "ungu"}
	reader := bufio.NewReader(os.Stdin)

	for {
		var hadError bool

		// Membaca input untuk 5 percobaan
		for i := 1; i <= 5; i++ {
			fmt.Printf("Percobaan %d: ", i)

			// Membaca input dari pengguna
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			// Memisahkan input berdasarkan spasi
			colors := strings.Split(input, " ")

			// Mengecek apakah urutan warna sesuai
			for j := 0; j < 4; j++ {
				if j >= len(colors) || colors[j] != correctOrder[j] {
					hadError = true
					break
				}
			}
		}

		// Menampilkan hasil
		if !hadError {
			fmt.Println("BERHASIL: true")
		} else {
			fmt.Println("BERHASIL: false")
		}

		// Menanyakan pengguna apakah ingin mencoba lagi
		fmt.Print("Apakah Anda ingin mencoba lagi? (ya/tidak): ")
		retry, _ := reader.ReadString('\n')
		retry = strings.TrimSpace(retry)

		if strings.ToLower(retry) != "ya" {
			break // Keluar dari perulangan jika pengguna tidak ingin mencoba lagi
		}
	}

	fmt.Println("Terima kasih telah bermain!")
}
