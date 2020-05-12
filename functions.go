package main

import "fmt"

// обычное объявление
func singleIn(in int) int {
	return in
}

// много параметров
func multIn(a, b int, c int) int {
	return a + b + c
}

// именнованный результат
func namedReturn() (out int) {
	out = 2
	return
}

// несколько резултатов
func multipleReturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some error happend")
	}
	return in, nil
}

// несколько именнованных результатов
func multipleNamedReturn(ok bool) (res int, err error) {
	if ok {
		err = fmt.Errorf("some error happend")
		// аналогично return res, err
		// или return 1, fmt.Errorf("some error happend")
		return
	}
	res = 2
	return
}

// не фиксированное количесво параметров
func sum(in ...int) (result int) {
	fmt.Printf("in := %#v \n", in)
	for _, val := range in {
		result += val
	}
	return
}

func main() {
	sum(1, 23, 3)
}
