package main

import "fmt"

var number, tampung, hasil int

func isPrime(angka int) {
	if angka%2 == 0 && angka > 2 {
		fmt.Println("Bukan Bilangan Prima")
	} else {
		for bagi := 3; bagi <= 7; bagi += 2 {
			hasil := angka % bagi
			if hasil == 0 {
				tampung++
			}
		}
		if tampung >= 1 && angka >= 9 {
			fmt.Println("Bukan Bilangan Prima")
		} else {
			fmt.Println("Bilangan Prima")
		}
	}
}
func main() {
	fmt.Printf("Masukkan angka ")
	fmt.Scanln(&number)
	isPrime(number)
}
