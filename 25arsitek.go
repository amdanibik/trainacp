package main

import "fmt"

func main() {
	var angka int
	fmt.Printf("Masukkan angka ganjil : ")
	fmt.Scanln(&angka)
	if angka%2 != 0 {
		for x := 1; x <= angka; x++ {
			for y := angka - 1; y >= x; y-- {
				fmt.Printf(" ")
			}
			for z := 1; z <= x; z++ {
				fmt.Printf("*")
			}
			for a := 1; a <= x-1; a++ {
				fmt.Printf("*")
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Angka anda bukan ganjil ")
	}
}
