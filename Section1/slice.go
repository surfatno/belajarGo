package Section1

import "fmt"

func getMonth() [12]string {
	var months = [...]string{
		"Januari",   // 0
		"Februari",  // 1
		"Maret",     // 2
		"April",     // 3
		"Mei",       // 4
		"Juni",      // 5
		"Juli",      // 6
		"Agustus",   // 7
		"September", // 8
		"Oktober",   // 9
		"November",  // 10
		"Desember",  // 11
	}
	return months
}

func Slice() {
	var months = getMonth()
	fmt.Println(months)
	var slice1 = months[4:7]
	fmt.Println("Slice	: ", slice1)
	fmt.Println("Len	: ", len(slice1))
	fmt.Println("Cap	: ", cap(slice1))
}

func SliceEdit() {
	var months = getMonth()
	months[5] = "(edit Juni)"

	var slice1 = months[4:7]
	slice1[0] = "(edit Mei)"
	fmt.Println("Slice 	: ", slice1)
	fmt.Println("Months	: ", months)

}

func SliceMonth() {
	var months = getMonth()
	var slice2 = months[2:4]
	fmt.Println("Slice		: ", slice2)

	var slice3 = append(slice2, "Eko")
	fmt.Println("Append		: ", slice3)

	slice3[1] = "(bukan desember)"
	fmt.Println("Edit Append	: ", slice3)
	fmt.Println(months)

}

func MakeSlice() {

	//buat slice kosong. baru ditambahkan isi ny
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Eko"
	newSlice[1] = "Kurniawan"

	fmt.Println("Slice		: ", newSlice)
	fmt.Println("Len		: ", len(newSlice))
	fmt.Println("Cap		: ", cap(newSlice))
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("Copy Slice	: ", copySlice)

}

func PerbedaanArraySlice() {
	iniArray := [5]int{1, 2, 3, 4, 5} //[...]int{1,2,3,4,5}
	fmt.Println("Array	: ", iniArray)
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice 	: ", iniSlice)

}

/*
	slice adalah potongan dari data array
	ukuran slice bisa berubah, ukuran array tidak bisa berubah
	slice dan array selalu terkoneksi, dimana slice adalah data yang mengakses sebagian atau seluruh data di array

	tipe data slice : pointer , length dan capacity
	• pointer 	: penunjuk data pertama di array pada slice
	• length 	: panjang dari slice
	• capacity 	: kapasitas dari slice, dimana length tidak boleh lebih dari capacity

	membuat slice dari array
	• array [low:high] membuat slice dari array dimulai index low sampai index sebelum high
	• array [low:]     membuat slice dari array dimulai index low sampai index akhir di array
	• array [:high]    membuat slice dari array dimulai index 0 sampai index sebelum high
	• array [:] 	   membuat slice dari array dimulai index 0 sampai index akhir di array

	contoh:
	array =
	0. Ayam
	1. Bebek
	2. Cacing
	3. Domba
	4. Elang
	5. Flammingo
	6. Gajah
	7. Harimau

	slice = array[2:5]
	pointer = 2 (slice pertama adalah cacing)
	length = 3. dari array[2:5], memiliki 3 data, yang masuk : Cacing, Domba. Elang
	capacity = 6. capasitas sampai terakhir. Cacing, Domba, Elang,Flammingo, Gajah, Harimau
	slice ini bisa lebih besar. tetapi kapasitas hanya 6

	func Slice
	• len(slice)						Untuk mendapatkan panjang
	• cap(slice)						Untuk mendapatkan kapasitas
	• append(slice,data)				Membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
	• make([]TypeData,length,capcacity)	Membuat slice baru
	• copy(destination,source)			Menyalin slice dari source ke destination
*/
func Angka() [10]int {
	angka := [...]int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	return angka
}

func Tester() {

	//slice
	var angka = Angka()
	var angka1 []int = angka[4:7]
	fmt.Println(angka)
	fmt.Println(angka1)
	// capasitas sampe belakang ada brp
	fmt.Println(cap(angka1))
	// panjang slice yang di ambil
	fmt.Println(len(angka1))
	// edit slice
	angka[0] = 10
	fmt.Println(angka)
	angka1[0] = 44
	fmt.Println(angka1)
	fmt.Println(angka)
	// tambahin angka di belakang. yang di belakang ny kena timpa
	var angka2 = append(angka1, 55)
	fmt.Println(angka2)
	fmt.Println(angka)
	// bikin dari benar" awal.
	var ayam = make([]string, 1, 3)
	ayam[0] = "ayam"
	ayam = append(ayam, "bebek")

	fmt.Println(ayam)
	fmt.Println(len(ayam))
	fmt.Println(cap(ayam))
}
