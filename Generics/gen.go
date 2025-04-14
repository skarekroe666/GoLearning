package main

import "fmt"

func main() {
	// fmt.Println(add(235, 945))

	sum := Nums[int]{
		a: 24,
		b: 20,
	}

	fmt.Println("Sum of numbers is", add(sum.a, sum.b))
}

type Nums[T any] struct {
	a T
	b T
}

func add[T interface{ int | float64 }](a, b T) T {
	return a + b
}
