package main

import (
	"fmt"
	"learning-golang/person"
)

func main() {
	p := person.NewPerson(1, "rvasily", "secret")

	// p.secret undefined (can't refer to unexported field of method secret)
	//fmt.Printf("main.PrintPerson: %+v\n", p.secret)

	secret := person.GetSecret(p)
	fmt.Println("GetSecret", secret)
}
