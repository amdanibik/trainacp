package main

import "fmt"

var (
	t, r float64
	pi   float64 = 3.14
)

func main() {
	fmt.Printf("Masukkan tinggi tabung : ")
	fmt.Scanln(&t)
	fmt.Printf("Masukkan Jari-jari : ")
	fmt.Scanln(&r)
	lp := 2 * pi * r * (r + t)
	fmt.Println("Luas Permukaan tabung anda ", lp)
}
