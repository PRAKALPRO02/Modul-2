package main

import (
    "fmt"
    "strings"
)

func main() {
    urutanBenar := []string{"merah", "kuning", "hijau", "ungu"}
    percobaan := make([][]string, 5)
    berhasil := true

    // Input dan tampilkan 5 percobaan
    for i := 0; i < 5; i++ {
        percobaan[i] = make([]string, 4)
        fmt.Printf("Percobaan %d:", i+1)
        for j := 0; j < 4; j++ {
            //fmt.Printf("Masukkan warna gelas %d: ", j+1)
            fmt.Scan(&percobaan[i][j])
            percobaan[i][j] = strings.ToLower(percobaan[i][j])
        }
    }

    
    for i := 0; i < 5; i++ {
        for j := 0; j < 4; j++ {
           
            if percobaan[i][j] != urutanBenar[j] {
                berhasil = false
            }
        }
       
    }
    fmt.Printf("BERHASIL: %t\n", berhasil)
    fmt.Println()

    // Jika gagal, tampilkan percobaan kedua (gagal)
    if !berhasil {
        
        berhasil = true
       
        for i := 0; i < 5; i++ {
           
               berhasil = false
        }
       
    }
}