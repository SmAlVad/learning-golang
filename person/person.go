package person

var (
	Public  = 1 // если начинается с заглавной, то экспортируется
	private = 1 // только в пакете
)

type Person struct {
	ID     int
	Name   string
	secret string
}

func (p Person) UpdateSecret(secret string) {
	p.secret = secret
}
