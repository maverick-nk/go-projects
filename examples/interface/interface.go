package main

import "fmt"

/*
Reference: https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
Go's type system - Instead of desigining our abstractions in terms of 
what kind of data our types can hold, we design our abstractions in 
terms of what actions our types can execute.
*/

type Animal interface {
	Speak() string
}

type Cat struct {}
type Dog struct {}

// Value receiver
func (c Cat) Speak() string{
	return "Cat speaking"
}
func (d Dog) Speak() string{
	return "Dog speaking"
}

func main(){
	animals := []Animal{Cat{}, Dog{}}
	for _, animal :=range animals {
		fmt.Println(animal.Speak())
	}

	
}