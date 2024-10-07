package main

import "fmt"

func main() {
	var nam float64
	var nmk string

	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)

	switch {
	case nam > 80:
		nmk = "A"
	case nam > 72.5:
		nmk = "AB"
	case nam > 65:
		nmk = "B"
	case nam > 57.5:
		nmk = "BC"
	case nam > 50:
		nmk = "C"
	case nam > 40:
		nmk = "D"
	default:
		nmk = "E"
	}

	fmt.Println("Nilai mata kuliah: ", nmk)
}
