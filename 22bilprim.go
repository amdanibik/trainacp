package main

import "fmt"

var (
	angka, i, tampung int
)

func main() {
	fmt.Printf("Masukkan bilangan anda : ")
	fmt.Scanln(&angka)
	for i = 1; i <= angka; i++ {
		if angka%i == 0 {
			tampung++
		}
	}
	if tampung == 2 {
		fmt.Println(angka, " Merupakan bilangan prima")
	} else {
		fmt.Println(angka, " Bukan bilangan prima")
	}
}
