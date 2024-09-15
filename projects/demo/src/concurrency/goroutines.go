package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex = sync.Mutex{}
var waitGroup = sync.WaitGroup{}
var results = []int{}

func Goroutines() {
	t0 := time.Now()
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go dbCall(i)
	}
	waitGroup.Wait()
	fmt.Printf("Total execution time: %v\n", time.Since(t0))
	fmt.Printf("Results: %v\n", results)
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 10
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("Database query executed successfully: ", i)
	mutex.Lock()
	results = append(results, i)
	mutex.Unlock()
	waitGroup.Done()
}
