package main

import "fmt"

func main() {
	age := 32
	fmt.Println("Age:", age)
	getNumberOfYearsAsAdult(&age)
	fmt.Println(age)
}

func getNumberOfYearsAsAdult(age *int) {
	*age = *age - 18
}
