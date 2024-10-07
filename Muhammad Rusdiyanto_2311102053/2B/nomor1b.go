package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var result, prevResult string
	var conclusion bool = true
	input := bufio.NewScanner(os.Stdin)

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %v : ", i)
		input.Scan()
		result = input.Text()
		if i == 1 {
			prevResult = result
			continue
		}
		if result != prevResult {
			conclusion = false
		}
	}
	fmt.Printf("BERHASIL : %v", conclusion)
}
