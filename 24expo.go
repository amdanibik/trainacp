package main

import "fmt"

func main() {
	var angka, pangkat, i int
	var tampung int = 1
	fmt.Printf("Masukkan angka anda : ")
	fmt.Scanln(&angka)
	fmt.Printf("Masukkan pangkat yang anda inginkan : ")
	fmt.Scanln(&pangkat)
	for i = 1; i <= pangkat; i++ {
		tampung *= angka
	}
	fmt.Println("Hasil dari ", angka, " pangkat ", pangkat, " adalah ", tampung)
}
