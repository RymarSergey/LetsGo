package main

type human struct {
	firstName  string
	secondName string
	age        uint8
}

type worker interface {
	getProfession() string
	getPosition() string
	getName() string
}
type doctor struct {
	human
	Profession string `json:"specialty" xml:"specialty"`
	Position   string `json:"list_of_patients,omitempty" xml:"list_of_patients"`
}

func (resiver doctor) getProfession() string {
	return resiver.Profession
}
func (resiver doctor) getPosition() string {
	return resiver.Position
}
func (resiver doctor) getName() string {
	return resiver.firstName
}

type manager struct {
	human
	Profession string `json:"specialty" xml:"specialty"`
	Position   string `json:"list_of_patients,omitempty" xml:"list_of_patients"`
}

func (resiver manager) getProfession() string {
	return resiver.Profession
}
func (resiver manager) getPosition() string {
	return resiver.Position
}
func (resiver manager) getName() string {
	return resiver.firstName
}

type sportsmen struct {
	human
	Sport      string `json:"sport" xml:"sport"`
	Score      string `json:"score" xml:"score"`
	Profession string `json:"specialty" xml:"specialty"`
	Position   string `json:"list_of_patients,omitempty" xml:"list_of_patients"`
}

func (resiver sportsmen) getProfession() string {
	return resiver.Profession
}
func (resiver sportsmen) getPosition() string {
	return resiver.Position
}
func (resiver sportsmen) getName() string {
	return resiver.firstName
}

type programmer struct {
	human
	Languages  []string `json:"languages" xml:"languages"`
	Profession string   `json:"specialty" xml:"specialty"`
	Position   string   `json:"list_of_patients,omitempty" xml:"list_of_patients"`
}

func (resiver programmer) getProfession() string {
	return resiver.Profession
}
func (resiver programmer) getPosition() string {
	return resiver.Position
}
func (resiver programmer) getName() string {
	return resiver.firstName
}

type teacher struct {
	human
	Field            string `json:"field" xml:"field"`
	NumberOfStudents int    `json:"number_of_students" xml:"number_of_students"`
	Profession       string `json:"specialty" xml:"specialty"`
	Position         string `json:"list_of_patients,omitempty" xml:"list_of_patients"`
}

func (resiver teacher) getProfession() string {
	return resiver.Profession
}
func (resiver teacher) getPosition() string {
	return resiver.Position
}
func (resiver teacher) getName() string {
	return resiver.firstName
}

func main() {

}
