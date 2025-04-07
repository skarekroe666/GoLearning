package main

import "fmt"

func main() {

	m := map[int]string{
		1: "skarekroe",
		2: "zoya",
		3: "anisha",
	}

	students := map[string]map[string]int{
		"skarekroe": {"age": 26},
		"zoya":      {"age": 24},
		"anisha":    {"age": 24},
	}

	fmt.Println(m)

	for i, v := range students {
		fmt.Println(i, v)
	}
	// fmt.Println(students)
}
