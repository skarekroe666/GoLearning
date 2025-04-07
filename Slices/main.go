package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5}
	// s = append(s, 4, 5)

	fmt.Println(s)

	total := 0
	for i := range s {
		total += i
	}
	fmt.Println(total)
}
