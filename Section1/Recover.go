package Section1

import "fmt"

func endApp1() {
	fmt.Println("Selesai Memanggil Function")
	message := recover()
	fmt.Println("Terjadi Error", message)
}
func runApp1(error bool) {
	defer endApp1()
	if error {
		panic("App Error")
	}

}
func Recover() {
	runApp1(false)
}
