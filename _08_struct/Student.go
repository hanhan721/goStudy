package main

import "fmt"

type Student struct {
	name string
	age  int
}

func (s *Student) GetName() string {
	return s.name
}

func (s *Student) SetName(name string) {
	s.name = name
}

func (s *Student) GetAge() int {
	return s.age
}

func (s *Student) SetAge(age int) *Student {
	s.age = age
	return s
}

func (s *Student) String() string {
	return fmt.Sprintf("Student{Name: %s, Age: %d}", s.name, s.age)
}
