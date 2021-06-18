package main

import "fmt"

var (
	angka, i int
)

func main() {
	fmt.Printf("Masukkan bilangan anda : ")
	fmt.Scanln(&angka)
	for i = 1; i <= angka; i++ {
		if angka%i == 0 {
			fmt.Println(i)
		}
	}
}
