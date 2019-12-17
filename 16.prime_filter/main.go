package main

// https://github.com/chai2010/advanced-go-programming-book/blob/master/ch1-basic/ch1-06-goroutine.md

import (
	"fmt"
)

// 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	// // 2, 3, 4, 5, 6, 7...
	// originCh := make(chan int)
	// go func() {
	// 	for i := 2; ; i++ {
	// 		originCh <- i
	// 	}
	// }()

	// prime2 := <-originCh
	// fmt.Println(prime2)
	// filter2Ch := make(chan int)
	// go func() {
	// 	for {
	// 		i := <-originCh
	// 		if i%prime2 != 0 {
	// 			// without 2, 4, 6, 8...
	// 			// 3, 5, 7, 9, 11...
	// 			filter2Ch <- i
	// 		}
	// 	}
	// }()

	// prime3 := <-filter2Ch
	// fmt.Println(prime3)
	// filter3Ch := make(chan int)
	// go func() {
	// 	for {
	// 		i := <-filter2Ch
	// 		if i%prime3 != 0 {
	// 			// without 3, 9...
	// 			// 5, 7
	// 			filter3Ch <- i
	// 		}
	// 	}
	// }()

	// prime5 := <-filter3Ch
	// fmt.Println(prime5)
	// filter5Ch := make(chan int)
	// go func() {
	// 	for {
	// 		i := <-filter3Ch
	// 		if i%prime5 != 0 {
	// 			// without 5, 15
	// 			// 7
	// 			filter5Ch <- i
	// 		}
	// 	}
	// }()

	// prime7 := <-filter5Ch
	// fmt.Println(prime7)

	ch := GenerateNatural() // 自然数序列: 2, 3, 4, ...
	for i := 0; i < 10; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime) // 基于新素数构造的过滤器
	}
}
