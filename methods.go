package main

import "fmt"

type Person1 struct {
	Id   int
	Name string
}

// не изменит оригинальной структуры для которой вызван
// передача копии типа
func (p Person1) UpdateName(name string) {
	p.Name = name
}

// изменит оригинальную структуру
// передача адреса на тип
func (p *Person1) SetName(name string) {
	p.Name = name
}

type Account1 struct {
	Id   int
	Name string
	Person1
}

type MySlice []int

func (sl *MySlice) Add(val int) {
	*sl = append(*sl, val)
}

func (sl *MySlice) Count() int {
	return len(*sl)
}

func main() {
	pers := Person1{1, "Vasily"}
	//(&pers).SetName("Vasily Romanov")
	pers.SetName("Vasily Romanov")

	fmt.Printf("updated Person1: %#v\n", pers)

	var acc Account1 = Account1{
		Id:   1,
		Name: "rvadily",
		Person1: Person1{
			Id:   2,
			Name: "Vasily Romanov",
		},
	}
	acc.SetName("romanov.vasily")

	sl := MySlice([]int{1, 2})
	sl.Add(3)
	fmt.Println(sl.Count())
}
