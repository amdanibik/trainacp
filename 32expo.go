package main

import (
	"fmt"
)

func isPow(x, n int) {
	var sisaBagi int
	var tampung = x
	var hasil = x
	var tampungNol = 0
	var tampungSatu = 0
	var tampungSisabagi = 0
	for n >= 1 {
		sisaBagi = n % 2
		if sisaBagi == 0 {
			hasil *= hasil
			tampungNol++
			tampungSisabagi++
		} else {
			hasil = hasil * tampung
			tampungSatu++
			tampungSisabagi++
		}
		n /= 2
	}
	if tampungSatu == 1 {
		hasil = hasil / x
	} else if tampungNol == 0 && tampungSisabagi > 2 {
		hasil = hasil * (hasil / 2)
	}
	fmt.Println(x, hasil)
	fmt.Println()
}

func main() {
	// 6 pangkat
	isPow(6, 2)
	isPow(6, 3)
	isPow(6, 4)
	isPow(6, 5)
	isPow(6, 8)
	isPow(6, 9)
	// 7 pangkat
	isPow(7, 2)
	isPow(7, 3)
	isPow(7, 4)
	isPow(7, 5)
	isPow(7, 8)
	isPow(7, 9)
}
