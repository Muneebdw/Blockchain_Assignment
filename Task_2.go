package main

import (
	"fmt"
	"reflect"
)

type Employeeprofile struct {
	name   string
	salary int
	role   string
}

type CompanyProfile struct {
	company_name string
	employees    []Employeeprofile
}

func initiate() CompanyProfile {
	emp1 := Employeeprofile{
		name:   "Muneeb Bhalli",
		salary: 1234,
		role:   "Data analyst",
	}

	emp2 := Employeeprofile{
		name:   "Hammad Khan",
		salary: 5432,
		role:   "Lead QA Tester",
	}

	employees := []Employeeprofile{emp1, emp2}

	CompanyProfile := CompanyProfile{
		company_name: "AI FOR Everyone",
		employees:    employees,
	}
	return CompanyProfile
}

func main() {
	CompanyProfile := initiate()
	size := len(CompanyProfile.employees)
	fmt.Printf("%s\n", reflect.TypeOf(CompanyProfile).String())
	fmt.Printf("Company Name: %s\n", CompanyProfile.company_name)
	fmt.Println("   Employes: ", size)
	emp := CompanyProfile.employees
	for i := 0; i < size; i++ {
		fmt.Println("-Employee profile No. :", i+1)
		fmt.Println(" Emp Name : ", emp[i].name)
		fmt.Println(" Emp Salary : ", emp[i].salary)
		fmt.Println(" Emp role: ", emp[i].role)
	}
}
