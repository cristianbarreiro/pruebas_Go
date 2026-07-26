package main

import "fmt"

type days int

const (
	Sun days = iota
	Mon
	Tue
	Wed
	Thu
	Fri
	Sat
)

func nextDay(today days) days {
	return (days((int(today) + 1) % 7))
}

func name(d days) (nm string) {
	switch d {
	case Sun:
		nm = "Sun"
	case Mon:
		nm = "Mon"
	case Tue:
		nm = "Tue"
	case Wed:
		nm = "Wed"
	case Thu:
		nm = "Thu"
	case Fri:
		nm = "Fri"
	case Sat:
		nm = "Sat"
	}
	return nm
}

func main() {
	fmt.Println("Weekday Program")
	for i := Sun; i <= Sat; i++ {
		fmt.Println(days(i), "is ", name(days(i)))
		fmt.Println(name(nextDay(days(i))))
	}
}
