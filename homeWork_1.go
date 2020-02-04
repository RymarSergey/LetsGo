package main

import "fmt"

type human struct {
	firstName  string
	secondName string
	age        uint8
}

type worker struct {
	human
	salary int64
}
type doctor struct {
	worker
	Specialty      string   `json:"specialty" xml:"specialty"`
	ListOfPatients []string `json:"list_of_patients,omitempty" xml:"list_of_patients"`
}
type manager struct {
	worker
	Department  string `json:"department" xml:"department"`
	CompanyName string `json:"company_name" xml:"company_name"`
}

type sportsmen struct {
	worker
	Sport string `json:"sport" xml:"sport"`
	Score string `json:"score" xml:"score"`
}
type programmer struct {
	worker
	Languages []string `json:"languages" xml:"languages"`
}
type teacher struct {
	worker
	Field            string `json:"field" xml:"field"`
	NumberOfStudents int    `json:"number_of_students" xml:"number_of_students"`
}

func main() {
	callPanic()
	callPanic1()

}

func callPanic1() {
	var panicProgramer *programmer
	panicProgramer.Languages = []string{"Go", "Java", "C++"}
	fmt.Printf("I know the following languages %v", panicProgramer.Languages)
}

func callPanic() {
	var panicTeacher *teacher
	panicTeacher.NumberOfStudents = 2
}
