package homework

import (
	"fmt"
)

func ExecuteMapsTasks() {
	exerciseMaps1()
	exerciseMaps2()
	exerciseMaps3()
	exerciseMaps4()
	exerciseMaps5()
	exerciseMaps6()
	exerciseMapsHigh1()
	exerciseMapsHigh2()
	exerciseMapsMax1()
	exerciseMapsMax2()
}

func exerciseMaps1() {
	// Простое упражнение 1
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(m1["one"]) // 1
}

func exerciseMaps2() {
	// Простое упражнение 2
	m2 := map[int]bool{}
	fmt.Println(m2[2]) // false

}

func exerciseMaps3() {
	// Средней сложности упражнение 1
	m3 := map[string]*int{}
	val := 5
	m3["five"] = &val
	fmt.Println(*m3["five"]) // 5
}

func exerciseMaps4() {
	// Средней сложности упражнение 2
	m4 := map[int]string{1: "one", 2: "two", 3: "three"}
	delete(m4, 2)
	fmt.Println(len(m4)) // 2

}

func exerciseMaps5() {
	// Высокой сложности упражнение 1
	m5 := map[string]int{"a": 1, "b": 2, "c": 3}
	val, exists := m5["d"]
	fmt.Println(val, exists) // 0 false

}

func exerciseMaps6() {
	// Высокой сложности упражнение 2
	m6 := make(map[int]int)
	m6[1] = m6[1] + 1
	m6[2] = m6[2] + 2
	fmt.Println(m6)
	// Ответ [{1:1},{2:2}]
}

func exerciseMapsHigh1() {
	// Упражнение высокой сложности 1
	m1 := map[int]func() int{
		1: func() int { return 5 },
		2: func() int { return 10 },
	}
	fmt.Println(m1[1]() + m1[2]()) // 15

}

func exerciseMapsHigh2() {
	// Упражнение высокой сложности 2
	type person struct {
		age int
	}

	m2 := map[person]string{
		{age: 20}: "Alice",
		{age: 25}: "Bob",
	}
	p := person{age: 20}
	fmt.Println(m2[p]) // Alice

}

func exerciseMapsMax1() {
	// Упражнение максимальной сложности 1
	m3 := map[string]*int{}
	val := 5
	m3["key"] = &val
	val = 10
	fmt.Println(*m3["key"]) // 10
}

func exerciseMapsMax2() {
	// Упражнение максимальной сложности 2
	m4 := map[int]string{1: "a", 2: "b", 3: "c"}
	for k, v := range m4 {
		if k == 2 {
			m4[4] = "d"
		}
		if k == 3 {
			delete(m4, k)
		}
		fmt.Println(k, v)
	}
	fmt.Println(len(m4)) // 3
}
