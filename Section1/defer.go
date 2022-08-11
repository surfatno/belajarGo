package Section1

import "fmt"

func logging() {
	fmt.Println("Selesai Memanggil Function")
}
func runApplication(value int) {
	defer logging()
	fmt.Println("WOI")
	result := 10 / value
	fmt.Println(result)
}
func Defer() {
	runApplication(0)
}
