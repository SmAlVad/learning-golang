package main

import "fmt"

func main() {

	// Значение по умолчанию
	var num0 int

	// Значение при инициализации
	var num1 = 1

	// Пропуск типа
	var num2 = 20

	// Короткое объявление переменной
	// только для новых переменных
	num3 := 30

	// num1 += 1
	// Нет ++num1
	num1++

	// camelCase стиль
	userIndex := 10

	// Объявление нескольких переменных
	var weight, height = 11, 21

	// Приваивание в существующее значение
	weight, height = 12, 13

	// Короткое присваивание,
	// хотя бы одна переменная должна быть новой
	weight, age := 40, 59

	fmt.Println(num0, num1, num2, num3, userIndex, weight, height, age)
}
