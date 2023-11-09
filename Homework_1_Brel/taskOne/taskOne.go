package taskOne

type Person struct {
	Name string
	Age  int
}

func Execute(slice []Person) map[string]uint {
	result := make(map[string]uint)
	for _, value := range slice {
		result[value.Name] = uint(value.Age)
	}
	return result
}
