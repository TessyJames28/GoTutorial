// Channels in go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3 // track the price of tofu too, using a select statment

func main() {
	//declaring channel using the make method, chan keyword and the type of data it will hold
	// we can create buffer channel by adding a buffer after the type
	//Channels are used properly in conjunction with Go routine
	// the buffer is important as it will allow the process function to finsih it work quickly
	// and exist to allow main function to do its work, instead of waiting for main function to finish its work
	var c = make(chan int, 5) //holds a single integer

	go process(c) // call the go routine

	//we can iterate over the channel with a range keywords to print its result
	for i := range c {
		fmt.Println(i) //print value that's set by the go routine. Here we directly print the value popout from the channel rather then setting it to a variable first
		time.Sleep(time.Second * 1)
	}

	// //N.B - when you read from an unbuffered channel, you go execution will stuck there waiting for
	// // something to read from it. Go helps us end the program by throwing a deadlock error
	// c <- 1      // asign the value 1 to the channel
	// var i = <-c // pop the value from the channel and assign it to i leaving the channel empty
	// fmt.Println(i)

	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

// Function that go through the websites to check for chicken prices
func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

// Function that go through the websites to check for tofu prices
func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}

// Function that sends a message when a deal is found
func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	// fmt.Printf("Found a deal on chicken at %s\n", <-chickenChannel)

	//So instead of a text like the above, we send ourself an email using select statement
	select {
	case website := <-chickenChannel:
		fmt.Printf("Text Sent: Found a deal on chicken at %v\n", website)
	case website := <-tofuChannel:
		fmt.Printf("Text Sent: Found a deal on tofu at %v\n", website)
	}
}

func process(c chan int) {
	// c <- 123

	//looping and asign value to channel
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Existing Process...")
	close(c) //this notify the go routine and any other function that we are done and it can exit the loop
}
