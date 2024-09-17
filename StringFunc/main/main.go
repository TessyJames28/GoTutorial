package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Strings in go are indexed in utf-8 representation and returns the utf number value
// indexed can be skipped if a particular character is 2 bytes eg. 'é'. But to avoid the
// skipping of the index, we can cast the string as 'rune[]' which returns the exact unicode
// skipping an index. That's why the length of a string in go returns the length of the different
// byte codes instead of the actual length of the string
func main() {
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Println((myString))
	fmt.Println(indexed)         //This prints a number 114 which is the unicode for 'r'
	fmt.Println(string(indexed)) //typecasting allows us to print the exact value instead of the code

	//%T is used to print type
	fmt.Printf("%v, %T\n", indexed, indexed)

	//looping to seethe values
	for i, v := range myString {
		fmt.Println(i, v)
	}

	//Typecasting to rune allows us to print the exact index value without skipping
	var myString2 = []rune("résumé")
	//looping to seethe values
	for i, v := range myString2 {
		fmt.Println(i, v)
	}

	//getting the length
	var count = len(myString)                     // will give us 8 because it counts the number of bytes and 'é' is multi byte character
	var count2 = utf8.RuneCountInString(myString) //this return the exact rune count which is 6
	fmt.Printf("count using len: %v\ncount2 using RuneCountInString: %v\n", count, count2)

	// We can declare a rune type as
	var myRune = 'a'
	fmt.Printf("myRune = %v\n", myRune)

	//String building
	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var catStr = ""
	var stringBuilder strings.Builder
	for i := range strSlice {
		catStr += strSlice[i] // concantenation creates completly new string which is inefficient
		//instead we can use the stringBuilder from go inbuilt string package
		stringBuilder.WriteString(strSlice[i])
	}
	fmt.Printf("%v\n", catStr)
	//call the string method for the built string at the end
	catStr2 := stringBuilder.String() //Here arrays are allocated internally and values appended when called only at the end is a new string created
	fmt.Printf("%v\n", catStr2)

	//P.S: String are immutable in go, and thus can't be modified once created

}
