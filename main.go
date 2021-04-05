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
	monthToDays(string, error)
}

type Birth struct {
	dateOfBirth string
	days        Age
}

func (b Birth) ageCalculator() (string, error) {
	//reversing input
	reverseInput := strings.Split(b.dateOfBirth, "-")
	for i, j := 0, len(reverseInput)-1; i < j; i, j = i+1, j-1 {
		reverseInput[i], reverseInput[j] = reverseInput[j], reverseInput[i]
	}
	reversedInput := strings.Join(reverseInput, "-")

	//converting input to valid format
	date, err := time.Parse("2006-01-02 15 04", string(reversedInput))
	if err != nil {
		return "", err
	}

	//geting age from year of birth up to now
	age := time.Now().Year() - date.Year() - 1
	// getting left month
	// month := 0
	if date.Day() == time.Now().Day() && date.Month() == time.Now().Month() {
		age += 1
	} else {

	}

	return "", nil
}

func monthToDays(month string) int {
	switch month {
	case "1", "01", "Jan", "January":
		return 31
	case "2", "02", "Feb", "February":
		return 28
	case "3", "03", "Mar", "March":
		return 31
	case "4", "04", "Apr", "April":
		return 30
	case "5", "05", "May":
		return 31
	case "6", "06", "Jun", "June":
		return 30
	case "7", "07", "Jul", "July":
		return 31
	case "8", "08", "Aug", "August":
		return 31
	case "9", "09", "Sep", "September":
		return 30
	case "10", "Oct", "October":
		return 31
	case "11", "Nov", "November":
		return 30
	case "12", "Dec", "December":
		return 31
	default:
		return 0
	}
}
