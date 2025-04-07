package main

import "fmt"

func main() {

	person := Employee{
		Person: Person{Name: "Skarekreo", Age: 26},
		Role:   "DevOps Engineer",
		Salary: 500000,
	}

	updateAge(&person)
	// convertSalary(&person)
	person.displayInfo()
}

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	Role   string
	Salary int
}

func (e Employee) displayInfo() {
	fmt.Println("Name:", e.Name)
	fmt.Println("Age:", e.Age)
	fmt.Println("Role:", e.Role)
	fmt.Println("Salary:", e.Salary)
}
func updateAge(e *Employee) {
	e.Age = 30
}

func convertSalary(e *Employee) {
	e.Salary = 500000 * 86
}
