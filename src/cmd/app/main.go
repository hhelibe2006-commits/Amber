package main

import "fmt"

func main() {
	args, err := Args()
	if !args {
		fmt.Println(err)
	}
}
