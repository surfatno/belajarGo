package main

import (
	"BelajarGoLang/AllFunc"
	"BelajarGoLang/Section1"

	"fmt"
)

func main() {
	// allfunc()
	// section1()
	section2()
}
func allfunc() {
	// GetxCaharacter()
	// CountChar()
	// Constant()
	// Variable()
	// Conversion()
}
func section1() {
	// Array()
	// Slice()
	// Section1.Map()
	// Section1.Variadic()
	// Section1.FunctionAsParameter()
	// Section1.AnonymousFunc()
	// Section1.Closure()
	// Section1.Defer()
	// Section1.Panic()
	// Section1.Recover()
	// Struct()
	// Interface()
	// pointer()

	// ------------- Package
	packages()
}
func section2() {
	//go mod init nama-module
}

// ---------------------------------------------- All Function ----------------------------------------------
func Variable() {
	AllFunc.Variable()
	fmt.Println("-----------------------------")
	AllFunc.MultipleVariable()
}
func Conversion() {
	AllFunc.Conversion()

}
func GetxCaharacter() {
	fmt.Println("Byte si E : ", AllFunc.GetString())
}

func CountChar() {
	fmt.Print("Hitung Character : ")
	fmt.Println(AllFunc.LengthCharacter("Surfatno"))
}
func Constant() {
	fmt.Print("Constant	: ")
	fmt.Println(AllFunc.Constant())
	fmt.Print("Multi Constant	: ")
	fmt.Println(AllFunc.MultiConstant())
}

// ---------------------------------------------- Section 1 ----------------------------------------------
func Array() {
	// Array
	fmt.Print("Array  : ")
	fmt.Println(Section1.Array())
	fmt.Print("Array1 : ")
	fmt.Println(Section1.Array1())
	fmt.Print("Array2 : ")
	fmt.Println(Section1.Array2())
}

func Slice() {
	fmt.Println("---------------- 1 ----------------")
	Section1.Slice()

	fmt.Println("---------------- 2 ----------------")
	Section1.SliceEdit()

	fmt.Println("---------------- 3 ----------------")
	Section1.SliceMonth()

	fmt.Println("---------------- 4 ----------------")
	Section1.MakeSlice()

	fmt.Println("---------------- 5 ----------------")
	Section1.PerbedaanArraySlice()

}
func Struct() {
	Section1.Struct()
	fmt.Println("-------------------------------")
	Section1.Struct2()
	fmt.Println("-------------------------------")
	Section1.Struct3()
	fmt.Println("-------------------------------")
	Section1.Struct4()
}

func Interface() {
	Section1.Interface()
	fmt.Println("-------------------------------")
	Section1.InterfaceKosong()
	fmt.Println("-------------------------------")
	Section1.InterfaceError()
	fmt.Println("-------------------------------")
	Section1.TypeAssertions()
}

func pointer() {
	Section1.PointerValue()
	fmt.Println("-----------------------------")
	Section1.PointerinFunction()
	fmt.Println("-----------------------------")
	Section1.PointerinMethod()
}

func packages() {

	Section1.Os()
	fmt.Println("-----------------------------")
	Section1.Flag()
	fmt.Println("-----------------------------")
	Section1.StringPackage()
	fmt.Println("-----------------------------")
	Section1.StrconvPackage()
	fmt.Println("-----------------------------")
	Section1.MathPackage()
	fmt.Println("-----------------------------")
	Section1.ContainerSlashListPackage()
	fmt.Println("-----------------------------")
	Section1.ContainerSlashRingPackage()
	fmt.Println("-----------------------------")
	Section1.SortPackage()
	fmt.Println("-----------------------------")
	Section1.TimePackage()
	fmt.Println("-----------------------------")
	Section1.ReflectPackage()
	fmt.Println("-----------------------------")
	Section1.RegexPackage()

}

// ---------------------------------------------- Section 2 ----------------------------------------------
