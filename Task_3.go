package main

import (
	"fmt"
	"strings"
)

type Student struct {
	roll_number int
	name        string
	address     string
}

func add_student(roll_number int, name string, address string) *Student {
	s := new(Student)
	s.roll_number = roll_number
	s.name = name
	s.address = address
	return s
}

type Students_collection struct {
	list []*Student
}

func (ls *Students_collection) create_student(roll_number int, name string, address string) *Student {
	st := add_student(roll_number, name, address)
	ls.list = append(ls.list, st)
	return st
}

func print_student_l(st_ls *Students_collection) {
	student_list := st_ls.list
	for i := 0; i < len(st_ls.list); i++ {
		fmt.Printf("%s List %d %s \n", strings.Repeat("=", 10), i+1, strings.Repeat("=", 10))
		fmt.Println("Student Roll Number: ", student_list[i].roll_number)
		fmt.Println("Student Name: ", student_list[i].name)
		fmt.Println("Student Address: ", student_list[i].address)
	}
}

func main() {
	students := new(Students_collection)
	students.create_student(21, "Videl San", "Solaries")
	students.create_student(22, "Kal el", "Krypton")
	students.create_student(21, "doraemon", "Japan")
	students.create_student(22, "Toh Ruh", "OG Krypton")
	print_student_l(students)
}
