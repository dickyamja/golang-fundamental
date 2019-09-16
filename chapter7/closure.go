package main

import "fmt"

func main() {
	x := 0
	add := func(number1, number2 int) int {
		return number1 + number2
	}

	increment := func() int {
		x++
		return x
	}

	nextEven := makeEvenNumber()

	fmt.Println("add (1, 2)", add(1, 2))
	fmt.Println("increment : ", increment())
	fmt.Println("increment : ", increment())
	fmt.Println("nextEven : ", nextEven())
	fmt.Println("nextEven : ", nextEven())
	fmt.Println("nextEven : ", nextEven())
}

func makeEvenNumber() func() uint {
	i := uint(0)

	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
