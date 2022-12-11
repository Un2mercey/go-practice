package main

import (
	"fmt"
)

func main() {
	mapTesting()
}

func rangeTesting() {
	arr := []string{"a", "b", "c"}

	for i, el := range arr {
		fmt.Println("i:\t", i)
		fmt.Println("el:\t", el)
		fmt.Println()
	}
}

type Point struct {
	X, Y int
}

func mapTesting() {
	pointsMap := map[string]Point{}
	otherPointsMap := make(map[string]Point)
	var anotherPointsMap map[string]Point

	fmt.Println("pointsMap:\t\t", pointsMap)
	fmt.Println("otherPointsMap:\t\t", otherPointsMap)
	fmt.Println("anotherPointsMap:\t", anotherPointsMap)

	pointsMap["firstPoint"] = Point{1, 1}
	pointsMap["secondPoint"] = Point{2, 4}
	pointsMap["thirdPoint"] = Point{4, 16}
	pointsMap["fourthPoint"] = Point{16, 256}
	fmt.Println("pointsMap:\t\t", pointsMap)

	value, ok := pointsMap["firstPoint"]
	if ok {
		fmt.Println(value)
	}

	for key, value := range pointsMap {
		fmt.Println(key, value)
	}
}
