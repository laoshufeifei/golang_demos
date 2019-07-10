package main

import (
	"fmt"
	"time"
)

func main() {
	basic()
	channel1()
	channel2()
	pipeline()
}

func nowTime() string {
	return time.Now().Format("15:04:05.000")
}

////////////////////////////////////////////////////////////////////////////
func basic() {
	chWithNoBuff := make(chan string)
	fmt.Println("len:", len(chWithNoBuff), "cap:", cap(chWithNoBuff))

	chWithBuff := make(chan int, 3)
	chWithBuff <- 1
	chWithBuff <- 2
	fmt.Println("len:", len(chWithBuff), "cap:", cap(chWithBuff))
	fmt.Println(<-chWithBuff) // 1
	close(chWithBuff)
	fmt.Println(<-chWithBuff) // 2
	fmt.Println(<-chWithBuff) // 0
	fmt.Println(<-chWithBuff) // 0
}

////////////////////////////////////////////////////////////////////////////
func channel1() {
	messages := make(chan string)
	go func() {
		messages <- "ping"
	}()
	time.Sleep(2 * time.Second)
	msg := <-messages
	fmt.Println(msg)
}

////////////////////////////////////////////////////////////////////////////
func channel2() {
	done := make(chan struct{})
	fmt.Printf("%s main: begin...\n", nowTime())

	go func() {
		fmt.Printf("\t%s routine1: begin...\n", nowTime())
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("\t%s routine1: finish...\n", nowTime())
		done <- struct{}{}
	}()

	go func() {
		fmt.Printf("\t\t%s routine2: begin...\n", nowTime())
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("\t\t%s routine2: finish...\n", nowTime())
		done <- struct{}{}
	}()

	// time.Sleep(1 * time.Second)
	<-done
	<-done
	fmt.Printf("%s main: finish...\n", nowTime())
}

////////////////////////////////////////////////////////////////////////////
func pipeline() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals, 5)
	go squarer(squares, naturals)
	printer(squares)
}

func counter(out chan<- int, maxCounter int) {
	for x := 0; x < maxCounter; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
