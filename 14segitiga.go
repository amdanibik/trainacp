package main

import "fmt"

var (
	a, t int
)

func main() {
	fmt.Printf("Masukkan alas : ")
	fmt.Scanln(&a)
	fmt.Printf("Masukkan tinggi : ")
	fmt.Scanln(&t)
	luas := (a * t) / 2
	fmt.Println("Luas segitiga anda ", luas)
}
