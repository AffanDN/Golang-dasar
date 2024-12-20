package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("jow", blacklist)
	
	// Anonymous Function
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
