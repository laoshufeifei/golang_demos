package main

import (
	"fmt"
)

// People is struct have name and age
type People struct {
	Name string
	age  int
}

func newPeople(name string, age int) *People {
	p := &People{
		Name: name,
		age:  age, // must have comma(,)
	}
	return p
}

func (p *People) sayName() {
	fmt.Printf("Hi, my name is %s\n", p.Name)
}

func (p *People) sayAge() {
	fmt.Printf("Hi, I am %d\n", p.age)
}

// Student is inherit people and with school name
type Student struct {
	People
	school string
}

func newStudent(name string, age int, school string) *Student {
	s := &Student{
		People: People{name, age},
		school: school,
	}
	return s
}

func (s *Student) sayName() {
	fmt.Printf("Hello, my name is %s\n", s.Name)
}

func (s *Student) saySchool() {
	fmt.Printf("Hello, my school is %s\n", s.school)
}

func main() {
	p1 := newPeople("Tom", 3)
	p1.sayName()
	p1.sayAge()

	s1 := newStudent("Lili", 20, "TSU")
	s1.sayName()
	s1.People.sayName()
	s1.sayAge()
	s1.saySchool()
}
