// package main

// import "fmt"

// func main() {
// 	var bunga string
// 	var banyak int

// 	fmt.Print("N : ")
// 	fmt.Scan(&banyak)

// 	pita := ""

// 	for {
// 		fmt.Print("Bunga ", i+1, " : ")
// 		fmt.Scan(&bunga)

// 		if bunga == "SELESAI" {
// 			break
// 		}
// 		pita += bunga + " - "
// 		banyak++
// 	}

//		fmt.Println("Pita : ", pita)
//		fmt.PrintLn("Bunga: ", banyak)
//	}
package main

import "fmt"

func main() {
	var bunga string
	var banyak int

	pita := ""

	for {
		fmt.Print("Bunga ", banyak+1, " : ") // Memperbaiki penggunaan variabel banyak
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}
		pita += bunga + " - "
		banyak++ // Menambah jumlah bunga
	}

	fmt.Println("Pita: ", pita)
	fmt.Println("Jumlah bunga: ", banyak) // Memperbaiki Println
}
