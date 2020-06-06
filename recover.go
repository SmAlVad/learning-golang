package main

import "fmt"

func deferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happened", err)
		}
	}()
	fmt.Println("Some useful work")
	panic("something bad happened")
	return
}

func main() {
	deferTest()
	return
}
