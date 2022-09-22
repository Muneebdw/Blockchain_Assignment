package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type Student struct {
	roll_number int
	name        string
	address     string
	subjects    []string
}

type Student_collection struct {
	list []*Student
}

func add_student(roll_number int, name string, address string, subjects []string) *Student {
	s := new(Student)
	s.roll_number = roll_number
	s.name = name
	s.address = address
	s.subjects = subjects
	return s
}

func (ls *Student_collection) create_student(roll_number int, name string, address string, subjects []string) *Student {
	st := add_student(roll_number, name, address, subjects)
	ls.list = append(ls.list, st)
	return st
}

func print_student_l(st_ls *Student_collection) {
	student_list := st_ls.list
	for i := 0; i < len(st_ls.list); i++ {
		fmt.Printf("%s List %d %s \n", strings.Repeat("=", 10), i+1, strings.Repeat("=", 10))
		fmt.Println("Student Roll Number: ", student_list[i].roll_number)
		fmt.Println("Student Name: ", student_list[i].name)
		fmt.Println("Student Address: ", student_list[i].address)
		fmt.Println("Student Subjects: ", student_list[i].subjects)
	}
}

func CalculateHash(string_to_hash string) string {
	fmt.Printf("String To Calculate Hash: %s\n", string_to_hash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(string_to_hash)))
}

func main() {
	students := new(Student_collection)
	students.create_student(21, "Videl San", "Solaries", []string{"Physics", "QuantumMechanics"})
	students.create_student(22, "Kal el", "Krypton", []string{"Biochem: failed", " Computing: failed"})
	students.create_student(21, "doraemon", "Japan", []string{"Quantum maths: failed", "Quantum physics: failed"})
	students.create_student(22, "Toh Ruh", "OG Krypton", []string{"basic maths: failed", "elementary physics: failed"})

	print_student_l(students)
	fmt.Println("Calculating Hash `TheBallgoesrollingdownthestair`: ", CalculateHash("TheBallgoesrollingdownthestair"))

}
