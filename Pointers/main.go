package main

import "fmt"

func main() {
	x := 32
	fmt.Println("x:", x)
	pointer(&x)
	fmt.Println("x:", x)
}

func pointer(n *int) {
	*n = 42
}
