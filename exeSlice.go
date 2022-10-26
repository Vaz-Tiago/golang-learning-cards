package main

import "fmt"

func evenOdd (length int){
	var numberList []int
	for i := 1; i <= length; i++ {
		numberList = append(numberList, i)
	}

	for _, number := range numberList {
		whatIs := "Odd"
		if (number % 2 == 0) {
			whatIs = "Even"
		}

		fmt.Printf("%v[%v]\n", number, whatIs)
	}
}

