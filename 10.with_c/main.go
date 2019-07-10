package main

/*
	write C code before import "C";
	cd lib && gcc -c my_lib.c -o my_lib.o && ar -crs libmy_lib.a my_lib.o;
	don't use go run xxx.go;
	usage: go build -o bin && ./bin;
	refer: http://bastengao.com/blog/2017/12/go-cgo-c.html
*/

/*
	// test simple code
	int myAdd1(int a, int b) {
		return a + b;
	}

	// test myAdd2
	#include "header.h"

	// test myAdd3
	#cgo CFLAGS: -I./lib
	#cgo LDFLAGS: -L./lib -lmy_lib
	#include "lib/my_lib.h"

	int sumArray(int* arr, int num) {
		int sum = 0;
		for (int i = 0; i < num; i++) {
			sum += arr[i];
		}
		return sum;
	}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("test go with c")
	a := C.int(1)
	b := C.int(2)

	ret1 := C.myAdd1(a, b)
	fmt.Printf("test1: %T %v\n", ret1, ret1) //main._Ctype_int 3

	ret2 := C.myAdd2(a, b)
	fmt.Printf("test2: %T %v\n", ret2, ret2)

	ret3 := C.myAdd3(a, b)
	fmt.Printf("test3: %T %v\n", ret3, ret3)

	goSlice := []C.int{1, 2, 3, 4, 5}
	fmt.Printf("goSlice is %T\n", goSlice)
	goPointer := unsafe.Pointer(&goSlice[0])
	cArray := (*C.int)(goPointer)
	ret := C.sumArray(cArray, 5)
	fmt.Println(ret)
}
