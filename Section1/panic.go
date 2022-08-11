package Section1

import "fmt"

func endApp() {
	fmt.Println("Selesai Memanggil Function")
}
func runApp(error bool) {
	defer endApp()
	if error {
		panic("App Error")
	}
	fmt.Println("App Berjalan")
}
func Panic() {
	runApp(true)
}
