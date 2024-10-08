package main
import (
"fmt"
"strings"
)
func main() {
urutan := "merahkuninghijauungu"
var gelas1, gelas2, gelas3, gelas4 string
var hasil bool = true
for i := 0; i < 5; i++ {
fmt.Print("Percobaan ", i+1, ": ")
fmt.Scan(&gelas1, &gelas2, &gelas3, &gelas4)
if urutan != strings.ToLower(gelas1+gelas2+gelas3+gelas4) {
hasil = false
}
}
fmt.Print("Berhasil : ", hasil)
}