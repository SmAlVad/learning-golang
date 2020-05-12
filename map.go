package main

func main() {
	// инициализация при создании
	var user map[string]string = map[string]string{
		"name":     "Vasily",
		"lastName": "Romanov",
	}

	// сразу с нужной емкостью
	profile := make(map[string]string, 10)

	// количество элементов
	mapLength := len(user)

	// если ключа нет - вернет значение по умолчанию для типа
	nName := user["middleName"]

	// проверка на существование ключа
	mName, mNameExist := user["middleName"]

	// пустая переменная - только проверяем знаяение
	_, mNameExist := user["middleName"]

	// удаление ключа
	delete(user, "lastName")
}
