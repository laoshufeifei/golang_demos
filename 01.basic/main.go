package main

import (
	"fmt"
	"math"
	"os"
	"time"
	"unsafe"
)

func main() {
	fmt.Println("print in one line")

	// base
	s1 := ""
	var s2 string
	var s3 = ""
	// var s4 string = "123" // have warning
	var i, j, k = 0, 1, 2
	fmt.Println(s1, s2, s3)
	fmt.Println(i, j, k)

	var f float64
	f = math.Pi
	fmt.Printf("test %%f: %f\n", f)
	fmt.Printf("test %%g: %g\n", f)
	fmt.Printf("test %%v: %v\n", f)

	const c = unsafe.Sizeof(f)
	fmt.Println(c)

	// for loop
	var sum int
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Printf("sum is %d\n", sum)

	// array
	arr1 := [...]int{3}
	arr2 := [...]int{1, 3: 2, 9} // {1, 0, 0, 2, 9}
	fmt.Println(len(arr1), len(arr2))

	// unicode
	unicode := '国'
	fmt.Printf("%X %[1]d %[1]c %[1]q\n", unicode) // 56FD 22269 国 '国'
	var letters []rune
	unicodeCh := "中文汉字"
	for i, s := range unicodeCh {
		fmt.Printf("i: %d, s: %c\n", i, s)
		letters = append(letters, s)
	}

	// string
	s1 = "hello, world"
	for i, r := range s1 {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	s2 = "你好，世界"
	fmt.Println("len(s2) is", len(s2))
	for i, r := range s2 {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	// array and slice
	week := [...]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}
	workdays := week[1:6]
	weekends := week[6:8]

	fmt.Printf("week is %T, %[1]v\n", week)
	fmt.Printf("workdays is %T, %[1]v, len = %d, cap = %d\n", workdays, len(workdays), cap(workdays))
	fmt.Printf("weekends is %T, %[1]v\n", weekends)

	intSlice := make([]int, 3)
	intSlice = append(intSlice, 1)
	intSlice = append(intSlice, 2, 3)
	// [0 0 0 1 2 3]
	fmt.Println(intSlice)
	fmt.Printf("intSlice is %T, %[1]v\n", intSlice)

	var times [3][0]int
	for range times {
		fmt.Println("hello")
	}

	// map
	ages := map[string]int{
		"alice":   3,
		"charlie": 1,
	}
	ages["tony"] = 21
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	fmt.Println(ages["bob"])
	if age, ok := ages["bob"]; !ok {
		fmt.Println("bob is not in ages", age)
	}

	// struct
	type Employee struct {
		ID   int
		Name string
		Age  int
	}

	var e0 Employee
	e1 := Employee{1, "xxx", 1}
	e2 := new(Employee) // point
	fmt.Printf("%T %T %T\n", e0, e1, e2)

	// function
	succ := testWriteFile("test.txt")
	fmt.Println("testWriteFile return", succ)

	// time format
	fmt.Printf("%s\n", time.Now().Format("Mon Monday monday xxx2006-1-02 15:04:05.000"))
}

func testWriteFile(name string) bool {
	newFile, err := os.Create(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error is %v\n", err)
		os.Exit(1)
	}
	defer newFile.Close()

	newFile.WriteString("abc")
	return true
}
