package main

import (
	"fmt"
	homework "lesson4/exercises"
)

func main() {
	homework.ExecuteMapsTasks()
	homework.ExecuteSliceTasks()

	ex1test1 := []homework.Person{
		{Name: "Alex", Age: 25},
		{Name: "Bob", Age: 30}}
	ex1res := homework.NamesToAge(ex1test1)
	fmt.Println("Exercise 1:", ex1res)

	ex2test1 := []homework.Product{
		{Name: "Laptop", Category: "Electronics", Price: 1000.5},
		{Name: "Shirt", Category: "Clothes", Price: 25.5}}
	ex2res := homework.ProductsToCatalog(ex2test1)
	fmt.Println("Exercise 2:", ex2res)

	ex3test1 := map[string][]homework.Citizen{
		"CityA": {
			{Name: "Alex", Age: 25},
			{Name: "Bob", Age: 25},
		},
		"CityB": {
			{Name: "Charlie", Age: 30},
		},
	}

	ex3res := homework.CityCitizensToDemographyByAge(ex3test1)
	fmt.Println("Exercise 3:", ex3res)

	ex4test1 := []homework.Student{
		{Name: "Alex", Speciality: "Math", Enrollment: 2020, AverageMark: 4.5},
		{Name: "Bob", Speciality: "Math", Enrollment: 2021, AverageMark: 4.0},
		{Name: "Sarah", Speciality: "Math", Enrollment: 2020, AverageMark: 4.9},
		{Name: "Charlie", Speciality: "Physics", Enrollment: 2020, AverageMark: 4.8},
	}

	ex4res := homework.StudentsToSpecialityYearlyAverage(ex4test1)
	fmt.Println("Exercise 4:", ex4res)

}
