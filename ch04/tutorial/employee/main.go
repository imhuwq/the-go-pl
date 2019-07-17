package main

import "time"

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func EmployByID(id int) *Employee {
	e := &Employee{}
	e.ID = id
	return e
}

func EmployByID2(id int) Employee {
	e := Employee{}
	e.ID = id
	return e
}

func ModifyEmploy(e *Employee) {
	e.Salary = 10000000
}

func ModifyEmploy2(e Employee) {
	e.Salary = 1000
}

func main() {
	EmployByID(1).Salary = 10000000
	//EmployByID2(1).Salary = 1000
	e := EmployByID2(2)
	e.Salary = 1000
}
