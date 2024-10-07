package main
import "fmt"

func main (){
//DEKLARASI VARIABEL
var beratKanan float64
var beratKiri float64

//MEMERIKSA KONDISI
for beratKanan<9 && beratKanan>=0 && beratKiri<9 && beratKiri>=0 {
	fmt.Print("masukkan berat belanjaan di kedua kantong: ")
	fmt.Scan(&beratKanan, &beratKiri)
}
fmt.Print("Proses selesai")

//Modifikasi
fmt.Print("\n\nSETELAH MODIFIKASI\n\n")

for {
	//INPUT
    fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
	fmt.Scan(&beratKanan, &beratKiri)

	//CEK BERAT
    if beratKanan < 0 || beratKiri < 0 || beratKanan+beratKiri > 150 {
        break
    }

	//CEK KONDISI
    if beratKanan-beratKiri >= 9 || beratKiri-beratKanan >= 9 {
        fmt.Println("Sepeda motor pak Andi akan oleng: true")
    } else {
        fmt.Println("Sepeda motor pak Andi akan oleng: false")
    }
}
fmt.Print("Proses selesai.")
}
