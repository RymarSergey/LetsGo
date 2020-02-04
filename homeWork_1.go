package main

import (
	"fmt"
	"time"
)

type human struct {
	firstName  string
	secondName string
	age        int
}

func (receiver *human) happyBirthday() {
	fmt.Println("happyBirthday() from human ", time.Now().Year()-receiver.age)
}
func (receiver *human) sayName() {
	fmt.Printf("My name is %s %s", receiver.firstName, receiver.secondName)
}

type worker struct {
	*human
	salary int64
}

/*func (receiver *worker) happyBirthday() {
	fmt.Println("happyBirthday() from  worker ",time.Now().Year() - receiver.age)
}*/

type doctor struct {
	*worker
	Specialty      string   `json:"specialty" xml:"specialty"`
	ListOfPatients []string `json:"list_of_patients,omitempty" xml:"list_of_patients"`
}

type manager struct {
	*worker
	Department  string `json:"department" xml:"department"`
	CompanyName string `json:"company_name" xml:"company_name"`
}

type sportsmen struct {
	*worker
	Sport string `json:"sport" xml:"sport"`
	Score string `json:"score" xml:"score"`
}
type programmer struct {
	*worker
	Languages []string `json:"languages" xml:"languages"`
}
type teacher struct {
	*worker
	Field            string `json:"field" xml:"field"`
	NumberOfStudents int    `json:"number_of_students" xml:"number_of_students"`
}

func main() {
	mrBeen := doctor{
		worker: &worker{&human{
			firstName:  "Jac",
			secondName: "Been",
			age:        45,
		}, 4500},
		Specialty:      "surgeon",
		ListOfPatients: []string{"mrs.Tool"},
	}
	fmt.Printf("%+v\n", mrBeen.human.secondName)

	mrBeen.worker.happyBirthday()
	mrBeen.worker.human.happyBirthday()
	mrBeen.worker.sayName()
	mrBeen.worker.human.sayName()
}
