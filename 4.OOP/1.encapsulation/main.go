package main

import (
	"fmt"
)

// People is struct have name and age
type People struct {
	Name             string
	age              int
	introduceOneSelf func()
}

// NewPeople will init people, and return a pointer
func NewPeople(name string, age int) *People {
	// p1 := new(People)
	// p1.Name = name
	// p1.age = age
	p := &People{
		Name: name,
		age:  age, // must have comma(,)
	}
	return p
}

// sayName will print people.name
func (p *People) sayName() {
	fmt.Printf("Hi, my name is %s\n", p.Name)
}

// sayAge is private method out of this package
func (p *People) sayAge() {
	fmt.Printf("Hi, I am %d\n", p.age)
}

func (p *People) introduceNameAndAge() {
	fmt.Printf("Hi, my name is %s, my age is %d\n", p.Name, p.age)
}

func (p *People) introduceOnlyName() {
	fmt.Printf("Hi, my name is %s, my age is secret\n", p.Name)
}

func main() {
	p0 := People{Name: "Jack", age: 5}
	p0.sayName()
	p0.sayAge()

	p1 := NewPeople("Tom", 3)
	p1.sayName()
	p1.sayAge()

	p2 := NewPeople("Lili", 6)
	p2.Name = "Luck"
	p2.sayName()
	p2.age = 8 // if in other package, age is private
	p2.sayAge()

	p2.introduceOneSelf = p2.introduceNameAndAge
	p2.introduceOneSelf()
	p2.introduceOneSelf = p2.introduceOnlyName
	p2.introduceOneSelf()
}
