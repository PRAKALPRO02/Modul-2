package main

import (
	"fmt"
	"strings"
)

func main() {
	correctOrder := []string{"merah", "kuning", "hijau", "ungu"}

	numExperiments := 5

	for {
		success := true

		for i := 1; i <= numExperiments; i++ {
			var colors [4]string
			fmt.Printf("Percobaan %d: ", i)
			for j := 0; j < 4; j++ {
				fmt.Scan(&colors[j])
			}

			for j := 0; j < 4; j++ {
				if strings.ToLower(colors[j]) != correctOrder[j] {
					success = false
					break
				}
			}
		}

		if success {
			fmt.Println("Berhasil: true")
		} else {
			fmt.Println("Berhasil: false")
		}

		var again string
		fmt.Print("Lakukan percobaan lagi? (yes/no): ")
		fmt.Scan(&again)
		if strings.ToLower(again) != "yes" {
			break
		}
		fmt.Println() // Add a blank line for readability
	}
}
