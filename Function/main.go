package main

import "fmt"

func main() {
	fmt.Println(add(345, 53))

	p, d := product(24, 63)
	fmt.Println("Product:", p, "Double:", d)

	fmt.Println(fact(9, 8, 7, 6, 5, 4, 3, 2, 1))

	fmt.Println("Result:", subtract(24, 13))

	sum := apply(sum, 523, 388)
	fmt.Print(sum)
}

func add(a, b int) int {
	return a + b
}

func product(a, b int) (int, int) {
	p := a * b
	d := (a * b) * 2

	return p, d
}

func fact(nums ...int) int {
	total := 0

	for _, v := range nums {
		total += v
	}
	return total
}

func subtract(a, b int) (result int) {
	return a - b
}

func apply(f func(int, int) int, x, y int) int {
	return f(x, y)
}
func sum(x, y int) int {
	return x + y
}
