package main

import (
	"fmt"
)

func arraymerge(arrayA, arrayB []string) {
	fmt.Println()
	tampungArray := append(arrayA, arrayB...)
	mapArray := make(map[string]string)
	for i := 0; i < len(tampungArray); i++ {
		// fmt.Printf(tampungArray[i] + " ")
		// var urut = strconv.Itoa(i)
		var data = tampungArray[i]
		mapArray[data] = tampungArray[i]
		// fmt.Println(mapArray[i])
	}
	for _, val := range mapArray {
		fmt.Printf(val + " ")
	}
}
func main() {
	arraymerge([]string{"kazuya", "jin", "lee"}, []string{"kazuya", "feng"})
	arraymerge([]string{"jin", "lee"}, []string{"kazuya", "panda"})
	arraymerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"})
	arraymerge([]string{"sergie", "jin"}, []string{"jin", "steve"})
}
