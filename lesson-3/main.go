package main

import "fmt"

func main() {
	slices()
}

func pointers() {
	a := "hello world"

	p := &a
	fmt.Println("p_cell_link:", p)
	fmt.Println("p_cell_value:", *p)

	*p = fmt.Sprintf("new value of cell %p", p)
	fmt.Println("p_cell_value:", *p)
}

// structs

type Point struct {
	X int
	Y int
}

func (point *Point) printData() {
	data := fmt.Sprintf("Point Coords:\n\tX: %d\n\tY: %d\nMemory cell: %p\n", point.X, point.Y, point)
	fmt.Println(data)
}

func structs() {
	point := Point{
		X: 24,
		Y: 32,
	}

	point.printData()
	fmt.Println(&point)
}

type animal struct {
	name       string
	animalType string
}

func (receiver *animal) voice() {
	var voice string

	switch receiver.animalType {
	case "dog":
		voice = "bark"
	case "cat":
		voice = "meow"
	case "cow":
		voice = "moooo"
	}

	fmt.Println(fmt.Sprintf("%s is %sing!", receiver.name, voice))
}

func slices() {
	animalsArr := [4]animal{
		{
			name:       "Balto",
			animalType: "dog",
		},
		{
			name:       "Varvara",
			animalType: "cat",
		},
		{
			name:       "Burenka",
			animalType: "cow",
		},
		{
			name:       "Jessica",
			animalType: "dog",
		},
	}

	for _, animal := range animalsArr {
		animal.voice()
	}
}
