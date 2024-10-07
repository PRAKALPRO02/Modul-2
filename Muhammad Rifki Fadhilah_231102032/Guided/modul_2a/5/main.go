package main

import "fmt"

func main() {
    var num [5]int
    var char [3]rune


    for i := 0; i < 5; i++ {
        fmt.Scanf("%d", &num[i])
    }

    var ignore string
    fmt.Scanf("%s", &ignore)

    for i := 0; i < 3; i++ {
        fmt.Scanf("%c", &char[i])
    }
    fmt.Println()

    for i := 0; i < 5; i++ {
        fmt.Printf("%c", num[i])
    }
    fmt.Println()

    for i := 0; i < 3; i++ {
        fmt.Printf("%c", char[i]+1)
    }
    fmt.Println()
}
