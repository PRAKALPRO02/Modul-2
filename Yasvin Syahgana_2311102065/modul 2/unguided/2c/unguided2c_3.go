package main
//adamfaizal 2311102088
import (
	"fmt"
)

// Fungsi untuk mencari faktor dari bilangan bulat b
func cariFaktor(b int) []int {
	var faktor []int          // Slice untuk menyimpan faktor-faktor
	for i := 1; i <= b; i++ { // Iterasi dari 1 hingga b
		if b%i == 0 { // Cek apakah i adalah faktor dari b
			faktor = append(faktor, i) // Tambahkan faktor ke slice
		}
	}
	return faktor
}

// Fungsi untuk menentukan apakah bilangan bulat b adalah bilangan prima
func isBilanganPrima(b int) bool {
	if b <= 1 {
		return false // Bilangan kurang dari atau sama dengan 1 bukan prima
	}
	for i := 2; i*i <= b; i++ { // Cek hingga akar kuadrat dari b
		if b%i == 0 {
			return false // Jika ada pembagi lain, bukan prima
		}
	}
	return true // Jika tidak ada pembagi lain, maka bilangan prima
}

func main() {
	var b int

	// Input dari pengguna
	fmt.Print("Masukkan bilangan bulat b (b > 1): ")
	_, err := fmt.Scan(&b)
	if err != nil || b <= 1 {
		fmt.Println("Input tidak valid. Harap masukkan bilangan bulat yang lebih besar dari 1.")
		return
	}

	faktorB := cariFaktor(b)

	// Menampilkan hasil faktor
	fmt.Printf("Faktor-faktor dari %d adalah: %v\n", b, faktorB)

	// Menentukan dan menampilkan apakah bilangan tersebut adalah bilangan prima
	if isBilanganPrima(b) {
		fmt.Printf("%d adalah bilangan prima.\n", b)
	} else {
		fmt.Printf("%d bukan bilangan prima.\n", b)
	}
}