package main

import "fmt"

func MeanMedian(ar []float64) {
	var mean = 0.0
	var median = len(ar) / 2
	var tampungMedian = 0.0
	for i := 0; i <= len(ar)-1; i++ {
		mean += ar[i]
		if i == median {
			tampungMedian = ar[i]
			if median%2 != 0 {
				tampungMedian = (ar[i] + ar[i-1]) / 2
			}
		}
	}
	var hitungMeann = float64(mean) / float64(len(ar))
	fmt.Println("Mean", hitungMeann, "Median", tampungMedian, median)
}

func main() {
	MeanMedian([]float64{1, 3, 5, 7, 9, 11})
}
