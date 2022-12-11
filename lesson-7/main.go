package main

import (
	"fmt"
	"io"
	"strings"
)

type AppError struct {
	Err         error
	Status      uint8
	Code        uint8
	Description string
}

func (e *AppError) Error() string {
	e.Err = fmt.Errorf("\n-----\nStatus: %d\nCode: %d\nDescription: %s\n-----", e.Status, e.Code, e.Description)
	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func main() {
	readerTest()
}

func errTest() {
	err := generateError()
	fmt.Println(err.Unwrap())
	fmt.Println(err.Error())
}

func generateError() *AppError {
	return &AppError{
		Err:         fmt.Errorf("unathorized"),
		Status:      2,
		Code:        2,
		Description: "Something went wrong",
	}
}

func readerTest() {
	reader := strings.NewReader("some string on testing")
	byteArr := make([]byte, 4)

	for {
		n, err := reader.Read(byteArr)
		fmt.Printf("arr n bytes: %q\n", byteArr[:n])
		fmt.Printf("n: %d\nerr: %v\narr: %v\n-----\n", n, err, byteArr)
		if err == io.EOF {
			fmt.Println(byteArr)
			break
		}
	}
}
