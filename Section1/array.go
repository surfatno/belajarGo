package Section1

func Array() (string, string, string) {

	// jumlah maksimum array = 3 sudah di tetapkan. tidak bisa lebih dari 3. jika lebih dari 3 error
	var names [3]string
	// input data to array
	names[0] = "Eko"
	names[1] = "Kurniawan"
	names[2] = "Khannedy"

	return names[0], names[1], names[2]
}

func Array1() [3]int {
	// deklarasi array secara langsung
	// array memiliki tipe 1 data : length
	var values = [3]int{
		90,
		80,
		95,
	}
	return values
}

func Array2() (int, int, int) { // [3]int
	// deklarasi array secara langsung
	var values = [3]int{
		90,
		80,
		95,
	}
	return values[0], values[1], values[2]
}
