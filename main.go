package main

import (
	"fmt"
	"time"
)

type Student struct {
	firstName string
	lastName string
	Email string
}

// Value receiver | Pointer receiver
func (s *Student) email() string {
	s.Email = "alice@gmail.com"
	return s.Email
}

func main() {
	// bien, kieu du lieu, func, cau truc du lieu, oop, go routine, channel

	// go func1
	c := make(chan int)
	go func() {
		c <- 100
	}()

	go func() {
		fmt.Println(<-c)
	}()

	time.Sleep(time.Second)
}

func g1() {
	fmt.Println("1")
}

func demostruct() {
	st := Student{
		firstName: "Ryan",
		lastName:  "Nguyen",
		Email:     "ryan@gmail.com",
	}
	st.firstName = "Alice"
	fmt.Println(st)

	e := st.email()
	fmt.Println(e)

	fmt.Println(st.Email)
}

func demoStructure() {
	// slice, map, array
	s := make([]string, 0)
	s = append(s, "a")
	s = append(s, "b")
	s = append(s, "c")
	s = append(s, "d")

	fmt.Println(s)

	m := make(map[string] int)
	// k-v
	m["age"] = 90
	m["age1"] = 900

	fmt.Println(m["age1"])

	// array
	arr := [1] string{}
	arr[0] = "value1"
	fmt.Println(arr)
}

func demoFunc() {
	sayHello()

	total := sum(1, 2)
	fmt.Println(total)
}

func sayHello() {
	fmt.Println("inside sayHello function")
}

func sum(a int, b int) int {
	return a + b
}

func demoVairable() {
	var email = "ryan@code4func.com"

	email = "newEmail@Gmail.com"

	fmt.Println(email)

	fmt.Printf("%V", email)
	fmt.Println()

	fullName := "Ryan Nguyen"
	fmt.Println(fullName)
	fmt.Printf("%T", fullName)
	fmt.Println()

	var number = 10
	var float = 10.123
	fmt.Printf("%T", number)
	fmt.Println()
	fmt.Printf("%T", float)
	fmt.Println()
}




