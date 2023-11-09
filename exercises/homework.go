package homework

import (
	"math"
)

type Person struct {
	Name string
	Age  int
}

type Product struct {
	Name     string
	Category string
	Price    float64
}

func NamesToAge(persons []Person) map[string]int {
	out := make(map[string]int, len(persons))
	for _, p := range persons {
		out[p.Name] = p.Age
	}

	return out
}

func ProductsToCatalog(products []Product) map[string][]struct {
	Name  string
	Price float64
} {
	out := make(map[string][]struct {
		Name  string
		Price float64
	}, int(math.Min(1, float64(len(products))/2)))

	for _, p := range products {
		catalog, exists := out[p.Category]
		if !exists {
			catalog = make([]struct {
				Name  string
				Price float64
			}, 0, 4)
		}

		out[p.Category] = append(catalog, struct {
			Name  string
			Price float64
		}{
			Name:  p.Name,
			Price: p.Price,
		})
	}

	return out
}

type Citizen struct {
	Name string
	Age  int
}

func CityCitizensToDemographyByAge(citizensPerCity map[string][]Citizen) map[int]int {
	out := make(map[int]int, int(math.Min(1, float64(len(citizensPerCity))/2)))

	for _, citizens := range citizensPerCity {
		for _, citizen := range citizens {
			out[citizen.Age]++
		}
	}

	return out
}

type Student struct {
	Name        string
	Speciality  string
	Enrollment  int // Год поступления
	AverageMark float64
}

func StudentsToSpecialityYearlyAverage(students []Student) map[string]map[int]float64 {
	out := make(map[string]map[int]float64, 4)

	studentsPerYear := make(map[string]map[int]float64, 4)

	for _, student := range students {
		accMark, ok := out[student.Speciality]
		if !ok {
			accMark = map[int]float64{}
			out[student.Speciality] = accMark
		}

		accMark[student.Enrollment] += student.AverageMark
		enrolledStudents, ok := studentsPerYear[student.Speciality]
		if !ok {
			enrolledStudents = make(map[int]float64, 4)
			studentsPerYear[student.Speciality] = enrolledStudents
		}

		enrolledStudents[student.Enrollment]++
	}

	for spec, marksPerYear := range out {
		for year, mark := range marksPerYear {
			mark /= studentsPerYear[spec][year]
			marksPerYear[year] = mark
		}
	}

	return out
}
