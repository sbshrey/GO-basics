package main

import "fmt"

func main() {
	age := 32 // regular variable
	agePointer := &age
	fmt.Println("Age:", *agePointer) // de-referencing
	getAdultYears(agePointer)
	fmt.Println(age)

}

func getAdultYears(ageP *int) {
	*ageP = *ageP - 18
}
