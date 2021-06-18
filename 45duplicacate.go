package main

import (
	"fmt" "strings"
)

func cekJumlah(deretAngka []int) {
	tampungAngka := make(map[int]int)
	for i := 0; i < len(deretAngka); i++ {
		tampungAngka[i] = deretAngka[i]
	}
	for _, val := range tampungAngka {
		ubahAngka := strings.Atoi(val)
		fmt.Printf(ubahAngka)
	}
}
func main() {
	cekJumlah([]int{2, 3, 3, 3, 6, 9})
}
