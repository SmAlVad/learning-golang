package main

func main() {

	// int платформенный тип 32/64
	var i int = 60

	// автоматически выбранный int
	var autoInt = -60

	// int8, int16, int32, int64
	var bigInt int64 = 1<<32 - 1

	// платфомозависимый 32/64
	var unsignInt uint = 100500

	// int8, int16, int32, int64
	var unsignBigInt uint64 = 1<<64 - 1

	// fload32, float64
	var pi float32 = 3.141
	var e = 2.789
	goldenRation := 1.678

	// bool
	var b bool // false по-умолчанию
	var isOk bool = true
	var success = true
	cond := false

	// complex64, comples128
	var c complex128 = -1.1 + 7.12i
	c2 := -1.1 + 7.12i

}
