package main

import (
	"fmt"
	"strings"
)

func cekPercobaan(percobaan [5][4]string) bool {
	expected := [4]string{"merah", "kuning", "hijau", "ungu"}

	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if percobaan[i][j] != expected[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	percobaan1 := [5][4]string{
		{"merah", "kuning", "hijau", "ungu"},
		{"merah", "kuning", "hijau", "ungu"},
		{"merah", "kuning", "hijau", "ungu"},
		{"merah", "kuning", "hijau", "ungu"},
		{"merah", "kuning", "hijau", "ungu"},
	}

	percobaan2 := [5][4]string{
		{"merah", "kuning", "hijau", "ungu"},
		{"merah", "kuning", "hijau", "ungu"},
		{"merah", "kuning", "hijau", "ungu"},
		{"kuning", "merah", "hijau", "ungu"},
		{"merah", "kuning", "hijau", "ungu"},
	}

	fmt.Println("Percobaan 1:")
	for i, p := range percobaan1 {
		fmt.Printf("Percobaan %d: %s\n", i+1, strings.Join(p[:], " "))
	}
	if cekPercobaan(percobaan1) {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}

	fmt.Println("\nPercobaan 2:")
	for i, p := range percobaan2 {
		fmt.Printf("Percobaan %d: %s\n", i+1, strings.Join(p[:], " "))
	}
	if cekPercobaan(percobaan2) {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}
