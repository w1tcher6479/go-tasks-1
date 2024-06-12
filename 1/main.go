package main

import "fmt"

// person - структура, описывающая человека с именем и возрастом.
type person struct {
	name string
	age  int
}

// Метод describe структуры person, возвращающий отформатированную строку, описывающую человека.
func (p person) describe() string {
	return fmt.Sprintf("My name is %s and I am %d years old.", p.name, p.age)
}

// container - структура, описывающая person и строку str.
type container struct {
	person
	str string
}

func main() {
	c := container{
		person{"Anton", 22},
		"example",
	}
	fmt.Println(c.describe())
}
