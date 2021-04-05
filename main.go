package main

import (
	"fmt"
	"strings"
	"time"
)

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

type Age interface {
	ageCalculator(string, error)
}

type Birth struct {
	dateOfBirth string
}

func (b Birth) ageCalculator() (string, error) {
	//reversing input
	reverseInput := strings.Split(b.dateOfBirth, "-")
	for i, j := 0, len(reverseInput)-1; i < j; i, j = i+1, j-1 {
		reverseInput[i], reverseInput[j] = reverseInput[j], reverseInput[i]
	}
	reversedInput := strings.Join(reverseInput, "-")

	//converting input to valid format
	date, err := time.Parse("2006-01-02", string(reversedInput))
	if err != nil {
		return "", fmt.Errorf("Invalid date format")
	}

	// Geting age from year of birth up to now
	age := time.Now().Year() - date.Year()

	// Getting left days without leap
	days := 0
	if time.Now().YearDay()%365 == 0 {
		age += 1
	} else if time.Now().YearDay()%365 != 0 {
		days = time.Now().YearDay()
	}

	return "Age: " + fmt.Sprint(age) + " years " + fmt.Sprint(days) + " days", nil
}
