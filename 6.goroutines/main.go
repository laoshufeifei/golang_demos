package main

import (
	"fmt"
	"time"
)

func nowTime() string {
	return time.Now().Format("15:04:05.000")
}

func main() {
	fmt.Printf("%s main: begin...\n", nowTime())

	go func() {
		fmt.Printf("\t%s routine1: begin...\n", nowTime())
		time.Sleep(900 * time.Millisecond)
		fmt.Printf("\t%s routine1: finish...\n", nowTime())
	}()

	go func() {
		fmt.Printf("\t\t%s routine2: begin...\n", nowTime())
		time.Sleep(600 * time.Millisecond)
		fmt.Printf("\t\t%s routine2: finish...\n", nowTime())
	}()

	time.Sleep(1 * time.Second)
	fmt.Printf("%s main: finish...\n", nowTime())
}
