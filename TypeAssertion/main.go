package main

import "fmt"

func main() {

	var a any = 45
	if b, ok := a.(int); ok {
		fmt.Println("Type assertion successful:", b)
	} else {
		fmt.Println("Type assertion failed")
	}

	typeCheck(234)

	data := map[string]any{
		"name": "Alice",
		"age":  30,
	}
	name, ok := data["name"].(string)
	if ok {
		fmt.Println("Name:", name)
	} else {
		fmt.Println("Name is not a string")
	}

}

func typeCheck(a any) {
	switch v := a.(type) {
	case int:
		fmt.Println("value is int:", v)
	case string:
		fmt.Println("value is string:", v)
	case float64:
		fmt.Println("value is float:", v)
	case bool:
		fmt.Println("value is boolean")
	}
}
