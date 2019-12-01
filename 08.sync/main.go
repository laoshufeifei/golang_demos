package main

import (
	"fmt"
	"sync"
	"time"
)

func waitTest() {
	fmt.Println("wait test begin")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			// do some work
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Printf("%d", n)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Printf("\nwait test finish\n")
}

func mutexTest() {
	fmt.Println("mutex test begin")
	sum := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for i := 0; i < 10000; i++ {
			// atomic.AddUint64(&sum, 1)
			mu.Lock()
			sum++
			mu.Unlock()
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 10000; i++ {
			mu.Lock()
			sum++
			mu.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("sum is", sum)
	fmt.Println("mutex test finish")
}

func main() {
	waitTest()
	mutexTest()

	var rw sync.RWMutex
	rw.RLock()
	rw.RUnlock()
	rw.Lock()
	rw.Unlock()
}
