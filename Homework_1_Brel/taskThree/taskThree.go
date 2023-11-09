package taskThree

type Citizen struct {
	Name string
	Age  int
}

func Execute(inputMap map[string][]Citizen) map[uint]uint {
	result := make(map[uint]uint)

	for _, citizens := range inputMap {
		for _, citizen := range citizens {
			result[uint(citizen.Age)]++
		}
	}
	return result
}
