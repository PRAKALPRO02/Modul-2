package main
//Damara Galuh Pembayun _2311102110

import "fmt"

func main() {
        var warnaTujuan = []string{"merah", "kuning", "hijau", "ungu"}
        var percobaan int = 5

        for i := 1; i <= percobaan; i++ {
                fmt.Printf("Percobaan %d: ", i)
                var warnaInput [4]string
                fmt.Scanln(&warnaInput[0], &warnaInput[1], &warnaInput[2], &warnaInput[3])

                var berhasil bool = true
                for j := 0; j < len(warnaTujuan); j++ {
                        if warnaInput[j] != warnaTujuan[j] {
                                berhasil = false
                                break
                        }
                }

                fmt.Printf("BERHASIL: %t\n", berhasil)
        }
}