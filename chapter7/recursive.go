package main

import "fmt"

func main() {

	fmt.Println("factorial(10) : ", factorial(10))

}

func factorial(i uint) uint {
	if i == 0 {
		return 1
	}

	return i * factorial(i-1)
}
