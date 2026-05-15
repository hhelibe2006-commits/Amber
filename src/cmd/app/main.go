package main

import (
	"fmt"
)

func main() {
	cfg, err := Args()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", cfg)

}
