package main
import "fmt"

func main (){
	//deklarasi unsign integer (yang tidak bernilai negatif)
	var N uint64
	var i uint64
	var j uint64

	//membuat array
	fmt.Print("N: ")
	fmt.Scan(&N)
	bunga := make([]string,N)

	//memberikan peringatan jika inputan negatif atau kurang dari 0
	if N<=0{
		fmt.Printf("N: %v \nPita: \nPERINGATAN!\nINPUT BUKAN BILANGAN POSITIF ATAU SAMA DENGAN 0",N )
	}

	//melakukan input ke array
	for i = 0; i<N; i++ {
		fmt.Printf("Bunga %v: ", i+1)
		fmt.Scan(&bunga[i])
	}

	//mengoutputkan isi array sesuai format pita
	for j = 0 ; j<N ; j++{
		fmt.Print(bunga[j] + "-")
	}
	fmt.Println()

//Modifikasi
	fmt.Print("\n\nSETELAH MODIFIKASI\n\n")

	//membuat array
	var input string
	bungaMod := make([]string, 0)

	//melakukan input dan mengecek input != SELESAI
	for i = 0; input != "SELESAI"; i++ {
		fmt.Printf("Bunga %v: ", i+1)
		fmt.Scan(&input)
		if input!="SELESAI"{
		bungaMod = append(bungaMod, input)}
	}

	//mencetak output sesuai format
	fmt.Print("Pita: ")
	for _, b := range bungaMod {
		fmt.Print(b + "-")
	}
	fmt.Println()

	//mencetak jumlah input
	fmt.Printf("Jumlah bunga yang dimasukkan: %v\n", len(bungaMod))
}
