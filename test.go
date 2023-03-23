package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	b := a * 2 / 60
	c := a * 2 % 60

	fmt.Println("It is", b, "hours", c, "minutes.")
}
