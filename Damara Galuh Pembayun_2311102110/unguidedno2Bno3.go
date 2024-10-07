package main
//Damara Galuh Pembayun _2311102110

import "fmt"

func main() {
        var beratKiri, beratKanan float64
        var totalBerat float64

        for {
                fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
                fmt.Scan(&beratKiri, &beratKanan)

                totalBerat = beratKiri + beratKanan

                // Kondisi berhenti:
                if beratKiri >= 9 || beratKanan >= 9 || totalBerat > 150 || beratKiri < 0 || beratKanan < 0 {
                        fmt.Println("Proses selesai.")
                        break
                }

                // Cek apakah sepeda motor akan oleng
                selisih := beratKiri - beratKanan
                if selisih < 0 {
                        selisih = -selisih
                }
                fmt.Printf("Sepeda motor pak Andi akan oleng: %t\n", selisih >= 9)
        }
}