package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    pita_220 := ""
    var bunga_220List []string

    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("N: ")
    scanner.Scan()
    input := scanner.Text()
    if input == "0" {
        fmt.Println("Pita:")
        fmt.Println("Bunga: 0")
        return
    }

    count := 0
    for {
        count++
        fmt.Printf("Bunga %d: ", count)
        scanner.Scan()
        bunga_220 := scanner.Text()

        if strings.ToUpper(bunga_220) == "SELESAI" {
            count--
            break
        }

        if pita_220 == "" {
            pita_220 = bunga_220
        } else {
            pita_220 += "-" + bunga_220
        }

        bunga_220List = append(bunga_220List, bunga_220)
    }

    fmt.Println("Pita:", pita_220)
    fmt.Printf("Bunga: %d\n", len(bunga_220List))
}
