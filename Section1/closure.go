package Section1

import "fmt"

func Closure() {
	counter := 0
	increment := func() {
		fmt.Println("Increment")
		counter++
	}
	increment()
	fmt.Println(counter)
}
