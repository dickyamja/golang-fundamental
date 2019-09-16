package main

import "fmt"

func main() {
	arr := [5]float64{1, 2, 3, 4, 5}
	var x []float64
	y := make([]float64, 5, 10)
	sliceFromArr := arr[1:3]

	fmt.Println("x : ", x)
	fmt.Println("y ", y)
	fmt.Println("sliceFromArr : ", sliceFromArr)
}
