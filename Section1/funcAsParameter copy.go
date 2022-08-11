package Section1

import "fmt"

type Filter func(string) string // alias

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println(nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

/*
func testing(name string) {
	filter := test(name)
	fmt.Println(filter)
}
func test(name string) string {
	if name == "anjing" {
		return "LU YANG ANJING!"
	} else {
		return name
	}
}

func ayam(name string, bebek func(string) string) {
	filter := bebek(name)
	fmt.Println(filter)
}
func cacing(name string) string {
	if name == "BABI" {
		return "..."
	} else {
		return name
	}
}
*/
func FunctionAsParameter() {
	sayHelloWithFilter("Eko", spamFilter)
	//testing("anjing")
	//ayam("BABI", cacing)
}
