package main

import (
	"fmt"
)

// People is struct have name and age
// https://hackthology.com/object-oriented-inheritance-in-go.html
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

func (p *People) sayName() {
	fmt.Printf("Hi, my name is %s\n", p.Name)
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
	fmt.Printf("Hello, I'm a student, my name is %s, my age is %d, my school is %s\n", s.People.Name, s.age, s.school)
}

func (s *Student) sayName() {
	fmt.Printf("Hello, my name is %s\n", s.Name)
}

//TODO: use interface
func introLoudly(p *People) {
	fmt.Printf("louder: ")
	p.introduce()
}

func main() {
	p := newPeople("Tom", 3)
	p.introduce()
	introLoudly(p)

	s := newStudent("Lili", 20, "TSU")
	s.introduce()
	// introLoudly(s) had error, must use interface to fix
}
