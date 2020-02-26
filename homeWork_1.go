package main

import (
	"errors"
	"fmt"
	"sync"
)

type human struct {
	firstName  string
	secondName string
	age        uint8
}

type boss interface {
	bossPosition()
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
	Profession string `json:"specialty" xml:"specialty"`
	Position   string `json:"list_of_patients,omitempty" xml:"list_of_patients"`
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

type myCancurentMap struct {
	sync.RWMutex
	mapSync map[worker]bool
}

func main() {
	workerBoolMap := make(map[worker]bool)
	doctor1 := doctor{
		human: human{
			firstName:  "Sergey",
			secondName: "Rymar",
			age:        34,
		},
		Profession: "Doctor",
		Position:   "1",
	}
	workerBoolMap[doctor1] = true
	manager1 := manager{
		human: human{
			firstName:  "Denis",
			secondName: "Denisich",
			age:        18,
		},
		Profession: "manager",
		Position:   "boss",
	}
	workerBoolMap[manager1] = true
	sportsmen1 := sportsmen{
		human: human{
			firstName:  "Pavel",
			secondName: "Pavelovich",
			age:        25,
		},
		Sport:      "wrestling",
		Score:      "4",
		Profession: "sportsmen",
		Position:   "3",
	}
	workerBoolMap[sportsmen1] = true
	programmer1 := programmer{
		human: human{
			firstName:  "Sergey",
			secondName: "Sergeevich",
			age:        50,
		},
		Profession: "programmer",
		Position:   "4",
	}
	workerBoolMap[programmer1] = true
	teacher1 := teacher{
		human: human{
			firstName:  "Stas",
			secondName: "Stasovich",
			age:        23,
		},
		Field:            "Go",
		NumberOfStudents: 4,
		Profession:       "teacher",
		Position:         "boss",
	}
	workerBoolMap[teacher1] = true
	var wg sync.WaitGroup
	wg.Add(2) // в группе две горутины
	mWM := new(myCancurentMap)
	mWM.mapSync = workerBoolMap
	//var flag int
	work := func(workerBoolMap *myCancurentMap, indx int) {
		/*for {
			if indx == 1 {
				break
			}else {
				time.Sleep(time.Second*1)
				break
			}
		}*/
		defer wg.Done()
		workerBoolMap.Lock()
		for key := range workerBoolMap.mapSync {

			if key.getPosition() == "boss" && indx == 1 {
				fmt.Println("Gorutine -", indx, "  BOSS : ", key.getName(), key.getPosition())
			}
			if key.getPosition() != "boss" && indx == 2 {
				fmt.Println("Gorutine -", indx, "  Worker : ", key.getName(), key.getProfession())
			}
		}
		workerBoolMap.Unlock()

	}
	// вызываем горутины
	go work(mWM, 2)
	go work(mWM, 1)
	wg.Wait()

}

func workerToHuman(w worker) (h human, err error) {

	switch w.(type) {
	case doctor:
		d := w.(doctor)
		h.age = d.age
		h.firstName = d.firstName
		h.secondName = d.secondName
		err = nil
	case manager:
		d := w.(manager)
		h.age = d.age
		h.firstName = d.firstName
		h.secondName = d.secondName
		err = nil
		/*if d, ok := w.(manager); ok {
			h.age = d.age
			h.firstName = d.firstName
			h.secondName = d.secondName
		}*/
	case sportsmen:
		d := w.(sportsmen)
		h.age = d.age
		h.firstName = d.firstName
		h.secondName = d.secondName
		err = nil
		/*if d, ok := w.(sportsmen); ok {
			h.age = d.age
			h.firstName = d.firstName
			h.secondName = d.secondName
		}*/
	case programmer:
		d := w.(programmer)
		h.age = d.age
		h.firstName = d.firstName
		h.secondName = d.secondName
		err = nil
		/*if d, ok := w.(programmer); ok {
			h.age = d.age
			h.firstName = d.firstName
			h.secondName = d.secondName
		}*/
	case teacher:
		d := w.(teacher)
		h.age = d.age
		h.firstName = d.firstName
		h.secondName = d.secondName
		err = nil
		/*if d, ok := w.(teacher); ok {
			h.age = d.age
			h.firstName = d.firstName
			h.secondName = d.secondName
		}*/
	default:
		err = errors.New("Do not convert to human! ")
	}
	return
}
