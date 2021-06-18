package main

import "fmt"

var (
	angka, n int
)

func main() {
	fmt.Printf("Masukkan angka : ")
	fmt.Scanln(&angka)
	if angka%2 != 0 {
		fmt.Println("Bilangan ", angka, " merupakan bilangan prima")
	} else {
		fmt.Println("Bilangan ", angka, " bukan bilangan prima ")
	}
}
