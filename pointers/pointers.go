package main

import "fmt"

func main() {
	age := 32                // regular variable
	fmt.Println("Age:", age) // de-referencing
	editAgetoAdultYears(&age)
	fmt.Println(age)

}

func editAgetoAdultYears(ageP *int) {
	*ageP = *ageP - 18
}
