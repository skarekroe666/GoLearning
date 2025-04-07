package main

import (
	"fmt"
	"os"
)

func main() {

	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// return
	// }

	// defer file.Close()

	// fmt.Println("File creation successful")

	// text := "Writing in the file"

	// _, errors := file.WriteString(text)
	// if errors != nil {
	// 	fmt.Println("Error writing to the file:", errors)
	// 	return
	// }

	// fmt.Println("Writing to the file successful")

	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	fmt.Println("Error reading the file:", err)
	// 	return
	// }

	// fmt.Println("Contents of the file:", string(data))

	// file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	fmt.Println("Error appendig to the file:", err)
	// 	return
	// }

	// defer file.Close()

	// _, errors := file.WriteString(" ,Appended line")
	// if errors != nil {
	// 	fmt.Println("Error appending to the file:", errors)
	// 	return
	// }

	// fmt.Println("Apending to the file successful")

	// err := os.Remove("example.txt")
	// if err != nil {
	// 	fmt.Println("Error deleting the file:", err)
	// 	return
	// }

	// fmt.Println("File deletion succuessful")

	_, errors := os.Stat("example.txt")

	if os.IsNotExist(errors) {
		fmt.Println("File does not exist")
	} else {
		fmt.Println("File exists")
	}
}
