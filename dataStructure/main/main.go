package main

import (
	"fmt"
	"time"
)

func main() {
	//Array declaration. int32 stores 4bytes of memory so go initialize 12 bytes for this array
	var intArr [3]int32
	intArr[1] = 123

	//Array declaration and initialization
	var intArr2 [3]int32 = [3]int32{1, 2, 3}
	// intArr2 := [3]int32{1,2,3}
	//intArr2 := [...]int32{1,2,3}

	//assessing elements via index
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	fmt.Println(intArr2)

	//You print the memory location of the elements using the & symbol
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	//Slices - By ommitting the length value, we now have a slice
	//Slices are wrappers arrays with additional functions under the hood
	//You can values to slices using append unlike arrays
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)

	//cap function gets the capacity of the array
	fmt.Printf("The lengthe is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7) //Appends to the end of the slice
	fmt.Println(intSlice)
	fmt.Printf("The lengthe is %v with capacity %v\n", len(intSlice), cap(intSlice))

	//You can append to the slices through expansion
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	//We can create slice using the make function, allowing us to optionally specify the length and capacity therwise, the length and capacity will be the same
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	//Map - a set of key:value pair. You can look up a value with its key
	// You can create the map with make function
	var myMap map[string]uint8 = make(map[string]uint8) //this means the keys are type string and the values are unsigned int
	fmt.Println(myMap)

	//map initialization
	var myMap2 = map[string]uint8{"Adam": 32, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Sarah"])
	//N.B: If you try to access a key not int the map, it return the default value of value type, in this case 0

	//Map in go also returns an optional second value, which is a boolean that checks if thevalues in the map
	var age, ok = myMap2["Jason"]
	//Allows you to check the value before printing
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Printf("Invalid Name\n")
	}

	//Looping. We can use a range keyword within a for loop
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	//We can iterate over an array or slices like below
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	//Go do not have a while loop but this can be achieved with the for keyword
	var i int = 0
	for i < 10 {
		if i+1 == 10 {
			fmt.Printf("%v", i)
		} else {
			fmt.Printf("%v, ", i)
		}
		i = i + 1
	}

	//We can also omit the condition and use the break keyword
	var j int = 0
	for {
		if j >= 10 {
			break
		}
		fmt.Println(j)
		j += 1
	}

	//We can also achieve the above as the initailization, the condition, and the post
	for m := 0; m < 10; m++ {
		fmt.Printf("%v ", m)
	}

	//Function call
	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)
	fmt.Printf("\nTotal time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))
}

// Speed test showing the benefits of setting the capacity of a slice over realocation
func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}

	return time.Since(t0)
}
