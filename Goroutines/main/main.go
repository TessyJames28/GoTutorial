// Goroutines - This a way to run multiple functions and have them exceute concurrently
package main

import (
	"fmt"
	// "math/rand"
	"sync"
	"time"
	//for wait function
)

var mu = sync.RWMutex{}   //For reading and writing to Mutex
var m = sync.Mutex{}      //this allows us to write to our program in a way that is safe in a concurrent program to avoid memory corruption
var wg = sync.WaitGroup{} //creating a wait group - they are like counters
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

// Instead of printing the result, we could return them to the main function like below
var results = []string{} //create a slice database

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		//to make out program run concurrently, we add the 'go' keyword front of the funcation call
		wg.Add(1) //Whenever we creat a new go routine, we make sure to add 1 to the counter
		go dbCall(i)
	}
	wg.Wait() // Here we call the wait method - which waits to see that all the task has been completed, then the rest of the code will execute
	fmt.Printf("Total execution time: %v\n", time.Since(t0))
	fmt.Printf("The results are %v\n", results)
}

func dbCall(i int) {
	//simulate db call delay
	// var delay float32 = rand.Float32() * 2080
	var delay float32 = 2000
	// N.B: Where you put your lock matters. SO it's not advisable to place it here as it will destroy the program concurrency
	time.Sleep((time.Duration(delay) * time.Millisecond))
	// fmt.Println("The result from the database is: ", dbData[i])
	//Here we apprend the result to our fake database
	// m.Lock() // allows the go routine to lock it when writing to it and others will have to wait for it to be relase before locking it themselves
	// results = append(results, dbData[i])
	// m.Unlock() //unlocks it once the go routine is done writing to it

	save(dbData[i])
	log()
	wg.Done() //Call done method at the end - this decrement the counter
}

// More Mutex function - short for mutual exclusion
func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

// Read lock and unlock
func log() {
	mu.RLock()
	fmt.Printf("The current results are: %v\n", results)
	mu.RUnlock()
}
