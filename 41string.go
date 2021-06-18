package main

import (
	"fmt"
	"strings"
)

func isCompare(katasatu, katadua string) {
	if len(katadua) > len(katasatu) {
		sama := strings.Contains(katadua, katasatu)
		if sama == true {
			fmt.Println(katasatu)
		}
	} else {
		sama := strings.Contains(katasatu, katadua)
		if sama == true {
			fmt.Println(katadua)
		}
	}
}
func main() {
	isCompare("AKA", "AKASHI")
	isCompare("KANGGORO", "KANG")
}
