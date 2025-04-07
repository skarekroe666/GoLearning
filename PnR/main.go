package main

import "fmt"

func main() {
	fmt.Println("Before panic")											//1st STEP
	riskyFunction()														//2nd STEP   FUNCTION CALL
	fmt.Println("Program contiues after panic")							//LAST STEP
}

func handlePanic() {
	if r := recover(); r != nil {										//4TH STEP
		fmt.Println("Recovered from panic:", r)							//PRINTS IN THE 4TH STEP
	}
}
func riskyFunction() {
	defer handlePanic()													//THIS WILL EXECUTE AFTER 3RD STEP. CALLED AT LAST
	fmt.Println("Inside risky function")								//3rd STEP
	panic("Something went wrong")										//THIS IS THE PANIC MESSAGE GOES IN RECOVER
	fmt.Println("This will not be printed")
}
