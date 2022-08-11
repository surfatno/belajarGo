package Section1

import "fmt"

//variadic function bisa ditandai dengan ...variable (bisa di taruh banyak data)
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func nameAll(nama ...string) string {
	total := ""
	for _, merge := range nama {
		total += merge + " "
	}
	return total
}
func Variadic() {

	// variable
	total := sumAll(10, 10, 10)
	nama := nameAll("surfatno", "weng")
	fmt.Println(total)
	fmt.Println(nama)

	//jika sudah ada slice mw di jumlahkan
	slice := []int{20, 20}
	total = sumAll(slice...)
	fmt.Println(total)
}
