/*
This is the entry point into the tool. All commands are run
from here.
*/

package main

import (
	"GitGo/helpers"
	"os"
)

func main() {
	args := os.Args[1:]

	helpers.Parser(args[0], args[1:])
}
