package main

import "fmt"

func main() {
	// создание
	// не инициализированный слайс
	var buf0 []int // len=0, cap=0
	// инициализированный
	buf1 := []int{} // len=0, cap=0

	buf2 := []int{42} // len=1, cap=1

	buf3 := make([]int, 0)     // len=0, cap=0 []
	buf4 := make([]int, 5)     // len=5, cap=5 [0,0,0,0,0]
	buf5 := make([]int, 5, 10) // len=5, cap=10 [0,0,0,0,0]

	// обращение к елементам
	someInt := buf2[0]

	// ошибка при выполнении
	// panic: runtime error: index out of range
	someOtherInt = buf2[1]

	// добавление элементов
	var buf []int           // len=0, cap=0
	buf = append(buf, 1, 3) // len=2, cap=2
}
