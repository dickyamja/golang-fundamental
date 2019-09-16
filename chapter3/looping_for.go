package main

import "fmt"

func main() {
	i := 1
	fmt.Println("First Looping ")
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("Second Looping ")
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}
}
