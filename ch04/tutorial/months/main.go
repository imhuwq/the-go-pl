package main

import "fmt"

type Month int

const (
	January Month = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

var Months = [...]string{
	January:   "January",
	February:  "February",
	March:     "March",
	April:     "April",
	May:       "May",
	June:      "June",
	July:      "July",
	August:    "August",
	September: "September",
	October:   "October",
	November:  "November",
	December:  "December"}

func main() {
	months := Months[:]
	q2 := Months[4:7]
	summer := Months[6:9]
	longSummer := summer[:5]
	fmt.Println(months)
	fmt.Println(q2)
	fmt.Println(summer)
	fmt.Println(longSummer)
}
