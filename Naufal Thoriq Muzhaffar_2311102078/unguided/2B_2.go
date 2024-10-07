package main
import "fmt"

func main() {
	var bunga string
	var pita string
	var i int = 1
	for {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}

		pita += bunga + " - "
        i ++
	}
	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", i-1)
}
