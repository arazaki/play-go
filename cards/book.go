package main

import "fmt"

func book() {
	c := color("Red")

	fmt.Println(c.describe("is an awesome color"))
}

type color string

func (c color) describe(description string) string {
	return string(c) + " " + description
}
