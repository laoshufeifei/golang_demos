package main

import (
	"fmt"
)

// Introducer is a interface to introdece self
// https://hackthology.com/object-oriented-inheritance-in-go.html
type Introducer interface {
	introduce()
}

// People is struct have name and age
type People struct {
	Name            string
	age             int
	introduceMethod func()
}

func newPeople(name string, age int) *People {
	p := &People{
		Name: name,
		age:  age,
	}
	p.introduceMethod = p.peopleIntroduce
	return p
}

func (p *People) peopleIntroduce() {
	fmt.Printf("Hi, my name is %s, my age is %d\n", p.Name, p.age)
}

func (p *People) introduce() {
	p.introduceMethod()
}

// Student is inherit people and with school name
type Student struct {
	People
	school string
}

func newStudent(name string, age int, school string) *Student {
	s := &Student{
		People: People{name, age, nil},
		school: school,
	}
	s.introduceMethod = s.studentIntroduce
	return s
}

func (s *Student) studentIntroduce() {
	fmt.Printf("Hello, I'm a student, my name is %s, I'm is %d, my school is %s\n", s.Name, s.age, s.school)
}

func introLouder(p Introducer) {
	fmt.Printf("LOUDER: ")
	p.introduce()
}

func main() {
	p := newPeople("Tom", 3)
	// p.introduce()
	introLouder(p)

	s := newStudent("Lili", 20, "TSU")
	// s.introduce()
	introLouder(s)
}
