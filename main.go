package main

import (
	"fmt"
	"os"

	"github.com/yousifh/chrono/chrono"
)

func main() {
	args := os.Args[1:]

	c := chrono.NewChrono(args[0])

	err := c.BuildSite()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("done processing")
}
