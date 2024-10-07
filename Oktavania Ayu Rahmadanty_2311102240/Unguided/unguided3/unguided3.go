package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Loop to continuously ask for bag weights
	for {
		// Prompt the user to input the weights of the two bags
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")

		// Read user input
		scanner.Scan()
		input := scanner.Text()

		// Split the input into two values
		weights := strings.Split(input, " ")
		if len(weights) != 2 {
			fmt.Println("Input harus terdiri dari dua angka.")
			continue
		}

		// Convert the input strings to floats
		weight1, err1 := strconv.ParseFloat(weights[0], 64)
		weight2, err2 := strconv.ParseFloat(weights[1], 64)

		// Check if both inputs are valid numbers
		if err1 != nil || err2 != nil {
			fmt.Println("Masukan berat yang valid!")
			continue
		}

		// Stop if either bag's weight is negative
		if weight1 < 0 || weight2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		// Calculate the total weight of both bags
		totalWeight := weight1 + weight2

		// Stop if the total weight exceeds 150 kg
		if totalWeight > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		// Calculate the absolute difference between the weights
		weightDiff := math.Abs(weight1 - weight2)

		// Check if the difference is greater than or equal to 9 kg
		if weightDiff >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}
