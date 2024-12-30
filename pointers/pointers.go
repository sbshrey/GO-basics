package main

import "fmt"

func main() {
	age := 32 // regular variable
	agePointer := &age
	fmt.Println("Age:", *agePointer) // de-referencing
	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)

}

func getAdultYears(ageP *int) int {
	return *ageP - 18
}
