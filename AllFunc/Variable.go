package AllFunc

import "fmt"

func Variable() {
	var name string

	name = "Eko Kurniawan"
	fmt.Println(name)

	name = "Eko Khannedy"
	fmt.Println(name)

	var friendName = "Budi"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

}

func MultipleVariable() {
	var (
		firstName = "Eko"
		lastName  = "Khannedy"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
