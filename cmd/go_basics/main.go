package main

import (
	"fmt"
	"strconv"
)

var choice int = 1

const breaker string = "----------------"

var options = []Option{
	{
		Title: "Variables", // 1
		Path:  "variables.go",
	},
	{
		Title: "Exit",
		Path:  "",
	},
}

type (
	Option struct {
		Title string
		Path  string
	}
)

func main() {
	exit := len(options)
	for choice != exit {
		printOptions()
		fmt.Println(breaker)
		fmt.Println("What would you like to learn about today?")
		getResponse()
	}
}

func printOptions() {
	fmt.Println("Let's Learn Go!")

	for i, o := range options {
		current := i + 1
		fmt.Printf("%d. %s\n", current, o.Title)
	}
}
func getResponse() {
	var toCheck string
	fmt.Scan(&toCheck)
	choice = validateChoice(toCheck)
}

func validateChoice(s string) int {
	// var result := strconv.
	i, err := strconv.Atoi(s)
	if err != nil {
		// handle error
		fmt.Println("Sorry, please pick a valid option.")
		return 1
	} else {
		return i
	}
}
