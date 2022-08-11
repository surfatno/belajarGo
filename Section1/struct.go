package Section1

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	// Married       bool
}

func Struct() {
	var Surfatno Customer
	Surfatno.Name = "Surfatno"
	Surfatno.Address = "ASDASD"
	//   Surfatno.Age = 28
	fmt.Println(Surfatno.Name)
}

func Struct2() {
	Surfatno := Customer{
		Name:    "Surfatno",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(Surfatno)
	//tidak di rekomendasikan karena wajib sama dengan urutan struct
	budi := Customer{"Budi", "APALAH", 35}
	fmt.Println(budi)
}

func (customer Customer) sayHello() {
	fmt.Println("Hello", customer.Name)
}

func Struct3() {
	rully := Customer{Name: "Rully"}
	rully.sayHello()
}

type Ayam struct {
	bagian string
}

func (ayam Ayam) daging() {
	fmt.Println(ayam.bagian)
}

func Struct4() {
	dada := Ayam{bagian: "dada"}
	dada.daging()
}
