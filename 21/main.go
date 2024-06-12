package main

import "fmt"

// Интерфейс классов адаптера
type Target interface {
	Operation()
}

// Адаптируемый класс
type Adaptee struct {
}

// Метод адаптируемого класса, который нужно вызвать где-то
func (adaptee *Adaptee) AdaptedOperation() {
	fmt.Println("I am AdaptedOperation()")
}

// Класс конкретного адаптера
type ConcreteAdapter struct {
	*Adaptee
}

// Реализация метода интерфейса, реализующего вызов адаптируемого класса
func (adapter *ConcreteAdapter) Operation() {
	adapter.AdaptedOperation()
}

// Конструктор адаптера
func NewAdapter(adaptee *Adaptee) Target {
	return &ConcreteAdapter{adaptee}
}

func main() {
	adapter := NewAdapter(&Adaptee{})
	adapter.Operation()
}
