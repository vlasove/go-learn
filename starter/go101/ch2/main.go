package main

import (
	"log"
)

type Person struct {
	Name string
}

func NewPerson(name string) *Person {
	return &Person{
		Name: name,
	}
}

func (p *Person) Introduce() {
	log.Println("Hello, I'm person with name:", p.Name)
}

type Saiyan struct {
	*Person
	Power int
}

func NewSaiyan(p *Person, power int) *Saiyan {
	return &Saiyan{
		Person: p,
		Power:  power,
	}
}

func (s *Saiyan) Introduce() {
	log.Println("Hello, I'm sayian with SUPER name:", s.Name)
}

func (s *Saiyan) Super() {
	s.Power += 1234
}

func main() {
	goku := NewSaiyan(
		NewPerson("goku"),
		9000,
	)

	goku.Introduce()
	goku.Person.Introduce()
	goku.Super()
	log.Println("Goku's power:", goku.Power)

}

func Super(s *Saiyan) {
	s.Power += 100000
}
