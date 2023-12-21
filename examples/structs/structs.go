package main

import "fmt"

type Employee struct {
	name string
	salary int
}

// Creating instance using a function
func getEmployeeInstance(name string, salary int) *Employee{
	emp := Employee{name: name, salary:salary}
	return &emp
}

// Use of pointer receiver
// Struct methods 
func (emp *Employee) getName() string {
	return emp.name
}

// Ptrs allows mutating struct
func (emp *Employee) setName(newName string) {
	emp.name = newName
}

func main(){

	// Creating an instance of Employee
	emp := Employee{name: "Ram", salary: 100}
	fmt.Println(emp)

	employees := []*Employee{}
	employees = append(employees, getEmployeeInstance("Shayam", 0))
	fmt.Println(employees)
	for _, employee := range employees{
		fmt.Println(*employee)
	}

	// Anonymous structs
	stu := struct {
		name string 
		marks int
	} {
		"John Doe",
		123,
	}
	fmt.Println(stu)

	// Method invocation for a struct
	fmt.Println(emp.getName())
	emp.setName("Mark")
	fmt.Println(emp.getName())
}