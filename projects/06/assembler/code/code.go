package code

import (
	"fmt"
	"os"
)

func New(outputFile string) *os.File {
	file, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	return file
}

func Close(file *os.File) {
	file.Close()
}

func writeln(file *os.File, s string) {
	file.WriteString(s + "\n")
}

func WriteAInstruction(file *os.File, value int) {
	binaryToStringWithPad0 := fmt.Sprintf("%015b", value)

	trimmedBinary := binaryToStringWithPad0[len(binaryToStringWithPad0)-15:]

	hackInstructionWithPrefix0 := "0" + trimmedBinary
	writeln(file, hackInstructionWithPrefix0)
}

func WriteCInstruction(file *os.File, dest, comp, jump string) {
	if dest == "" {
		dest = "null"
	}

	if jump == "" {
		jump = "null"
	}

	hackInstruction := fmt.Sprintf("111%s%s%s", computations[comp], destinations[dest], jumps[jump])
	writeln(file, hackInstruction)
}
