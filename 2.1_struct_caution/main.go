// ref: http://blueskykong.com/2019/01/25/go-mistakes-2/
// ref: https://books.studygolang.com/gopl-zh/ch3/ch3.html

package main

import "fmt"

type data struct {
	num   int
	key   *string
	items map[string]bool
	str   string
}

func (this *data) pointerFunc() {
	this.num = 7
	this.str = "a"
	*this.key = "valueFunc.key1"
	this.items["valueFunc"] = false
}

// num and str will not change by this function
func (this data) valueFunc() {
	this.num = 8
	this.str = "b"
	*this.key = "valueFunc.key2"
	this.items["valueFunc"] = true
}

func main() {
	key := "key1"

	d := data{1, &key, make(map[string]bool), "0"}
	fmt.Printf("num=%v  key=%v  items=%v  str=%v\n", d.num, *d.key, d.items, d.str)

	d.pointerFunc() // 都能修改
	fmt.Printf("num=%v  key=%v  items=%v  str=%v\n", d.num, *d.key, d.items, d.str)

	d.valueFunc() // 只能修改 key 和 items 的值
	fmt.Printf("num=%v  key=%v  items=%v  str=%v\n", d.num, *d.key, d.items, d.str)
}
