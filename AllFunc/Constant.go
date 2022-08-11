package AllFunc

func Constant() (string, string) {
	const firstName = "Eko"
	const lastName = "Khannedy"

	// firstName = "a"
	return firstName, lastName
}

func MultiConstant() (string, string) {
	const (
		firstName = "Eko"
		lastName  = "Khannedy"
	)
	// firstName = "a"
	return firstName, lastName
}
