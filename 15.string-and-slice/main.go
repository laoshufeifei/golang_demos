package main

// https://books.studygolang.com/advanced-go-programming-book/ch1-basic/ch1-03-array-string-and-slice.html
// 参考以上网址的动手代码，版权归原作者所有

import (
	"fmt"
	"unicode/utf8"
)

func delOneFromHeader() {
	fmt.Printf("\n")
	// a = a[1:]
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("s[:0] is", s[:0])

	fmt.Println("s is", s)
	s = append(s[:0], s[1:]...)
	fmt.Println("s is", s)

	fmt.Printf("\n")

	s = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ret := copy(s, s[1:])
	fmt.Println("copy(s, s[1:]) return is", ret)
	fmt.Println("s is", s)
	s = s[:ret]
	fmt.Println("s is", s)
}

func delNFromHeader() {
	fmt.Printf("\n")

	N := 3
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = append(s[:0], s[N:]...)
	fmt.Println("s is", s)

	fmt.Printf("\n")

	s = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = s[:copy(s, s[N:])]
	fmt.Println("s is", s)
}

func delNFromMdiddle() {
	fmt.Printf("\n")
	i := 3
	N := 2
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("s is", s)
	s = append(s[:i], s[i+N:]...)
	fmt.Println("s is", s)

	s = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ret := copy(s[i:], s[i+N:])
	fmt.Println("copy(s[i:], s[i+N:]) return is", ret)
	fmt.Println("s is", s)
	s = s[:i+ret]
	fmt.Println("s is", s)
}

func inserOne() {
	i := 2
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = 22
	fmt.Println("s is", s)
}

func inserN() {
	i := 2
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := []int{22, 23, 24}
	N := len(s2)

	s = append(s, s2...)
	copy(s[i+N:], s[i:])
	copy(s[i:], s2)
	fmt.Println("s is", s)
}

func string2bytes(s string) []byte {
	fmt.Printf("[]byte(%s) is %v\n", s, []byte(s))

	bs := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		bs[i] = s[i]
	}

	fmt.Println("self convert result is", bs)
	return bs
}

func string2runes(s []byte) []rune {
	var p []rune
	for len(s) > 0 {
		r, size := utf8.DecodeRune(s)
		p = append(p, rune(r))
		s = s[size:]
	}
	return p
}

func runes2string(s []int32) string {
	var p []byte
	buf := make([]byte, 3)
	for _, r := range s {
		n := utf8.EncodeRune(buf, r)
		p = append(p, buf[:n]...)
	}
	return string(p)
}

func main() {
	// 暂时不考虑越界的情况
	delOneFromHeader()
	delNFromHeader()
	delNFromMdiddle()

	inserOne()
	inserN()

	string2bytes("abc中文")
	string2runes([]byte("abc中文"))
	runes2string([]rune("abc中文"))
}
