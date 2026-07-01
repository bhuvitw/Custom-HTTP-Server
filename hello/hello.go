package main

import (
	// "errors"
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
	// var printValue string = "Hello World"
	// printMe(printValue)

	// var numerator int = 11
	// var denominator int = 0
	// var result, remainder, err = intDivision(numerator, denominator)
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// }else if remainder == 0{
	// 	fmt.Printf("The result fo tintefer division %v", result)
	// }else {
	// 	fmt.Printf("The result fo integer divisin is %v with remainder %v", result, remainder)
	// }
	// switch {
	// case err != nil:
	// 	fmt.Printf(err.Error())
	// case remainder == 0:
	// 	fmt.Printf("The result fo tintefer division %v", result)
	// default:
	// 	fmt.Printf("The result fo integer divisin is %v with remainder %v", result, remainder)
	// }

	// switch remainder {
	// case 0:
	// 	fmt.Println("The divison was exact")
	// case 1, 2:
	// 	fmt.Println("The division was close")
	// default:
	// 	fmt.Println("The divison was not close")
	// }

	// var intArr [3]int32
	// intArr[1] = 123
	// fmt.Println(intArr[0])
	// fmt.Println(intArr[1:3])

	// fmt.Println((&intArr[0]))
	// fmt.Println((&intArr[1]))
	// fmt.Println((&intArr[2]))

	// var intArr [3]int32 = [3]int32{1, 2, 3}
	// intArr := [3]int32{1, 2, 3}

	// intArr := [...]int32{1, 2, 3}
	// fmt.Println(intArr)

	// var intSlice []int32 = []int32{4, 5, 6}
	// fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	// intSlice = append(intSlice, 7)
	// fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	// var intSlice2 []int32 = []int32{8, 9}
	// intSlice = append(intSlice, intSlice2...)
	// fmt.Println(intSlice)

	// var intSlice3 []int32 = make(int32[], 3, 8)

	// var myMap map[string]uint8 = make(map[string]uint8)
	// fmt.Println(myMap)

	// var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	// fmt.Println(myMap2)

	// fmt.Println(myMap2["Adam"])
	// var age, ok = myMap2["Jason"]
	// // delete(myMap2, "Adam")
	// if ok {
	// 	fmt.Printf("The age is %v", age)
	// } else {
	// 	fmt.Println("Invalid Name")
	// }

	// for name := range myMap2 {
	// 	fmt.Printf("Name: %v\n", name)
	// }

	// var i int = 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for {
	// 	if i >= 10 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	// var myString = "resume"
	// var indexed = myString[0]
	// fmt.Printf("%v, %T", indexed, indexed)
	fmt.Println("start")
	defer fmt.Println("delayed by defer")
	fmt.Println("end")
}

// func printMe(printValue string) {
// 	fmt.Println(printValue)
// }

// func intDivision(numerator int, denominator int) (int, int, error) {
// 	var err error
// 	if denominator == 0 {
// 		err = errors.New("Cannot Divide by Zero")
// 		return 0, 0, err
// 	}
// 	var result int = numerator / denominator
// 	var remainder int = numerator % denominator
// 	return result, remainder, err
// }
