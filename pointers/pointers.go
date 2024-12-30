package main

import "fmt"

func main() {
	age := 32 // regular variable
	agePointer := &age
	fmt.Println("Age:", *agePointer) // de-referencing
	editAgetoAdultYears(agePointer)
	fmt.Println(age)

}

func editAgetoAdultYears(ageP *int) {
	*ageP = *ageP - 18
}
