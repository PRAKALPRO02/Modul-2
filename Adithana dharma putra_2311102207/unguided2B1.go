package main
import "fmt"

//fungsi Pengecekan
func cek(Input [][]string, urutan []string) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < len(urutan); j++ {
			if Input[i][j] != urutan[j] {
				return false
			}     
		}
	} 
	return true
}

func main() {
	//deklarasi array
	warnaInput := make([][]string, 5)
	for i := range warnaInput {
		warnaInput[i] = make([]string, 5)
	}
	//deklarasi susunan yang benar
	urutanWarna := []string{"merah", "kuning", "hijau", "ungu"}

	//meminta input
	for i := 0; i < 5; i++ {
		fmt.Printf("Percobaan %v: ", i+1)
		fmt.Scan(&warnaInput[i][0],&warnaInput[i][1],&warnaInput[i][2],&warnaInput[i][3])
	}

	//memanggil fungsi cek dan mengeluarkan Output
	Hasil := cek(warnaInput, urutanWarna)
	fmt.Print("BERHASI: ",Hasil)
}
