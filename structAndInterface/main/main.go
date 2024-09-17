// Structs and Interfcaes
package main

import "fmt"

//Struct is a way of defining our own type and we do this using type
//struct can hold mixed types, including a follow struct in the form of fields

//Gas engine
type gasEngine struct {
	mpg       uint8
	gallon    uint8
	ownerInfo owner // we can just use the struct type owner without the name ownerInfo and access it using myEngine.name
}

//Struct also has a method which has a direct access to struct instances itself
func (e gasEngine) milesLeft() uint8 {
	return e.gallon * e.mpg
}

//Electric Engine
type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

//Electric engine method
func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type owner struct {
	name string
}

//We can also define an anonymous struct without a name. With an anonymous struct,
//we can define and initialize it at the same location. However, it's not reusable
var myEngine2 = struct {
	mpg    uint8
	gallon uint8
}{25, 15}

//passing the new type to functions
// func canMakeIt(e gasEngine, miles uint8) {
// 	if miles < e.milesLeft() {
// 		fmt.Println("You can make it there!")
// 	} else {
// 		fmt.Println("Need to fuel up first!")
// 	}
// }

//Interfaces
//To make the canMakeIt function more general to take in any engine, we make use of interface
type engine interface {
	//it takes in the function signature eg milesLeft. Thus, any struct with this methods
	//can be accessed via the interface
	milesLeft() uint8
}

//We can now modify the canMakeIt function by using the interface instead of the gasEngine struct
func canMakeIt(e engine, miles uint8) {
	if miles < e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to charge up first!")
	}
}

func main() {
	//You can set values for structs using the struct literal syntax like below or we
	//can omit the fields name and they are assined values in order. eg gasEngine{25, 15}
	//or using the dot notation.
	var myEngine gasEngine = gasEngine{mpg: 25, gallon: 15, ownerInfo: owner{"Alex"}}
	var myElecEngine electricEngine = electricEngine{25, 15}
	fmt.Println(myEngine.mpg, myEngine.gallon, myEngine.ownerInfo.name)

	// for anonymous struct
	fmt.Println(myEngine2.mpg, myEngine2.gallon)

	//calling struct method
	fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())

	//Calling the function with the interface
	canMakeIt(myEngine, 50)
	canMakeIt(myElecEngine, 200)
}
