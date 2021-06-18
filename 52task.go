package main

import "fmt"

func getMinMax(numbers ...*int) (min int, max int) {
	tempMax := 0
	tempMin := 0
	for i := 0; i < len(numbers)-1; i++ {
		if *(numbers[i]) < *(numbers[i+1]) && tempMax == 0 && tempMin == 0 {
			tempMax = *(numbers[i+1])
			tempMin = *(numbers[i])
		} else if *(numbers[i]) > *(numbers[i+1]) && tempMax == 0 && tempMin == 0 {
			tempMax = *(numbers[i])
			tempMin = *(numbers[i+1])
		} else {
			if tempMax < *(numbers[i]) {
				tempMax = *(numbers[i])
			} else if tempMin > *(numbers[i]) {
				tempMin = *(numbers[i])
			} else if tempMax < *(numbers[len(numbers)-1]) {
				tempMax = *(numbers[len(numbers)-1])
			} else if tempMin > *(numbers[len(numbers)-1]) {
				tempMin = *(numbers[len(numbers)-1])
			}
		}
	}
	min = tempMin
	max = tempMax
	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai Min : ", min)
	fmt.Println("Nilai Max : ", max)
}
