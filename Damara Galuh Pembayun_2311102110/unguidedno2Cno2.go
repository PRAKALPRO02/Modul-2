package modul
import "fmt"
//Damara Galuh Pembayun _2311102110

func main () {
	var nam float 64
	var nmk string
	fmt.Print ("Nilai akhir mata kuliah: ")
	fmt.Scanln (&nam)
	if nam > 80 {
		nam = "A"
	}
	if nam > 72.5 {
		nam = "AB"
	}
	if nam > 65 {
		nam = "B"
	}
	if nam > 57.5 {
		nam = "BC"
	}
	if nam > 50 {
		nam = "C"
	}
	if nam > 40 {
		nam = "D"
	}else if nam <= 40 {
		nam = "E"
	}
	fmt.Println ("Nilai mata kuliah : " , nmk)
}