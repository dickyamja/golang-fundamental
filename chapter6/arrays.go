package main

import "fmt"

func main() {
	var x [5]float64
	var total, totalRange float64

	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	for _, value := range x {
		totalRange += value
	}
	fmt.Println("x : ", x)
	fmt.Println("Average using standard loop : ", total/float64(len(x)))
	fmt.Println("Average using loop in range : ", totalRange/float64(len(x)))
}
