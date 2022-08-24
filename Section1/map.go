package Section1

import "fmt"

func Map() {
	person := map[string]string{
		// map [name] Eko
		"name":    "Eko",
		"address": "Subang",
	}

	person["title"] = "Programmer" // map[key] = value

	fmt.Println(person)
	fmt.Println(person["name"]) // map [key]
	fmt.Println(person["address"])

	book := make(map[string]string) //make(map[TypeKey]TypeValue)
	book["title"] = "Buku Go Lang"
	book["author"] = "Surfatno"
	book["wrong"] = "ups"
	delete(book, "wrong") //delete (map,key)
	fmt.Println(book)
}
