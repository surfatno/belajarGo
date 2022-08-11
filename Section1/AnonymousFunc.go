package Section1

import "fmt"

type Blacklist func(string) bool // alias

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func kambeng(name string) bool {
	return name == "admin"
}
func AnonymousFunc() {
	blacklist := func(name string) bool {
		return name == "admin" // yang bikin di block
	}
	registerUser("admin", blacklist)
	registerUser("Eko", blacklist)
	registerUser("admin", kambeng)
}
