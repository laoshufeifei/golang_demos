// ref: http://blueskykong.com/2019/01/25/go-mistakes-2/
// ref: http://blueskykong.com/2019/01/29/go-mistakes-3/
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

func testStruct() {
	key := "key1"

	d := data{1, &key, make(map[string]bool), "0"}
	fmt.Printf("num=%v  key=%v  items=%v  str=%v\n", d.num, *d.key, d.items, d.str)

	d.pointerFunc() // 都能修改
	fmt.Printf("num=%v  key=%v  items=%v  str=%v\n", d.num, *d.key, d.items, d.str)

	d.valueFunc() // 只能修改 key 和 items 的值
	fmt.Printf("num=%v  key=%v  items=%v  str=%v\n", d.num, *d.key, d.items, d.str)
}

func main() {

	// array or slice to function param
	testArrayAndSlice()
	testStruct()
}

func testArrayAndSlice() {
	// have error
	arr1 := [3]int{1, 2, 3}
	func(arr [3]int) {
		arr[0] = 9
		fmt.Println("in function", arr)
	}(arr1)
	fmt.Println("in main", arr1)

	arr2 := [3]int{1, 2, 3}
	func(arr *[3]int) {
		arr[0] = 9
		fmt.Println("in function", arr)
	}(&arr2)
	fmt.Println("in main", arr2)

	// slc is a slice
	slc := []int{1, 2, 3}
	func(slc []int) {
		slc[0] = 9
		fmt.Println("in function", slc)
	}(slc)
	fmt.Println("in main", slc)

	fmt.Printf("%T %T %T\n", arr1, arr2, slc)

	// v is index
	arr := []string{"a", "b", "c"}
	for v := range arr {
		fmt.Println(v) // 1 2 3
	}

	// v is value
	for _, v := range arr {
		fmt.Println(v) // "a", "b", "c"
	}

	// old slice and new slice
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1) // 3 3 [1 2 3 ]

	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2) // 2 2 [2 3]

	for i := range s2 {
		s2[i] += 20
	}
	// 此时的 s1 与 s2 是指向同一个底层数组的
	fmt.Println(s1) // [1 22 23]
	fmt.Println(s2) // [22 23]

	s2 = append(s2, 4) // 向容量为 2 的 s2 中再追加元素，此时将分配新数组来存

	for i := range s2 {
		s2[i] += 10
	}
	fmt.Println(s1) // [1 22 23]	// 此时的 s1 不再更新，为旧数据
	fmt.Println(s2) // [32 33 14]
}
