package main

import "fmt"

func main() {

	var a int

	fmt.Scan(&a) // считаем переменную 'a' с консоли

	h := a % 10
	m := a % 60

	fmt.Println("It is", h, "hours", m, "minutes.")
}
