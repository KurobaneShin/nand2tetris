package main

import (
	"os"
)

func main() {
	asmSource := os.Args[1]
	hackOutput := os.Args[2]

	println("Assembling " + asmSource + " to " + hackOutput)
}
