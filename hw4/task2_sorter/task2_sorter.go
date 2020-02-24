package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name  string
	Phone int
}

//создание псевдонима
type addressbook []person

func (a addressbook) Len() int {
	return len(a)
}

func (a addressbook) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func (a addressbook) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	var persons addressbook
	persons = append(persons, person{
		Name:  "Миша",
		Phone: 78293467382,
	})
	persons = append(persons, person{
		Name:  "Никита",
		Phone: 89167253764,
	})
	persons = append(persons, person{
		Name:  "Маша",
		Phone: 89635437382,
	})
	persons = append(persons, person{
		Name:  "Марина",
		Phone: 78293467381,
	})
	persons = append(persons, person{
		Name:  "Генадий",
		Phone: 78293467384,
	})

	fmt.Println(persons)
	sort.Sort(persons)
	fmt.Println(persons)
}
