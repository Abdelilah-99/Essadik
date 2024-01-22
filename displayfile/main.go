package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	lent := 0
	for i := range args {
		lent = i + 1
	}

	if lent > 1 {
		fmt.Println("Too many arguments")
	} else if lent == 0 {
		fmt.Println("File name missing")
	} else if args[0] == "quest8.txt" {

		cont, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Print(string(cont))

	}
}
