// Go supports _methods_ defined on struct types.
// refer: https://gobyexample.com/methods

package main

import "fmt"

type rect struct {
	width, height int
}

// This `area` method has a _receiver type_ of `*rect`.
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value
// receiver types. Here's an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

type coloredRect struct {
	rect
	colored string
}

func (r coloredRect) color() string {
	return r.colored
}

func main() {
	r := rect{width: 10, height: 5}

	// Here we call the 2 methods defined for our struct.
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values
	// and pointers for method calls. You may want to use
	// a pointer receiver type to avoid copying on method
	// calls or to allow the method to mutate the
	// receiving struct.
	rp := &rect{width: 2, height: 5}
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())

	cr := coloredRect{rect{2, 3}, "red"}
	fmt.Println("area:", cr.area())
	fmt.Println("perim:", cr.perim())
	fmt.Println("color:", cr.color())
}
