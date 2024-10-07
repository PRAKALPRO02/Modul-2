package main

import "fmt"

func main() {
	var nmk float64
	var nam string

	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nmk)

	if nmk > 80 && nmk < 100 {
		nam = "A"
	} else if nmk > 72.5 && nmk < 80 {
		nam = "AB"
	} else if nmk > 65 && nmk < 72.5 {
		nam = "B"
	} else if nmk > 57.5 && nmk < 65 {
		nam = "BC"
	} else if nmk > 50 && nmk < 57.5 {
		nam = "C"
	} else if nmk > 40 && nmk < 50 {
		nam = "D"
	} else if nmk <= 40 && nmk >= 0 {
		nam = "E"
	}

	fmt.Print("Nilai mata kuliah: ", nam)
}
