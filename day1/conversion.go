package main

import "fmt"

func main() {
	var numberOfLife int16 = 1
	var toNumberType int32 = int32(numberOfLife)
	var name = "Aditya"
	var getFirstIndex = name[0]
	var convertToString = string(getFirstIndex)
	fmt.Println(toNumberType)
	fmt.Println(convertToString)
}
