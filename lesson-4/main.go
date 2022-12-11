package main

import "fmt"

func main() {
	structMethods()
}

// functions
func functions() {
	sum := func(num1, num2 int) int {
		return num1 + num2
	}

	multiple := func(num1, num2 int) int {
		return num1 * num2
	}

	sqr1 := squareOfResult(sum, "Sum of numbers")
	fmt.Println("Square of result is: ", sqr1)
	sqr2 := squareOfResult(multiple, "Multiple of numbers")
	fmt.Println("Square of result is: ", sqr2)

	price := totalPrice(100)
	fmt.Println(price(10))
	fmt.Println(price(20))
}

func squareOfResult(callback func(int, int) int, s string) int {
	result := callback(5, 5)
	fmt.Println(fmt.Sprintf("Result of %s is %d", s, result))
	return result * result
}

func totalPrice(initValue int) func(int) int {
	sum := initValue

	return func(x int) int {
		sum += x
		return sum
	}
}

// Struct methods
func structMethods() {
	point := Point{
		X: 4,
		Y: 8,
	}

	point.move(24, 28)
	fmt.Println("After move:", point)
}

type Point struct {
	X, Y int
}

func (p *Point) move(x, y int) {
	fmt.Println("Point: ", p)
	fmt.Println("new X: ", x)
	fmt.Println("new Y: ", y)
	p.X += x
	p.Y += y
}
