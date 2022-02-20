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
	bobby := personConstructor("bobby", 15)
	john := personConstructor("john", 19)
	peter := personConstructor("peter", 27)
	joan := personConstructor("joan", 45)
	alibaba := personConstructor("alibaba", 60)
	checkPerson(bobby)
	checkPerson(john)
	checkPerson(peter)
	checkPerson(joan)
	checkPerson(alibaba)

	//for i := range [7]int{} {
	printPrediction(10)
	//}
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
