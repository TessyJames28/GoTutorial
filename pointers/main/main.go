// Pointer - points to addresses in memory location
package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32

	//To get the value stored at the memory location, we can use the * symbol

	fmt.Printf("The value p points to is: %v\n", *p) //*p is said to be de-referencing the pointer
	fmt.Printf("The value of i is: %v\n", i)

	//To change the initial value -default value of the type - save in the pointer, we can initailize as below
	*p = 10
	fmt.Printf("The value p points to is: %v\n", *p)

	//We can also create a pointer from the address of another variable using '&'
	p = &i //This means we want the memory address of the variable and not its value
	*p = 1
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)

	//Using pointers in function
	var things1 = [5]float64{1, 2, 3, 4, 5}

	//Print memory location of thing1
	fmt.Printf("The memory location of the thing1 array is: %p\n", &things1)
	var result [5]float64 = square(&things1)
	fmt.Printf("The result is: %v\n", result)
}

//Using a pointer in function
func square(things2 *[5]float64) [5]float64 {
	//Print memory location of thing2
	fmt.Printf("The memory location of the thing2 array is: %p\n", things2)
	for i := range things2 {
		things2[i] = things2[i] * things2[i]
	}
	return *things2
}
