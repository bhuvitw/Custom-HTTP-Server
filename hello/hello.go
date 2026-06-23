package main

import (
	"errors"
	"fmt"
)

func main() {
	// var intNum int = 32767+1
	// fmt.Println(intNum)

	// var floatNum float64 = 10.1

	// var myString string = "Hello" + " " + "World"
	// fmt.Println(myString)

	// fmt.Println(utf8.RuneCountInString("Y"))

	// var myRune rune = 'a'
	// fmt.Println(myRune)

	// var myBoolean bool = false
	// fmt.Println(myBoolean)

	// var intNum int
	// fmt.Println(intNum) // 0

	// myVar := "text"
	// fmt.printLn(myVar)

	// var1, var2 := 1, 3
	// fmt.Println(var1, var2)

	// const myConst string = "const value"
	// fmt.Println(myConst)

	// const pi float = 3.1451
	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 0
	var result, remainder, err = intDivision(numerator, denominator)
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// }else if remainder == 0{
	// 	fmt.Printf("The result fo tintefer division %v", result)
	// }else {
	// 	fmt.Printf("The result fo integer divisin is %v with remainder %v", result, remainder)
	// }
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result fo tintefer division %v", result)
	default:
		fmt.Printf("The result fo integer divisin is %v with remainder %v", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Println("The divison was exact")
	case 1, 2:
		fmt.Println("The division was close")
	default:
		fmt.Println("The divison was not close")
	}

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
