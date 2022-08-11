package Section1

import (
	"errors"
	"fmt"
)

//kontrak bernama hasname
type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

//struct
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func Interface() {
	person := Person{Name: "Eko"}
	SayHello(person)

}

// Interface Kosong
func Ups(i interface{}) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}
func InterfaceKosong() {
	var data interface{} = Ups(true)
	fmt.Println(data)
}

//error Interface

func Pembagian(nilai int, pembagian int) (int, error) {
	if pembagian == 0 {
		return 0, errors.New("Pembagian dengan 0")
	} else {
		return nilai / pembagian, nil
	}
}

func InterfaceError() {
	hasil, err := Pembagian(1, 0)
	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println(err.Error())
	}
}
func random() interface{} {
	return "oke"
}

func TypeAssertions() {
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("STRING ", value)
	case int:
		fmt.Println("INT", value)
	default:
		fmt.Println("UNKNOWN")
	}
}
