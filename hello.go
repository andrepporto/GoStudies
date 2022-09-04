package main

// imports
import (
	"fmt"
)

func main(){
	//first print
	fmt.Println("Hello world")
	
	//variables
	//string, int, float32, bool
	var firstName string = "Andre"

	fmt.Println(firstName)

	// := shorthand
	// acts like js
	// cannot be used outside a function
	lastName := "Porto"

	// Create w/o assign
	var fullName string
	fullName = (firstName + lastName)

	// multiple variable declaration
	var name1, name2 string = "andre", "gabriel"
	fmt.Println(name1, name2)

	//Constants
	const PIZZA string = "Pepperoni"
	fmt.Println(fullName, "loves", PIZZA)

}