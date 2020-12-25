package main

import (
	"fmt"
	"os"

	"github.com/yousifh/chrono/chrono"
)

func main() {
	args := os.Args[1:]

	err := chrono.BuildSite(args[0])
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("done processing")
}
