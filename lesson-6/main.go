package main

import "fmt"

var _ Authentication = &user{}
var _ Authentication = &admin{}

type Authentication interface {
	Login()
	Logout()
	ChangePassword(newPassword string)
}

type user struct {
	Age          uint8
	Name         string
	Email        string
	Password     string
	IsAuthorized bool
}

func (u *user) Login() {
	u.IsAuthorized = true
}
func (u *user) Logout() {
	u.IsAuthorized = false
}
func (u *user) ChangePassword(newPassword string) {
	u.Password = newPassword
	u.Logout()
}

func NewUser(age uint8, name string, email string, password string) Authentication {
	newUser := user{
		Age:          age,
		Name:         name,
		Email:        email,
		Password:     password,
		IsAuthorized: false,
	}

	return &newUser
}

type admin struct {
	Email        string
	Password     string
	IsAuthorized bool
}

func (a admin) Login() {
	a.IsAuthorized = true
}
func (a admin) Logout() {
	a.IsAuthorized = false
}
func (a admin) ChangePassword(newPassword string) {
	a.Password = newPassword
	a.Logout()
}

func NewAdmin(email string, password string) Authentication {
	admin := admin{
		Email:        email,
		Password:     password,
		IsAuthorized: false,
	}

	return admin
}

func main() {
	casts()
}

func fromCounstructor() {
	amarkov := NewAdmin("admin", "admin")
	vasiliy := NewUser(21, "Vasiliy", "vasiliy@gmail.com", "1214")

	amarkov.Login()
	vasiliy.Login()

	fmt.Println(amarkov)
	fmt.Println(vasiliy)

	amarkov.ChangePassword("admin1")
	vasiliy.ChangePassword("1215")

	fmt.Println(amarkov)
	fmt.Println(vasiliy)
}

func fromInterface() {
	var auth Authentication
	u := user{
		Age:          28,
		Name:         "Arseniy",
		Email:        "zamberg42@gmail.com",
		Password:     "1213",
		IsAuthorized: false,
	}

	a := admin{
		Email:        "admin",
		Password:     "admin",
		IsAuthorized: false,
	}
	fmt.Println(a)

	auth = &u

	fmt.Printf("(%v, %T)\n", auth, auth)
}

func anyInterface() {
	// equals any
	var a interface{}
	fmt.Printf("(%v, %T)\n", a, a)

	a = "jelly"
	fmt.Printf("(%v, %T)\n", a, a)

}

// Type Assertion
func casts() {
	var a interface{} = "my dream string"

	s := a.(string)

	fmt.Println(s)

	v, ok := a.(int)
	fmt.Println("string to int cast:\n\tv: ", v, "\n\tok: ", ok)
	if ok {
		// some logic
	}

	var b interface{} = "45"

	switch b.(type) {
	case int:
		fmt.Println("b is int")
	case float32:
		fmt.Println("b is float32")
	case float64:
		fmt.Println("b is float64")
	case uint:
		fmt.Println("b is uint")
	case string:
		fmt.Println("b is string")
	default:
		fmt.Println("unknown type")
	}

}
