package main

import "fmt"

func main() {
	var input Birth
	fmt.Print("Input date of birth with format dd-mm-yyyy: ")
	fmt.Scan(&input.dateOfBirth)

	age, err := input.ageCalculator()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(age)
}

type Birth struct {
	dateOfBirth string
}

func (b Birth) ageCalculator() (string, error) {

	return "", nil
}
