package main

import "fmt"

func weekDay(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid Number"
	}

}

func weekDay2(number int) string {
	switch {
	case number == 1:
		return "Sunday"
	case number == 2:
		return "Monday"
	case number == 3:
		return "Tuesday"
	case number == 4:
		return "Wednesday"
	case number == 5:
		return "Thursday"
	case number == 6:
		return "Friday"
	case number == 7:
		return "Saturday"
	default:
		return "Ivalid Number"

	}
}

func weekDay3(number int) string {
	var weekDay string

	switch {
	case number == 1:
		weekDay = "Sunday"
	case number == 2:
		weekDay = "Monday"
	case number == 3:
		weekDay = "Tuesday"
	case number == 4:
		weekDay = "Wednesday"
	case number == 5:
		weekDay = "Thursday"
	case number == 6:
		weekDay = "Friday"
	case number == 7:
		weekDay = "Saturday"
	default:
		weekDay = "Ivalid Number"

	}
	return weekDay
}

func main() {
	fmt.Println("Switch")
	day1 := weekDay(1)
	fmt.Println(day1)

	fmt.Println("------------")
	day2 := weekDay2(2)
	fmt.Println(day2)

	fmt.Println("------------")
	day3 := weekDay3(3)
	fmt.Println(day3)

}
