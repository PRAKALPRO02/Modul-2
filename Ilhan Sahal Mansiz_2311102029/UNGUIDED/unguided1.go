package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cekPercobaan(percobaan [][]string) bool {
	susunanBenar := []string{"merah", "kuning", "hijau", "ungu"}

	for _, p := range percobaan {
		for i, warna := range p {
			if warna != susunanBenar[i] {
				return false
			}
		}
	}
	return true
}
func main() {
	var percobaan [][]string
	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		scanner.Scan()
		input := scanner.Text()
		warna := strings.Split(input, " ")
		percobaan = append(percobaan, warna)
	}

	hasil := cekPercobaan(percobaan)
	fmt.Println("BERHASIL:", hasil)
}