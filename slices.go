package main

import "fmt"

func main() {
	arraysWorks()
	fmt.Printf("\n")
	fmt.Printf("\n")
	slicesWorks()
}

func arraysWorks() {
	simpleArray()
	multipleArray()
}

func slicesWorks() {
	testSlices()
	modifySlice()
	testSliceCapacity()
	testSliceViaMake()
	memoryOptimization()
}

func simpleArray() {
	a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a {
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)
}

func multipleArray() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"},
	}
	printMultipleArray(a)
	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printMultipleArray(b)
}

func printMultipleArray(array [3][2]string) {
	for _, v1 := range array {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func testSlices() {
	a := [...]int{76, 77, 78, 79, 80}
	fmt.Println(getSliceFromArray(a, 1, 2))
}

func getSliceFromArray(arr [5]int, start int, end int) []int {
	return arr[start:end]
}

func modifySlice() {
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)
}

func testSliceCapacity() {
	fruitArray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitSlice := fruitArray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitSlice), cap(fruitSlice))
	fruitSlice = fruitSlice[:cap(fruitSlice)]
	fmt.Println("After re-slicing length is", len(fruitSlice), "and capacity is", cap(fruitSlice))
}

func testSliceViaMake() {
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))
}

func getCountries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries)
	return countriesCpy
}

func memoryOptimization() {
	countriesNeeded := getCountries()
	fmt.Println(countriesNeeded)
}
