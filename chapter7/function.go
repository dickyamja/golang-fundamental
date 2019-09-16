package main

import "fmt"

func main() {
	xs := []float64{2, 4, 5.6, 2.3, 5.6, 8.9, 6.7, 9.2, 5.2, 8.5}
	fmt.Println("Average : ", average(xs))
}

func average(xs []float64) float64 {
	total := 0.0
	for _, value := range xs {
		total += value
	}

	return total / float64(len(xs))
}
