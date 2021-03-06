package main

const pi = 3.141

const (
	hello = "Привет"
	e = 2.718
)

const (
	zero = iota
	_          	// пустая переменная, пропуск iota
	three		// = 3
)

const (
	_ = iota						// пропускаем первое занчение
	KB uint64 = 1 << (10 * iota) 	// 1024
	MB								// 1048576
)

const (
	// нетипизированная котстанта
	year = 2017
	// типизированная константа
	yearTyped int = 2017
)

func main()  {

}
