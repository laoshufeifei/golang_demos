package main

// https://books.studygolang.com/gopl-zh/ch8/ch8-08.html

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("running go routine number", runtime.NumGoroutine())
		fmt.Fprintf(os.Stderr, "read dir %s have error: %v\n", dir, err)
		return nil
	}
	return entries
}

func walkDir(dir string, fileSizes chan<- int64, waitGroup *sync.WaitGroup) {
	// fmt.Println("will work dir for ", dir)
	defer waitGroup.Done()

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			waitGroup.Add(1)
			go walkDir(subdir, fileSizes, waitGroup)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func main() {
	roots := []string{"../../../"}

	fileSizes := make(chan int64)
	var waitGroup sync.WaitGroup
	for _, root := range roots {
		waitGroup.Add(1)
		go walkDir(root, fileSizes, &waitGroup)
	}

	go func() {
		waitGroup.Wait()
		close(fileSizes)
	}()

	var fileCount, totalSize int64
	var tick <-chan time.Time
	tick = time.Tick(500 * time.Millisecond)

loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			fileCount++
			totalSize += size
		case <-tick:
			showSize(fileCount, totalSize)
		}
	}

	showSize(fileCount, totalSize)
}

func showSize(count, size int64) {
	fmt.Println("running go routine number", runtime.NumGoroutine())
	fmt.Printf("%d files %.2f MB\n", count, float64(size)/1e6)
}
