package main

import "fmt"

func main() {
	// fmt.Println("Hello World!")
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	//print an approximation unlike float64 that gives exact number
	var floatNum float32 = 12345678.9
	fmt.Println(floatNum)

	var floatNum1 float64 = 12345678.9
	fmt.Println(floatNum1)

	//N.B: can't perform arithmetic operation with 2 different types, howeever, you can typecast
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	// Division or operation between two type returns the type and rounded if needed
	var intNum1 int = 3
	var intNum2 = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	//String uses "" and ``. backquote allows for multilime text
	var myString string = "Hello world"
	var myString2 string = `Hello
World`
	var myString3 string = "Hello \nWorld"

	fmt.Println(myString)
	fmt.Println(myString2)
	fmt.Println(myString3)

	//A single quote represents a rune. Rune are weird
	var myRune rune = 'a'
	fmt.Println(myRune)

	//booleans
	var myBoolean bool = false
	fmt.Println(myBoolean)

	fancyShorthand := true
	fmt.Println(fancyShorthand)

	//for constance - doesn't change once created
	const myConst string = "const value"
	fmt.Println(myConst)

}
