package main

type UserId int

func main()  {
	idx := 1
	var uid UserId = 43

	// даже если базовый тип одинаковый, разные типы не совместимы
	// myID := idx

	myID := UserId(idx)

	println(uid, myID)
}
