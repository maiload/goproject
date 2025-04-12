package main

import "fmt"

var g = 1

func main() {
	fmt.Println("Hello World!")

	var a int64 = 10
	var b = 20
	{
		c := 30
		fmt.Println(c)
	}

	fmt.Println(int(a) + b)
}
