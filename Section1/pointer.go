package Section1

import "fmt"

type Address struct {
	City, Province, Country string
}

func PointerValue() {
	fmt.Println("----- 1 -----")
	address1 := Address{"Medan", "SUMUT", "Indonesia"}
	//var address1 Address = Address{"Medan", "SUMUT", "Indonesia"}
	fmt.Println(address1)

	fmt.Println("----- 2 -----")
	address2 := address1
	address2.City = "Palembang"
	fmt.Println(address1)
	fmt.Println(address2)

	fmt.Println("----- 3 -----")
	address3 := &address1 //pointer
	// var address3 *Address= &address1
	address3.City = "Aceh"
	fmt.Println(address1)
	fmt.Println(address3)

	fmt.Println("----- 4 -----")
	address4 := Address{"Jakarta", "DKI JAKARTA", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address4)

	fmt.Println("----- 5 -----")
	address5 := &address1
	*address5 = Address{"Bali", "Bali", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address5)

	fmt.Println("----- 6 -----")
	address6 := new(Address)
	address6.City = "Medan"
	fmt.Println(address6)
}
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}
func PointerinFunction() {

	alamat := Address{
		City:     "Subang",
		Province: "Java Barat",
		Country:  "",
	}
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}
func PointerinMethod() {
	eko := Man{"eko"}
	eko.Married()
	fmt.Println(eko.Name)
}
