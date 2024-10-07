package main

import "fmt"

func isPrime(number int) bool {
    if number < 2 {
        return false
    }
    for i := 2; i*i <= number; i++ {
        if number%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    var number int
    fmt.Print("Masukkan Bilangan: ")
    fmt.Scanln(&number)

    fmt.Printf("Faktor dari %d adalah: ", number)
    for i := 1; i <= number; i++ {
        if number%i == 0 {
            fmt.Printf("%d ", i)
        }
    }
    fmt.Println()

    primeStatus := isPrime(number)
    fmt.Printf("Prima: %t\n", primeStatus)
}