package main

import "fmt"

func main() {
	var emptyPointer *int
	number := 5
	emptyPointer = &number

	*emptyPointer = 10

	message := "pointers"
	fmt.Println(message)
	changeMessage(&message)
	fmt.Println(message)
	fmt.Println(*emptyPointer)
	fmt.Println(&number)
}

func changeMessage(message *string) {
	*message += " (from printMessage() func)"
}
