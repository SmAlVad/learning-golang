package main

import "fmt"

func main() {
	// размер массива является частью его типа

	// инициализация значениями по умолчанию
	var a1 [3]int // [0,0,0]

	const size = 2
	var a2 [2 * size]bool // [false,false,false,false]

	// определение размера при инициализации
	a3 := [...]int{1, 2, 3}

	// проверка при компиляции или при выполнении
	// invalid array index 4 (out of bounds for 3-element array)
	// a3[4] = 12
}
