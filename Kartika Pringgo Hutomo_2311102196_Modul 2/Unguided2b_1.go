package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	correctOrder := []string{"merah", "kuning", "hijau", "ungu"}

	reader := bufio.NewReader(os.Stdin)
	success := false

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		colors := strings.Split(input, " ")

		attemptSuccess := true // flag for the current attempt

		if len(colors) == 4 { // make sure input has 4 colors
			for j := 0; j < 4; j++ {
				if colors[j] != correctOrder[j] {
					attemptSuccess = false
					break
				}
			}
		} else {
			attemptSuccess = false
		}

		if attemptSuccess {
			success = true
		} else {
			success = false
		}
	}

	if success {
		fmt.Println("BERHASIL : true")
	} else {
		fmt.Println("BERHASIL : false")
	}
}
