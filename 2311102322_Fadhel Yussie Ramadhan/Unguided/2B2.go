package main

import ("fmt")

func main() {
    var n int
    var bunga [20] string
    n = 0

    for {
        fmt.Print ("Bunga ", n + 1, ": ")
        fmt.Scan (&bunga[n])
        if bunga [n] == "SELESAI"{
            break
        }
        n++
    }
    fmt.Print ("Pita :")
    for i := 0; i < n; i++ {
        fmt.Print (bunga [i], " - ")
    }
    fmt.Println ("\nBunga : ", n)
}
