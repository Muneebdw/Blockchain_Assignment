package main

import (
	"fmt"
	"reflect"
)

type PAddr struct {
	country string
	city    string
	street  string
}

type Person struct {
	f_name   string
	l_name   string
	phone    string
	age      int
	is_alive bool
	PAddr  PAddr
}

func initiate() Person {
	person := Person{
		f_name: "Muneeb",
		l_name: "Bhalli",
		phone:  "+0317-7243088",
		age:    21, is_alive: true,
		PAddr: PAddr{ city: "Sialkot"},
	}
	return person
}

func reciever_func(person Person) {
	fmt.Printf("Recieved struct of type %s \n", reflect.TypeOf(person).String())
}

func main() {
	var person = initiate()
	fmt.Printf("Recieved a struct of type %s \n", reflect.TypeOf(person).String())
	fmt.Printf("First Name: %s \n", person.f_name)
	fmt.Printf("Last Name: %s \n", person.l_name)
	fmt.Printf("Passing the created structure to another function\n")
	reciever_func(person)
}
