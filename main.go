package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// Person

type person struct {
	name string
	age  int
}

func personConstructor(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

// Main

func main() {
	//bobby := personConstructor("bobby", 15)
	//john := personConstructor("john", 19)
	//peter := personConstructor("peter", 27)
	//joan := personConstructor("joan", 45)
	//alibaba := personConstructor("alibaba", 60)
	//checkPerson(bobby)
	//checkPerson(john)
	//checkPerson(peter)
	//checkPerson(joan)
	//checkPerson(alibaba)
	//for i := range [7]int{} {
	//printPrediction(10)
	//}

	// min finding
	//numbers := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -255, 254, -444, 666}
	//printMessage(fmt.Sprintf(
	//	"The min number of \n%d\nis: %d",
	//	numbers,
	//	findMin(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -255, 254, -444, 666),
	//))

	// anonymous func
	inc := increment()
	dec := decrement()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())
}

func logError(err error) {
	printMessage("---------------------------------------------")
	log.Println(err)
	printMessage("---------------------------------------------")
}
func printMessage(message string) {
	fmt.Println(message)
}
func checkPerson(person person) {
	var canEnterMessage, err = canEnter(person)
	printMessage(sayHello(person))
	printMessage("Checking your age...")
	if err != nil {
		logError(err)
		return
	}
	printMessage(canEnterMessage)
	printMessage("")
}
func sayHello(person person) string {
	return fmt.Sprintf("Hello, %s!\nAre your age is %d?",
		toTitle(person.name),
		person.age,
	)
}
func toTitle(str string) string {
	return strings.Title(strings.TrimSpace(str))
}
func canEnter(person person) (string, error) {
	var message string
	var err error
	age := person.age

	if age < 21 {
		message = ""
		err = errors.New("you are too young")
	} else if age >= 21 && age < 45 {
		message = "You can enter!"
		err = nil
	} else if age >= 45 && age < 60 {
		message = "You can enter, but be careful."
		err = nil
	} else {
		message = ""
		err = errors.New("you are too old")
	}

	return message, err
}

// Weekday

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (weekday Weekday) String() string {
	return []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}[weekday]
}
func getWeekdays() [7]Weekday {
	return [7]Weekday{
		Sunday,
		Monday,
		Tuesday,
		Wednesday,
		Thursday,
		Friday,
		Saturday,
	}
}
func getPredictions() [7]string {
	weekdays := getWeekdays()
	predictions := [7]string{}

	for i, v := range weekdays {
		predictions[i] = fmt.Sprintf("Good %s to you!", v)
	}

	return predictions
}
func printWeekday(index int, day Weekday) {
	fmt.Println(fmt.Sprintf("%d: %s", index, day))
}
func printWeekdays() {
	weekdays := getWeekdays()

	for i, v := range weekdays {
		printWeekday(i, v)
	}
}
func getPrediction(weekdayIndex int) (string, error) {
	if weekdayIndex > 7 {
		return "", errors.New("getPrediction call out of range")
	}
	return getPredictions()[weekdayIndex], nil
}
func printPrediction(weekdayIndex int) {
	var prediction, err = getPrediction(weekdayIndex)
	if err != nil {
		logError(err)
		printMessage("You can choose weekday from: ")
		printWeekdays()
		printMessage("---------------------------------------------")
		return
	}
	fmt.Println(prediction)
}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}

	return min
}

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func decrement() func() int {
	count := 10
	return func() int {
		count--
		return count
	}
}
