package main

import (
	"errors"
	"fmt"
)

// Handles function and flow statements
func main() {
	//calling function with parameter
	var printValue = "Hello World"
	printMe(printValue)

	var numerator = 11
	var denominator = 2
	var output, remainder, err = intDivision(numerator, denominator)

	//check if an error was encountered with if ... else if statement
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		//with printf function, we can format the string using variables
		fmt.Printf("The result of the integer division is %v\n", output)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v\n", output, remainder)
	}

	//checking if an error was encountered using the switchcase statement
	// in Go, the break is implied and don't need to be explicitly written
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		//with printf function, we can format the string using variables
		fmt.Printf("The result of the integer division is %v\n", output)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v\n", output, remainder)
	}

	//There is also the conditional switch statement where the condition is inserted after the switch keyword
	switch remainder {
	case 0:
		fmt.Printf("The division was exact\n")
	case 1, 2:
		fmt.Printf("The division was close\n")
	default:
		fmt.Printf("The division was not close\n")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

// If a function is returning a value, you have to indicate the return type
// You can specify multiple return value by putting them in a parenthesis
// with the added error return type, if the denominator is not zero, the code will continue and return nil
func intDivision(numerator int, denominator int) (int, int, error) {
	//All errors are returned as error type which is a built in type in go and can be imported
	var err error
	if denominator == 0 {
		//creates an error type we can initialize with an error message
		err = errors.New("cannot divide by 0") //Note, error string should be capitalize
		return 0, 0, err
	}
	var result = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
