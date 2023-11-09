package main

import (
	"Homework_1_Brel/taskFour"
	"Homework_1_Brel/taskOne"
	"Homework_1_Brel/taskThree"
	"Homework_1_Brel/taskTwo"
	"fmt"
)

func main() {
	executeFirst()
	executeSecond()
	executeThird()
	executeFourth()
}

func executeFirst() {
	slice := []taskOne.Person{{Name: "Alex", Age: 25}, {Name: "Bob", Age: 30}}
	result := taskOne.Execute(slice)
	fmt.Println(result)
}

func executeSecond() {
	slice := []taskTwo.Product{{Name: "Laptop", Category: "Electronics", Price: 1000.5}, {Name: "Shirt", Category: "Clothes", Price: 25.5}}
	result := taskTwo.Execute(slice)
	fmt.Println(result)
}

func executeThird() {
	var inputMap = map[string][]taskThree.Citizen{
		"CityA": {
			{Name: "Alex", Age: 25},
			{Name: "Bob", Age: 25},
		},
		"CityB": {
			{Name: "Charlie", Age: 30},
		},
	}
	result := taskThree.Execute(inputMap)
	fmt.Println(result)
}

func executeFourth() {
	slice := []taskFour.Student{
		{Name: "Alex", Speciality: "Math", Enrollment: 2020, AverageMark: 4.5},
		{Name: "Bob", Speciality: "Math", Enrollment: 2021, AverageMark: 4.0},
		{Name: "Charlie", Speciality: "Physics", Enrollment: 2020, AverageMark: 4.8},
	}
	result := taskFour.Execute(slice)
	fmt.Println(result)
}
