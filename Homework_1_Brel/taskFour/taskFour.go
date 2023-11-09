package taskFour

type Student struct {
	Name        string
	Speciality  string
	Enrollment  int
	AverageMark float64
}

type studentsAverageMark struct {
	count   float64
	sum     float64
	average float64
}

func (s *studentsAverageMark) addStudent(student Student) {
	s.count++
	s.sum += student.AverageMark
	s.average = s.sum / s.count
}

func Execute(slice []Student) map[string]map[int]float64 {
	averageMap := make(map[string]map[int]studentsAverageMark)
	resultMap := make(map[string]map[int]float64)
	for _, value := range slice {
		if averageMap[value.Speciality] == nil {
			averageMap[value.Speciality] = make(map[int]studentsAverageMark)
			resultMap[value.Speciality] = make(map[int]float64)
		}
		temp := averageMap[value.Speciality][value.Enrollment]
		(&temp).addStudent(value)
		averageMap[value.Speciality][value.Enrollment] = temp
		resultMap[value.Speciality][value.Enrollment] = averageMap[value.Speciality][value.Enrollment].average
	}
	return resultMap
}
